// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"context"
	"database/sql"
	_ "expvar" // register /debug/vars
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof" // register /debug/pprof
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_gateway "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/api/inventorypb"
	"github.com/percona/pmm/api/managementpb"
	"github.com/percona/pmm/api/serverpb"
	"github.com/percona/pmm/utils/sqlmetrics"
	"github.com/percona/pmm/version"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/agents"
	agentgrpc "github.com/percona/pmm-managed/services/agents/grpc"
	"github.com/percona/pmm-managed/services/grafana"
	"github.com/percona/pmm-managed/services/inventory"
	inventorygrpc "github.com/percona/pmm-managed/services/inventory/grpc"
	"github.com/percona/pmm-managed/services/management"
	managementgrpc "github.com/percona/pmm-managed/services/management/grpc"
	"github.com/percona/pmm-managed/services/prometheus"
	"github.com/percona/pmm-managed/services/qan"
	"github.com/percona/pmm-managed/services/server"
	"github.com/percona/pmm-managed/services/supervisord"
	"github.com/percona/pmm-managed/services/telemetry"
	"github.com/percona/pmm-managed/utils/interceptors"
	"github.com/percona/pmm-managed/utils/logger"
)

const (
	shutdownTimeout = 3 * time.Second

	gRPCAddr  = "127.0.0.1:7771"
	http1Addr = "127.0.0.1:7772"
	debugAddr = "127.0.0.1:7773"

	defaultAlertManagerFile  = "/srv/prometheus/rules/pmm.rules.yml"
	prometheusBaseConfigFile = "/srv/prometheus/prometheus.base.yml"
)

func addLogsHandler(mux *http.ServeMux, logs *supervisord.Logs) {
	l := logrus.WithField("component", "logs.zip")

	mux.HandleFunc("/logs.zip", func(rw http.ResponseWriter, req *http.Request) {
		// fail-safe
		ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
		defer cancel()

		filename := fmt.Sprintf("pmm-server_%s.zip", time.Now().UTC().Format("2006-01-02_15-04"))
		rw.Header().Set(`Access-Control-Allow-Origin`, `*`)
		rw.Header().Set(`Content-Type`, `application/zip`)
		rw.Header().Set(`Content-Disposition`, `attachment; filename="`+filename+`"`)

		ctx = logger.Set(ctx, "logs")
		if err := logs.Zip(ctx, rw); err != nil {
			l.Error(err)
		}
	})
}

type gRPCServerDeps struct {
	db             *reform.DB
	prometheus     *prometheus.Service
	server         *server.Server
	agentsRegistry *agents.Registry
}

// runGRPCServer runs gRPC server until context is canceled, then gracefully stops it.
func runGRPCServer(ctx context.Context, deps *gRPCServerDeps) {
	l := logrus.WithField("component", "gRPC")
	l.Infof("Starting server on http://%s/ ...", gRPCAddr)

	gRPCServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(10*1024*1024), //nolint:gomnd

		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptors.Unary,
			grpc_validator.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			interceptors.Stream,
			grpc_validator.StreamServerInterceptor(),
		)),
	)

	serverpb.RegisterServerServer(gRPCServer, deps.server)

	agentpb.RegisterAgentServer(gRPCServer, agentgrpc.NewAgentServer(deps.agentsRegistry))

	nodesSvc := inventory.NewNodesService(deps.db)
	servicesSvc := inventory.NewServicesService(deps.db, deps.agentsRegistry)
	agentsSvc := inventory.NewAgentsService(deps.db, deps.agentsRegistry)

	inventorypb.RegisterNodesServer(gRPCServer, inventorygrpc.NewNodesServer(nodesSvc))
	inventorypb.RegisterServicesServer(gRPCServer, inventorygrpc.NewServicesServer(servicesSvc))
	inventorypb.RegisterAgentsServer(gRPCServer, inventorygrpc.NewAgentsServer(agentsSvc))

	nodeSvc := management.NewNodeService(deps.db, deps.agentsRegistry)
	serviceSvc := management.NewServiceService(deps.db, deps.agentsRegistry)
	mysqlSvc := management.NewMySQLService(deps.db, deps.agentsRegistry)
	mongodbSvc := management.NewMongoDBService(deps.db, deps.agentsRegistry)
	postgresqlSvc := management.NewPostgreSQLService(deps.db, deps.agentsRegistry)
	proxysqlSvc := management.NewProxySQLService(deps.db, deps.agentsRegistry)

	managementpb.RegisterNodeServer(gRPCServer, managementgrpc.NewManagementNodeServer(nodeSvc))
	managementpb.RegisterServiceServer(gRPCServer, managementgrpc.NewManagementServiceServer(serviceSvc))
	managementpb.RegisterMySQLServer(gRPCServer, managementgrpc.NewManagementMySQLServer(mysqlSvc))
	managementpb.RegisterMongoDBServer(gRPCServer, managementgrpc.NewManagementMongoDBServer(mongodbSvc))
	managementpb.RegisterPostgreSQLServer(gRPCServer, managementgrpc.NewManagementPostgreSQLServer(postgresqlSvc))
	managementpb.RegisterProxySQLServer(gRPCServer, managementgrpc.NewManagementProxySQLServer(proxysqlSvc))
	managementpb.RegisterActionsServer(gRPCServer, managementgrpc.NewActionsServer(deps.agentsRegistry, deps.db))
	managementpb.RegisterRDSServer(gRPCServer, management.NewRDSService(deps.db, deps.agentsRegistry))

	if l.Logger.GetLevel() >= logrus.DebugLevel {
		l.Debug("Reflection and channelz are enabled.")
		reflection.Register(gRPCServer)
		channelz.RegisterChannelzServiceToServer(gRPCServer)

		l.Debug("RPC response latency histogram enabled.")
		grpc_prometheus.EnableHandlingTimeHistogram()
	}

	grpc_prometheus.Register(gRPCServer)

	// run server until it is stopped gracefully or not
	listener, err := net.Listen("tcp", gRPCAddr)
	if err != nil {
		l.Panic(err)
	}
	go func() {
		for {
			err = gRPCServer.Serve(listener)
			if err == nil || err == grpc.ErrServerStopped {
				break
			}
			l.Errorf("Failed to serve: %s", err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()

	// try to stop server gracefully, then not
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	go func() {
		<-ctx.Done()
		gRPCServer.Stop()
	}()
	gRPCServer.GracefulStop()
	cancel()
}

type http1ServerDeps struct {
	logs       *supervisord.Logs
	authServer *grafana.AuthServer
}

// runHTTP1Server runs grpc-gateway and other HTTP 1.1 APIs (like auth_request and logs.zip)
// until context is canceled, then gracefully stops it.
func runHTTP1Server(ctx context.Context, deps *http1ServerDeps) {
	l := logrus.WithField("component", "JSON")
	l.Infof("Starting server on http://%s/ ...", http1Addr)

	proxyMux := grpc_gateway.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// TODO switch from RegisterXXXHandlerFromEndpoint to RegisterXXXHandler to avoid extra dials
	// (even if they dial to localhost)
	// https://jira.percona.com/browse/PMM-4326
	type registrar func(context.Context, *grpc_gateway.ServeMux, string, []grpc.DialOption) error
	for _, r := range []registrar{
		serverpb.RegisterServerHandlerFromEndpoint,

		inventorypb.RegisterNodesHandlerFromEndpoint,
		inventorypb.RegisterServicesHandlerFromEndpoint,
		inventorypb.RegisterAgentsHandlerFromEndpoint,

		managementpb.RegisterNodeHandlerFromEndpoint,
		managementpb.RegisterServiceHandlerFromEndpoint,
		managementpb.RegisterMySQLHandlerFromEndpoint,
		managementpb.RegisterMongoDBHandlerFromEndpoint,
		managementpb.RegisterPostgreSQLHandlerFromEndpoint,
		managementpb.RegisterProxySQLHandlerFromEndpoint,
		managementpb.RegisterActionsHandlerFromEndpoint,
		managementpb.RegisterRDSHandlerFromEndpoint,
	} {
		if err := r(ctx, proxyMux, gRPCAddr, opts); err != nil {
			l.Panic(err)
		}
	}

	mux := http.NewServeMux()
	addLogsHandler(mux, deps.logs)
	mux.Handle("/auth_request", deps.authServer)
	mux.Handle("/", proxyMux)

	server := &http.Server{
		Addr:     http1Addr,
		ErrorLog: log.New(os.Stderr, "runJSONServer: ", 0),
		Handler:  mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Panic(err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Failed to shutdown gracefully: %s", err)
	}
	cancel()
}

// runDebugServer runs debug server until context is canceled, then gracefully stops it.
// TODO merge with HTTP1 server? https://jira.percona.com/browse/PMM-4326
func runDebugServer(ctx context.Context) {
	handler := promhttp.HandlerFor(prom.DefaultGatherer, promhttp.HandlerOpts{
		ErrorLog:      logrus.WithField("component", "metrics"),
		ErrorHandling: promhttp.ContinueOnError,
	})
	http.Handle("/debug/metrics", promhttp.InstrumentMetricHandler(prom.DefaultRegisterer, handler))

	l := logrus.WithField("component", "debug")

	handlers := []string{
		"/debug/metrics",  // by http.Handle above
		"/debug/vars",     // by expvar
		"/debug/requests", // by golang.org/x/net/trace imported by google.golang.org/grpc
		"/debug/events",   // by golang.org/x/net/trace imported by google.golang.org/grpc
		"/debug/pprof",    // by net/http/pprof
	}
	for i, h := range handlers {
		handlers[i] = "http://" + debugAddr + h
	}

	var buf bytes.Buffer
	err := template.Must(template.New("debug").Parse(`
	<html>
	<body>
	<ul>
	{{ range . }}
		<li><a href="{{ . }}">{{ . }}</a></li>
	{{ end }}
	</ul>
	</body>
	</html>
	`)).Execute(&buf, handlers)
	if err != nil {
		l.Panic(err)
	}
	http.HandleFunc("/debug", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(buf.Bytes())
	})
	l.Infof("Starting server on http://%s/debug\nRegistered handlers:\n\t%s", debugAddr, strings.Join(handlers, "\n\t"))

	server := &http.Server{
		Addr:     debugAddr,
		ErrorLog: log.New(os.Stderr, "runDebugServer: ", 0),
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Panic(err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Failed to shutdown gracefully: %s", err)
	}
	cancel()
}

type setupDeps struct {
	sqlDB       *sql.DB
	dbUsername  string
	dbPassword  string
	supervisord *supervisord.Service
	prometheus  *prometheus.Service
	server      *server.Server
	l           *logrus.Entry
}

// setup migrates database and performs other setup tasks that depend on database.
func setup(ctx context.Context, deps *setupDeps) bool {
	deps.l.Infof("Migrating database...")
	db, err := models.SetupDB(deps.sqlDB, &models.SetupDBParams{
		Logf:          deps.l.Debugf,
		Username:      deps.dbUsername,
		Password:      deps.dbPassword,
		SetupFixtures: models.SetupFixtures,
	})
	if err != nil {
		deps.l.Warnf("Failed to migrate database: %s.", err)
		return false
	}

	// log and ignore validation errors; fail on other errors
	deps.l.Infof("Updating settings...")
	env := os.Environ()
	sort.Strings(env)
	if errs := deps.server.UpdateSettingsFromEnv(env); len(errs) != 0 {
		// This should be impossible in the normal workflow.
		// An invalid environment variable must be caught with pmm-managed-init
		// and the docker run must be terminated.
		deps.l.Errorln("Failed to update settings from environment:")
		for _, e := range errs {
			deps.l.Errorln(e)
		}
		return false
	}

	deps.l.Infof("Updating supervisord configuration...")
	settings, err := models.GetSettings(db.Querier)
	if err != nil {
		deps.l.Warnf("Failed to get settings: %+v.", err)
		return false
	}
	if err = deps.supervisord.UpdateConfiguration(settings); err != nil {
		deps.l.Warnf("Failed to update supervisord configuration: %+v.", err)
		return false
	}

	deps.l.Infof("Checking Prometheus...")
	if err = deps.prometheus.IsReady(ctx); err != nil {
		deps.l.Warnf("Prometheus problem: %+v.", err)
		return false
	}
	deps.prometheus.RequestConfigurationUpdate()

	deps.l.Info("Setup completed.")
	return true
}

func getQANClient(ctx context.Context, sqlDB *sql.DB, dbName, qanAPIAddr string) *qan.Client {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second),
		grpc.WithUserAgent("pmm-managed/" + version.Version),
	}

	// Without grpc.WithBlock() DialContext returns an error only if something very wrong with address or options;
	// it does not return an error of connection failure but tries to reconnect in the background.
	conn, err := grpc.DialContext(ctx, qanAPIAddr, opts...)
	if err != nil {
		logrus.Fatalf("Failed to connect QAN API %s: %s.", qanAPIAddr, err)
	}

	l := logrus.WithField("component", "reform/qan")
	reformL := sqlmetrics.NewReform("postgres", dbName+"/qan", l.Tracef)
	prom.MustRegister(reformL)
	db := reform.NewDB(sqlDB, postgresql.Dialect, reformL)
	return qan.NewClient(conn, db)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("stdlog: ")

	kingpin.Version(version.FullInfo())
	kingpin.HelpFlag.Short('h')

	prometheusConfigF := kingpin.Flag("prometheus-config", "Prometheus configuration file path").Required().String()
	prometheusURLF := kingpin.Flag("prometheus-url", "Prometheus base URL").Default("http://127.0.0.1:9090/prometheus/").String()
	promtoolF := kingpin.Flag("promtool", "promtool path").Default("promtool").String()

	grafanaAddrF := kingpin.Flag("grafana-addr", "Grafana HTTP API address").Default("127.0.0.1:3000").String()
	qanAPIAddrF := kingpin.Flag("qan-api-addr", "QAN API gRPC API address").Default("127.0.0.1:9911").String()

	postgresAddrF := kingpin.Flag("postgres-addr", "PostgreSQL address").Default("127.0.0.1:5432").String()
	postgresDBNameF := kingpin.Flag("postgres-name", "PostgreSQL database name").Required().String()
	postgresDBUsernameF := kingpin.Flag("postgres-username", "PostgreSQL database username").Default("pmm-managed").String()
	postgresDBPasswordF := kingpin.Flag("postgres-password", "PostgreSQL database password").Default("pmm-managed").String()

	supervisordConfigDirF := kingpin.Flag("supervisord-config-dir", "Supervisord configuration directory").Required().String()

	alertManagerRulesFileF := kingpin.Flag("alert-manager-rules-file", "Path to the Alert Manager Rules file").
		Default(defaultAlertManagerFile).String()

	debugF := kingpin.Flag("debug", "Enable debug logging").Envar("PMM_DEBUG").Bool()
	traceF := kingpin.Flag("trace", "Enable trace logging (implies debug)").Envar("PMM_TRACE").Bool()

	kingpin.Parse()

	logrus.SetFormatter(&logrus.TextFormatter{
		// Enable multiline-friendly formatter in both development (with terminal) and production (without terminal):
		// https://github.com/sirupsen/logrus/blob/839c75faf7f98a33d445d181f3018b5c3409a45e/text_formatter.go#L176-L178
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000-07:00",

		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			_, function = filepath.Split(f.Function)

			// keep a single directory name as a compromise between brevity and unambiguity
			var dir string
			dir, file = filepath.Split(f.File)
			dir = filepath.Base(dir)
			file = fmt.Sprintf("%s/%s:%d", dir, file, f.Line)

			return
		},
	})
	if *debugF {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if *traceF {
		logrus.SetLevel(logrus.TraceLevel)
		grpclog.SetLoggerV2(&logger.GRPC{Entry: logrus.WithField("component", "grpclog")})
		logrus.SetReportCaller(true)
	}
	logrus.Infof("Log level: %s.", logrus.GetLevel())

	l := logrus.WithField("component", "main")
	ctx, cancel := context.WithCancel(context.Background())
	ctx = logger.Set(ctx, "main")
	defer l.Info("Done.")

	// handle termination signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT)
	go func() {
		s := <-signals
		signal.Stop(signals)
		logrus.Warnf("Got %s, shutting down...", unix.SignalName(s.(unix.Signal)))
		cancel()
	}()

	sqlDB, err := models.OpenDB(*postgresAddrF, *postgresDBNameF, *postgresDBUsernameF, *postgresDBPasswordF)
	if err != nil {
		l.Panicf("Failed to connect to database: %+v", err)
	}
	defer sqlDB.Close() //nolint:errcheck
	prom.MustRegister(sqlmetrics.NewCollector("postgres", *postgresDBNameF, sqlDB))
	reformL := sqlmetrics.NewReform("postgres", *postgresDBNameF, logrus.WithField("component", "reform").Tracef)
	prom.MustRegister(reformL)
	db := reform.NewDB(sqlDB, postgresql.Dialect, reformL)

	prometheus, err := prometheus.NewService(*prometheusConfigF, prometheusBaseConfigFile, *promtoolF, db, *prometheusURLF)
	if err != nil {
		l.Panicf("Prometheus service problem: %+v", err)
	}

	pmmUpdateCheck := supervisord.NewPMMUpdateChecker(logrus.WithField("component", "supervisord/pmm-update-checker"))

	logs := supervisord.NewLogs(version.FullInfo(), pmmUpdateCheck)
	supervisord := supervisord.New(*supervisordConfigDirF, pmmUpdateCheck)
	telemetry := telemetry.NewService(db, version.Version)
	awsInstanceChecker := server.NewAWSInstanceChecker(db, telemetry)
	grafanaClient := grafana.NewClient(*grafanaAddrF)
	prom.MustRegister(grafanaClient)

	serverParams := &server.ServerParams{
		DB:                 db,
		Prometheus:         prometheus,
		Supervisord:        supervisord,
		TelemetryService:   telemetry,
		AwsInstanceChecker: awsInstanceChecker,
		AlertManagerFile:   *alertManagerRulesFileF,
		GrafanaClient:      grafanaClient,
	}
	server, err := server.NewServer(serverParams)
	if err != nil {
		l.Panicf("Server problem: %+v", err)
	}

	// try synchronously once, then retry in the background
	deps := &setupDeps{
		sqlDB:       sqlDB,
		dbUsername:  *postgresDBUsernameF,
		dbPassword:  *postgresDBPasswordF,
		supervisord: supervisord,
		prometheus:  prometheus,
		server:      server,
		l:           logrus.WithField("component", "setup"),
	}
	if !setup(ctx, deps) {
		go func() {
			const delay = 2 * time.Second
			for {
				deps.l.Warnf("Retrying in %s.", delay)
				sleepCtx, sleepCancel := context.WithTimeout(ctx, delay)
				<-sleepCtx.Done()
				sleepCancel()

				if ctx.Err() != nil {
					return
				}

				if setup(ctx, deps) {
					return
				}
			}
		}()
	}

	qanClient := getQANClient(ctx, sqlDB, *postgresDBNameF, *qanAPIAddrF)

	agentsRegistry := agents.NewRegistry(db, prometheus, qanClient)
	prom.MustRegister(agentsRegistry)

	authServer := grafana.NewAuthServer(grafanaClient, awsInstanceChecker)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		prometheus.Run(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		supervisord.Run(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		telemetry.Run(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runGRPCServer(ctx, &gRPCServerDeps{
			db:             db,
			prometheus:     prometheus,
			server:         server,
			agentsRegistry: agentsRegistry,
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runHTTP1Server(ctx, &http1ServerDeps{
			logs:       logs,
			authServer: authServer,
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runDebugServer(ctx)
	}()

	wg.Wait()
}
