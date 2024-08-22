// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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
)

// NewReadDashboardParams creates a new ReadDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadDashboardParams() *ReadDashboardParams {
	return &ReadDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadDashboardParamsWithTimeout creates a new ReadDashboardParams object
// with the ability to set a timeout on a request.
func NewReadDashboardParamsWithTimeout(timeout time.Duration) *ReadDashboardParams {
	return &ReadDashboardParams{
		timeout: timeout,
	}
}

// NewReadDashboardParamsWithContext creates a new ReadDashboardParams object
// with the ability to set a context for a request.
func NewReadDashboardParamsWithContext(ctx context.Context) *ReadDashboardParams {
	return &ReadDashboardParams{
		Context: ctx,
	}
}

// NewReadDashboardParamsWithHTTPClient creates a new ReadDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadDashboardParamsWithHTTPClient(client *http.Client) *ReadDashboardParams {
	return &ReadDashboardParams{
		HTTPClient: client,
	}
}

/*
ReadDashboardParams contains all the parameters to send to the API endpoint

	for the read dashboard operation.

	Typically these are written to a http.Request.
*/
type ReadDashboardParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDashboardParams) WithDefaults() *ReadDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read dashboard params
func (o *ReadDashboardParams) WithTimeout(timeout time.Duration) *ReadDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read dashboard params
func (o *ReadDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read dashboard params
func (o *ReadDashboardParams) WithContext(ctx context.Context) *ReadDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read dashboard params
func (o *ReadDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read dashboard params
func (o *ReadDashboardParams) WithHTTPClient(client *http.Client) *ReadDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read dashboard params
func (o *ReadDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read dashboard params
func (o *ReadDashboardParams) WithSlug(slug string) *ReadDashboardParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read dashboard params
func (o *ReadDashboardParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
