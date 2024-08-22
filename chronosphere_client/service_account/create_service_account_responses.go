// Code generated by go-swagger; DO NOT EDIT.

package service_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// CreateServiceAccountReader is a Reader for the CreateServiceAccount structure.
type CreateServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateServiceAccountConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateServiceAccountOK creates a CreateServiceAccountOK with default headers values
func NewCreateServiceAccountOK() *CreateServiceAccountOK {
	return &CreateServiceAccountOK{}
}

/*
CreateServiceAccountOK describes a response with status code 200, with default header values.

A successful response containing the created ServiceAccount.
*/
type CreateServiceAccountOK struct {
	Payload *models.Configv1CreateServiceAccountResponse
}

// IsSuccess returns true when this create service account o k response has a 2xx status code
func (o *CreateServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create service account o k response has a 3xx status code
func (o *CreateServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account o k response has a 4xx status code
func (o *CreateServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service account o k response has a 5xx status code
func (o *CreateServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account o k response a status code equal to that given
func (o *CreateServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create service account o k response
func (o *CreateServiceAccountOK) Code() int {
	return 200
}

func (o *CreateServiceAccountOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountOK  %+v", 200, o.Payload)
}

func (o *CreateServiceAccountOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountOK  %+v", 200, o.Payload)
}

func (o *CreateServiceAccountOK) GetPayload() *models.Configv1CreateServiceAccountResponse {
	return o.Payload
}

func (o *CreateServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateServiceAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountBadRequest creates a CreateServiceAccountBadRequest with default headers values
func NewCreateServiceAccountBadRequest() *CreateServiceAccountBadRequest {
	return &CreateServiceAccountBadRequest{}
}

/*
CreateServiceAccountBadRequest describes a response with status code 400, with default header values.

Cannot create the ServiceAccount because the request is invalid.
*/
type CreateServiceAccountBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create service account bad request response has a 2xx status code
func (o *CreateServiceAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account bad request response has a 3xx status code
func (o *CreateServiceAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account bad request response has a 4xx status code
func (o *CreateServiceAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account bad request response has a 5xx status code
func (o *CreateServiceAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account bad request response a status code equal to that given
func (o *CreateServiceAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create service account bad request response
func (o *CreateServiceAccountBadRequest) Code() int {
	return 400
}

func (o *CreateServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceAccountBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateServiceAccountBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountConflict creates a CreateServiceAccountConflict with default headers values
func NewCreateServiceAccountConflict() *CreateServiceAccountConflict {
	return &CreateServiceAccountConflict{}
}

/*
CreateServiceAccountConflict describes a response with status code 409, with default header values.

Cannot create the ServiceAccount because there is a conflict with an existing ServiceAccount.
*/
type CreateServiceAccountConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create service account conflict response has a 2xx status code
func (o *CreateServiceAccountConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account conflict response has a 3xx status code
func (o *CreateServiceAccountConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account conflict response has a 4xx status code
func (o *CreateServiceAccountConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account conflict response has a 5xx status code
func (o *CreateServiceAccountConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account conflict response a status code equal to that given
func (o *CreateServiceAccountConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create service account conflict response
func (o *CreateServiceAccountConflict) Code() int {
	return 409
}

func (o *CreateServiceAccountConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountConflict  %+v", 409, o.Payload)
}

func (o *CreateServiceAccountConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountConflict  %+v", 409, o.Payload)
}

func (o *CreateServiceAccountConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateServiceAccountConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountInternalServerError creates a CreateServiceAccountInternalServerError with default headers values
func NewCreateServiceAccountInternalServerError() *CreateServiceAccountInternalServerError {
	return &CreateServiceAccountInternalServerError{}
}

/*
CreateServiceAccountInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateServiceAccountInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create service account internal server error response has a 2xx status code
func (o *CreateServiceAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account internal server error response has a 3xx status code
func (o *CreateServiceAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account internal server error response has a 4xx status code
func (o *CreateServiceAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service account internal server error response has a 5xx status code
func (o *CreateServiceAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create service account internal server error response a status code equal to that given
func (o *CreateServiceAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create service account internal server error response
func (o *CreateServiceAccountInternalServerError) Code() int {
	return 500
}

func (o *CreateServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateServiceAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] createServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateServiceAccountInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountDefault creates a CreateServiceAccountDefault with default headers values
func NewCreateServiceAccountDefault(code int) *CreateServiceAccountDefault {
	return &CreateServiceAccountDefault{
		_statusCode: code,
	}
}

/*
CreateServiceAccountDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateServiceAccountDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create service account default response has a 2xx status code
func (o *CreateServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create service account default response has a 3xx status code
func (o *CreateServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create service account default response has a 4xx status code
func (o *CreateServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create service account default response has a 5xx status code
func (o *CreateServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create service account default response a status code equal to that given
func (o *CreateServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create service account default response
func (o *CreateServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *CreateServiceAccountDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] CreateServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServiceAccountDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/service-accounts][%d] CreateServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateServiceAccountDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
