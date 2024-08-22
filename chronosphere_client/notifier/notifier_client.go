// Code generated by go-swagger; DO NOT EDIT.

package notifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new notifier API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for notifier API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNotifier(params *CreateNotifierParams, opts ...ClientOption) (*CreateNotifierOK, error)

	DeleteNotifier(params *DeleteNotifierParams, opts ...ClientOption) (*DeleteNotifierOK, error)

	ListNotifiers(params *ListNotifiersParams, opts ...ClientOption) (*ListNotifiersOK, error)

	ReadNotifier(params *ReadNotifierParams, opts ...ClientOption) (*ReadNotifierOK, error)

	UpdateNotifier(params *UpdateNotifierParams, opts ...ClientOption) (*UpdateNotifierOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateNotifier create notifier API
*/
func (a *Client) CreateNotifier(params *CreateNotifierParams, opts ...ClientOption) (*CreateNotifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNotifierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateNotifier",
		Method:             "POST",
		PathPattern:        "/api/v1/config/notifiers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNotifierReader{formats: a.formats},
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
	success, ok := result.(*CreateNotifierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateNotifierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteNotifier delete notifier API
*/
func (a *Client) DeleteNotifier(params *DeleteNotifierParams, opts ...ClientOption) (*DeleteNotifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNotifierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNotifier",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/notifiers/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNotifierReader{formats: a.formats},
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
	success, ok := result.(*DeleteNotifierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNotifierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNotifiers list notifiers API
*/
func (a *Client) ListNotifiers(params *ListNotifiersParams, opts ...ClientOption) (*ListNotifiersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNotifiersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNotifiers",
		Method:             "GET",
		PathPattern:        "/api/v1/config/notifiers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNotifiersReader{formats: a.formats},
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
	success, ok := result.(*ListNotifiersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNotifiersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadNotifier read notifier API
*/
func (a *Client) ReadNotifier(params *ReadNotifierParams, opts ...ClientOption) (*ReadNotifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadNotifierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadNotifier",
		Method:             "GET",
		PathPattern:        "/api/v1/config/notifiers/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadNotifierReader{formats: a.formats},
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
	success, ok := result.(*ReadNotifierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadNotifierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateNotifier update notifier API
*/
func (a *Client) UpdateNotifier(params *UpdateNotifierParams, opts ...ClientOption) (*UpdateNotifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNotifierParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateNotifier",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/notifiers/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNotifierReader{formats: a.formats},
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
	success, ok := result.(*UpdateNotifierOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateNotifierDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
