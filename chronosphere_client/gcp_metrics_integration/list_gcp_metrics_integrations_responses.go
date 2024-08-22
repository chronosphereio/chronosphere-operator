// Code generated by go-swagger; DO NOT EDIT.

package gcp_metrics_integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// ListGcpMetricsIntegrationsReader is a Reader for the ListGcpMetricsIntegrations structure.
type ListGcpMetricsIntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGcpMetricsIntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGcpMetricsIntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListGcpMetricsIntegrationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListGcpMetricsIntegrationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGcpMetricsIntegrationsOK creates a ListGcpMetricsIntegrationsOK with default headers values
func NewListGcpMetricsIntegrationsOK() *ListGcpMetricsIntegrationsOK {
	return &ListGcpMetricsIntegrationsOK{}
}

/*
ListGcpMetricsIntegrationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListGcpMetricsIntegrationsOK struct {
	Payload *models.Configv1ListGcpMetricsIntegrationsResponse
}

// IsSuccess returns true when this list gcp metrics integrations o k response has a 2xx status code
func (o *ListGcpMetricsIntegrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list gcp metrics integrations o k response has a 3xx status code
func (o *ListGcpMetricsIntegrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list gcp metrics integrations o k response has a 4xx status code
func (o *ListGcpMetricsIntegrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list gcp metrics integrations o k response has a 5xx status code
func (o *ListGcpMetricsIntegrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list gcp metrics integrations o k response a status code equal to that given
func (o *ListGcpMetricsIntegrationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list gcp metrics integrations o k response
func (o *ListGcpMetricsIntegrationsOK) Code() int {
	return 200
}

func (o *ListGcpMetricsIntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/gcp-metrics-integrations][%d] listGcpMetricsIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListGcpMetricsIntegrationsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/gcp-metrics-integrations][%d] listGcpMetricsIntegrationsOK  %+v", 200, o.Payload)
}

func (o *ListGcpMetricsIntegrationsOK) GetPayload() *models.Configv1ListGcpMetricsIntegrationsResponse {
	return o.Payload
}

func (o *ListGcpMetricsIntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListGcpMetricsIntegrationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGcpMetricsIntegrationsInternalServerError creates a ListGcpMetricsIntegrationsInternalServerError with default headers values
func NewListGcpMetricsIntegrationsInternalServerError() *ListGcpMetricsIntegrationsInternalServerError {
	return &ListGcpMetricsIntegrationsInternalServerError{}
}

/*
ListGcpMetricsIntegrationsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListGcpMetricsIntegrationsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list gcp metrics integrations internal server error response has a 2xx status code
func (o *ListGcpMetricsIntegrationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list gcp metrics integrations internal server error response has a 3xx status code
func (o *ListGcpMetricsIntegrationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list gcp metrics integrations internal server error response has a 4xx status code
func (o *ListGcpMetricsIntegrationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list gcp metrics integrations internal server error response has a 5xx status code
func (o *ListGcpMetricsIntegrationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list gcp metrics integrations internal server error response a status code equal to that given
func (o *ListGcpMetricsIntegrationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list gcp metrics integrations internal server error response
func (o *ListGcpMetricsIntegrationsInternalServerError) Code() int {
	return 500
}

func (o *ListGcpMetricsIntegrationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/gcp-metrics-integrations][%d] listGcpMetricsIntegrationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGcpMetricsIntegrationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/gcp-metrics-integrations][%d] listGcpMetricsIntegrationsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGcpMetricsIntegrationsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListGcpMetricsIntegrationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGcpMetricsIntegrationsDefault creates a ListGcpMetricsIntegrationsDefault with default headers values
func NewListGcpMetricsIntegrationsDefault(code int) *ListGcpMetricsIntegrationsDefault {
	return &ListGcpMetricsIntegrationsDefault{
		_statusCode: code,
	}
}

/*
ListGcpMetricsIntegrationsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListGcpMetricsIntegrationsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list gcp metrics integrations default response has a 2xx status code
func (o *ListGcpMetricsIntegrationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list gcp metrics integrations default response has a 3xx status code
func (o *ListGcpMetricsIntegrationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list gcp metrics integrations default response has a 4xx status code
func (o *ListGcpMetricsIntegrationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list gcp metrics integrations default response has a 5xx status code
func (o *ListGcpMetricsIntegrationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list gcp metrics integrations default response a status code equal to that given
func (o *ListGcpMetricsIntegrationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list gcp metrics integrations default response
func (o *ListGcpMetricsIntegrationsDefault) Code() int {
	return o._statusCode
}

func (o *ListGcpMetricsIntegrationsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/gcp-metrics-integrations][%d] ListGcpMetricsIntegrations default  %+v", o._statusCode, o.Payload)
}

func (o *ListGcpMetricsIntegrationsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/gcp-metrics-integrations][%d] ListGcpMetricsIntegrations default  %+v", o._statusCode, o.Payload)
}

func (o *ListGcpMetricsIntegrationsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListGcpMetricsIntegrationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
