// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// NewCreateMonitorParams creates a new CreateMonitorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMonitorParams() *CreateMonitorParams {
	return &CreateMonitorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMonitorParamsWithTimeout creates a new CreateMonitorParams object
// with the ability to set a timeout on a request.
func NewCreateMonitorParamsWithTimeout(timeout time.Duration) *CreateMonitorParams {
	return &CreateMonitorParams{
		timeout: timeout,
	}
}

// NewCreateMonitorParamsWithContext creates a new CreateMonitorParams object
// with the ability to set a context for a request.
func NewCreateMonitorParamsWithContext(ctx context.Context) *CreateMonitorParams {
	return &CreateMonitorParams{
		Context: ctx,
	}
}

// NewCreateMonitorParamsWithHTTPClient creates a new CreateMonitorParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMonitorParamsWithHTTPClient(client *http.Client) *CreateMonitorParams {
	return &CreateMonitorParams{
		HTTPClient: client,
	}
}

/*
CreateMonitorParams contains all the parameters to send to the API endpoint

	for the create monitor operation.

	Typically these are written to a http.Request.
*/
type CreateMonitorParams struct {

	// Body.
	Body *models.Configv1CreateMonitorRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create monitor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMonitorParams) WithDefaults() *CreateMonitorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create monitor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMonitorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create monitor params
func (o *CreateMonitorParams) WithTimeout(timeout time.Duration) *CreateMonitorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create monitor params
func (o *CreateMonitorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create monitor params
func (o *CreateMonitorParams) WithContext(ctx context.Context) *CreateMonitorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create monitor params
func (o *CreateMonitorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create monitor params
func (o *CreateMonitorParams) WithHTTPClient(client *http.Client) *CreateMonitorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create monitor params
func (o *CreateMonitorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create monitor params
func (o *CreateMonitorParams) WithBody(body *models.Configv1CreateMonitorRequest) *CreateMonitorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create monitor params
func (o *CreateMonitorParams) SetBody(body *models.Configv1CreateMonitorRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMonitorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
