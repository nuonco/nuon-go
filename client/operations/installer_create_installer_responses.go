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

// InstallerCreateInstallerReader is a Reader for the InstallerCreateInstaller structure.
type InstallerCreateInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallerCreateInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInstallerCreateInstallerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstallerCreateInstallerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInstallerCreateInstallerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstallerCreateInstallerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstallerCreateInstallerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInstallerCreateInstallerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/installer/{installer_slug}/installs] InstallerCreateInstaller", response, response.Code())
	}
}

// NewInstallerCreateInstallerCreated creates a InstallerCreateInstallerCreated with default headers values
func NewInstallerCreateInstallerCreated() *InstallerCreateInstallerCreated {
	return &InstallerCreateInstallerCreated{}
}

/*
InstallerCreateInstallerCreated describes a response with status code 201, with default header values.

Created
*/
type InstallerCreateInstallerCreated struct {
	Payload *models.AppInstall
}

// IsSuccess returns true when this installer create installer created response has a 2xx status code
func (o *InstallerCreateInstallerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this installer create installer created response has a 3xx status code
func (o *InstallerCreateInstallerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this installer create installer created response has a 4xx status code
func (o *InstallerCreateInstallerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this installer create installer created response has a 5xx status code
func (o *InstallerCreateInstallerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this installer create installer created response a status code equal to that given
func (o *InstallerCreateInstallerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the installer create installer created response
func (o *InstallerCreateInstallerCreated) Code() int {
	return 201
}

func (o *InstallerCreateInstallerCreated) Error() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerCreated  %+v", 201, o.Payload)
}

func (o *InstallerCreateInstallerCreated) String() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerCreated  %+v", 201, o.Payload)
}

func (o *InstallerCreateInstallerCreated) GetPayload() *models.AppInstall {
	return o.Payload
}

func (o *InstallerCreateInstallerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallerCreateInstallerBadRequest creates a InstallerCreateInstallerBadRequest with default headers values
func NewInstallerCreateInstallerBadRequest() *InstallerCreateInstallerBadRequest {
	return &InstallerCreateInstallerBadRequest{}
}

/*
InstallerCreateInstallerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InstallerCreateInstallerBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this installer create installer bad request response has a 2xx status code
func (o *InstallerCreateInstallerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this installer create installer bad request response has a 3xx status code
func (o *InstallerCreateInstallerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this installer create installer bad request response has a 4xx status code
func (o *InstallerCreateInstallerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this installer create installer bad request response has a 5xx status code
func (o *InstallerCreateInstallerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this installer create installer bad request response a status code equal to that given
func (o *InstallerCreateInstallerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the installer create installer bad request response
func (o *InstallerCreateInstallerBadRequest) Code() int {
	return 400
}

func (o *InstallerCreateInstallerBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *InstallerCreateInstallerBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *InstallerCreateInstallerBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *InstallerCreateInstallerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallerCreateInstallerUnauthorized creates a InstallerCreateInstallerUnauthorized with default headers values
func NewInstallerCreateInstallerUnauthorized() *InstallerCreateInstallerUnauthorized {
	return &InstallerCreateInstallerUnauthorized{}
}

/*
InstallerCreateInstallerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InstallerCreateInstallerUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this installer create installer unauthorized response has a 2xx status code
func (o *InstallerCreateInstallerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this installer create installer unauthorized response has a 3xx status code
func (o *InstallerCreateInstallerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this installer create installer unauthorized response has a 4xx status code
func (o *InstallerCreateInstallerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this installer create installer unauthorized response has a 5xx status code
func (o *InstallerCreateInstallerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this installer create installer unauthorized response a status code equal to that given
func (o *InstallerCreateInstallerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the installer create installer unauthorized response
func (o *InstallerCreateInstallerUnauthorized) Code() int {
	return 401
}

func (o *InstallerCreateInstallerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *InstallerCreateInstallerUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *InstallerCreateInstallerUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *InstallerCreateInstallerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallerCreateInstallerForbidden creates a InstallerCreateInstallerForbidden with default headers values
func NewInstallerCreateInstallerForbidden() *InstallerCreateInstallerForbidden {
	return &InstallerCreateInstallerForbidden{}
}

/*
InstallerCreateInstallerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InstallerCreateInstallerForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this installer create installer forbidden response has a 2xx status code
func (o *InstallerCreateInstallerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this installer create installer forbidden response has a 3xx status code
func (o *InstallerCreateInstallerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this installer create installer forbidden response has a 4xx status code
func (o *InstallerCreateInstallerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this installer create installer forbidden response has a 5xx status code
func (o *InstallerCreateInstallerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this installer create installer forbidden response a status code equal to that given
func (o *InstallerCreateInstallerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the installer create installer forbidden response
func (o *InstallerCreateInstallerForbidden) Code() int {
	return 403
}

func (o *InstallerCreateInstallerForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerForbidden  %+v", 403, o.Payload)
}

func (o *InstallerCreateInstallerForbidden) String() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerForbidden  %+v", 403, o.Payload)
}

func (o *InstallerCreateInstallerForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *InstallerCreateInstallerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallerCreateInstallerNotFound creates a InstallerCreateInstallerNotFound with default headers values
func NewInstallerCreateInstallerNotFound() *InstallerCreateInstallerNotFound {
	return &InstallerCreateInstallerNotFound{}
}

/*
InstallerCreateInstallerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InstallerCreateInstallerNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this installer create installer not found response has a 2xx status code
func (o *InstallerCreateInstallerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this installer create installer not found response has a 3xx status code
func (o *InstallerCreateInstallerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this installer create installer not found response has a 4xx status code
func (o *InstallerCreateInstallerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this installer create installer not found response has a 5xx status code
func (o *InstallerCreateInstallerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this installer create installer not found response a status code equal to that given
func (o *InstallerCreateInstallerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the installer create installer not found response
func (o *InstallerCreateInstallerNotFound) Code() int {
	return 404
}

func (o *InstallerCreateInstallerNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerNotFound  %+v", 404, o.Payload)
}

func (o *InstallerCreateInstallerNotFound) String() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerNotFound  %+v", 404, o.Payload)
}

func (o *InstallerCreateInstallerNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *InstallerCreateInstallerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallerCreateInstallerInternalServerError creates a InstallerCreateInstallerInternalServerError with default headers values
func NewInstallerCreateInstallerInternalServerError() *InstallerCreateInstallerInternalServerError {
	return &InstallerCreateInstallerInternalServerError{}
}

/*
InstallerCreateInstallerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InstallerCreateInstallerInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this installer create installer internal server error response has a 2xx status code
func (o *InstallerCreateInstallerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this installer create installer internal server error response has a 3xx status code
func (o *InstallerCreateInstallerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this installer create installer internal server error response has a 4xx status code
func (o *InstallerCreateInstallerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this installer create installer internal server error response has a 5xx status code
func (o *InstallerCreateInstallerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this installer create installer internal server error response a status code equal to that given
func (o *InstallerCreateInstallerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the installer create installer internal server error response
func (o *InstallerCreateInstallerInternalServerError) Code() int {
	return 500
}

func (o *InstallerCreateInstallerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *InstallerCreateInstallerInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/installer/{installer_slug}/installs][%d] installerCreateInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *InstallerCreateInstallerInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *InstallerCreateInstallerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
