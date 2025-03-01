// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-go/models"
)

// UpdateInstallerReader is a Reader for the UpdateInstaller structure.
type UpdateInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateInstallerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInstallerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateInstallerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInstallerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInstallerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInstallerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/installers/{installer_id}] UpdateInstaller", response, response.Code())
	}
}

// NewUpdateInstallerCreated creates a UpdateInstallerCreated with default headers values
func NewUpdateInstallerCreated() *UpdateInstallerCreated {
	return &UpdateInstallerCreated{}
}

/*
UpdateInstallerCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateInstallerCreated struct {
	Payload *models.AppInstaller
}

// IsSuccess returns true when this update installer created response has a 2xx status code
func (o *UpdateInstallerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update installer created response has a 3xx status code
func (o *UpdateInstallerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update installer created response has a 4xx status code
func (o *UpdateInstallerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update installer created response has a 5xx status code
func (o *UpdateInstallerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update installer created response a status code equal to that given
func (o *UpdateInstallerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update installer created response
func (o *UpdateInstallerCreated) Code() int {
	return 201
}

func (o *UpdateInstallerCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerCreated %s", 201, payload)
}

func (o *UpdateInstallerCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerCreated %s", 201, payload)
}

func (o *UpdateInstallerCreated) GetPayload() *models.AppInstaller {
	return o.Payload
}

func (o *UpdateInstallerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstaller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallerBadRequest creates a UpdateInstallerBadRequest with default headers values
func NewUpdateInstallerBadRequest() *UpdateInstallerBadRequest {
	return &UpdateInstallerBadRequest{}
}

/*
UpdateInstallerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateInstallerBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update installer bad request response has a 2xx status code
func (o *UpdateInstallerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update installer bad request response has a 3xx status code
func (o *UpdateInstallerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update installer bad request response has a 4xx status code
func (o *UpdateInstallerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update installer bad request response has a 5xx status code
func (o *UpdateInstallerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update installer bad request response a status code equal to that given
func (o *UpdateInstallerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update installer bad request response
func (o *UpdateInstallerBadRequest) Code() int {
	return 400
}

func (o *UpdateInstallerBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerBadRequest %s", 400, payload)
}

func (o *UpdateInstallerBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerBadRequest %s", 400, payload)
}

func (o *UpdateInstallerBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallerUnauthorized creates a UpdateInstallerUnauthorized with default headers values
func NewUpdateInstallerUnauthorized() *UpdateInstallerUnauthorized {
	return &UpdateInstallerUnauthorized{}
}

/*
UpdateInstallerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateInstallerUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update installer unauthorized response has a 2xx status code
func (o *UpdateInstallerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update installer unauthorized response has a 3xx status code
func (o *UpdateInstallerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update installer unauthorized response has a 4xx status code
func (o *UpdateInstallerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update installer unauthorized response has a 5xx status code
func (o *UpdateInstallerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update installer unauthorized response a status code equal to that given
func (o *UpdateInstallerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update installer unauthorized response
func (o *UpdateInstallerUnauthorized) Code() int {
	return 401
}

func (o *UpdateInstallerUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerUnauthorized %s", 401, payload)
}

func (o *UpdateInstallerUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerUnauthorized %s", 401, payload)
}

func (o *UpdateInstallerUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallerForbidden creates a UpdateInstallerForbidden with default headers values
func NewUpdateInstallerForbidden() *UpdateInstallerForbidden {
	return &UpdateInstallerForbidden{}
}

/*
UpdateInstallerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateInstallerForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update installer forbidden response has a 2xx status code
func (o *UpdateInstallerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update installer forbidden response has a 3xx status code
func (o *UpdateInstallerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update installer forbidden response has a 4xx status code
func (o *UpdateInstallerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update installer forbidden response has a 5xx status code
func (o *UpdateInstallerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update installer forbidden response a status code equal to that given
func (o *UpdateInstallerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update installer forbidden response
func (o *UpdateInstallerForbidden) Code() int {
	return 403
}

func (o *UpdateInstallerForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerForbidden %s", 403, payload)
}

func (o *UpdateInstallerForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerForbidden %s", 403, payload)
}

func (o *UpdateInstallerForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallerNotFound creates a UpdateInstallerNotFound with default headers values
func NewUpdateInstallerNotFound() *UpdateInstallerNotFound {
	return &UpdateInstallerNotFound{}
}

/*
UpdateInstallerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateInstallerNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update installer not found response has a 2xx status code
func (o *UpdateInstallerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update installer not found response has a 3xx status code
func (o *UpdateInstallerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update installer not found response has a 4xx status code
func (o *UpdateInstallerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update installer not found response has a 5xx status code
func (o *UpdateInstallerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update installer not found response a status code equal to that given
func (o *UpdateInstallerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update installer not found response
func (o *UpdateInstallerNotFound) Code() int {
	return 404
}

func (o *UpdateInstallerNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerNotFound %s", 404, payload)
}

func (o *UpdateInstallerNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerNotFound %s", 404, payload)
}

func (o *UpdateInstallerNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInstallerInternalServerError creates a UpdateInstallerInternalServerError with default headers values
func NewUpdateInstallerInternalServerError() *UpdateInstallerInternalServerError {
	return &UpdateInstallerInternalServerError{}
}

/*
UpdateInstallerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateInstallerInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update installer internal server error response has a 2xx status code
func (o *UpdateInstallerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update installer internal server error response has a 3xx status code
func (o *UpdateInstallerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update installer internal server error response has a 4xx status code
func (o *UpdateInstallerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update installer internal server error response has a 5xx status code
func (o *UpdateInstallerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update installer internal server error response a status code equal to that given
func (o *UpdateInstallerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update installer internal server error response
func (o *UpdateInstallerInternalServerError) Code() int {
	return 500
}

func (o *UpdateInstallerInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerInternalServerError %s", 500, payload)
}

func (o *UpdateInstallerInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/installers/{installer_id}][%d] updateInstallerInternalServerError %s", 500, payload)
}

func (o *UpdateInstallerInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateInstallerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
