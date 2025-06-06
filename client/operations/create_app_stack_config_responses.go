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

// CreateAppStackConfigReader is a Reader for the CreateAppStackConfig structure.
type CreateAppStackConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppStackConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppStackConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppStackConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppStackConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppStackConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppStackConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppStackConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/stack-configs] CreateAppStackConfig", response, response.Code())
	}
}

// NewCreateAppStackConfigCreated creates a CreateAppStackConfigCreated with default headers values
func NewCreateAppStackConfigCreated() *CreateAppStackConfigCreated {
	return &CreateAppStackConfigCreated{}
}

/*
CreateAppStackConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppStackConfigCreated struct {
	Payload *models.AppAppStackConfig
}

// IsSuccess returns true when this create app stack config created response has a 2xx status code
func (o *CreateAppStackConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app stack config created response has a 3xx status code
func (o *CreateAppStackConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app stack config created response has a 4xx status code
func (o *CreateAppStackConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app stack config created response has a 5xx status code
func (o *CreateAppStackConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app stack config created response a status code equal to that given
func (o *CreateAppStackConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app stack config created response
func (o *CreateAppStackConfigCreated) Code() int {
	return 201
}

func (o *CreateAppStackConfigCreated) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppStackConfigCreated) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppStackConfigCreated) GetPayload() *models.AppAppStackConfig {
	return o.Payload
}

func (o *CreateAppStackConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppStackConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppStackConfigBadRequest creates a CreateAppStackConfigBadRequest with default headers values
func NewCreateAppStackConfigBadRequest() *CreateAppStackConfigBadRequest {
	return &CreateAppStackConfigBadRequest{}
}

/*
CreateAppStackConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppStackConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app stack config bad request response has a 2xx status code
func (o *CreateAppStackConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app stack config bad request response has a 3xx status code
func (o *CreateAppStackConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app stack config bad request response has a 4xx status code
func (o *CreateAppStackConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app stack config bad request response has a 5xx status code
func (o *CreateAppStackConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app stack config bad request response a status code equal to that given
func (o *CreateAppStackConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app stack config bad request response
func (o *CreateAppStackConfigBadRequest) Code() int {
	return 400
}

func (o *CreateAppStackConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppStackConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppStackConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppStackConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppStackConfigUnauthorized creates a CreateAppStackConfigUnauthorized with default headers values
func NewCreateAppStackConfigUnauthorized() *CreateAppStackConfigUnauthorized {
	return &CreateAppStackConfigUnauthorized{}
}

/*
CreateAppStackConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppStackConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app stack config unauthorized response has a 2xx status code
func (o *CreateAppStackConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app stack config unauthorized response has a 3xx status code
func (o *CreateAppStackConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app stack config unauthorized response has a 4xx status code
func (o *CreateAppStackConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app stack config unauthorized response has a 5xx status code
func (o *CreateAppStackConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app stack config unauthorized response a status code equal to that given
func (o *CreateAppStackConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app stack config unauthorized response
func (o *CreateAppStackConfigUnauthorized) Code() int {
	return 401
}

func (o *CreateAppStackConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppStackConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppStackConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppStackConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppStackConfigForbidden creates a CreateAppStackConfigForbidden with default headers values
func NewCreateAppStackConfigForbidden() *CreateAppStackConfigForbidden {
	return &CreateAppStackConfigForbidden{}
}

/*
CreateAppStackConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppStackConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app stack config forbidden response has a 2xx status code
func (o *CreateAppStackConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app stack config forbidden response has a 3xx status code
func (o *CreateAppStackConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app stack config forbidden response has a 4xx status code
func (o *CreateAppStackConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app stack config forbidden response has a 5xx status code
func (o *CreateAppStackConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app stack config forbidden response a status code equal to that given
func (o *CreateAppStackConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app stack config forbidden response
func (o *CreateAppStackConfigForbidden) Code() int {
	return 403
}

func (o *CreateAppStackConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppStackConfigForbidden) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppStackConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppStackConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppStackConfigNotFound creates a CreateAppStackConfigNotFound with default headers values
func NewCreateAppStackConfigNotFound() *CreateAppStackConfigNotFound {
	return &CreateAppStackConfigNotFound{}
}

/*
CreateAppStackConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppStackConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app stack config not found response has a 2xx status code
func (o *CreateAppStackConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app stack config not found response has a 3xx status code
func (o *CreateAppStackConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app stack config not found response has a 4xx status code
func (o *CreateAppStackConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app stack config not found response has a 5xx status code
func (o *CreateAppStackConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app stack config not found response a status code equal to that given
func (o *CreateAppStackConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app stack config not found response
func (o *CreateAppStackConfigNotFound) Code() int {
	return 404
}

func (o *CreateAppStackConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppStackConfigNotFound) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppStackConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppStackConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppStackConfigInternalServerError creates a CreateAppStackConfigInternalServerError with default headers values
func NewCreateAppStackConfigInternalServerError() *CreateAppStackConfigInternalServerError {
	return &CreateAppStackConfigInternalServerError{}
}

/*
CreateAppStackConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppStackConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app stack config internal server error response has a 2xx status code
func (o *CreateAppStackConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app stack config internal server error response has a 3xx status code
func (o *CreateAppStackConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app stack config internal server error response has a 4xx status code
func (o *CreateAppStackConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app stack config internal server error response has a 5xx status code
func (o *CreateAppStackConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app stack config internal server error response a status code equal to that given
func (o *CreateAppStackConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app stack config internal server error response
func (o *CreateAppStackConfigInternalServerError) Code() int {
	return 500
}

func (o *CreateAppStackConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppStackConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/stack-configs][%d] createAppStackConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppStackConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppStackConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
