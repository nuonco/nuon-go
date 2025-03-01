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

// RenderInstallerReader is a Reader for the RenderInstaller structure.
type RenderInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderInstallerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenderInstallerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRenderInstallerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRenderInstallerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRenderInstallerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenderInstallerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installer/{installer_id}/render] RenderInstaller", response, response.Code())
	}
}

// NewRenderInstallerOK creates a RenderInstallerOK with default headers values
func NewRenderInstallerOK() *RenderInstallerOK {
	return &RenderInstallerOK{}
}

/*
RenderInstallerOK describes a response with status code 200, with default header values.

OK
*/
type RenderInstallerOK struct {
	Payload *models.ServiceRenderedInstaller
}

// IsSuccess returns true when this render installer o k response has a 2xx status code
func (o *RenderInstallerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this render installer o k response has a 3xx status code
func (o *RenderInstallerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render installer o k response has a 4xx status code
func (o *RenderInstallerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this render installer o k response has a 5xx status code
func (o *RenderInstallerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this render installer o k response a status code equal to that given
func (o *RenderInstallerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the render installer o k response
func (o *RenderInstallerOK) Code() int {
	return 200
}

func (o *RenderInstallerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerOK %s", 200, payload)
}

func (o *RenderInstallerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerOK %s", 200, payload)
}

func (o *RenderInstallerOK) GetPayload() *models.ServiceRenderedInstaller {
	return o.Payload
}

func (o *RenderInstallerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceRenderedInstaller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderInstallerBadRequest creates a RenderInstallerBadRequest with default headers values
func NewRenderInstallerBadRequest() *RenderInstallerBadRequest {
	return &RenderInstallerBadRequest{}
}

/*
RenderInstallerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RenderInstallerBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this render installer bad request response has a 2xx status code
func (o *RenderInstallerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render installer bad request response has a 3xx status code
func (o *RenderInstallerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render installer bad request response has a 4xx status code
func (o *RenderInstallerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this render installer bad request response has a 5xx status code
func (o *RenderInstallerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this render installer bad request response a status code equal to that given
func (o *RenderInstallerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the render installer bad request response
func (o *RenderInstallerBadRequest) Code() int {
	return 400
}

func (o *RenderInstallerBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerBadRequest %s", 400, payload)
}

func (o *RenderInstallerBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerBadRequest %s", 400, payload)
}

func (o *RenderInstallerBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RenderInstallerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderInstallerUnauthorized creates a RenderInstallerUnauthorized with default headers values
func NewRenderInstallerUnauthorized() *RenderInstallerUnauthorized {
	return &RenderInstallerUnauthorized{}
}

/*
RenderInstallerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RenderInstallerUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this render installer unauthorized response has a 2xx status code
func (o *RenderInstallerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render installer unauthorized response has a 3xx status code
func (o *RenderInstallerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render installer unauthorized response has a 4xx status code
func (o *RenderInstallerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this render installer unauthorized response has a 5xx status code
func (o *RenderInstallerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this render installer unauthorized response a status code equal to that given
func (o *RenderInstallerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the render installer unauthorized response
func (o *RenderInstallerUnauthorized) Code() int {
	return 401
}

func (o *RenderInstallerUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerUnauthorized %s", 401, payload)
}

func (o *RenderInstallerUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerUnauthorized %s", 401, payload)
}

func (o *RenderInstallerUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RenderInstallerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderInstallerForbidden creates a RenderInstallerForbidden with default headers values
func NewRenderInstallerForbidden() *RenderInstallerForbidden {
	return &RenderInstallerForbidden{}
}

/*
RenderInstallerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RenderInstallerForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this render installer forbidden response has a 2xx status code
func (o *RenderInstallerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render installer forbidden response has a 3xx status code
func (o *RenderInstallerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render installer forbidden response has a 4xx status code
func (o *RenderInstallerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this render installer forbidden response has a 5xx status code
func (o *RenderInstallerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this render installer forbidden response a status code equal to that given
func (o *RenderInstallerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the render installer forbidden response
func (o *RenderInstallerForbidden) Code() int {
	return 403
}

func (o *RenderInstallerForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerForbidden %s", 403, payload)
}

func (o *RenderInstallerForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerForbidden %s", 403, payload)
}

func (o *RenderInstallerForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RenderInstallerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderInstallerNotFound creates a RenderInstallerNotFound with default headers values
func NewRenderInstallerNotFound() *RenderInstallerNotFound {
	return &RenderInstallerNotFound{}
}

/*
RenderInstallerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RenderInstallerNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this render installer not found response has a 2xx status code
func (o *RenderInstallerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render installer not found response has a 3xx status code
func (o *RenderInstallerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render installer not found response has a 4xx status code
func (o *RenderInstallerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this render installer not found response has a 5xx status code
func (o *RenderInstallerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this render installer not found response a status code equal to that given
func (o *RenderInstallerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the render installer not found response
func (o *RenderInstallerNotFound) Code() int {
	return 404
}

func (o *RenderInstallerNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerNotFound %s", 404, payload)
}

func (o *RenderInstallerNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerNotFound %s", 404, payload)
}

func (o *RenderInstallerNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RenderInstallerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderInstallerInternalServerError creates a RenderInstallerInternalServerError with default headers values
func NewRenderInstallerInternalServerError() *RenderInstallerInternalServerError {
	return &RenderInstallerInternalServerError{}
}

/*
RenderInstallerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RenderInstallerInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this render installer internal server error response has a 2xx status code
func (o *RenderInstallerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render installer internal server error response has a 3xx status code
func (o *RenderInstallerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render installer internal server error response has a 4xx status code
func (o *RenderInstallerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this render installer internal server error response has a 5xx status code
func (o *RenderInstallerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this render installer internal server error response a status code equal to that given
func (o *RenderInstallerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the render installer internal server error response
func (o *RenderInstallerInternalServerError) Code() int {
	return 500
}

func (o *RenderInstallerInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerInternalServerError %s", 500, payload)
}

func (o *RenderInstallerInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installer/{installer_id}/render][%d] renderInstallerInternalServerError %s", 500, payload)
}

func (o *RenderInstallerInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RenderInstallerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
