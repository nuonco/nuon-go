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

// CreateAppInputConfigReader is a Reader for the CreateAppInputConfig structure.
type CreateAppInputConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppInputConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppInputConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppInputConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppInputConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppInputConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppInputConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppInputConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/input-config] CreateAppInputConfig", response, response.Code())
	}
}

// NewCreateAppInputConfigCreated creates a CreateAppInputConfigCreated with default headers values
func NewCreateAppInputConfigCreated() *CreateAppInputConfigCreated {
	return &CreateAppInputConfigCreated{}
}

/*
CreateAppInputConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppInputConfigCreated struct {
	Payload *models.AppAppInputConfig
}

// IsSuccess returns true when this create app input config created response has a 2xx status code
func (o *CreateAppInputConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app input config created response has a 3xx status code
func (o *CreateAppInputConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app input config created response has a 4xx status code
func (o *CreateAppInputConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app input config created response has a 5xx status code
func (o *CreateAppInputConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app input config created response a status code equal to that given
func (o *CreateAppInputConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app input config created response
func (o *CreateAppInputConfigCreated) Code() int {
	return 201
}

func (o *CreateAppInputConfigCreated) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppInputConfigCreated) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigCreated  %+v", 201, o.Payload)
}

func (o *CreateAppInputConfigCreated) GetPayload() *models.AppAppInputConfig {
	return o.Payload
}

func (o *CreateAppInputConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppInputConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppInputConfigBadRequest creates a CreateAppInputConfigBadRequest with default headers values
func NewCreateAppInputConfigBadRequest() *CreateAppInputConfigBadRequest {
	return &CreateAppInputConfigBadRequest{}
}

/*
CreateAppInputConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppInputConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app input config bad request response has a 2xx status code
func (o *CreateAppInputConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app input config bad request response has a 3xx status code
func (o *CreateAppInputConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app input config bad request response has a 4xx status code
func (o *CreateAppInputConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app input config bad request response has a 5xx status code
func (o *CreateAppInputConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app input config bad request response a status code equal to that given
func (o *CreateAppInputConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app input config bad request response
func (o *CreateAppInputConfigBadRequest) Code() int {
	return 400
}

func (o *CreateAppInputConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppInputConfigBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAppInputConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppInputConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppInputConfigUnauthorized creates a CreateAppInputConfigUnauthorized with default headers values
func NewCreateAppInputConfigUnauthorized() *CreateAppInputConfigUnauthorized {
	return &CreateAppInputConfigUnauthorized{}
}

/*
CreateAppInputConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppInputConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app input config unauthorized response has a 2xx status code
func (o *CreateAppInputConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app input config unauthorized response has a 3xx status code
func (o *CreateAppInputConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app input config unauthorized response has a 4xx status code
func (o *CreateAppInputConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app input config unauthorized response has a 5xx status code
func (o *CreateAppInputConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app input config unauthorized response a status code equal to that given
func (o *CreateAppInputConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app input config unauthorized response
func (o *CreateAppInputConfigUnauthorized) Code() int {
	return 401
}

func (o *CreateAppInputConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppInputConfigUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAppInputConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppInputConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppInputConfigForbidden creates a CreateAppInputConfigForbidden with default headers values
func NewCreateAppInputConfigForbidden() *CreateAppInputConfigForbidden {
	return &CreateAppInputConfigForbidden{}
}

/*
CreateAppInputConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppInputConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app input config forbidden response has a 2xx status code
func (o *CreateAppInputConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app input config forbidden response has a 3xx status code
func (o *CreateAppInputConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app input config forbidden response has a 4xx status code
func (o *CreateAppInputConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app input config forbidden response has a 5xx status code
func (o *CreateAppInputConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app input config forbidden response a status code equal to that given
func (o *CreateAppInputConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app input config forbidden response
func (o *CreateAppInputConfigForbidden) Code() int {
	return 403
}

func (o *CreateAppInputConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppInputConfigForbidden) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigForbidden  %+v", 403, o.Payload)
}

func (o *CreateAppInputConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppInputConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppInputConfigNotFound creates a CreateAppInputConfigNotFound with default headers values
func NewCreateAppInputConfigNotFound() *CreateAppInputConfigNotFound {
	return &CreateAppInputConfigNotFound{}
}

/*
CreateAppInputConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppInputConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app input config not found response has a 2xx status code
func (o *CreateAppInputConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app input config not found response has a 3xx status code
func (o *CreateAppInputConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app input config not found response has a 4xx status code
func (o *CreateAppInputConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app input config not found response has a 5xx status code
func (o *CreateAppInputConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app input config not found response a status code equal to that given
func (o *CreateAppInputConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app input config not found response
func (o *CreateAppInputConfigNotFound) Code() int {
	return 404
}

func (o *CreateAppInputConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppInputConfigNotFound) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigNotFound  %+v", 404, o.Payload)
}

func (o *CreateAppInputConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppInputConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppInputConfigInternalServerError creates a CreateAppInputConfigInternalServerError with default headers values
func NewCreateAppInputConfigInternalServerError() *CreateAppInputConfigInternalServerError {
	return &CreateAppInputConfigInternalServerError{}
}

/*
CreateAppInputConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppInputConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app input config internal server error response has a 2xx status code
func (o *CreateAppInputConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app input config internal server error response has a 3xx status code
func (o *CreateAppInputConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app input config internal server error response has a 4xx status code
func (o *CreateAppInputConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app input config internal server error response has a 5xx status code
func (o *CreateAppInputConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app input config internal server error response a status code equal to that given
func (o *CreateAppInputConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app input config internal server error response
func (o *CreateAppInputConfigInternalServerError) Code() int {
	return 500
}

func (o *CreateAppInputConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppInputConfigInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/apps/{app_id}/input-config][%d] createAppInputConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAppInputConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppInputConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}