// Code generated by go-swagger; DO NOT EDIT.

package trace_tail_sampling_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// ReadTraceTailSamplingRulesReader is a Reader for the ReadTraceTailSamplingRules structure.
type ReadTraceTailSamplingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadTraceTailSamplingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadTraceTailSamplingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadTraceTailSamplingRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadTraceTailSamplingRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadTraceTailSamplingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadTraceTailSamplingRulesOK creates a ReadTraceTailSamplingRulesOK with default headers values
func NewReadTraceTailSamplingRulesOK() *ReadTraceTailSamplingRulesOK {
	return &ReadTraceTailSamplingRulesOK{}
}

/*
ReadTraceTailSamplingRulesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadTraceTailSamplingRulesOK struct {
	Payload *models.Configv1ReadTraceTailSamplingRulesResponse
}

// IsSuccess returns true when this read trace tail sampling rules o k response has a 2xx status code
func (o *ReadTraceTailSamplingRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read trace tail sampling rules o k response has a 3xx status code
func (o *ReadTraceTailSamplingRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace tail sampling rules o k response has a 4xx status code
func (o *ReadTraceTailSamplingRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace tail sampling rules o k response has a 5xx status code
func (o *ReadTraceTailSamplingRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace tail sampling rules o k response a status code equal to that given
func (o *ReadTraceTailSamplingRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read trace tail sampling rules o k response
func (o *ReadTraceTailSamplingRulesOK) Code() int {
	return 200
}

func (o *ReadTraceTailSamplingRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] readTraceTailSamplingRulesOK  %+v", 200, o.Payload)
}

func (o *ReadTraceTailSamplingRulesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] readTraceTailSamplingRulesOK  %+v", 200, o.Payload)
}

func (o *ReadTraceTailSamplingRulesOK) GetPayload() *models.Configv1ReadTraceTailSamplingRulesResponse {
	return o.Payload
}

func (o *ReadTraceTailSamplingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadTraceTailSamplingRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceTailSamplingRulesNotFound creates a ReadTraceTailSamplingRulesNotFound with default headers values
func NewReadTraceTailSamplingRulesNotFound() *ReadTraceTailSamplingRulesNotFound {
	return &ReadTraceTailSamplingRulesNotFound{}
}

/*
ReadTraceTailSamplingRulesNotFound describes a response with status code 404, with default header values.

Cannot read the TraceTailSamplingRules because TraceTailSamplingRules has not been created.
*/
type ReadTraceTailSamplingRulesNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace tail sampling rules not found response has a 2xx status code
func (o *ReadTraceTailSamplingRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace tail sampling rules not found response has a 3xx status code
func (o *ReadTraceTailSamplingRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace tail sampling rules not found response has a 4xx status code
func (o *ReadTraceTailSamplingRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read trace tail sampling rules not found response has a 5xx status code
func (o *ReadTraceTailSamplingRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace tail sampling rules not found response a status code equal to that given
func (o *ReadTraceTailSamplingRulesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read trace tail sampling rules not found response
func (o *ReadTraceTailSamplingRulesNotFound) Code() int {
	return 404
}

func (o *ReadTraceTailSamplingRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] readTraceTailSamplingRulesNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceTailSamplingRulesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] readTraceTailSamplingRulesNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceTailSamplingRulesNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceTailSamplingRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceTailSamplingRulesInternalServerError creates a ReadTraceTailSamplingRulesInternalServerError with default headers values
func NewReadTraceTailSamplingRulesInternalServerError() *ReadTraceTailSamplingRulesInternalServerError {
	return &ReadTraceTailSamplingRulesInternalServerError{}
}

/*
ReadTraceTailSamplingRulesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadTraceTailSamplingRulesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace tail sampling rules internal server error response has a 2xx status code
func (o *ReadTraceTailSamplingRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace tail sampling rules internal server error response has a 3xx status code
func (o *ReadTraceTailSamplingRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace tail sampling rules internal server error response has a 4xx status code
func (o *ReadTraceTailSamplingRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace tail sampling rules internal server error response has a 5xx status code
func (o *ReadTraceTailSamplingRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read trace tail sampling rules internal server error response a status code equal to that given
func (o *ReadTraceTailSamplingRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read trace tail sampling rules internal server error response
func (o *ReadTraceTailSamplingRulesInternalServerError) Code() int {
	return 500
}

func (o *ReadTraceTailSamplingRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] readTraceTailSamplingRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceTailSamplingRulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] readTraceTailSamplingRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceTailSamplingRulesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceTailSamplingRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceTailSamplingRulesDefault creates a ReadTraceTailSamplingRulesDefault with default headers values
func NewReadTraceTailSamplingRulesDefault(code int) *ReadTraceTailSamplingRulesDefault {
	return &ReadTraceTailSamplingRulesDefault{
		_statusCode: code,
	}
}

/*
ReadTraceTailSamplingRulesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadTraceTailSamplingRulesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read trace tail sampling rules default response has a 2xx status code
func (o *ReadTraceTailSamplingRulesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read trace tail sampling rules default response has a 3xx status code
func (o *ReadTraceTailSamplingRulesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read trace tail sampling rules default response has a 4xx status code
func (o *ReadTraceTailSamplingRulesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read trace tail sampling rules default response has a 5xx status code
func (o *ReadTraceTailSamplingRulesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read trace tail sampling rules default response a status code equal to that given
func (o *ReadTraceTailSamplingRulesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read trace tail sampling rules default response
func (o *ReadTraceTailSamplingRulesDefault) Code() int {
	return o._statusCode
}

func (o *ReadTraceTailSamplingRulesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] ReadTraceTailSamplingRules default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceTailSamplingRulesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/trace-tail-sampling-rules][%d] ReadTraceTailSamplingRules default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceTailSamplingRulesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadTraceTailSamplingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
