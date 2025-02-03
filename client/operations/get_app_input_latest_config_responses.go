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

// GetAppInputLatestConfigReader is a Reader for the GetAppInputLatestConfig structure.
type GetAppInputLatestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppInputLatestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppInputLatestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppInputLatestConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppInputLatestConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppInputLatestConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppInputLatestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppInputLatestConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/input-latest-config] GetAppInputLatestConfig", response, response.Code())
	}
}

// NewGetAppInputLatestConfigOK creates a GetAppInputLatestConfigOK with default headers values
func NewGetAppInputLatestConfigOK() *GetAppInputLatestConfigOK {
	return &GetAppInputLatestConfigOK{}
}

/*
GetAppInputLatestConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppInputLatestConfigOK struct {
	Payload *models.AppAppInputConfig
}

// IsSuccess returns true when this get app input latest config o k response has a 2xx status code
func (o *GetAppInputLatestConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app input latest config o k response has a 3xx status code
func (o *GetAppInputLatestConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app input latest config o k response has a 4xx status code
func (o *GetAppInputLatestConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app input latest config o k response has a 5xx status code
func (o *GetAppInputLatestConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app input latest config o k response a status code equal to that given
func (o *GetAppInputLatestConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app input latest config o k response
func (o *GetAppInputLatestConfigOK) Code() int {
	return 200
}

func (o *GetAppInputLatestConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigOK %s", 200, payload)
}

func (o *GetAppInputLatestConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigOK %s", 200, payload)
}

func (o *GetAppInputLatestConfigOK) GetPayload() *models.AppAppInputConfig {
	return o.Payload
}

func (o *GetAppInputLatestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppInputConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInputLatestConfigBadRequest creates a GetAppInputLatestConfigBadRequest with default headers values
func NewGetAppInputLatestConfigBadRequest() *GetAppInputLatestConfigBadRequest {
	return &GetAppInputLatestConfigBadRequest{}
}

/*
GetAppInputLatestConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppInputLatestConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app input latest config bad request response has a 2xx status code
func (o *GetAppInputLatestConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app input latest config bad request response has a 3xx status code
func (o *GetAppInputLatestConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app input latest config bad request response has a 4xx status code
func (o *GetAppInputLatestConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app input latest config bad request response has a 5xx status code
func (o *GetAppInputLatestConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app input latest config bad request response a status code equal to that given
func (o *GetAppInputLatestConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app input latest config bad request response
func (o *GetAppInputLatestConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppInputLatestConfigBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigBadRequest %s", 400, payload)
}

func (o *GetAppInputLatestConfigBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigBadRequest %s", 400, payload)
}

func (o *GetAppInputLatestConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInputLatestConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInputLatestConfigUnauthorized creates a GetAppInputLatestConfigUnauthorized with default headers values
func NewGetAppInputLatestConfigUnauthorized() *GetAppInputLatestConfigUnauthorized {
	return &GetAppInputLatestConfigUnauthorized{}
}

/*
GetAppInputLatestConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppInputLatestConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app input latest config unauthorized response has a 2xx status code
func (o *GetAppInputLatestConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app input latest config unauthorized response has a 3xx status code
func (o *GetAppInputLatestConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app input latest config unauthorized response has a 4xx status code
func (o *GetAppInputLatestConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app input latest config unauthorized response has a 5xx status code
func (o *GetAppInputLatestConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app input latest config unauthorized response a status code equal to that given
func (o *GetAppInputLatestConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app input latest config unauthorized response
func (o *GetAppInputLatestConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppInputLatestConfigUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigUnauthorized %s", 401, payload)
}

func (o *GetAppInputLatestConfigUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigUnauthorized %s", 401, payload)
}

func (o *GetAppInputLatestConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInputLatestConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInputLatestConfigForbidden creates a GetAppInputLatestConfigForbidden with default headers values
func NewGetAppInputLatestConfigForbidden() *GetAppInputLatestConfigForbidden {
	return &GetAppInputLatestConfigForbidden{}
}

/*
GetAppInputLatestConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppInputLatestConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app input latest config forbidden response has a 2xx status code
func (o *GetAppInputLatestConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app input latest config forbidden response has a 3xx status code
func (o *GetAppInputLatestConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app input latest config forbidden response has a 4xx status code
func (o *GetAppInputLatestConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app input latest config forbidden response has a 5xx status code
func (o *GetAppInputLatestConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app input latest config forbidden response a status code equal to that given
func (o *GetAppInputLatestConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app input latest config forbidden response
func (o *GetAppInputLatestConfigForbidden) Code() int {
	return 403
}

func (o *GetAppInputLatestConfigForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigForbidden %s", 403, payload)
}

func (o *GetAppInputLatestConfigForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigForbidden %s", 403, payload)
}

func (o *GetAppInputLatestConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInputLatestConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInputLatestConfigNotFound creates a GetAppInputLatestConfigNotFound with default headers values
func NewGetAppInputLatestConfigNotFound() *GetAppInputLatestConfigNotFound {
	return &GetAppInputLatestConfigNotFound{}
}

/*
GetAppInputLatestConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppInputLatestConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app input latest config not found response has a 2xx status code
func (o *GetAppInputLatestConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app input latest config not found response has a 3xx status code
func (o *GetAppInputLatestConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app input latest config not found response has a 4xx status code
func (o *GetAppInputLatestConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app input latest config not found response has a 5xx status code
func (o *GetAppInputLatestConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app input latest config not found response a status code equal to that given
func (o *GetAppInputLatestConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app input latest config not found response
func (o *GetAppInputLatestConfigNotFound) Code() int {
	return 404
}

func (o *GetAppInputLatestConfigNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigNotFound %s", 404, payload)
}

func (o *GetAppInputLatestConfigNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigNotFound %s", 404, payload)
}

func (o *GetAppInputLatestConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInputLatestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInputLatestConfigInternalServerError creates a GetAppInputLatestConfigInternalServerError with default headers values
func NewGetAppInputLatestConfigInternalServerError() *GetAppInputLatestConfigInternalServerError {
	return &GetAppInputLatestConfigInternalServerError{}
}

/*
GetAppInputLatestConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppInputLatestConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app input latest config internal server error response has a 2xx status code
func (o *GetAppInputLatestConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app input latest config internal server error response has a 3xx status code
func (o *GetAppInputLatestConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app input latest config internal server error response has a 4xx status code
func (o *GetAppInputLatestConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app input latest config internal server error response has a 5xx status code
func (o *GetAppInputLatestConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app input latest config internal server error response a status code equal to that given
func (o *GetAppInputLatestConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app input latest config internal server error response
func (o *GetAppInputLatestConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppInputLatestConfigInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigInternalServerError %s", 500, payload)
}

func (o *GetAppInputLatestConfigInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/input-latest-config][%d] getAppInputLatestConfigInternalServerError %s", 500, payload)
}

func (o *GetAppInputLatestConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInputLatestConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
