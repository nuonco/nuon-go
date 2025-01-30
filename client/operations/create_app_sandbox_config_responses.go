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

// CreateAppSandboxConfigReader is a Reader for the CreateAppSandboxConfig structure.
type CreateAppSandboxConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppSandboxConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppSandboxConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppSandboxConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppSandboxConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppSandboxConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppSandboxConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppSandboxConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/sandbox-config] CreateAppSandboxConfig", response, response.Code())
	}
}

// NewCreateAppSandboxConfigCreated creates a CreateAppSandboxConfigCreated with default headers values
func NewCreateAppSandboxConfigCreated() *CreateAppSandboxConfigCreated {
	return &CreateAppSandboxConfigCreated{}
}

/*
CreateAppSandboxConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppSandboxConfigCreated struct {
	Payload *models.AppAppSandboxConfig
}

// IsSuccess returns true when this create app sandbox config created response has a 2xx status code
func (o *CreateAppSandboxConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app sandbox config created response has a 3xx status code
func (o *CreateAppSandboxConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app sandbox config created response has a 4xx status code
func (o *CreateAppSandboxConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app sandbox config created response has a 5xx status code
func (o *CreateAppSandboxConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app sandbox config created response a status code equal to that given
func (o *CreateAppSandboxConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app sandbox config created response
func (o *CreateAppSandboxConfigCreated) Code() int {
	return 201
}

func (o *CreateAppSandboxConfigCreated) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppSandboxConfigCreated) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppSandboxConfigCreated) GetPayload() *models.AppAppSandboxConfig {
	return o.Payload
}

func (o *CreateAppSandboxConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppSandboxConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSandboxConfigBadRequest creates a CreateAppSandboxConfigBadRequest with default headers values
func NewCreateAppSandboxConfigBadRequest() *CreateAppSandboxConfigBadRequest {
	return &CreateAppSandboxConfigBadRequest{}
}

/*
CreateAppSandboxConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppSandboxConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app sandbox config bad request response has a 2xx status code
func (o *CreateAppSandboxConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app sandbox config bad request response has a 3xx status code
func (o *CreateAppSandboxConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app sandbox config bad request response has a 4xx status code
func (o *CreateAppSandboxConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app sandbox config bad request response has a 5xx status code
func (o *CreateAppSandboxConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app sandbox config bad request response a status code equal to that given
func (o *CreateAppSandboxConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app sandbox config bad request response
func (o *CreateAppSandboxConfigBadRequest) Code() int {
	return 400
}

func (o *CreateAppSandboxConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppSandboxConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppSandboxConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSandboxConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSandboxConfigUnauthorized creates a CreateAppSandboxConfigUnauthorized with default headers values
func NewCreateAppSandboxConfigUnauthorized() *CreateAppSandboxConfigUnauthorized {
	return &CreateAppSandboxConfigUnauthorized{}
}

/*
CreateAppSandboxConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppSandboxConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app sandbox config unauthorized response has a 2xx status code
func (o *CreateAppSandboxConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app sandbox config unauthorized response has a 3xx status code
func (o *CreateAppSandboxConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app sandbox config unauthorized response has a 4xx status code
func (o *CreateAppSandboxConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app sandbox config unauthorized response has a 5xx status code
func (o *CreateAppSandboxConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app sandbox config unauthorized response a status code equal to that given
func (o *CreateAppSandboxConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app sandbox config unauthorized response
func (o *CreateAppSandboxConfigUnauthorized) Code() int {
	return 401
}

func (o *CreateAppSandboxConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppSandboxConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppSandboxConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSandboxConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSandboxConfigForbidden creates a CreateAppSandboxConfigForbidden with default headers values
func NewCreateAppSandboxConfigForbidden() *CreateAppSandboxConfigForbidden {
	return &CreateAppSandboxConfigForbidden{}
}

/*
CreateAppSandboxConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppSandboxConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app sandbox config forbidden response has a 2xx status code
func (o *CreateAppSandboxConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app sandbox config forbidden response has a 3xx status code
func (o *CreateAppSandboxConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app sandbox config forbidden response has a 4xx status code
func (o *CreateAppSandboxConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app sandbox config forbidden response has a 5xx status code
func (o *CreateAppSandboxConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app sandbox config forbidden response a status code equal to that given
func (o *CreateAppSandboxConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app sandbox config forbidden response
func (o *CreateAppSandboxConfigForbidden) Code() int {
	return 403
}

func (o *CreateAppSandboxConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppSandboxConfigForbidden) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppSandboxConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSandboxConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSandboxConfigNotFound creates a CreateAppSandboxConfigNotFound with default headers values
func NewCreateAppSandboxConfigNotFound() *CreateAppSandboxConfigNotFound {
	return &CreateAppSandboxConfigNotFound{}
}

/*
CreateAppSandboxConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppSandboxConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app sandbox config not found response has a 2xx status code
func (o *CreateAppSandboxConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app sandbox config not found response has a 3xx status code
func (o *CreateAppSandboxConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app sandbox config not found response has a 4xx status code
func (o *CreateAppSandboxConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app sandbox config not found response has a 5xx status code
func (o *CreateAppSandboxConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app sandbox config not found response a status code equal to that given
func (o *CreateAppSandboxConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app sandbox config not found response
func (o *CreateAppSandboxConfigNotFound) Code() int {
	return 404
}

func (o *CreateAppSandboxConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppSandboxConfigNotFound) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppSandboxConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSandboxConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSandboxConfigInternalServerError creates a CreateAppSandboxConfigInternalServerError with default headers values
func NewCreateAppSandboxConfigInternalServerError() *CreateAppSandboxConfigInternalServerError {
	return &CreateAppSandboxConfigInternalServerError{}
}

/*
CreateAppSandboxConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppSandboxConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app sandbox config internal server error response has a 2xx status code
func (o *CreateAppSandboxConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app sandbox config internal server error response has a 3xx status code
func (o *CreateAppSandboxConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app sandbox config internal server error response has a 4xx status code
func (o *CreateAppSandboxConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app sandbox config internal server error response has a 5xx status code
func (o *CreateAppSandboxConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app sandbox config internal server error response a status code equal to that given
func (o *CreateAppSandboxConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app sandbox config internal server error response
func (o *CreateAppSandboxConfigInternalServerError) Code() int {
	return 500
}

func (o *CreateAppSandboxConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppSandboxConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/sandbox-config][%d] createAppSandboxConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppSandboxConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSandboxConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
