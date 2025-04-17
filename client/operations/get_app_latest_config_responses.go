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

// GetAppLatestConfigReader is a Reader for the GetAppLatestConfig structure.
type GetAppLatestConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppLatestConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppLatestConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppLatestConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppLatestConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppLatestConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppLatestConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppLatestConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/latest-config] GetAppLatestConfig", response, response.Code())
	}
}

// NewGetAppLatestConfigOK creates a GetAppLatestConfigOK with default headers values
func NewGetAppLatestConfigOK() *GetAppLatestConfigOK {
	return &GetAppLatestConfigOK{}
}

/*
GetAppLatestConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppLatestConfigOK struct {
	Payload *models.AppAppConfig
}

// IsSuccess returns true when this get app latest config o k response has a 2xx status code
func (o *GetAppLatestConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app latest config o k response has a 3xx status code
func (o *GetAppLatestConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app latest config o k response has a 4xx status code
func (o *GetAppLatestConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app latest config o k response has a 5xx status code
func (o *GetAppLatestConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app latest config o k response a status code equal to that given
func (o *GetAppLatestConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app latest config o k response
func (o *GetAppLatestConfigOK) Code() int {
	return 200
}

func (o *GetAppLatestConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppLatestConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppLatestConfigOK) GetPayload() *models.AppAppConfig {
	return o.Payload
}

func (o *GetAppLatestConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppLatestConfigBadRequest creates a GetAppLatestConfigBadRequest with default headers values
func NewGetAppLatestConfigBadRequest() *GetAppLatestConfigBadRequest {
	return &GetAppLatestConfigBadRequest{}
}

/*
GetAppLatestConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppLatestConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app latest config bad request response has a 2xx status code
func (o *GetAppLatestConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app latest config bad request response has a 3xx status code
func (o *GetAppLatestConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app latest config bad request response has a 4xx status code
func (o *GetAppLatestConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app latest config bad request response has a 5xx status code
func (o *GetAppLatestConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app latest config bad request response a status code equal to that given
func (o *GetAppLatestConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app latest config bad request response
func (o *GetAppLatestConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppLatestConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppLatestConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppLatestConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppLatestConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppLatestConfigUnauthorized creates a GetAppLatestConfigUnauthorized with default headers values
func NewGetAppLatestConfigUnauthorized() *GetAppLatestConfigUnauthorized {
	return &GetAppLatestConfigUnauthorized{}
}

/*
GetAppLatestConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppLatestConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app latest config unauthorized response has a 2xx status code
func (o *GetAppLatestConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app latest config unauthorized response has a 3xx status code
func (o *GetAppLatestConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app latest config unauthorized response has a 4xx status code
func (o *GetAppLatestConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app latest config unauthorized response has a 5xx status code
func (o *GetAppLatestConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app latest config unauthorized response a status code equal to that given
func (o *GetAppLatestConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app latest config unauthorized response
func (o *GetAppLatestConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppLatestConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppLatestConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppLatestConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppLatestConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppLatestConfigForbidden creates a GetAppLatestConfigForbidden with default headers values
func NewGetAppLatestConfigForbidden() *GetAppLatestConfigForbidden {
	return &GetAppLatestConfigForbidden{}
}

/*
GetAppLatestConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppLatestConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app latest config forbidden response has a 2xx status code
func (o *GetAppLatestConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app latest config forbidden response has a 3xx status code
func (o *GetAppLatestConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app latest config forbidden response has a 4xx status code
func (o *GetAppLatestConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app latest config forbidden response has a 5xx status code
func (o *GetAppLatestConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app latest config forbidden response a status code equal to that given
func (o *GetAppLatestConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app latest config forbidden response
func (o *GetAppLatestConfigForbidden) Code() int {
	return 403
}

func (o *GetAppLatestConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppLatestConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppLatestConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppLatestConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppLatestConfigNotFound creates a GetAppLatestConfigNotFound with default headers values
func NewGetAppLatestConfigNotFound() *GetAppLatestConfigNotFound {
	return &GetAppLatestConfigNotFound{}
}

/*
GetAppLatestConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppLatestConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app latest config not found response has a 2xx status code
func (o *GetAppLatestConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app latest config not found response has a 3xx status code
func (o *GetAppLatestConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app latest config not found response has a 4xx status code
func (o *GetAppLatestConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app latest config not found response has a 5xx status code
func (o *GetAppLatestConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app latest config not found response a status code equal to that given
func (o *GetAppLatestConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app latest config not found response
func (o *GetAppLatestConfigNotFound) Code() int {
	return 404
}

func (o *GetAppLatestConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppLatestConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppLatestConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppLatestConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppLatestConfigInternalServerError creates a GetAppLatestConfigInternalServerError with default headers values
func NewGetAppLatestConfigInternalServerError() *GetAppLatestConfigInternalServerError {
	return &GetAppLatestConfigInternalServerError{}
}

/*
GetAppLatestConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppLatestConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app latest config internal server error response has a 2xx status code
func (o *GetAppLatestConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app latest config internal server error response has a 3xx status code
func (o *GetAppLatestConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app latest config internal server error response has a 4xx status code
func (o *GetAppLatestConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app latest config internal server error response has a 5xx status code
func (o *GetAppLatestConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app latest config internal server error response a status code equal to that given
func (o *GetAppLatestConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app latest config internal server error response
func (o *GetAppLatestConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppLatestConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppLatestConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-config][%d] getAppLatestConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppLatestConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppLatestConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
