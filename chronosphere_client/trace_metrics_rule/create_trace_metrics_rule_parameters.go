// Code generated by go-swagger; DO NOT EDIT.

package trace_metrics_rule

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

// NewCreateTraceMetricsRuleParams creates a new CreateTraceMetricsRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTraceMetricsRuleParams() *CreateTraceMetricsRuleParams {
	return &CreateTraceMetricsRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTraceMetricsRuleParamsWithTimeout creates a new CreateTraceMetricsRuleParams object
// with the ability to set a timeout on a request.
func NewCreateTraceMetricsRuleParamsWithTimeout(timeout time.Duration) *CreateTraceMetricsRuleParams {
	return &CreateTraceMetricsRuleParams{
		timeout: timeout,
	}
}

// NewCreateTraceMetricsRuleParamsWithContext creates a new CreateTraceMetricsRuleParams object
// with the ability to set a context for a request.
func NewCreateTraceMetricsRuleParamsWithContext(ctx context.Context) *CreateTraceMetricsRuleParams {
	return &CreateTraceMetricsRuleParams{
		Context: ctx,
	}
}

// NewCreateTraceMetricsRuleParamsWithHTTPClient creates a new CreateTraceMetricsRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTraceMetricsRuleParamsWithHTTPClient(client *http.Client) *CreateTraceMetricsRuleParams {
	return &CreateTraceMetricsRuleParams{
		HTTPClient: client,
	}
}

/*
CreateTraceMetricsRuleParams contains all the parameters to send to the API endpoint

	for the create trace metrics rule operation.

	Typically these are written to a http.Request.
*/
type CreateTraceMetricsRuleParams struct {

	// Body.
	Body *models.Configv1CreateTraceMetricsRuleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create trace metrics rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTraceMetricsRuleParams) WithDefaults() *CreateTraceMetricsRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create trace metrics rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTraceMetricsRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) WithTimeout(timeout time.Duration) *CreateTraceMetricsRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) WithContext(ctx context.Context) *CreateTraceMetricsRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) WithHTTPClient(client *http.Client) *CreateTraceMetricsRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) WithBody(body *models.Configv1CreateTraceMetricsRuleRequest) *CreateTraceMetricsRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create trace metrics rule params
func (o *CreateTraceMetricsRuleParams) SetBody(body *models.Configv1CreateTraceMetricsRuleRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTraceMetricsRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
