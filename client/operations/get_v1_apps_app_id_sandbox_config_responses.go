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

// GetV1AppsAppIDSandboxConfigReader is a Reader for the GetV1AppsAppIDSandboxConfig structure.
type GetV1AppsAppIDSandboxConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AppsAppIDSandboxConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AppsAppIDSandboxConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1AppsAppIDSandboxConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1AppsAppIDSandboxConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1AppsAppIDSandboxConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1AppsAppIDSandboxConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1AppsAppIDSandboxConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/sandbox-config] GetV1AppsAppIDSandboxConfig", response, response.Code())
	}
}

// NewGetV1AppsAppIDSandboxConfigOK creates a GetV1AppsAppIDSandboxConfigOK with default headers values
func NewGetV1AppsAppIDSandboxConfigOK() *GetV1AppsAppIDSandboxConfigOK {
	return &GetV1AppsAppIDSandboxConfigOK{}
}

/*
GetV1AppsAppIDSandboxConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetV1AppsAppIDSandboxConfigOK struct {
	Payload *models.AppAppSandboxConfig
}

// IsSuccess returns true when this get v1 apps app Id sandbox config o k response has a 2xx status code
func (o *GetV1AppsAppIDSandboxConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 apps app Id sandbox config o k response has a 3xx status code
func (o *GetV1AppsAppIDSandboxConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id sandbox config o k response has a 4xx status code
func (o *GetV1AppsAppIDSandboxConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 apps app Id sandbox config o k response has a 5xx status code
func (o *GetV1AppsAppIDSandboxConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id sandbox config o k response a status code equal to that given
func (o *GetV1AppsAppIDSandboxConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 apps app Id sandbox config o k response
func (o *GetV1AppsAppIDSandboxConfigOK) Code() int {
	return 200
}

func (o *GetV1AppsAppIDSandboxConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigOK  %+v", 200, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigOK  %+v", 200, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigOK) GetPayload() *models.AppAppSandboxConfig {
	return o.Payload
}

func (o *GetV1AppsAppIDSandboxConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppSandboxConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDSandboxConfigBadRequest creates a GetV1AppsAppIDSandboxConfigBadRequest with default headers values
func NewGetV1AppsAppIDSandboxConfigBadRequest() *GetV1AppsAppIDSandboxConfigBadRequest {
	return &GetV1AppsAppIDSandboxConfigBadRequest{}
}

/*
GetV1AppsAppIDSandboxConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1AppsAppIDSandboxConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id sandbox config bad request response has a 2xx status code
func (o *GetV1AppsAppIDSandboxConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id sandbox config bad request response has a 3xx status code
func (o *GetV1AppsAppIDSandboxConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id sandbox config bad request response has a 4xx status code
func (o *GetV1AppsAppIDSandboxConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id sandbox config bad request response has a 5xx status code
func (o *GetV1AppsAppIDSandboxConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id sandbox config bad request response a status code equal to that given
func (o *GetV1AppsAppIDSandboxConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 apps app Id sandbox config bad request response
func (o *GetV1AppsAppIDSandboxConfigBadRequest) Code() int {
	return 400
}

func (o *GetV1AppsAppIDSandboxConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDSandboxConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDSandboxConfigUnauthorized creates a GetV1AppsAppIDSandboxConfigUnauthorized with default headers values
func NewGetV1AppsAppIDSandboxConfigUnauthorized() *GetV1AppsAppIDSandboxConfigUnauthorized {
	return &GetV1AppsAppIDSandboxConfigUnauthorized{}
}

/*
GetV1AppsAppIDSandboxConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1AppsAppIDSandboxConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id sandbox config unauthorized response has a 2xx status code
func (o *GetV1AppsAppIDSandboxConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id sandbox config unauthorized response has a 3xx status code
func (o *GetV1AppsAppIDSandboxConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id sandbox config unauthorized response has a 4xx status code
func (o *GetV1AppsAppIDSandboxConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id sandbox config unauthorized response has a 5xx status code
func (o *GetV1AppsAppIDSandboxConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id sandbox config unauthorized response a status code equal to that given
func (o *GetV1AppsAppIDSandboxConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 apps app Id sandbox config unauthorized response
func (o *GetV1AppsAppIDSandboxConfigUnauthorized) Code() int {
	return 401
}

func (o *GetV1AppsAppIDSandboxConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDSandboxConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDSandboxConfigForbidden creates a GetV1AppsAppIDSandboxConfigForbidden with default headers values
func NewGetV1AppsAppIDSandboxConfigForbidden() *GetV1AppsAppIDSandboxConfigForbidden {
	return &GetV1AppsAppIDSandboxConfigForbidden{}
}

/*
GetV1AppsAppIDSandboxConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1AppsAppIDSandboxConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id sandbox config forbidden response has a 2xx status code
func (o *GetV1AppsAppIDSandboxConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id sandbox config forbidden response has a 3xx status code
func (o *GetV1AppsAppIDSandboxConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id sandbox config forbidden response has a 4xx status code
func (o *GetV1AppsAppIDSandboxConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id sandbox config forbidden response has a 5xx status code
func (o *GetV1AppsAppIDSandboxConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id sandbox config forbidden response a status code equal to that given
func (o *GetV1AppsAppIDSandboxConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 apps app Id sandbox config forbidden response
func (o *GetV1AppsAppIDSandboxConfigForbidden) Code() int {
	return 403
}

func (o *GetV1AppsAppIDSandboxConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDSandboxConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDSandboxConfigNotFound creates a GetV1AppsAppIDSandboxConfigNotFound with default headers values
func NewGetV1AppsAppIDSandboxConfigNotFound() *GetV1AppsAppIDSandboxConfigNotFound {
	return &GetV1AppsAppIDSandboxConfigNotFound{}
}

/*
GetV1AppsAppIDSandboxConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1AppsAppIDSandboxConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id sandbox config not found response has a 2xx status code
func (o *GetV1AppsAppIDSandboxConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id sandbox config not found response has a 3xx status code
func (o *GetV1AppsAppIDSandboxConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id sandbox config not found response has a 4xx status code
func (o *GetV1AppsAppIDSandboxConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id sandbox config not found response has a 5xx status code
func (o *GetV1AppsAppIDSandboxConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id sandbox config not found response a status code equal to that given
func (o *GetV1AppsAppIDSandboxConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 apps app Id sandbox config not found response
func (o *GetV1AppsAppIDSandboxConfigNotFound) Code() int {
	return 404
}

func (o *GetV1AppsAppIDSandboxConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDSandboxConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDSandboxConfigInternalServerError creates a GetV1AppsAppIDSandboxConfigInternalServerError with default headers values
func NewGetV1AppsAppIDSandboxConfigInternalServerError() *GetV1AppsAppIDSandboxConfigInternalServerError {
	return &GetV1AppsAppIDSandboxConfigInternalServerError{}
}

/*
GetV1AppsAppIDSandboxConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1AppsAppIDSandboxConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id sandbox config internal server error response has a 2xx status code
func (o *GetV1AppsAppIDSandboxConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id sandbox config internal server error response has a 3xx status code
func (o *GetV1AppsAppIDSandboxConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id sandbox config internal server error response has a 4xx status code
func (o *GetV1AppsAppIDSandboxConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 apps app Id sandbox config internal server error response has a 5xx status code
func (o *GetV1AppsAppIDSandboxConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 apps app Id sandbox config internal server error response a status code equal to that given
func (o *GetV1AppsAppIDSandboxConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 apps app Id sandbox config internal server error response
func (o *GetV1AppsAppIDSandboxConfigInternalServerError) Code() int {
	return 500
}

func (o *GetV1AppsAppIDSandboxConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/sandbox-config][%d] getV1AppsAppIdSandboxConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1AppsAppIDSandboxConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDSandboxConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
