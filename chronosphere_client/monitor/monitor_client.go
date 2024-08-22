// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new monitor API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for monitor API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMonitor(params *CreateMonitorParams, opts ...ClientOption) (*CreateMonitorOK, error)

	DeleteMonitor(params *DeleteMonitorParams, opts ...ClientOption) (*DeleteMonitorOK, error)

	ListMonitors(params *ListMonitorsParams, opts ...ClientOption) (*ListMonitorsOK, error)

	ReadMonitor(params *ReadMonitorParams, opts ...ClientOption) (*ReadMonitorOK, error)

	UpdateMonitor(params *UpdateMonitorParams, opts ...ClientOption) (*UpdateMonitorOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateMonitor create monitor API
*/
func (a *Client) CreateMonitor(params *CreateMonitorParams, opts ...ClientOption) (*CreateMonitorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMonitorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateMonitor",
		Method:             "POST",
		PathPattern:        "/api/v1/config/monitors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMonitorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMonitorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateMonitorDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteMonitor delete monitor API
*/
func (a *Client) DeleteMonitor(params *DeleteMonitorParams, opts ...ClientOption) (*DeleteMonitorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMonitorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteMonitor",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/monitors/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMonitorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteMonitorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteMonitorDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListMonitors list monitors API
*/
func (a *Client) ListMonitors(params *ListMonitorsParams, opts ...ClientOption) (*ListMonitorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMonitorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListMonitors",
		Method:             "GET",
		PathPattern:        "/api/v1/config/monitors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListMonitorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMonitorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListMonitorsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadMonitor read monitor API
*/
func (a *Client) ReadMonitor(params *ReadMonitorParams, opts ...ClientOption) (*ReadMonitorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadMonitorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadMonitor",
		Method:             "GET",
		PathPattern:        "/api/v1/config/monitors/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadMonitorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadMonitorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadMonitorDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateMonitor update monitor API
*/
func (a *Client) UpdateMonitor(params *UpdateMonitorParams, opts ...ClientOption) (*UpdateMonitorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMonitorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateMonitor",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/monitors/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMonitorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMonitorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateMonitorDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
