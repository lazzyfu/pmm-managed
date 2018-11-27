// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/agents.proto

package inventory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AgentProcessStatus represents agent process state.
type AgentProcessStatus int32

const (
	AgentProcessStatus_AGENT_PROCESS_STATUS_INVALID AgentProcessStatus = 0
	AgentProcessStatus_DISABLED                     AgentProcessStatus = 1
	AgentProcessStatus_RUNNING                      AgentProcessStatus = 2
)

var AgentProcessStatus_name = map[int32]string{
	0: "AGENT_PROCESS_STATUS_INVALID",
	1: "DISABLED",
	2: "RUNNING",
}
var AgentProcessStatus_value = map[string]int32{
	"AGENT_PROCESS_STATUS_INVALID": 0,
	"DISABLED":                     1,
	"RUNNING":                      2,
}

func (x AgentProcessStatus) String() string {
	return proto.EnumName(AgentProcessStatus_name, int32(x))
}
func (AgentProcessStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{0}
}

// MySQLdExporter represents mysqld_exporter Agent configuration.
type MySQLdExporter struct {
	// Unique Agent identifier.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Node identifier where Agent runs.
	RunsOnNodeId uint32 `protobuf:"varint,2,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	// Node identifiers for which insights are provided by that Agent.
	NodeIds []uint32 `protobuf:"varint,3,rep,packed,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	// Service identifiers for which insights are provided by that Agent.
	ServiceIds []uint32 `protobuf:"varint,4,rep,packed,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
	// Agent process status.
	ProcessStatus AgentProcessStatus `protobuf:"varint,5,opt,name=process_status,json=processStatus,proto3,enum=inventory.AgentProcessStatus" json:"process_status,omitempty"`
	// HTTP listen port for exposing metrics.
	ListenPort uint32 `protobuf:"varint,6,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	// MySQL username for extracting metrics.
	Username string `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	// MySQL password for extracting metrics.
	Password             string   `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLdExporter) Reset()         { *m = MySQLdExporter{} }
func (m *MySQLdExporter) String() string { return proto.CompactTextString(m) }
func (*MySQLdExporter) ProtoMessage()    {}
func (*MySQLdExporter) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{0}
}
func (m *MySQLdExporter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLdExporter.Unmarshal(m, b)
}
func (m *MySQLdExporter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLdExporter.Marshal(b, m, deterministic)
}
func (dst *MySQLdExporter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLdExporter.Merge(dst, src)
}
func (m *MySQLdExporter) XXX_Size() int {
	return xxx_messageInfo_MySQLdExporter.Size(m)
}
func (m *MySQLdExporter) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLdExporter.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLdExporter proto.InternalMessageInfo

func (m *MySQLdExporter) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MySQLdExporter) GetRunsOnNodeId() uint32 {
	if m != nil {
		return m.RunsOnNodeId
	}
	return 0
}

func (m *MySQLdExporter) GetNodeIds() []uint32 {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *MySQLdExporter) GetServiceIds() []uint32 {
	if m != nil {
		return m.ServiceIds
	}
	return nil
}

func (m *MySQLdExporter) GetProcessStatus() AgentProcessStatus {
	if m != nil {
		return m.ProcessStatus
	}
	return AgentProcessStatus_AGENT_PROCESS_STATUS_INVALID
}

func (m *MySQLdExporter) GetListenPort() uint32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

func (m *MySQLdExporter) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MySQLdExporter) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ListAgentsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAgentsRequest) Reset()         { *m = ListAgentsRequest{} }
func (m *ListAgentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAgentsRequest) ProtoMessage()    {}
func (*ListAgentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{1}
}
func (m *ListAgentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentsRequest.Unmarshal(m, b)
}
func (m *ListAgentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentsRequest.Marshal(b, m, deterministic)
}
func (dst *ListAgentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentsRequest.Merge(dst, src)
}
func (m *ListAgentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAgentsRequest.Size(m)
}
func (m *ListAgentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentsRequest proto.InternalMessageInfo

type ListAgentsResponse struct {
	MysqldExporter       []*MySQLdExporter `protobuf:"bytes,1,rep,name=mysqld_exporter,json=mysqldExporter,proto3" json:"mysqld_exporter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListAgentsResponse) Reset()         { *m = ListAgentsResponse{} }
func (m *ListAgentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAgentsResponse) ProtoMessage()    {}
func (*ListAgentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{2}
}
func (m *ListAgentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentsResponse.Unmarshal(m, b)
}
func (m *ListAgentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentsResponse.Marshal(b, m, deterministic)
}
func (dst *ListAgentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentsResponse.Merge(dst, src)
}
func (m *ListAgentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAgentsResponse.Size(m)
}
func (m *ListAgentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentsResponse proto.InternalMessageInfo

func (m *ListAgentsResponse) GetMysqldExporter() []*MySQLdExporter {
	if m != nil {
		return m.MysqldExporter
	}
	return nil
}

type GetAgentRequest struct {
	// Unique Agent identifier.
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentRequest) Reset()         { *m = GetAgentRequest{} }
func (m *GetAgentRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgentRequest) ProtoMessage()    {}
func (*GetAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{3}
}
func (m *GetAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentRequest.Unmarshal(m, b)
}
func (m *GetAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentRequest.Marshal(b, m, deterministic)
}
func (dst *GetAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentRequest.Merge(dst, src)
}
func (m *GetAgentRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgentRequest.Size(m)
}
func (m *GetAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentRequest proto.InternalMessageInfo

func (m *GetAgentRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAgentResponse struct {
	// Types that are valid to be assigned to Agent:
	//	*GetAgentResponse_MysqldExporter
	Agent                isGetAgentResponse_Agent `protobuf_oneof:"agent"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetAgentResponse) Reset()         { *m = GetAgentResponse{} }
func (m *GetAgentResponse) String() string { return proto.CompactTextString(m) }
func (*GetAgentResponse) ProtoMessage()    {}
func (*GetAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{4}
}
func (m *GetAgentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentResponse.Unmarshal(m, b)
}
func (m *GetAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentResponse.Marshal(b, m, deterministic)
}
func (dst *GetAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentResponse.Merge(dst, src)
}
func (m *GetAgentResponse) XXX_Size() int {
	return xxx_messageInfo_GetAgentResponse.Size(m)
}
func (m *GetAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentResponse proto.InternalMessageInfo

type isGetAgentResponse_Agent interface {
	isGetAgentResponse_Agent()
}

type GetAgentResponse_MysqldExporter struct {
	MysqldExporter *MySQLdExporter `protobuf:"bytes,1,opt,name=mysqld_exporter,json=mysqldExporter,proto3,oneof"`
}

func (*GetAgentResponse_MysqldExporter) isGetAgentResponse_Agent() {}

func (m *GetAgentResponse) GetAgent() isGetAgentResponse_Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *GetAgentResponse) GetMysqldExporter() *MySQLdExporter {
	if x, ok := m.GetAgent().(*GetAgentResponse_MysqldExporter); ok {
		return x.MysqldExporter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetAgentResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetAgentResponse_OneofMarshaler, _GetAgentResponse_OneofUnmarshaler, _GetAgentResponse_OneofSizer, []interface{}{
		(*GetAgentResponse_MysqldExporter)(nil),
	}
}

func _GetAgentResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetAgentResponse)
	// agent
	switch x := m.Agent.(type) {
	case *GetAgentResponse_MysqldExporter:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MysqldExporter); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GetAgentResponse.Agent has unexpected type %T", x)
	}
	return nil
}

func _GetAgentResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetAgentResponse)
	switch tag {
	case 1: // agent.mysqld_exporter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MySQLdExporter)
		err := b.DecodeMessage(msg)
		m.Agent = &GetAgentResponse_MysqldExporter{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GetAgentResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetAgentResponse)
	// agent
	switch x := m.Agent.(type) {
	case *GetAgentResponse_MysqldExporter:
		s := proto.Size(x.MysqldExporter)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AddMySQLdExporterAgentRequest struct {
	// Node identifier where Agent runs.
	RunsOnNodeId uint32 `protobuf:"varint,2,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	// Node identifiers for which insights are provided by that Agent.
	NodeIds []uint32 `protobuf:"varint,3,rep,packed,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	// Service identifiers for which insights are provided by that Agent.
	ServiceIds []uint32 `protobuf:"varint,4,rep,packed,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
	// MySQL username for extracting metrics.
	Username string `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	// MySQL password for extracting metrics.
	Password             string   `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLdExporterAgentRequest) Reset()         { *m = AddMySQLdExporterAgentRequest{} }
func (m *AddMySQLdExporterAgentRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLdExporterAgentRequest) ProtoMessage()    {}
func (*AddMySQLdExporterAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{5}
}
func (m *AddMySQLdExporterAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLdExporterAgentRequest.Unmarshal(m, b)
}
func (m *AddMySQLdExporterAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLdExporterAgentRequest.Marshal(b, m, deterministic)
}
func (dst *AddMySQLdExporterAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLdExporterAgentRequest.Merge(dst, src)
}
func (m *AddMySQLdExporterAgentRequest) XXX_Size() int {
	return xxx_messageInfo_AddMySQLdExporterAgentRequest.Size(m)
}
func (m *AddMySQLdExporterAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLdExporterAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLdExporterAgentRequest proto.InternalMessageInfo

func (m *AddMySQLdExporterAgentRequest) GetRunsOnNodeId() uint32 {
	if m != nil {
		return m.RunsOnNodeId
	}
	return 0
}

func (m *AddMySQLdExporterAgentRequest) GetNodeIds() []uint32 {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *AddMySQLdExporterAgentRequest) GetServiceIds() []uint32 {
	if m != nil {
		return m.ServiceIds
	}
	return nil
}

func (m *AddMySQLdExporterAgentRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddMySQLdExporterAgentRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AddMySQLdExporterAgentResponse struct {
	MysqldExporter       *MySQLdExporter `protobuf:"bytes,1,opt,name=mysqld_exporter,json=mysqldExporter,proto3" json:"mysqld_exporter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddMySQLdExporterAgentResponse) Reset()         { *m = AddMySQLdExporterAgentResponse{} }
func (m *AddMySQLdExporterAgentResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLdExporterAgentResponse) ProtoMessage()    {}
func (*AddMySQLdExporterAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{6}
}
func (m *AddMySQLdExporterAgentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLdExporterAgentResponse.Unmarshal(m, b)
}
func (m *AddMySQLdExporterAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLdExporterAgentResponse.Marshal(b, m, deterministic)
}
func (dst *AddMySQLdExporterAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLdExporterAgentResponse.Merge(dst, src)
}
func (m *AddMySQLdExporterAgentResponse) XXX_Size() int {
	return xxx_messageInfo_AddMySQLdExporterAgentResponse.Size(m)
}
func (m *AddMySQLdExporterAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLdExporterAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLdExporterAgentResponse proto.InternalMessageInfo

func (m *AddMySQLdExporterAgentResponse) GetMysqldExporter() *MySQLdExporter {
	if m != nil {
		return m.MysqldExporter
	}
	return nil
}

type RemoveAgentRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAgentRequest) Reset()         { *m = RemoveAgentRequest{} }
func (m *RemoveAgentRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveAgentRequest) ProtoMessage()    {}
func (*RemoveAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{7}
}
func (m *RemoveAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAgentRequest.Unmarshal(m, b)
}
func (m *RemoveAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAgentRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAgentRequest.Merge(dst, src)
}
func (m *RemoveAgentRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveAgentRequest.Size(m)
}
func (m *RemoveAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAgentRequest proto.InternalMessageInfo

func (m *RemoveAgentRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RemoveAgentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAgentResponse) Reset()         { *m = RemoveAgentResponse{} }
func (m *RemoveAgentResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveAgentResponse) ProtoMessage()    {}
func (*RemoveAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_68e561a4371a2dd5, []int{8}
}
func (m *RemoveAgentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAgentResponse.Unmarshal(m, b)
}
func (m *RemoveAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAgentResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAgentResponse.Merge(dst, src)
}
func (m *RemoveAgentResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveAgentResponse.Size(m)
}
func (m *RemoveAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAgentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MySQLdExporter)(nil), "inventory.MySQLdExporter")
	proto.RegisterType((*ListAgentsRequest)(nil), "inventory.ListAgentsRequest")
	proto.RegisterType((*ListAgentsResponse)(nil), "inventory.ListAgentsResponse")
	proto.RegisterType((*GetAgentRequest)(nil), "inventory.GetAgentRequest")
	proto.RegisterType((*GetAgentResponse)(nil), "inventory.GetAgentResponse")
	proto.RegisterType((*AddMySQLdExporterAgentRequest)(nil), "inventory.AddMySQLdExporterAgentRequest")
	proto.RegisterType((*AddMySQLdExporterAgentResponse)(nil), "inventory.AddMySQLdExporterAgentResponse")
	proto.RegisterType((*RemoveAgentRequest)(nil), "inventory.RemoveAgentRequest")
	proto.RegisterType((*RemoveAgentResponse)(nil), "inventory.RemoveAgentResponse")
	proto.RegisterEnum("inventory.AgentProcessStatus", AgentProcessStatus_name, AgentProcessStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentsClient is the client API for Agents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentsClient interface {
	// ListAgents returns a list of all Agents.
	ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
	// GetAgent returns a single Agent by ID.
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentResponse, error)
	// AddMySQLdExporterAgent adds mysqld_exporter Agent.
	AddMySQLdExporterAgent(ctx context.Context, in *AddMySQLdExporterAgentRequest, opts ...grpc.CallOption) (*AddMySQLdExporterAgentResponse, error)
	// RemoveAgent removes Agent.
	RemoveAgent(ctx context.Context, in *RemoveAgentRequest, opts ...grpc.CallOption) (*RemoveAgentResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, "/inventory.Agents/ListAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentResponse, error) {
	out := new(GetAgentResponse)
	err := c.cc.Invoke(ctx, "/inventory.Agents/GetAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) AddMySQLdExporterAgent(ctx context.Context, in *AddMySQLdExporterAgentRequest, opts ...grpc.CallOption) (*AddMySQLdExporterAgentResponse, error) {
	out := new(AddMySQLdExporterAgentResponse)
	err := c.cc.Invoke(ctx, "/inventory.Agents/AddMySQLdExporterAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) RemoveAgent(ctx context.Context, in *RemoveAgentRequest, opts ...grpc.CallOption) (*RemoveAgentResponse, error) {
	out := new(RemoveAgentResponse)
	err := c.cc.Invoke(ctx, "/inventory.Agents/RemoveAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentsServer is the server API for Agents service.
type AgentsServer interface {
	// ListAgents returns a list of all Agents.
	ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error)
	// GetAgent returns a single Agent by ID.
	GetAgent(context.Context, *GetAgentRequest) (*GetAgentResponse, error)
	// AddMySQLdExporterAgent adds mysqld_exporter Agent.
	AddMySQLdExporterAgent(context.Context, *AddMySQLdExporterAgentRequest) (*AddMySQLdExporterAgentResponse, error)
	// RemoveAgent removes Agent.
	RemoveAgent(context.Context, *RemoveAgentRequest) (*RemoveAgentResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Agents/ListAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).ListAgents(ctx, req.(*ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Agents/GetAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_AddMySQLdExporterAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLdExporterAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).AddMySQLdExporterAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Agents/AddMySQLdExporterAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).AddMySQLdExporterAgent(ctx, req.(*AddMySQLdExporterAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_RemoveAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).RemoveAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Agents/RemoveAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).RemoveAgent(ctx, req.(*RemoveAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAgents",
			Handler:    _Agents_ListAgents_Handler,
		},
		{
			MethodName: "GetAgent",
			Handler:    _Agents_GetAgent_Handler,
		},
		{
			MethodName: "AddMySQLdExporterAgent",
			Handler:    _Agents_AddMySQLdExporterAgent_Handler,
		},
		{
			MethodName: "RemoveAgent",
			Handler:    _Agents_RemoveAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/agents.proto",
}

func init() { proto.RegisterFile("inventory/agents.proto", fileDescriptor_agents_68e561a4371a2dd5) }

var fileDescriptor_agents_68e561a4371a2dd5 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcb, 0x4e, 0xdb, 0x40,
	0x14, 0x65, 0x4c, 0x48, 0xcc, 0x0d, 0x04, 0x77, 0x50, 0x91, 0x71, 0x09, 0xb8, 0x23, 0x90, 0xdc,
	0x20, 0x11, 0x44, 0xa5, 0x2e, 0xd8, 0x85, 0x26, 0xa2, 0x41, 0xa9, 0x01, 0x1b, 0xaa, 0xee, 0x2c,
	0x17, 0x8f, 0x90, 0x25, 0x32, 0x63, 0x3c, 0x4e, 0xda, 0x6c, 0xfb, 0x07, 0x55, 0xff, 0xa0, 0x52,
	0xff, 0xa6, 0xbb, 0xfe, 0x42, 0x3f, 0xa4, 0xf2, 0x23, 0xc4, 0x01, 0x13, 0xb5, 0x5d, 0x74, 0x39,
	0xf7, 0x9c, 0x3b, 0xe7, 0xdc, 0xb9, 0x47, 0x03, 0x6b, 0x3e, 0x1b, 0x52, 0x16, 0xf1, 0x70, 0xd4,
	0x74, 0xaf, 0x29, 0x8b, 0xc4, 0x5e, 0x10, 0xf2, 0x88, 0xe3, 0xc5, 0xbb, 0xba, 0xb6, 0x71, 0xcd,
	0xf9, 0xf5, 0x0d, 0x6d, 0xba, 0x81, 0xdf, 0x74, 0x19, 0xe3, 0x91, 0x1b, 0xf9, 0x9c, 0x65, 0x44,
	0xf2, 0x4d, 0x82, 0xda, 0xdb, 0x91, 0x7d, 0xde, 0xf3, 0x3a, 0x9f, 0x02, 0x1e, 0x46, 0x34, 0xc4,
	0x35, 0x90, 0x7c, 0x4f, 0x45, 0x3a, 0x32, 0x96, 0x2d, 0xc9, 0xf7, 0xf0, 0x0e, 0xac, 0x84, 0x03,
	0x26, 0x1c, 0xce, 0x1c, 0xc6, 0x3d, 0xea, 0xf8, 0x9e, 0x2a, 0x25, 0xe0, 0x52, 0x5c, 0x3e, 0x65,
	0x26, 0xf7, 0x68, 0xd7, 0xc3, 0xeb, 0x20, 0x67, 0xb0, 0x50, 0xe7, 0xf5, 0x79, 0x63, 0xd9, 0xaa,
	0xb0, 0x04, 0x11, 0x78, 0x0b, 0xaa, 0x82, 0x86, 0x43, 0xff, 0x2a, 0x45, 0x4b, 0x09, 0x0a, 0x59,
	0x29, 0x26, 0xb4, 0xa1, 0x16, 0x84, 0xfc, 0x8a, 0x0a, 0xe1, 0x88, 0xc8, 0x8d, 0x06, 0x42, 0x5d,
	0xd0, 0x91, 0x51, 0x3b, 0xa8, 0xef, 0xdd, 0xcd, 0xb1, 0xd7, 0x8a, 0xe7, 0x3b, 0x4b, 0x59, 0x76,
	0x42, 0xb2, 0x96, 0x83, 0xfc, 0x31, 0x96, 0xb9, 0xf1, 0x45, 0x44, 0x99, 0x13, 0x4f, 0xa2, 0x96,
	0x13, 0x93, 0x90, 0x96, 0xce, 0x78, 0x18, 0x61, 0x0d, 0xe4, 0x81, 0xa0, 0x21, 0x73, 0xfb, 0x54,
	0xad, 0xe8, 0xc8, 0x58, 0xb4, 0xee, 0xce, 0x31, 0x16, 0xb8, 0x42, 0x7c, 0xe4, 0xa1, 0xa7, 0xca,
	0x29, 0x36, 0x3e, 0x93, 0x55, 0x78, 0xd2, 0xf3, 0x45, 0x94, 0x38, 0x10, 0x16, 0xbd, 0x1d, 0x50,
	0x11, 0x91, 0xf7, 0x80, 0xf3, 0x45, 0x11, 0x70, 0x26, 0x28, 0x3e, 0x82, 0x95, 0xfe, 0x48, 0xdc,
	0xde, 0x78, 0x0e, 0xcd, 0xde, 0x53, 0x45, 0xfa, 0xbc, 0x51, 0x3d, 0x58, 0xcf, 0x8d, 0x32, 0xfd,
	0xe0, 0x56, 0x2d, 0xed, 0x18, 0x9f, 0xc9, 0x73, 0x58, 0x39, 0xa6, 0xe9, 0xc5, 0x99, 0xd8, 0xfd,
	0x9d, 0x10, 0x17, 0x94, 0x09, 0x25, 0x93, 0x6e, 0x17, 0x49, 0xa3, 0x99, 0xd2, 0x6f, 0xe6, 0xee,
	0x8b, 0x1f, 0x55, 0x60, 0x21, 0x49, 0x12, 0xf9, 0x81, 0xa0, 0xde, 0xf2, 0xbc, 0xe9, 0x86, 0x29,
	0x53, 0xff, 0x21, 0x18, 0xff, 0xb8, 0xb1, 0x93, 0x92, 0x8c, 0x14, 0xe9, 0xa4, 0x24, 0x2f, 0x28,
	0xe5, 0x93, 0x92, 0x5c, 0x56, 0x2a, 0xc4, 0x83, 0xcd, 0xc7, 0xa6, 0x99, 0xb5, 0x3a, 0xf4, 0x77,
	0xab, 0xdb, 0x06, 0x6c, 0xd1, 0x3e, 0x1f, 0xd2, 0x99, 0xdb, 0x7b, 0x0a, 0xab, 0x53, 0xac, 0xd4,
	0x40, 0xe3, 0x1c, 0xf0, 0xc3, 0x90, 0x63, 0x1d, 0x36, 0x5a, 0xc7, 0x1d, 0xf3, 0xc2, 0x39, 0xb3,
	0x4e, 0x5f, 0x77, 0x6c, 0xdb, 0xb1, 0x2f, 0x5a, 0x17, 0x97, 0xb6, 0xd3, 0x35, 0xdf, 0xb5, 0x7a,
	0xdd, 0xb6, 0x32, 0x87, 0x97, 0x40, 0x6e, 0x77, 0xed, 0xd6, 0x51, 0xaf, 0xd3, 0x56, 0x10, 0xae,
	0x42, 0xc5, 0xba, 0x34, 0xcd, 0xae, 0x79, 0xac, 0x48, 0x07, 0x5f, 0x4a, 0x50, 0x4e, 0x13, 0x8a,
	0x07, 0x00, 0x93, 0xbc, 0xe2, 0x8d, 0xdc, 0x4c, 0x0f, 0xb2, 0xad, 0xd5, 0x1f, 0x41, 0x53, 0xa3,
	0xa4, 0xf1, 0xf9, 0xe7, 0xaf, 0xaf, 0xd2, 0x36, 0xd9, 0x6a, 0x0e, 0xf7, 0x9b, 0x93, 0x1f, 0x28,
	0x65, 0x35, 0x27, 0x0d, 0x87, 0xa8, 0x81, 0xfb, 0x20, 0x8f, 0x93, 0x8a, 0xb5, 0xdc, 0xb5, 0xf7,
	0x12, 0xae, 0x3d, 0x2b, 0xc4, 0x32, 0x41, 0x23, 0x11, 0x24, 0xa4, 0x5e, 0x28, 0x38, 0xa6, 0xc7,
	0x72, 0xdf, 0x11, 0xac, 0x15, 0xef, 0x19, 0x1b, 0xf9, 0xcf, 0x64, 0x56, 0xb0, 0xb5, 0x17, 0x7f,
	0xc0, 0xcc, 0x9c, 0xbd, 0x4a, 0x9c, 0xed, 0x93, 0xdd, 0x42, 0x67, 0xc5, 0xcd, 0xb1, 0xcf, 0x11,
	0x54, 0x73, 0x11, 0xc0, 0xf9, 0x07, 0x7f, 0x18, 0x20, 0x6d, 0xf3, 0x31, 0x38, 0x73, 0xb1, 0x9b,
	0xb8, 0xd8, 0x21, 0x7a, 0xa1, 0x8b, 0x5c, 0xc7, 0x21, 0x6a, 0x7c, 0x28, 0x27, 0x3f, 0xff, 0xcb,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xd2, 0x68, 0x4d, 0x3c, 0x06, 0x00, 0x00,
}