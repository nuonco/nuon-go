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

// GetAppSandboxLatestConfigReader is a Reader for the GetAppSandboxLatestConfig structure.
type GetAppSandboxLatestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppSandboxLatestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppSandboxLatestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppSandboxLatestConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppSandboxLatestConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppSandboxLatestConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppSandboxLatestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppSandboxLatestConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/sandbox-latest-config] GetAppSandboxLatestConfig", response, response.Code())
	}
}

// NewGetAppSandboxLatestConfigOK creates a GetAppSandboxLatestConfigOK with default headers values
func NewGetAppSandboxLatestConfigOK() *GetAppSandboxLatestConfigOK {
	return &GetAppSandboxLatestConfigOK{}
}

/*
GetAppSandboxLatestConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppSandboxLatestConfigOK struct {
	Payload *models.AppAppSandboxConfig
}

// IsSuccess returns true when this get app sandbox latest config o k response has a 2xx status code
func (o *GetAppSandboxLatestConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app sandbox latest config o k response has a 3xx status code
func (o *GetAppSandboxLatestConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app sandbox latest config o k response has a 4xx status code
func (o *GetAppSandboxLatestConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app sandbox latest config o k response has a 5xx status code
func (o *GetAppSandboxLatestConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app sandbox latest config o k response a status code equal to that given
func (o *GetAppSandboxLatestConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app sandbox latest config o k response
func (o *GetAppSandboxLatestConfigOK) Code() int {
	return 200
}

func (o *GetAppSandboxLatestConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigOK %s", 200, payload)
}

func (o *GetAppSandboxLatestConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigOK %s", 200, payload)
}

func (o *GetAppSandboxLatestConfigOK) GetPayload() *models.AppAppSandboxConfig {
	return o.Payload
}

func (o *GetAppSandboxLatestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppSandboxConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSandboxLatestConfigBadRequest creates a GetAppSandboxLatestConfigBadRequest with default headers values
func NewGetAppSandboxLatestConfigBadRequest() *GetAppSandboxLatestConfigBadRequest {
	return &GetAppSandboxLatestConfigBadRequest{}
}

/*
GetAppSandboxLatestConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppSandboxLatestConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app sandbox latest config bad request response has a 2xx status code
func (o *GetAppSandboxLatestConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app sandbox latest config bad request response has a 3xx status code
func (o *GetAppSandboxLatestConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app sandbox latest config bad request response has a 4xx status code
func (o *GetAppSandboxLatestConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app sandbox latest config bad request response has a 5xx status code
func (o *GetAppSandboxLatestConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app sandbox latest config bad request response a status code equal to that given
func (o *GetAppSandboxLatestConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app sandbox latest config bad request response
func (o *GetAppSandboxLatestConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppSandboxLatestConfigBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigBadRequest %s", 400, payload)
}

func (o *GetAppSandboxLatestConfigBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigBadRequest %s", 400, payload)
}

func (o *GetAppSandboxLatestConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSandboxLatestConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSandboxLatestConfigUnauthorized creates a GetAppSandboxLatestConfigUnauthorized with default headers values
func NewGetAppSandboxLatestConfigUnauthorized() *GetAppSandboxLatestConfigUnauthorized {
	return &GetAppSandboxLatestConfigUnauthorized{}
}

/*
GetAppSandboxLatestConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppSandboxLatestConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app sandbox latest config unauthorized response has a 2xx status code
func (o *GetAppSandboxLatestConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app sandbox latest config unauthorized response has a 3xx status code
func (o *GetAppSandboxLatestConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app sandbox latest config unauthorized response has a 4xx status code
func (o *GetAppSandboxLatestConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app sandbox latest config unauthorized response has a 5xx status code
func (o *GetAppSandboxLatestConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app sandbox latest config unauthorized response a status code equal to that given
func (o *GetAppSandboxLatestConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app sandbox latest config unauthorized response
func (o *GetAppSandboxLatestConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppSandboxLatestConfigUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigUnauthorized %s", 401, payload)
}

func (o *GetAppSandboxLatestConfigUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigUnauthorized %s", 401, payload)
}

func (o *GetAppSandboxLatestConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSandboxLatestConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSandboxLatestConfigForbidden creates a GetAppSandboxLatestConfigForbidden with default headers values
func NewGetAppSandboxLatestConfigForbidden() *GetAppSandboxLatestConfigForbidden {
	return &GetAppSandboxLatestConfigForbidden{}
}

/*
GetAppSandboxLatestConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppSandboxLatestConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app sandbox latest config forbidden response has a 2xx status code
func (o *GetAppSandboxLatestConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app sandbox latest config forbidden response has a 3xx status code
func (o *GetAppSandboxLatestConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app sandbox latest config forbidden response has a 4xx status code
func (o *GetAppSandboxLatestConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app sandbox latest config forbidden response has a 5xx status code
func (o *GetAppSandboxLatestConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app sandbox latest config forbidden response a status code equal to that given
func (o *GetAppSandboxLatestConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app sandbox latest config forbidden response
func (o *GetAppSandboxLatestConfigForbidden) Code() int {
	return 403
}

func (o *GetAppSandboxLatestConfigForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigForbidden %s", 403, payload)
}

func (o *GetAppSandboxLatestConfigForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigForbidden %s", 403, payload)
}

func (o *GetAppSandboxLatestConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSandboxLatestConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSandboxLatestConfigNotFound creates a GetAppSandboxLatestConfigNotFound with default headers values
func NewGetAppSandboxLatestConfigNotFound() *GetAppSandboxLatestConfigNotFound {
	return &GetAppSandboxLatestConfigNotFound{}
}

/*
GetAppSandboxLatestConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppSandboxLatestConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app sandbox latest config not found response has a 2xx status code
func (o *GetAppSandboxLatestConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app sandbox latest config not found response has a 3xx status code
func (o *GetAppSandboxLatestConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app sandbox latest config not found response has a 4xx status code
func (o *GetAppSandboxLatestConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app sandbox latest config not found response has a 5xx status code
func (o *GetAppSandboxLatestConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app sandbox latest config not found response a status code equal to that given
func (o *GetAppSandboxLatestConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app sandbox latest config not found response
func (o *GetAppSandboxLatestConfigNotFound) Code() int {
	return 404
}

func (o *GetAppSandboxLatestConfigNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigNotFound %s", 404, payload)
}

func (o *GetAppSandboxLatestConfigNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigNotFound %s", 404, payload)
}

func (o *GetAppSandboxLatestConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSandboxLatestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSandboxLatestConfigInternalServerError creates a GetAppSandboxLatestConfigInternalServerError with default headers values
func NewGetAppSandboxLatestConfigInternalServerError() *GetAppSandboxLatestConfigInternalServerError {
	return &GetAppSandboxLatestConfigInternalServerError{}
}

/*
GetAppSandboxLatestConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppSandboxLatestConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app sandbox latest config internal server error response has a 2xx status code
func (o *GetAppSandboxLatestConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app sandbox latest config internal server error response has a 3xx status code
func (o *GetAppSandboxLatestConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app sandbox latest config internal server error response has a 4xx status code
func (o *GetAppSandboxLatestConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app sandbox latest config internal server error response has a 5xx status code
func (o *GetAppSandboxLatestConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app sandbox latest config internal server error response a status code equal to that given
func (o *GetAppSandboxLatestConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app sandbox latest config internal server error response
func (o *GetAppSandboxLatestConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppSandboxLatestConfigInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigInternalServerError %s", 500, payload)
}

func (o *GetAppSandboxLatestConfigInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-latest-config][%d] getAppSandboxLatestConfigInternalServerError %s", 500, payload)
}

func (o *GetAppSandboxLatestConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSandboxLatestConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
