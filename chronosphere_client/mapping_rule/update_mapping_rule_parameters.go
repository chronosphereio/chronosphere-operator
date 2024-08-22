// Code generated by go-swagger; DO NOT EDIT.

package mapping_rule

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

// NewUpdateMappingRuleParams creates a new UpdateMappingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMappingRuleParams() *UpdateMappingRuleParams {
	return &UpdateMappingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMappingRuleParamsWithTimeout creates a new UpdateMappingRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateMappingRuleParamsWithTimeout(timeout time.Duration) *UpdateMappingRuleParams {
	return &UpdateMappingRuleParams{
		timeout: timeout,
	}
}

// NewUpdateMappingRuleParamsWithContext creates a new UpdateMappingRuleParams object
// with the ability to set a context for a request.
func NewUpdateMappingRuleParamsWithContext(ctx context.Context) *UpdateMappingRuleParams {
	return &UpdateMappingRuleParams{
		Context: ctx,
	}
}

// NewUpdateMappingRuleParamsWithHTTPClient creates a new UpdateMappingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMappingRuleParamsWithHTTPClient(client *http.Client) *UpdateMappingRuleParams {
	return &UpdateMappingRuleParams{
		HTTPClient: client,
	}
}

/*
UpdateMappingRuleParams contains all the parameters to send to the API endpoint

	for the update mapping rule operation.

	Typically these are written to a http.Request.
*/
type UpdateMappingRuleParams struct {

	// Body.
	Body *models.ConfigV1UpdateMappingRuleBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update mapping rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMappingRuleParams) WithDefaults() *UpdateMappingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mapping rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMappingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mapping rule params
func (o *UpdateMappingRuleParams) WithTimeout(timeout time.Duration) *UpdateMappingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mapping rule params
func (o *UpdateMappingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mapping rule params
func (o *UpdateMappingRuleParams) WithContext(ctx context.Context) *UpdateMappingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mapping rule params
func (o *UpdateMappingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mapping rule params
func (o *UpdateMappingRuleParams) WithHTTPClient(client *http.Client) *UpdateMappingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mapping rule params
func (o *UpdateMappingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update mapping rule params
func (o *UpdateMappingRuleParams) WithBody(body *models.ConfigV1UpdateMappingRuleBody) *UpdateMappingRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update mapping rule params
func (o *UpdateMappingRuleParams) SetBody(body *models.ConfigV1UpdateMappingRuleBody) {
	o.Body = body
}

// WithSlug adds the slug to the update mapping rule params
func (o *UpdateMappingRuleParams) WithSlug(slug string) *UpdateMappingRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update mapping rule params
func (o *UpdateMappingRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMappingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
