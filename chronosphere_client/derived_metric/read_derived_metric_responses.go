// Code generated by go-swagger; DO NOT EDIT.

package derived_metric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// ReadDerivedMetricReader is a Reader for the ReadDerivedMetric structure.
type ReadDerivedMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadDerivedMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadDerivedMetricOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadDerivedMetricNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadDerivedMetricInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadDerivedMetricDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadDerivedMetricOK creates a ReadDerivedMetricOK with default headers values
func NewReadDerivedMetricOK() *ReadDerivedMetricOK {
	return &ReadDerivedMetricOK{}
}

/*
ReadDerivedMetricOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadDerivedMetricOK struct {
	Payload *models.Configv1ReadDerivedMetricResponse
}

// IsSuccess returns true when this read derived metric o k response has a 2xx status code
func (o *ReadDerivedMetricOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read derived metric o k response has a 3xx status code
func (o *ReadDerivedMetricOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read derived metric o k response has a 4xx status code
func (o *ReadDerivedMetricOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read derived metric o k response has a 5xx status code
func (o *ReadDerivedMetricOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read derived metric o k response a status code equal to that given
func (o *ReadDerivedMetricOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read derived metric o k response
func (o *ReadDerivedMetricOK) Code() int {
	return 200
}

func (o *ReadDerivedMetricOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] readDerivedMetricOK  %+v", 200, o.Payload)
}

func (o *ReadDerivedMetricOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] readDerivedMetricOK  %+v", 200, o.Payload)
}

func (o *ReadDerivedMetricOK) GetPayload() *models.Configv1ReadDerivedMetricResponse {
	return o.Payload
}

func (o *ReadDerivedMetricOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadDerivedMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDerivedMetricNotFound creates a ReadDerivedMetricNotFound with default headers values
func NewReadDerivedMetricNotFound() *ReadDerivedMetricNotFound {
	return &ReadDerivedMetricNotFound{}
}

/*
ReadDerivedMetricNotFound describes a response with status code 404, with default header values.

Cannot read the DerivedMetric because the slug does not exist.
*/
type ReadDerivedMetricNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read derived metric not found response has a 2xx status code
func (o *ReadDerivedMetricNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read derived metric not found response has a 3xx status code
func (o *ReadDerivedMetricNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read derived metric not found response has a 4xx status code
func (o *ReadDerivedMetricNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read derived metric not found response has a 5xx status code
func (o *ReadDerivedMetricNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read derived metric not found response a status code equal to that given
func (o *ReadDerivedMetricNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read derived metric not found response
func (o *ReadDerivedMetricNotFound) Code() int {
	return 404
}

func (o *ReadDerivedMetricNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] readDerivedMetricNotFound  %+v", 404, o.Payload)
}

func (o *ReadDerivedMetricNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] readDerivedMetricNotFound  %+v", 404, o.Payload)
}

func (o *ReadDerivedMetricNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadDerivedMetricNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDerivedMetricInternalServerError creates a ReadDerivedMetricInternalServerError with default headers values
func NewReadDerivedMetricInternalServerError() *ReadDerivedMetricInternalServerError {
	return &ReadDerivedMetricInternalServerError{}
}

/*
ReadDerivedMetricInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadDerivedMetricInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read derived metric internal server error response has a 2xx status code
func (o *ReadDerivedMetricInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read derived metric internal server error response has a 3xx status code
func (o *ReadDerivedMetricInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read derived metric internal server error response has a 4xx status code
func (o *ReadDerivedMetricInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read derived metric internal server error response has a 5xx status code
func (o *ReadDerivedMetricInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read derived metric internal server error response a status code equal to that given
func (o *ReadDerivedMetricInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read derived metric internal server error response
func (o *ReadDerivedMetricInternalServerError) Code() int {
	return 500
}

func (o *ReadDerivedMetricInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] readDerivedMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDerivedMetricInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] readDerivedMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDerivedMetricInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadDerivedMetricInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDerivedMetricDefault creates a ReadDerivedMetricDefault with default headers values
func NewReadDerivedMetricDefault(code int) *ReadDerivedMetricDefault {
	return &ReadDerivedMetricDefault{
		_statusCode: code,
	}
}

/*
ReadDerivedMetricDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadDerivedMetricDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read derived metric default response has a 2xx status code
func (o *ReadDerivedMetricDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read derived metric default response has a 3xx status code
func (o *ReadDerivedMetricDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read derived metric default response has a 4xx status code
func (o *ReadDerivedMetricDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read derived metric default response has a 5xx status code
func (o *ReadDerivedMetricDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read derived metric default response a status code equal to that given
func (o *ReadDerivedMetricDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read derived metric default response
func (o *ReadDerivedMetricDefault) Code() int {
	return o._statusCode
}

func (o *ReadDerivedMetricDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] ReadDerivedMetric default  %+v", o._statusCode, o.Payload)
}

func (o *ReadDerivedMetricDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-metrics/{slug}][%d] ReadDerivedMetric default  %+v", o._statusCode, o.Payload)
}

func (o *ReadDerivedMetricDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadDerivedMetricDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
