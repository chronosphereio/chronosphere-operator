// Code generated by go-swagger; DO NOT EDIT.

package recording_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// CreateRecordingRuleReader is a Reader for the CreateRecordingRule structure.
type CreateRecordingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRecordingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRecordingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRecordingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRecordingRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRecordingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateRecordingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRecordingRuleOK creates a CreateRecordingRuleOK with default headers values
func NewCreateRecordingRuleOK() *CreateRecordingRuleOK {
	return &CreateRecordingRuleOK{}
}

/*
CreateRecordingRuleOK describes a response with status code 200, with default header values.

A successful response containing the created RecordingRule.
*/
type CreateRecordingRuleOK struct {
	Payload *models.Configv1CreateRecordingRuleResponse
}

// IsSuccess returns true when this create recording rule o k response has a 2xx status code
func (o *CreateRecordingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create recording rule o k response has a 3xx status code
func (o *CreateRecordingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create recording rule o k response has a 4xx status code
func (o *CreateRecordingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create recording rule o k response has a 5xx status code
func (o *CreateRecordingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create recording rule o k response a status code equal to that given
func (o *CreateRecordingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create recording rule o k response
func (o *CreateRecordingRuleOK) Code() int {
	return 200
}

func (o *CreateRecordingRuleOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleOK  %+v", 200, o.Payload)
}

func (o *CreateRecordingRuleOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleOK  %+v", 200, o.Payload)
}

func (o *CreateRecordingRuleOK) GetPayload() *models.Configv1CreateRecordingRuleResponse {
	return o.Payload
}

func (o *CreateRecordingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateRecordingRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecordingRuleBadRequest creates a CreateRecordingRuleBadRequest with default headers values
func NewCreateRecordingRuleBadRequest() *CreateRecordingRuleBadRequest {
	return &CreateRecordingRuleBadRequest{}
}

/*
CreateRecordingRuleBadRequest describes a response with status code 400, with default header values.

Cannot create the RecordingRule because the request is invalid.
*/
type CreateRecordingRuleBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create recording rule bad request response has a 2xx status code
func (o *CreateRecordingRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create recording rule bad request response has a 3xx status code
func (o *CreateRecordingRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create recording rule bad request response has a 4xx status code
func (o *CreateRecordingRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create recording rule bad request response has a 5xx status code
func (o *CreateRecordingRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create recording rule bad request response a status code equal to that given
func (o *CreateRecordingRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create recording rule bad request response
func (o *CreateRecordingRuleBadRequest) Code() int {
	return 400
}

func (o *CreateRecordingRuleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRecordingRuleBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRecordingRuleBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateRecordingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecordingRuleConflict creates a CreateRecordingRuleConflict with default headers values
func NewCreateRecordingRuleConflict() *CreateRecordingRuleConflict {
	return &CreateRecordingRuleConflict{}
}

/*
CreateRecordingRuleConflict describes a response with status code 409, with default header values.

Cannot create the RecordingRule because there is a conflict with an existing RecordingRule.
*/
type CreateRecordingRuleConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create recording rule conflict response has a 2xx status code
func (o *CreateRecordingRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create recording rule conflict response has a 3xx status code
func (o *CreateRecordingRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create recording rule conflict response has a 4xx status code
func (o *CreateRecordingRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create recording rule conflict response has a 5xx status code
func (o *CreateRecordingRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create recording rule conflict response a status code equal to that given
func (o *CreateRecordingRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create recording rule conflict response
func (o *CreateRecordingRuleConflict) Code() int {
	return 409
}

func (o *CreateRecordingRuleConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateRecordingRuleConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleConflict  %+v", 409, o.Payload)
}

func (o *CreateRecordingRuleConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateRecordingRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecordingRuleInternalServerError creates a CreateRecordingRuleInternalServerError with default headers values
func NewCreateRecordingRuleInternalServerError() *CreateRecordingRuleInternalServerError {
	return &CreateRecordingRuleInternalServerError{}
}

/*
CreateRecordingRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateRecordingRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create recording rule internal server error response has a 2xx status code
func (o *CreateRecordingRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create recording rule internal server error response has a 3xx status code
func (o *CreateRecordingRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create recording rule internal server error response has a 4xx status code
func (o *CreateRecordingRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create recording rule internal server error response has a 5xx status code
func (o *CreateRecordingRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create recording rule internal server error response a status code equal to that given
func (o *CreateRecordingRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create recording rule internal server error response
func (o *CreateRecordingRuleInternalServerError) Code() int {
	return 500
}

func (o *CreateRecordingRuleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRecordingRuleInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] createRecordingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRecordingRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateRecordingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecordingRuleDefault creates a CreateRecordingRuleDefault with default headers values
func NewCreateRecordingRuleDefault(code int) *CreateRecordingRuleDefault {
	return &CreateRecordingRuleDefault{
		_statusCode: code,
	}
}

/*
CreateRecordingRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateRecordingRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create recording rule default response has a 2xx status code
func (o *CreateRecordingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create recording rule default response has a 3xx status code
func (o *CreateRecordingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create recording rule default response has a 4xx status code
func (o *CreateRecordingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create recording rule default response has a 5xx status code
func (o *CreateRecordingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create recording rule default response a status code equal to that given
func (o *CreateRecordingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create recording rule default response
func (o *CreateRecordingRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateRecordingRuleDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] CreateRecordingRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRecordingRuleDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/recording-rules][%d] CreateRecordingRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRecordingRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateRecordingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
