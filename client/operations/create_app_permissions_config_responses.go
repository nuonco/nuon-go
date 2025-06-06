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

// CreateAppPermissionsConfigReader is a Reader for the CreateAppPermissionsConfig structure.
type CreateAppPermissionsConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppPermissionsConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppPermissionsConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppPermissionsConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppPermissionsConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppPermissionsConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppPermissionsConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppPermissionsConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/permissions-configs] CreateAppPermissionsConfig", response, response.Code())
	}
}

// NewCreateAppPermissionsConfigCreated creates a CreateAppPermissionsConfigCreated with default headers values
func NewCreateAppPermissionsConfigCreated() *CreateAppPermissionsConfigCreated {
	return &CreateAppPermissionsConfigCreated{}
}

/*
CreateAppPermissionsConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppPermissionsConfigCreated struct {
	Payload *models.AppAppPermissionsConfig
}

// IsSuccess returns true when this create app permissions config created response has a 2xx status code
func (o *CreateAppPermissionsConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app permissions config created response has a 3xx status code
func (o *CreateAppPermissionsConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app permissions config created response has a 4xx status code
func (o *CreateAppPermissionsConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app permissions config created response has a 5xx status code
func (o *CreateAppPermissionsConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app permissions config created response a status code equal to that given
func (o *CreateAppPermissionsConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app permissions config created response
func (o *CreateAppPermissionsConfigCreated) Code() int {
	return 201
}

func (o *CreateAppPermissionsConfigCreated) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppPermissionsConfigCreated) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppPermissionsConfigCreated) GetPayload() *models.AppAppPermissionsConfig {
	return o.Payload
}

func (o *CreateAppPermissionsConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppPermissionsConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppPermissionsConfigBadRequest creates a CreateAppPermissionsConfigBadRequest with default headers values
func NewCreateAppPermissionsConfigBadRequest() *CreateAppPermissionsConfigBadRequest {
	return &CreateAppPermissionsConfigBadRequest{}
}

/*
CreateAppPermissionsConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppPermissionsConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app permissions config bad request response has a 2xx status code
func (o *CreateAppPermissionsConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app permissions config bad request response has a 3xx status code
func (o *CreateAppPermissionsConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app permissions config bad request response has a 4xx status code
func (o *CreateAppPermissionsConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app permissions config bad request response has a 5xx status code
func (o *CreateAppPermissionsConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app permissions config bad request response a status code equal to that given
func (o *CreateAppPermissionsConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app permissions config bad request response
func (o *CreateAppPermissionsConfigBadRequest) Code() int {
	return 400
}

func (o *CreateAppPermissionsConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppPermissionsConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppPermissionsConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppPermissionsConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppPermissionsConfigUnauthorized creates a CreateAppPermissionsConfigUnauthorized with default headers values
func NewCreateAppPermissionsConfigUnauthorized() *CreateAppPermissionsConfigUnauthorized {
	return &CreateAppPermissionsConfigUnauthorized{}
}

/*
CreateAppPermissionsConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppPermissionsConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app permissions config unauthorized response has a 2xx status code
func (o *CreateAppPermissionsConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app permissions config unauthorized response has a 3xx status code
func (o *CreateAppPermissionsConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app permissions config unauthorized response has a 4xx status code
func (o *CreateAppPermissionsConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app permissions config unauthorized response has a 5xx status code
func (o *CreateAppPermissionsConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app permissions config unauthorized response a status code equal to that given
func (o *CreateAppPermissionsConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app permissions config unauthorized response
func (o *CreateAppPermissionsConfigUnauthorized) Code() int {
	return 401
}

func (o *CreateAppPermissionsConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppPermissionsConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppPermissionsConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppPermissionsConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppPermissionsConfigForbidden creates a CreateAppPermissionsConfigForbidden with default headers values
func NewCreateAppPermissionsConfigForbidden() *CreateAppPermissionsConfigForbidden {
	return &CreateAppPermissionsConfigForbidden{}
}

/*
CreateAppPermissionsConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppPermissionsConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app permissions config forbidden response has a 2xx status code
func (o *CreateAppPermissionsConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app permissions config forbidden response has a 3xx status code
func (o *CreateAppPermissionsConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app permissions config forbidden response has a 4xx status code
func (o *CreateAppPermissionsConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app permissions config forbidden response has a 5xx status code
func (o *CreateAppPermissionsConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app permissions config forbidden response a status code equal to that given
func (o *CreateAppPermissionsConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app permissions config forbidden response
func (o *CreateAppPermissionsConfigForbidden) Code() int {
	return 403
}

func (o *CreateAppPermissionsConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppPermissionsConfigForbidden) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppPermissionsConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppPermissionsConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppPermissionsConfigNotFound creates a CreateAppPermissionsConfigNotFound with default headers values
func NewCreateAppPermissionsConfigNotFound() *CreateAppPermissionsConfigNotFound {
	return &CreateAppPermissionsConfigNotFound{}
}

/*
CreateAppPermissionsConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppPermissionsConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app permissions config not found response has a 2xx status code
func (o *CreateAppPermissionsConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app permissions config not found response has a 3xx status code
func (o *CreateAppPermissionsConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app permissions config not found response has a 4xx status code
func (o *CreateAppPermissionsConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app permissions config not found response has a 5xx status code
func (o *CreateAppPermissionsConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app permissions config not found response a status code equal to that given
func (o *CreateAppPermissionsConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app permissions config not found response
func (o *CreateAppPermissionsConfigNotFound) Code() int {
	return 404
}

func (o *CreateAppPermissionsConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppPermissionsConfigNotFound) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppPermissionsConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppPermissionsConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppPermissionsConfigInternalServerError creates a CreateAppPermissionsConfigInternalServerError with default headers values
func NewCreateAppPermissionsConfigInternalServerError() *CreateAppPermissionsConfigInternalServerError {
	return &CreateAppPermissionsConfigInternalServerError{}
}

/*
CreateAppPermissionsConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppPermissionsConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app permissions config internal server error response has a 2xx status code
func (o *CreateAppPermissionsConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app permissions config internal server error response has a 3xx status code
func (o *CreateAppPermissionsConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app permissions config internal server error response has a 4xx status code
func (o *CreateAppPermissionsConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app permissions config internal server error response has a 5xx status code
func (o *CreateAppPermissionsConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app permissions config internal server error response a status code equal to that given
func (o *CreateAppPermissionsConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app permissions config internal server error response
func (o *CreateAppPermissionsConfigInternalServerError) Code() int {
	return 500
}

func (o *CreateAppPermissionsConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppPermissionsConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/permissions-configs][%d] createAppPermissionsConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppPermissionsConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppPermissionsConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
