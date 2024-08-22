// Code generated by go-swagger; DO NOT EDIT.

package team

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
	"github.com/go-openapi/swag"
)

// NewListTeamsParams creates a new ListTeamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTeamsParams() *ListTeamsParams {
	return &ListTeamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTeamsParamsWithTimeout creates a new ListTeamsParams object
// with the ability to set a timeout on a request.
func NewListTeamsParamsWithTimeout(timeout time.Duration) *ListTeamsParams {
	return &ListTeamsParams{
		timeout: timeout,
	}
}

// NewListTeamsParamsWithContext creates a new ListTeamsParams object
// with the ability to set a context for a request.
func NewListTeamsParamsWithContext(ctx context.Context) *ListTeamsParams {
	return &ListTeamsParams{
		Context: ctx,
	}
}

// NewListTeamsParamsWithHTTPClient creates a new ListTeamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTeamsParamsWithHTTPClient(client *http.Client) *ListTeamsParams {
	return &ListTeamsParams{
		HTTPClient: client,
	}
}

/*
ListTeamsParams contains all the parameters to send to the API endpoint

	for the list teams operation.

	Typically these are written to a http.Request.
*/
type ListTeamsParams struct {

	/* Names.

	   Filters results by name, where any Team with a matching name in the given list (and matches all other filters) is returned.
	*/
	Names []string

	/* PageMaxSize.

	     Page size preference (i.e. how many items are returned in the next
	page). If zero, the server will use a default. Regardless of what size
	is given, clients must never assume how many items will be returned.

	     Format: int64
	*/
	PageMaxSize *int64

	/* PageToken.

	     Opaque page token identifying which page to request. An empty token
	identifies the first page.
	*/
	PageToken *string

	/* Slugs.

	   Filters results by slug, where any Team with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTeamsParams) WithDefaults() *ListTeamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTeamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list teams params
func (o *ListTeamsParams) WithTimeout(timeout time.Duration) *ListTeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list teams params
func (o *ListTeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list teams params
func (o *ListTeamsParams) WithContext(ctx context.Context) *ListTeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list teams params
func (o *ListTeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list teams params
func (o *ListTeamsParams) WithHTTPClient(client *http.Client) *ListTeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list teams params
func (o *ListTeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the list teams params
func (o *ListTeamsParams) WithNames(names []string) *ListTeamsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list teams params
func (o *ListTeamsParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list teams params
func (o *ListTeamsParams) WithPageMaxSize(pageMaxSize *int64) *ListTeamsParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list teams params
func (o *ListTeamsParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list teams params
func (o *ListTeamsParams) WithPageToken(pageToken *string) *ListTeamsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list teams params
func (o *ListTeamsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list teams params
func (o *ListTeamsParams) WithSlugs(slugs []string) *ListTeamsParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list teams params
func (o *ListTeamsParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListTeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.PageMaxSize != nil {

		// query param page.max_size
		var qrPageMaxSize int64

		if o.PageMaxSize != nil {
			qrPageMaxSize = *o.PageMaxSize
		}
		qPageMaxSize := swag.FormatInt64(qrPageMaxSize)
		if qPageMaxSize != "" {

			if err := r.SetQueryParam("page.max_size", qPageMaxSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page.token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("page.token", qPageToken); err != nil {
				return err
			}
		}
	}

	if o.Slugs != nil {

		// binding items for slugs
		joinedSlugs := o.bindParamSlugs(reg)

		// query array param slugs
		if err := r.SetQueryParam("slugs", joinedSlugs...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListTeams binds the parameter names
func (o *ListTeamsParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: "multi"
	namesIS := swag.JoinByFormat(namesIC, "multi")

	return namesIS
}

// bindParamListTeams binds the parameter slugs
func (o *ListTeamsParams) bindParamSlugs(formats strfmt.Registry) []string {
	slugsIR := o.Slugs

	var slugsIC []string
	for _, slugsIIR := range slugsIR { // explode []string

		slugsIIV := slugsIIR // string as string
		slugsIC = append(slugsIC, slugsIIV)
	}

	// items.CollectionFormat: "multi"
	slugsIS := swag.JoinByFormat(slugsIC, "multi")

	return slugsIS
}
