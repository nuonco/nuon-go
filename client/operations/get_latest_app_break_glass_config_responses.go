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

// GetLatestAppBreakGlassConfigReader is a Reader for the GetLatestAppBreakGlassConfig structure.
type GetLatestAppBreakGlassConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestAppBreakGlassConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestAppBreakGlassConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLatestAppBreakGlassConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLatestAppBreakGlassConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestAppBreakGlassConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLatestAppBreakGlassConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestAppBreakGlassConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/latest-app-break-glass-config] GetLatestAppBreakGlassConfig", response, response.Code())
	}
}

// NewGetLatestAppBreakGlassConfigOK creates a GetLatestAppBreakGlassConfigOK with default headers values
func NewGetLatestAppBreakGlassConfigOK() *GetLatestAppBreakGlassConfigOK {
	return &GetLatestAppBreakGlassConfigOK{}
}

/*
GetLatestAppBreakGlassConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetLatestAppBreakGlassConfigOK struct {
	Payload *models.AppAppBreakGlassConfig
}

// IsSuccess returns true when this get latest app break glass config o k response has a 2xx status code
func (o *GetLatestAppBreakGlassConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest app break glass config o k response has a 3xx status code
func (o *GetLatestAppBreakGlassConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest app break glass config o k response has a 4xx status code
func (o *GetLatestAppBreakGlassConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest app break glass config o k response has a 5xx status code
func (o *GetLatestAppBreakGlassConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest app break glass config o k response a status code equal to that given
func (o *GetLatestAppBreakGlassConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get latest app break glass config o k response
func (o *GetLatestAppBreakGlassConfigOK) Code() int {
	return 200
}

func (o *GetLatestAppBreakGlassConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigOK  %+v", 200, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigOK  %+v", 200, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigOK) GetPayload() *models.AppAppBreakGlassConfig {
	return o.Payload
}

func (o *GetLatestAppBreakGlassConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppBreakGlassConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestAppBreakGlassConfigBadRequest creates a GetLatestAppBreakGlassConfigBadRequest with default headers values
func NewGetLatestAppBreakGlassConfigBadRequest() *GetLatestAppBreakGlassConfigBadRequest {
	return &GetLatestAppBreakGlassConfigBadRequest{}
}

/*
GetLatestAppBreakGlassConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLatestAppBreakGlassConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get latest app break glass config bad request response has a 2xx status code
func (o *GetLatestAppBreakGlassConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest app break glass config bad request response has a 3xx status code
func (o *GetLatestAppBreakGlassConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest app break glass config bad request response has a 4xx status code
func (o *GetLatestAppBreakGlassConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest app break glass config bad request response has a 5xx status code
func (o *GetLatestAppBreakGlassConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest app break glass config bad request response a status code equal to that given
func (o *GetLatestAppBreakGlassConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get latest app break glass config bad request response
func (o *GetLatestAppBreakGlassConfigBadRequest) Code() int {
	return 400
}

func (o *GetLatestAppBreakGlassConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetLatestAppBreakGlassConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestAppBreakGlassConfigUnauthorized creates a GetLatestAppBreakGlassConfigUnauthorized with default headers values
func NewGetLatestAppBreakGlassConfigUnauthorized() *GetLatestAppBreakGlassConfigUnauthorized {
	return &GetLatestAppBreakGlassConfigUnauthorized{}
}

/*
GetLatestAppBreakGlassConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLatestAppBreakGlassConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get latest app break glass config unauthorized response has a 2xx status code
func (o *GetLatestAppBreakGlassConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest app break glass config unauthorized response has a 3xx status code
func (o *GetLatestAppBreakGlassConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest app break glass config unauthorized response has a 4xx status code
func (o *GetLatestAppBreakGlassConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest app break glass config unauthorized response has a 5xx status code
func (o *GetLatestAppBreakGlassConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest app break glass config unauthorized response a status code equal to that given
func (o *GetLatestAppBreakGlassConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get latest app break glass config unauthorized response
func (o *GetLatestAppBreakGlassConfigUnauthorized) Code() int {
	return 401
}

func (o *GetLatestAppBreakGlassConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetLatestAppBreakGlassConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestAppBreakGlassConfigForbidden creates a GetLatestAppBreakGlassConfigForbidden with default headers values
func NewGetLatestAppBreakGlassConfigForbidden() *GetLatestAppBreakGlassConfigForbidden {
	return &GetLatestAppBreakGlassConfigForbidden{}
}

/*
GetLatestAppBreakGlassConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLatestAppBreakGlassConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get latest app break glass config forbidden response has a 2xx status code
func (o *GetLatestAppBreakGlassConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest app break glass config forbidden response has a 3xx status code
func (o *GetLatestAppBreakGlassConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest app break glass config forbidden response has a 4xx status code
func (o *GetLatestAppBreakGlassConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest app break glass config forbidden response has a 5xx status code
func (o *GetLatestAppBreakGlassConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest app break glass config forbidden response a status code equal to that given
func (o *GetLatestAppBreakGlassConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get latest app break glass config forbidden response
func (o *GetLatestAppBreakGlassConfigForbidden) Code() int {
	return 403
}

func (o *GetLatestAppBreakGlassConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetLatestAppBreakGlassConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestAppBreakGlassConfigNotFound creates a GetLatestAppBreakGlassConfigNotFound with default headers values
func NewGetLatestAppBreakGlassConfigNotFound() *GetLatestAppBreakGlassConfigNotFound {
	return &GetLatestAppBreakGlassConfigNotFound{}
}

/*
GetLatestAppBreakGlassConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLatestAppBreakGlassConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get latest app break glass config not found response has a 2xx status code
func (o *GetLatestAppBreakGlassConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest app break glass config not found response has a 3xx status code
func (o *GetLatestAppBreakGlassConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest app break glass config not found response has a 4xx status code
func (o *GetLatestAppBreakGlassConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest app break glass config not found response has a 5xx status code
func (o *GetLatestAppBreakGlassConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest app break glass config not found response a status code equal to that given
func (o *GetLatestAppBreakGlassConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get latest app break glass config not found response
func (o *GetLatestAppBreakGlassConfigNotFound) Code() int {
	return 404
}

func (o *GetLatestAppBreakGlassConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetLatestAppBreakGlassConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestAppBreakGlassConfigInternalServerError creates a GetLatestAppBreakGlassConfigInternalServerError with default headers values
func NewGetLatestAppBreakGlassConfigInternalServerError() *GetLatestAppBreakGlassConfigInternalServerError {
	return &GetLatestAppBreakGlassConfigInternalServerError{}
}

/*
GetLatestAppBreakGlassConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLatestAppBreakGlassConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get latest app break glass config internal server error response has a 2xx status code
func (o *GetLatestAppBreakGlassConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest app break glass config internal server error response has a 3xx status code
func (o *GetLatestAppBreakGlassConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest app break glass config internal server error response has a 4xx status code
func (o *GetLatestAppBreakGlassConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest app break glass config internal server error response has a 5xx status code
func (o *GetLatestAppBreakGlassConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get latest app break glass config internal server error response a status code equal to that given
func (o *GetLatestAppBreakGlassConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get latest app break glass config internal server error response
func (o *GetLatestAppBreakGlassConfigInternalServerError) Code() int {
	return 500
}

func (o *GetLatestAppBreakGlassConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/latest-app-break-glass-config][%d] getLatestAppBreakGlassConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestAppBreakGlassConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetLatestAppBreakGlassConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
