// Code generated by go-swagger; DO NOT EDIT.

package bucket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// ListBucketsReader is a Reader for the ListBuckets structure.
type ListBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListBucketsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListBucketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBucketsOK creates a ListBucketsOK with default headers values
func NewListBucketsOK() *ListBucketsOK {
	return &ListBucketsOK{}
}

/*
ListBucketsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListBucketsOK struct {
	Payload *models.Configv1ListBucketsResponse
}

// IsSuccess returns true when this list buckets o k response has a 2xx status code
func (o *ListBucketsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list buckets o k response has a 3xx status code
func (o *ListBucketsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list buckets o k response has a 4xx status code
func (o *ListBucketsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list buckets o k response has a 5xx status code
func (o *ListBucketsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list buckets o k response a status code equal to that given
func (o *ListBucketsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list buckets o k response
func (o *ListBucketsOK) Code() int {
	return 200
}

func (o *ListBucketsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/buckets][%d] listBucketsOK  %+v", 200, o.Payload)
}

func (o *ListBucketsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/buckets][%d] listBucketsOK  %+v", 200, o.Payload)
}

func (o *ListBucketsOK) GetPayload() *models.Configv1ListBucketsResponse {
	return o.Payload
}

func (o *ListBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListBucketsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBucketsInternalServerError creates a ListBucketsInternalServerError with default headers values
func NewListBucketsInternalServerError() *ListBucketsInternalServerError {
	return &ListBucketsInternalServerError{}
}

/*
ListBucketsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListBucketsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list buckets internal server error response has a 2xx status code
func (o *ListBucketsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list buckets internal server error response has a 3xx status code
func (o *ListBucketsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list buckets internal server error response has a 4xx status code
func (o *ListBucketsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list buckets internal server error response has a 5xx status code
func (o *ListBucketsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list buckets internal server error response a status code equal to that given
func (o *ListBucketsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list buckets internal server error response
func (o *ListBucketsInternalServerError) Code() int {
	return 500
}

func (o *ListBucketsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/buckets][%d] listBucketsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListBucketsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/buckets][%d] listBucketsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListBucketsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListBucketsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBucketsDefault creates a ListBucketsDefault with default headers values
func NewListBucketsDefault(code int) *ListBucketsDefault {
	return &ListBucketsDefault{
		_statusCode: code,
	}
}

/*
ListBucketsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListBucketsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list buckets default response has a 2xx status code
func (o *ListBucketsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list buckets default response has a 3xx status code
func (o *ListBucketsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list buckets default response has a 4xx status code
func (o *ListBucketsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list buckets default response has a 5xx status code
func (o *ListBucketsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list buckets default response a status code equal to that given
func (o *ListBucketsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list buckets default response
func (o *ListBucketsDefault) Code() int {
	return o._statusCode
}

func (o *ListBucketsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/buckets][%d] ListBuckets default  %+v", o._statusCode, o.Payload)
}

func (o *ListBucketsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/buckets][%d] ListBuckets default  %+v", o._statusCode, o.Payload)
}

func (o *ListBucketsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListBucketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
