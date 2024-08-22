// Code generated by go-swagger; DO NOT EDIT.

package gcp_metrics_integration

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

// NewCreateGcpMetricsIntegrationParams creates a new CreateGcpMetricsIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGcpMetricsIntegrationParams() *CreateGcpMetricsIntegrationParams {
	return &CreateGcpMetricsIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGcpMetricsIntegrationParamsWithTimeout creates a new CreateGcpMetricsIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateGcpMetricsIntegrationParamsWithTimeout(timeout time.Duration) *CreateGcpMetricsIntegrationParams {
	return &CreateGcpMetricsIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateGcpMetricsIntegrationParamsWithContext creates a new CreateGcpMetricsIntegrationParams object
// with the ability to set a context for a request.
func NewCreateGcpMetricsIntegrationParamsWithContext(ctx context.Context) *CreateGcpMetricsIntegrationParams {
	return &CreateGcpMetricsIntegrationParams{
		Context: ctx,
	}
}

// NewCreateGcpMetricsIntegrationParamsWithHTTPClient creates a new CreateGcpMetricsIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGcpMetricsIntegrationParamsWithHTTPClient(client *http.Client) *CreateGcpMetricsIntegrationParams {
	return &CreateGcpMetricsIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateGcpMetricsIntegrationParams contains all the parameters to send to the API endpoint

	for the create gcp metrics integration operation.

	Typically these are written to a http.Request.
*/
type CreateGcpMetricsIntegrationParams struct {

	// Body.
	Body *models.Configv1CreateGcpMetricsIntegrationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create gcp metrics integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGcpMetricsIntegrationParams) WithDefaults() *CreateGcpMetricsIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create gcp metrics integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGcpMetricsIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) WithTimeout(timeout time.Duration) *CreateGcpMetricsIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) WithContext(ctx context.Context) *CreateGcpMetricsIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) WithHTTPClient(client *http.Client) *CreateGcpMetricsIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) WithBody(body *models.Configv1CreateGcpMetricsIntegrationRequest) *CreateGcpMetricsIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create gcp metrics integration params
func (o *CreateGcpMetricsIntegrationParams) SetBody(body *models.Configv1CreateGcpMetricsIntegrationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGcpMetricsIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
