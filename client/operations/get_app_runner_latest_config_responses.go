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

// GetAppRunnerLatestConfigReader is a Reader for the GetAppRunnerLatestConfig structure.
type GetAppRunnerLatestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppRunnerLatestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppRunnerLatestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppRunnerLatestConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppRunnerLatestConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppRunnerLatestConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppRunnerLatestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppRunnerLatestConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/runner-latest-config] GetAppRunnerLatestConfig", response, response.Code())
	}
}

// NewGetAppRunnerLatestConfigOK creates a GetAppRunnerLatestConfigOK with default headers values
func NewGetAppRunnerLatestConfigOK() *GetAppRunnerLatestConfigOK {
	return &GetAppRunnerLatestConfigOK{}
}

/*
GetAppRunnerLatestConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppRunnerLatestConfigOK struct {
	Payload *models.AppAppRunnerConfig
}

// IsSuccess returns true when this get app runner latest config o k response has a 2xx status code
func (o *GetAppRunnerLatestConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app runner latest config o k response has a 3xx status code
func (o *GetAppRunnerLatestConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app runner latest config o k response has a 4xx status code
func (o *GetAppRunnerLatestConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app runner latest config o k response has a 5xx status code
func (o *GetAppRunnerLatestConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app runner latest config o k response a status code equal to that given
func (o *GetAppRunnerLatestConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app runner latest config o k response
func (o *GetAppRunnerLatestConfigOK) Code() int {
	return 200
}

func (o *GetAppRunnerLatestConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppRunnerLatestConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppRunnerLatestConfigOK) GetPayload() *models.AppAppRunnerConfig {
	return o.Payload
}

func (o *GetAppRunnerLatestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppRunnerConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppRunnerLatestConfigBadRequest creates a GetAppRunnerLatestConfigBadRequest with default headers values
func NewGetAppRunnerLatestConfigBadRequest() *GetAppRunnerLatestConfigBadRequest {
	return &GetAppRunnerLatestConfigBadRequest{}
}

/*
GetAppRunnerLatestConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppRunnerLatestConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app runner latest config bad request response has a 2xx status code
func (o *GetAppRunnerLatestConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app runner latest config bad request response has a 3xx status code
func (o *GetAppRunnerLatestConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app runner latest config bad request response has a 4xx status code
func (o *GetAppRunnerLatestConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app runner latest config bad request response has a 5xx status code
func (o *GetAppRunnerLatestConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app runner latest config bad request response a status code equal to that given
func (o *GetAppRunnerLatestConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app runner latest config bad request response
func (o *GetAppRunnerLatestConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppRunnerLatestConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppRunnerLatestConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppRunnerLatestConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppRunnerLatestConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppRunnerLatestConfigUnauthorized creates a GetAppRunnerLatestConfigUnauthorized with default headers values
func NewGetAppRunnerLatestConfigUnauthorized() *GetAppRunnerLatestConfigUnauthorized {
	return &GetAppRunnerLatestConfigUnauthorized{}
}

/*
GetAppRunnerLatestConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppRunnerLatestConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app runner latest config unauthorized response has a 2xx status code
func (o *GetAppRunnerLatestConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app runner latest config unauthorized response has a 3xx status code
func (o *GetAppRunnerLatestConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app runner latest config unauthorized response has a 4xx status code
func (o *GetAppRunnerLatestConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app runner latest config unauthorized response has a 5xx status code
func (o *GetAppRunnerLatestConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app runner latest config unauthorized response a status code equal to that given
func (o *GetAppRunnerLatestConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app runner latest config unauthorized response
func (o *GetAppRunnerLatestConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppRunnerLatestConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppRunnerLatestConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppRunnerLatestConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppRunnerLatestConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppRunnerLatestConfigForbidden creates a GetAppRunnerLatestConfigForbidden with default headers values
func NewGetAppRunnerLatestConfigForbidden() *GetAppRunnerLatestConfigForbidden {
	return &GetAppRunnerLatestConfigForbidden{}
}

/*
GetAppRunnerLatestConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppRunnerLatestConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app runner latest config forbidden response has a 2xx status code
func (o *GetAppRunnerLatestConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app runner latest config forbidden response has a 3xx status code
func (o *GetAppRunnerLatestConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app runner latest config forbidden response has a 4xx status code
func (o *GetAppRunnerLatestConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app runner latest config forbidden response has a 5xx status code
func (o *GetAppRunnerLatestConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app runner latest config forbidden response a status code equal to that given
func (o *GetAppRunnerLatestConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app runner latest config forbidden response
func (o *GetAppRunnerLatestConfigForbidden) Code() int {
	return 403
}

func (o *GetAppRunnerLatestConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppRunnerLatestConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppRunnerLatestConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppRunnerLatestConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppRunnerLatestConfigNotFound creates a GetAppRunnerLatestConfigNotFound with default headers values
func NewGetAppRunnerLatestConfigNotFound() *GetAppRunnerLatestConfigNotFound {
	return &GetAppRunnerLatestConfigNotFound{}
}

/*
GetAppRunnerLatestConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppRunnerLatestConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app runner latest config not found response has a 2xx status code
func (o *GetAppRunnerLatestConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app runner latest config not found response has a 3xx status code
func (o *GetAppRunnerLatestConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app runner latest config not found response has a 4xx status code
func (o *GetAppRunnerLatestConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app runner latest config not found response has a 5xx status code
func (o *GetAppRunnerLatestConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app runner latest config not found response a status code equal to that given
func (o *GetAppRunnerLatestConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app runner latest config not found response
func (o *GetAppRunnerLatestConfigNotFound) Code() int {
	return 404
}

func (o *GetAppRunnerLatestConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppRunnerLatestConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppRunnerLatestConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppRunnerLatestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppRunnerLatestConfigInternalServerError creates a GetAppRunnerLatestConfigInternalServerError with default headers values
func NewGetAppRunnerLatestConfigInternalServerError() *GetAppRunnerLatestConfigInternalServerError {
	return &GetAppRunnerLatestConfigInternalServerError{}
}

/*
GetAppRunnerLatestConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppRunnerLatestConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app runner latest config internal server error response has a 2xx status code
func (o *GetAppRunnerLatestConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app runner latest config internal server error response has a 3xx status code
func (o *GetAppRunnerLatestConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app runner latest config internal server error response has a 4xx status code
func (o *GetAppRunnerLatestConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app runner latest config internal server error response has a 5xx status code
func (o *GetAppRunnerLatestConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app runner latest config internal server error response a status code equal to that given
func (o *GetAppRunnerLatestConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app runner latest config internal server error response
func (o *GetAppRunnerLatestConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppRunnerLatestConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppRunnerLatestConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/runner-latest-config][%d] getAppRunnerLatestConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppRunnerLatestConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppRunnerLatestConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
