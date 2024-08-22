// Code generated by go-swagger; DO NOT EDIT.

package derived_label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// ListDerivedLabelsReader is a Reader for the ListDerivedLabels structure.
type ListDerivedLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDerivedLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDerivedLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListDerivedLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListDerivedLabelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDerivedLabelsOK creates a ListDerivedLabelsOK with default headers values
func NewListDerivedLabelsOK() *ListDerivedLabelsOK {
	return &ListDerivedLabelsOK{}
}

/*
ListDerivedLabelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListDerivedLabelsOK struct {
	Payload *models.Configv1ListDerivedLabelsResponse
}

// IsSuccess returns true when this list derived labels o k response has a 2xx status code
func (o *ListDerivedLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list derived labels o k response has a 3xx status code
func (o *ListDerivedLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list derived labels o k response has a 4xx status code
func (o *ListDerivedLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list derived labels o k response has a 5xx status code
func (o *ListDerivedLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list derived labels o k response a status code equal to that given
func (o *ListDerivedLabelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list derived labels o k response
func (o *ListDerivedLabelsOK) Code() int {
	return 200
}

func (o *ListDerivedLabelsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels][%d] listDerivedLabelsOK  %+v", 200, o.Payload)
}

func (o *ListDerivedLabelsOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels][%d] listDerivedLabelsOK  %+v", 200, o.Payload)
}

func (o *ListDerivedLabelsOK) GetPayload() *models.Configv1ListDerivedLabelsResponse {
	return o.Payload
}

func (o *ListDerivedLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListDerivedLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDerivedLabelsInternalServerError creates a ListDerivedLabelsInternalServerError with default headers values
func NewListDerivedLabelsInternalServerError() *ListDerivedLabelsInternalServerError {
	return &ListDerivedLabelsInternalServerError{}
}

/*
ListDerivedLabelsInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListDerivedLabelsInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list derived labels internal server error response has a 2xx status code
func (o *ListDerivedLabelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list derived labels internal server error response has a 3xx status code
func (o *ListDerivedLabelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list derived labels internal server error response has a 4xx status code
func (o *ListDerivedLabelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list derived labels internal server error response has a 5xx status code
func (o *ListDerivedLabelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list derived labels internal server error response a status code equal to that given
func (o *ListDerivedLabelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list derived labels internal server error response
func (o *ListDerivedLabelsInternalServerError) Code() int {
	return 500
}

func (o *ListDerivedLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels][%d] listDerivedLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDerivedLabelsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels][%d] listDerivedLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDerivedLabelsInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListDerivedLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDerivedLabelsDefault creates a ListDerivedLabelsDefault with default headers values
func NewListDerivedLabelsDefault(code int) *ListDerivedLabelsDefault {
	return &ListDerivedLabelsDefault{
		_statusCode: code,
	}
}

/*
ListDerivedLabelsDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListDerivedLabelsDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list derived labels default response has a 2xx status code
func (o *ListDerivedLabelsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list derived labels default response has a 3xx status code
func (o *ListDerivedLabelsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list derived labels default response has a 4xx status code
func (o *ListDerivedLabelsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list derived labels default response has a 5xx status code
func (o *ListDerivedLabelsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list derived labels default response a status code equal to that given
func (o *ListDerivedLabelsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list derived labels default response
func (o *ListDerivedLabelsDefault) Code() int {
	return o._statusCode
}

func (o *ListDerivedLabelsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels][%d] ListDerivedLabels default  %+v", o._statusCode, o.Payload)
}

func (o *ListDerivedLabelsDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/derived-labels][%d] ListDerivedLabels default  %+v", o._statusCode, o.Payload)
}

func (o *ListDerivedLabelsDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListDerivedLabelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
