// Code generated by go-swagger; DO NOT EDIT.

package classic_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// CreateClassicDashboardReader is a Reader for the CreateClassicDashboard structure.
type CreateClassicDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClassicDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateClassicDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateClassicDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateClassicDashboardConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateClassicDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateClassicDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClassicDashboardOK creates a CreateClassicDashboardOK with default headers values
func NewCreateClassicDashboardOK() *CreateClassicDashboardOK {
	return &CreateClassicDashboardOK{}
}

/*
CreateClassicDashboardOK describes a response with status code 200, with default header values.

A successful response containing the created GrafanaDashboard.
*/
type CreateClassicDashboardOK struct {
	Payload *models.Configv1CreateClassicDashboardResponse
}

// IsSuccess returns true when this create classic dashboard o k response has a 2xx status code
func (o *CreateClassicDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create classic dashboard o k response has a 3xx status code
func (o *CreateClassicDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create classic dashboard o k response has a 4xx status code
func (o *CreateClassicDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create classic dashboard o k response has a 5xx status code
func (o *CreateClassicDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create classic dashboard o k response a status code equal to that given
func (o *CreateClassicDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create classic dashboard o k response
func (o *CreateClassicDashboardOK) Code() int {
	return 200
}

func (o *CreateClassicDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateClassicDashboardOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateClassicDashboardOK) GetPayload() *models.Configv1CreateClassicDashboardResponse {
	return o.Payload
}

func (o *CreateClassicDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateClassicDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClassicDashboardBadRequest creates a CreateClassicDashboardBadRequest with default headers values
func NewCreateClassicDashboardBadRequest() *CreateClassicDashboardBadRequest {
	return &CreateClassicDashboardBadRequest{}
}

/*
CreateClassicDashboardBadRequest describes a response with status code 400, with default header values.

Cannot create the GrafanaDashboard because the request is invalid.
*/
type CreateClassicDashboardBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create classic dashboard bad request response has a 2xx status code
func (o *CreateClassicDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create classic dashboard bad request response has a 3xx status code
func (o *CreateClassicDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create classic dashboard bad request response has a 4xx status code
func (o *CreateClassicDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create classic dashboard bad request response has a 5xx status code
func (o *CreateClassicDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create classic dashboard bad request response a status code equal to that given
func (o *CreateClassicDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create classic dashboard bad request response
func (o *CreateClassicDashboardBadRequest) Code() int {
	return 400
}

func (o *CreateClassicDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClassicDashboardBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *CreateClassicDashboardBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateClassicDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClassicDashboardConflict creates a CreateClassicDashboardConflict with default headers values
func NewCreateClassicDashboardConflict() *CreateClassicDashboardConflict {
	return &CreateClassicDashboardConflict{}
}

/*
CreateClassicDashboardConflict describes a response with status code 409, with default header values.

Cannot create the GrafanaDashboard because there is a conflict with an existing GrafanaDashboard.
*/
type CreateClassicDashboardConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create classic dashboard conflict response has a 2xx status code
func (o *CreateClassicDashboardConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create classic dashboard conflict response has a 3xx status code
func (o *CreateClassicDashboardConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create classic dashboard conflict response has a 4xx status code
func (o *CreateClassicDashboardConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create classic dashboard conflict response has a 5xx status code
func (o *CreateClassicDashboardConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create classic dashboard conflict response a status code equal to that given
func (o *CreateClassicDashboardConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create classic dashboard conflict response
func (o *CreateClassicDashboardConflict) Code() int {
	return 409
}

func (o *CreateClassicDashboardConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardConflict  %+v", 409, o.Payload)
}

func (o *CreateClassicDashboardConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardConflict  %+v", 409, o.Payload)
}

func (o *CreateClassicDashboardConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateClassicDashboardConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClassicDashboardInternalServerError creates a CreateClassicDashboardInternalServerError with default headers values
func NewCreateClassicDashboardInternalServerError() *CreateClassicDashboardInternalServerError {
	return &CreateClassicDashboardInternalServerError{}
}

/*
CreateClassicDashboardInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateClassicDashboardInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create classic dashboard internal server error response has a 2xx status code
func (o *CreateClassicDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create classic dashboard internal server error response has a 3xx status code
func (o *CreateClassicDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create classic dashboard internal server error response has a 4xx status code
func (o *CreateClassicDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create classic dashboard internal server error response has a 5xx status code
func (o *CreateClassicDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create classic dashboard internal server error response a status code equal to that given
func (o *CreateClassicDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create classic dashboard internal server error response
func (o *CreateClassicDashboardInternalServerError) Code() int {
	return 500
}

func (o *CreateClassicDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateClassicDashboardInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] createClassicDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateClassicDashboardInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateClassicDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClassicDashboardDefault creates a CreateClassicDashboardDefault with default headers values
func NewCreateClassicDashboardDefault(code int) *CreateClassicDashboardDefault {
	return &CreateClassicDashboardDefault{
		_statusCode: code,
	}
}

/*
CreateClassicDashboardDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateClassicDashboardDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create classic dashboard default response has a 2xx status code
func (o *CreateClassicDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create classic dashboard default response has a 3xx status code
func (o *CreateClassicDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create classic dashboard default response has a 4xx status code
func (o *CreateClassicDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create classic dashboard default response has a 5xx status code
func (o *CreateClassicDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create classic dashboard default response a status code equal to that given
func (o *CreateClassicDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create classic dashboard default response
func (o *CreateClassicDashboardDefault) Code() int {
	return o._statusCode
}

func (o *CreateClassicDashboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] CreateClassicDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClassicDashboardDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/classic-dashboards][%d] CreateClassicDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClassicDashboardDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateClassicDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
