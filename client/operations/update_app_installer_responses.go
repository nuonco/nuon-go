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

// UpdateAppInstallerReader is a Reader for the UpdateAppInstaller structure.
type UpdateAppInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAppInstallerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAppInstallerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAppInstallerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAppInstallerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAppInstallerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAppInstallerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/installers/{installer_id}] UpdateAppInstaller", response, response.Code())
	}
}

// NewUpdateAppInstallerCreated creates a UpdateAppInstallerCreated with default headers values
func NewUpdateAppInstallerCreated() *UpdateAppInstallerCreated {
	return &UpdateAppInstallerCreated{}
}

/*
UpdateAppInstallerCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateAppInstallerCreated struct {
	Payload *models.AppAppInstaller
}

// IsSuccess returns true when this update app installer created response has a 2xx status code
func (o *UpdateAppInstallerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update app installer created response has a 3xx status code
func (o *UpdateAppInstallerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app installer created response has a 4xx status code
func (o *UpdateAppInstallerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update app installer created response has a 5xx status code
func (o *UpdateAppInstallerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update app installer created response a status code equal to that given
func (o *UpdateAppInstallerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update app installer created response
func (o *UpdateAppInstallerCreated) Code() int {
	return 201
}

func (o *UpdateAppInstallerCreated) Error() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerCreated  %+v", 201, o.Payload)
}

func (o *UpdateAppInstallerCreated) String() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerCreated  %+v", 201, o.Payload)
}

func (o *UpdateAppInstallerCreated) GetPayload() *models.AppAppInstaller {
	return o.Payload
}

func (o *UpdateAppInstallerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppInstaller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInstallerBadRequest creates a UpdateAppInstallerBadRequest with default headers values
func NewUpdateAppInstallerBadRequest() *UpdateAppInstallerBadRequest {
	return &UpdateAppInstallerBadRequest{}
}

/*
UpdateAppInstallerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAppInstallerBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app installer bad request response has a 2xx status code
func (o *UpdateAppInstallerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app installer bad request response has a 3xx status code
func (o *UpdateAppInstallerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app installer bad request response has a 4xx status code
func (o *UpdateAppInstallerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app installer bad request response has a 5xx status code
func (o *UpdateAppInstallerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update app installer bad request response a status code equal to that given
func (o *UpdateAppInstallerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update app installer bad request response
func (o *UpdateAppInstallerBadRequest) Code() int {
	return 400
}

func (o *UpdateAppInstallerBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAppInstallerBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAppInstallerBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppInstallerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInstallerUnauthorized creates a UpdateAppInstallerUnauthorized with default headers values
func NewUpdateAppInstallerUnauthorized() *UpdateAppInstallerUnauthorized {
	return &UpdateAppInstallerUnauthorized{}
}

/*
UpdateAppInstallerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateAppInstallerUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app installer unauthorized response has a 2xx status code
func (o *UpdateAppInstallerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app installer unauthorized response has a 3xx status code
func (o *UpdateAppInstallerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app installer unauthorized response has a 4xx status code
func (o *UpdateAppInstallerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app installer unauthorized response has a 5xx status code
func (o *UpdateAppInstallerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update app installer unauthorized response a status code equal to that given
func (o *UpdateAppInstallerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update app installer unauthorized response
func (o *UpdateAppInstallerUnauthorized) Code() int {
	return 401
}

func (o *UpdateAppInstallerUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAppInstallerUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAppInstallerUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppInstallerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInstallerForbidden creates a UpdateAppInstallerForbidden with default headers values
func NewUpdateAppInstallerForbidden() *UpdateAppInstallerForbidden {
	return &UpdateAppInstallerForbidden{}
}

/*
UpdateAppInstallerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateAppInstallerForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app installer forbidden response has a 2xx status code
func (o *UpdateAppInstallerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app installer forbidden response has a 3xx status code
func (o *UpdateAppInstallerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app installer forbidden response has a 4xx status code
func (o *UpdateAppInstallerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app installer forbidden response has a 5xx status code
func (o *UpdateAppInstallerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update app installer forbidden response a status code equal to that given
func (o *UpdateAppInstallerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update app installer forbidden response
func (o *UpdateAppInstallerForbidden) Code() int {
	return 403
}

func (o *UpdateAppInstallerForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAppInstallerForbidden) String() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAppInstallerForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppInstallerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInstallerNotFound creates a UpdateAppInstallerNotFound with default headers values
func NewUpdateAppInstallerNotFound() *UpdateAppInstallerNotFound {
	return &UpdateAppInstallerNotFound{}
}

/*
UpdateAppInstallerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAppInstallerNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app installer not found response has a 2xx status code
func (o *UpdateAppInstallerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app installer not found response has a 3xx status code
func (o *UpdateAppInstallerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app installer not found response has a 4xx status code
func (o *UpdateAppInstallerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app installer not found response has a 5xx status code
func (o *UpdateAppInstallerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update app installer not found response a status code equal to that given
func (o *UpdateAppInstallerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update app installer not found response
func (o *UpdateAppInstallerNotFound) Code() int {
	return 404
}

func (o *UpdateAppInstallerNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAppInstallerNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAppInstallerNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppInstallerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInstallerInternalServerError creates a UpdateAppInstallerInternalServerError with default headers values
func NewUpdateAppInstallerInternalServerError() *UpdateAppInstallerInternalServerError {
	return &UpdateAppInstallerInternalServerError{}
}

/*
UpdateAppInstallerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateAppInstallerInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app installer internal server error response has a 2xx status code
func (o *UpdateAppInstallerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app installer internal server error response has a 3xx status code
func (o *UpdateAppInstallerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app installer internal server error response has a 4xx status code
func (o *UpdateAppInstallerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update app installer internal server error response has a 5xx status code
func (o *UpdateAppInstallerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update app installer internal server error response a status code equal to that given
func (o *UpdateAppInstallerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update app installer internal server error response
func (o *UpdateAppInstallerInternalServerError) Code() int {
	return 500
}

func (o *UpdateAppInstallerInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAppInstallerInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateAppInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAppInstallerInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppInstallerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
