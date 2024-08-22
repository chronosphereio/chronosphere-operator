// Code generated by go-swagger; DO NOT EDIT.

package dataset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"chronosphere.io/chronosphere-operator/models"
)

// UpdateDatasetReader is a Reader for the UpdateDataset structure.
type UpdateDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDatasetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDatasetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDatasetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDatasetOK creates a UpdateDatasetOK with default headers values
func NewUpdateDatasetOK() *UpdateDatasetOK {
	return &UpdateDatasetOK{}
}

/*
UpdateDatasetOK describes a response with status code 200, with default header values.

A successful response containing the updated Dataset.
*/
type UpdateDatasetOK struct {
	Payload *models.Configv1UpdateDatasetResponse
}

// IsSuccess returns true when this update dataset o k response has a 2xx status code
func (o *UpdateDatasetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update dataset o k response has a 3xx status code
func (o *UpdateDatasetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dataset o k response has a 4xx status code
func (o *UpdateDatasetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dataset o k response has a 5xx status code
func (o *UpdateDatasetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update dataset o k response a status code equal to that given
func (o *UpdateDatasetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update dataset o k response
func (o *UpdateDatasetOK) Code() int {
	return 200
}

func (o *UpdateDatasetOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetOK  %+v", 200, o.Payload)
}

func (o *UpdateDatasetOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetOK  %+v", 200, o.Payload)
}

func (o *UpdateDatasetOK) GetPayload() *models.Configv1UpdateDatasetResponse {
	return o.Payload
}

func (o *UpdateDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateDatasetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatasetBadRequest creates a UpdateDatasetBadRequest with default headers values
func NewUpdateDatasetBadRequest() *UpdateDatasetBadRequest {
	return &UpdateDatasetBadRequest{}
}

/*
UpdateDatasetBadRequest describes a response with status code 400, with default header values.

Cannot update the Dataset because the request is invalid.
*/
type UpdateDatasetBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dataset bad request response has a 2xx status code
func (o *UpdateDatasetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dataset bad request response has a 3xx status code
func (o *UpdateDatasetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dataset bad request response has a 4xx status code
func (o *UpdateDatasetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dataset bad request response has a 5xx status code
func (o *UpdateDatasetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update dataset bad request response a status code equal to that given
func (o *UpdateDatasetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update dataset bad request response
func (o *UpdateDatasetBadRequest) Code() int {
	return 400
}

func (o *UpdateDatasetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDatasetBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDatasetBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatasetNotFound creates a UpdateDatasetNotFound with default headers values
func NewUpdateDatasetNotFound() *UpdateDatasetNotFound {
	return &UpdateDatasetNotFound{}
}

/*
UpdateDatasetNotFound describes a response with status code 404, with default header values.

Cannot update the Dataset because the slug does not exist.
*/
type UpdateDatasetNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dataset not found response has a 2xx status code
func (o *UpdateDatasetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dataset not found response has a 3xx status code
func (o *UpdateDatasetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dataset not found response has a 4xx status code
func (o *UpdateDatasetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dataset not found response has a 5xx status code
func (o *UpdateDatasetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update dataset not found response a status code equal to that given
func (o *UpdateDatasetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update dataset not found response
func (o *UpdateDatasetNotFound) Code() int {
	return 404
}

func (o *UpdateDatasetNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDatasetNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDatasetNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDatasetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatasetConflict creates a UpdateDatasetConflict with default headers values
func NewUpdateDatasetConflict() *UpdateDatasetConflict {
	return &UpdateDatasetConflict{}
}

/*
UpdateDatasetConflict describes a response with status code 409, with default header values.

Cannot update the Dataset because there is a conflict with an existing Dataset.
*/
type UpdateDatasetConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dataset conflict response has a 2xx status code
func (o *UpdateDatasetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dataset conflict response has a 3xx status code
func (o *UpdateDatasetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dataset conflict response has a 4xx status code
func (o *UpdateDatasetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update dataset conflict response has a 5xx status code
func (o *UpdateDatasetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update dataset conflict response a status code equal to that given
func (o *UpdateDatasetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update dataset conflict response
func (o *UpdateDatasetConflict) Code() int {
	return 409
}

func (o *UpdateDatasetConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetConflict  %+v", 409, o.Payload)
}

func (o *UpdateDatasetConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetConflict  %+v", 409, o.Payload)
}

func (o *UpdateDatasetConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDatasetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatasetInternalServerError creates a UpdateDatasetInternalServerError with default headers values
func NewUpdateDatasetInternalServerError() *UpdateDatasetInternalServerError {
	return &UpdateDatasetInternalServerError{}
}

/*
UpdateDatasetInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateDatasetInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update dataset internal server error response has a 2xx status code
func (o *UpdateDatasetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update dataset internal server error response has a 3xx status code
func (o *UpdateDatasetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update dataset internal server error response has a 4xx status code
func (o *UpdateDatasetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update dataset internal server error response has a 5xx status code
func (o *UpdateDatasetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update dataset internal server error response a status code equal to that given
func (o *UpdateDatasetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update dataset internal server error response
func (o *UpdateDatasetInternalServerError) Code() int {
	return 500
}

func (o *UpdateDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDatasetInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] updateDatasetInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDatasetInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDatasetDefault creates a UpdateDatasetDefault with default headers values
func NewUpdateDatasetDefault(code int) *UpdateDatasetDefault {
	return &UpdateDatasetDefault{
		_statusCode: code,
	}
}

/*
UpdateDatasetDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateDatasetDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update dataset default response has a 2xx status code
func (o *UpdateDatasetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update dataset default response has a 3xx status code
func (o *UpdateDatasetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update dataset default response has a 4xx status code
func (o *UpdateDatasetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update dataset default response has a 5xx status code
func (o *UpdateDatasetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update dataset default response a status code equal to that given
func (o *UpdateDatasetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update dataset default response
func (o *UpdateDatasetDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDatasetDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] UpdateDataset default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDatasetDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/datasets/{slug}][%d] UpdateDataset default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDatasetDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateDatasetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
