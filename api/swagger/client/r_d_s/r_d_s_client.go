// Code generated by go-swagger; DO NOT EDIT.

package r_d_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new r d s API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for r d s API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Add add API
*/
func (a *Client) Add(params *AddParams) (*AddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Add",
		Method:             "POST",
		PathPattern:        "/v0/rds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOK), nil

}

/*
Discover queries parameters aws access key id aws secret access key
*/
func (a *Client) Discover(params *DiscoverParams) (*DiscoverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiscoverParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Discover",
		Method:             "GET",
		PathPattern:        "/v0/rds/discover",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DiscoverReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DiscoverOK), nil

}

/*
ListMixin2 list mixin2 API
*/
func (a *Client) ListMixin2(params *ListMixin2Params) (*ListMixin2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMixin2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListMixin2",
		Method:             "GET",
		PathPattern:        "/v0/rds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListMixin2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListMixin2OK), nil

}

/*
Remove remove API
*/
func (a *Client) Remove(params *RemoveParams) (*RemoveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Remove",
		Method:             "DELETE",
		PathPattern:        "/v0/rds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}