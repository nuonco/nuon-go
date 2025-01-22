// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-go/models"
)

// UpdateInstallInputReader is a Reader for the UpdateInstallInput structure.
type UpdateInstallInputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstallInputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInstallInputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInstallInputBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateInstallInputUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInstallInputForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInstallInputNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInstallInputInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/installs/{install_id}/inputs] UpdateInstallInput", response, response.Code())
	}
}

// NewUpdateInstallInputOK creates a UpdateInstallInputOK with default headers values
func NewUpdateInstallInputOK() *UpdateInstallInputOK {
	return &UpdateInstallInputOK{}
}

/*
UpdateInstallInputOK describes a response with status code 200, with default header values.

OK
*/
type UpdateInstallInputOK struct {
	Payload *models.AppInstallInputs
}

// IsSuccess returns true when this update install input o k response has a 2xx status code
func (o *UpdateInstallInputOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update install input o k response has a 3xx status code
func (o *UpdateInstallInputOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update install input o k response has a 4xx status code
func (o *UpdateInstallInputOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update install input o k response has a 5xx status code
func (o *UpdateInstallInputOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update install input o k response a status code equal to that given
func (o *UpdateInstallInputOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update install input o k response
func (o *UpdateInstallInputOK) Code() int {
	return 200
}

func (o *UpdateInstallInputOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputOK  %+v", 200, o.Payload)
}

func (o *UpdateInstallInputOK) String() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputOK  %+v", 200, o.Payload)
}

func (o *UpdateInstallInputOK) GetPayload() *models.AppInstallInputs {
	return o.Payload
}

func (o *UpdateInstallInputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallInputs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallInputBadRequest creates a UpdateInstallInputBadRequest with default headers values
func NewUpdateInstallInputBadRequest() *UpdateInstallInputBadRequest {
	return &UpdateInstallInputBadRequest{}
}

/*
UpdateInstallInputBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateInstallInputBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update install input bad request response has a 2xx status code
func (o *UpdateInstallInputBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update install input bad request response has a 3xx status code
func (o *UpdateInstallInputBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update install input bad request response has a 4xx status code
func (o *UpdateInstallInputBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update install input bad request response has a 5xx status code
func (o *UpdateInstallInputBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update install input bad request response a status code equal to that given
func (o *UpdateInstallInputBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update install input bad request response
func (o *UpdateInstallInputBadRequest) Code() int {
	return 400
}

func (o *UpdateInstallInputBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInstallInputBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInstallInputBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallInputBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallInputUnauthorized creates a UpdateInstallInputUnauthorized with default headers values
func NewUpdateInstallInputUnauthorized() *UpdateInstallInputUnauthorized {
	return &UpdateInstallInputUnauthorized{}
}

/*
UpdateInstallInputUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateInstallInputUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update install input unauthorized response has a 2xx status code
func (o *UpdateInstallInputUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update install input unauthorized response has a 3xx status code
func (o *UpdateInstallInputUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update install input unauthorized response has a 4xx status code
func (o *UpdateInstallInputUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update install input unauthorized response has a 5xx status code
func (o *UpdateInstallInputUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update install input unauthorized response a status code equal to that given
func (o *UpdateInstallInputUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update install input unauthorized response
func (o *UpdateInstallInputUnauthorized) Code() int {
	return 401
}

func (o *UpdateInstallInputUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInstallInputUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateInstallInputUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallInputUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallInputForbidden creates a UpdateInstallInputForbidden with default headers values
func NewUpdateInstallInputForbidden() *UpdateInstallInputForbidden {
	return &UpdateInstallInputForbidden{}
}

/*
UpdateInstallInputForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateInstallInputForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update install input forbidden response has a 2xx status code
func (o *UpdateInstallInputForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update install input forbidden response has a 3xx status code
func (o *UpdateInstallInputForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update install input forbidden response has a 4xx status code
func (o *UpdateInstallInputForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update install input forbidden response has a 5xx status code
func (o *UpdateInstallInputForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update install input forbidden response a status code equal to that given
func (o *UpdateInstallInputForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update install input forbidden response
func (o *UpdateInstallInputForbidden) Code() int {
	return 403
}

func (o *UpdateInstallInputForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInstallInputForbidden) String() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInstallInputForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallInputForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallInputNotFound creates a UpdateInstallInputNotFound with default headers values
func NewUpdateInstallInputNotFound() *UpdateInstallInputNotFound {
	return &UpdateInstallInputNotFound{}
}

/*
UpdateInstallInputNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateInstallInputNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update install input not found response has a 2xx status code
func (o *UpdateInstallInputNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update install input not found response has a 3xx status code
func (o *UpdateInstallInputNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update install input not found response has a 4xx status code
func (o *UpdateInstallInputNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update install input not found response has a 5xx status code
func (o *UpdateInstallInputNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update install input not found response a status code equal to that given
func (o *UpdateInstallInputNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update install input not found response
func (o *UpdateInstallInputNotFound) Code() int {
	return 404
}

func (o *UpdateInstallInputNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInstallInputNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInstallInputNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallInputNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallInputInternalServerError creates a UpdateInstallInputInternalServerError with default headers values
func NewUpdateInstallInputInternalServerError() *UpdateInstallInputInternalServerError {
	return &UpdateInstallInputInternalServerError{}
}

/*
UpdateInstallInputInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateInstallInputInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update install input internal server error response has a 2xx status code
func (o *UpdateInstallInputInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update install input internal server error response has a 3xx status code
func (o *UpdateInstallInputInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update install input internal server error response has a 4xx status code
func (o *UpdateInstallInputInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update install input internal server error response has a 5xx status code
func (o *UpdateInstallInputInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update install input internal server error response a status code equal to that given
func (o *UpdateInstallInputInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update install input internal server error response
func (o *UpdateInstallInputInternalServerError) Code() int {
	return 500
}

func (o *UpdateInstallInputInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInstallInputInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/installs/{install_id}/inputs][%d] updateInstallInputInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInstallInputInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallInputInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
