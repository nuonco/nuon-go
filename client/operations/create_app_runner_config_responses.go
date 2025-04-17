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

// CreateAppRunnerConfigReader is a Reader for the CreateAppRunnerConfig structure.
type CreateAppRunnerConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppRunnerConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppRunnerConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppRunnerConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppRunnerConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppRunnerConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppRunnerConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppRunnerConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/runner-configs] CreateAppRunnerConfig", response, response.Code())
	}
}

// NewCreateAppRunnerConfigCreated creates a CreateAppRunnerConfigCreated with default headers values
func NewCreateAppRunnerConfigCreated() *CreateAppRunnerConfigCreated {
	return &CreateAppRunnerConfigCreated{}
}

/*
CreateAppRunnerConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppRunnerConfigCreated struct {
	Payload *models.AppAppRunnerConfig
}

// IsSuccess returns true when this create app runner config created response has a 2xx status code
func (o *CreateAppRunnerConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app runner config created response has a 3xx status code
func (o *CreateAppRunnerConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app runner config created response has a 4xx status code
func (o *CreateAppRunnerConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app runner config created response has a 5xx status code
func (o *CreateAppRunnerConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app runner config created response a status code equal to that given
func (o *CreateAppRunnerConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app runner config created response
func (o *CreateAppRunnerConfigCreated) Code() int {
	return 201
}

func (o *CreateAppRunnerConfigCreated) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppRunnerConfigCreated) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppRunnerConfigCreated) GetPayload() *models.AppAppRunnerConfig {
	return o.Payload
}

func (o *CreateAppRunnerConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppRunnerConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppRunnerConfigBadRequest creates a CreateAppRunnerConfigBadRequest with default headers values
func NewCreateAppRunnerConfigBadRequest() *CreateAppRunnerConfigBadRequest {
	return &CreateAppRunnerConfigBadRequest{}
}

/*
CreateAppRunnerConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppRunnerConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app runner config bad request response has a 2xx status code
func (o *CreateAppRunnerConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app runner config bad request response has a 3xx status code
func (o *CreateAppRunnerConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app runner config bad request response has a 4xx status code
func (o *CreateAppRunnerConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app runner config bad request response has a 5xx status code
func (o *CreateAppRunnerConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app runner config bad request response a status code equal to that given
func (o *CreateAppRunnerConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app runner config bad request response
func (o *CreateAppRunnerConfigBadRequest) Code() int {
	return 400
}

func (o *CreateAppRunnerConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppRunnerConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppRunnerConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppRunnerConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppRunnerConfigUnauthorized creates a CreateAppRunnerConfigUnauthorized with default headers values
func NewCreateAppRunnerConfigUnauthorized() *CreateAppRunnerConfigUnauthorized {
	return &CreateAppRunnerConfigUnauthorized{}
}

/*
CreateAppRunnerConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppRunnerConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app runner config unauthorized response has a 2xx status code
func (o *CreateAppRunnerConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app runner config unauthorized response has a 3xx status code
func (o *CreateAppRunnerConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app runner config unauthorized response has a 4xx status code
func (o *CreateAppRunnerConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app runner config unauthorized response has a 5xx status code
func (o *CreateAppRunnerConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app runner config unauthorized response a status code equal to that given
func (o *CreateAppRunnerConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app runner config unauthorized response
func (o *CreateAppRunnerConfigUnauthorized) Code() int {
	return 401
}

func (o *CreateAppRunnerConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppRunnerConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppRunnerConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppRunnerConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppRunnerConfigForbidden creates a CreateAppRunnerConfigForbidden with default headers values
func NewCreateAppRunnerConfigForbidden() *CreateAppRunnerConfigForbidden {
	return &CreateAppRunnerConfigForbidden{}
}

/*
CreateAppRunnerConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppRunnerConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app runner config forbidden response has a 2xx status code
func (o *CreateAppRunnerConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app runner config forbidden response has a 3xx status code
func (o *CreateAppRunnerConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app runner config forbidden response has a 4xx status code
func (o *CreateAppRunnerConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app runner config forbidden response has a 5xx status code
func (o *CreateAppRunnerConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app runner config forbidden response a status code equal to that given
func (o *CreateAppRunnerConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app runner config forbidden response
func (o *CreateAppRunnerConfigForbidden) Code() int {
	return 403
}

func (o *CreateAppRunnerConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppRunnerConfigForbidden) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppRunnerConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppRunnerConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppRunnerConfigNotFound creates a CreateAppRunnerConfigNotFound with default headers values
func NewCreateAppRunnerConfigNotFound() *CreateAppRunnerConfigNotFound {
	return &CreateAppRunnerConfigNotFound{}
}

/*
CreateAppRunnerConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppRunnerConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app runner config not found response has a 2xx status code
func (o *CreateAppRunnerConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app runner config not found response has a 3xx status code
func (o *CreateAppRunnerConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app runner config not found response has a 4xx status code
func (o *CreateAppRunnerConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app runner config not found response has a 5xx status code
func (o *CreateAppRunnerConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app runner config not found response a status code equal to that given
func (o *CreateAppRunnerConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app runner config not found response
func (o *CreateAppRunnerConfigNotFound) Code() int {
	return 404
}

func (o *CreateAppRunnerConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppRunnerConfigNotFound) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppRunnerConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppRunnerConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppRunnerConfigInternalServerError creates a CreateAppRunnerConfigInternalServerError with default headers values
func NewCreateAppRunnerConfigInternalServerError() *CreateAppRunnerConfigInternalServerError {
	return &CreateAppRunnerConfigInternalServerError{}
}

/*
CreateAppRunnerConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppRunnerConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app runner config internal server error response has a 2xx status code
func (o *CreateAppRunnerConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app runner config internal server error response has a 3xx status code
func (o *CreateAppRunnerConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app runner config internal server error response has a 4xx status code
func (o *CreateAppRunnerConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app runner config internal server error response has a 5xx status code
func (o *CreateAppRunnerConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app runner config internal server error response a status code equal to that given
func (o *CreateAppRunnerConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app runner config internal server error response
func (o *CreateAppRunnerConfigInternalServerError) Code() int {
	return 500
}

func (o *CreateAppRunnerConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppRunnerConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/runner-configs][%d] createAppRunnerConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppRunnerConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppRunnerConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
