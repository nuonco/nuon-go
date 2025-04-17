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

// GetAppPermissionsConfigReader is a Reader for the GetAppPermissionsConfig structure.
type GetAppPermissionsConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppPermissionsConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppPermissionsConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppPermissionsConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppPermissionsConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppPermissionsConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppPermissionsConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppPermissionsConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}] GetAppPermissionsConfig", response, response.Code())
	}
}

// NewGetAppPermissionsConfigOK creates a GetAppPermissionsConfigOK with default headers values
func NewGetAppPermissionsConfigOK() *GetAppPermissionsConfigOK {
	return &GetAppPermissionsConfigOK{}
}

/*
GetAppPermissionsConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppPermissionsConfigOK struct {
	Payload *models.AppAppPermissionsConfig
}

// IsSuccess returns true when this get app permissions config o k response has a 2xx status code
func (o *GetAppPermissionsConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app permissions config o k response has a 3xx status code
func (o *GetAppPermissionsConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app permissions config o k response has a 4xx status code
func (o *GetAppPermissionsConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app permissions config o k response has a 5xx status code
func (o *GetAppPermissionsConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app permissions config o k response a status code equal to that given
func (o *GetAppPermissionsConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app permissions config o k response
func (o *GetAppPermissionsConfigOK) Code() int {
	return 200
}

func (o *GetAppPermissionsConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppPermissionsConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppPermissionsConfigOK) GetPayload() *models.AppAppPermissionsConfig {
	return o.Payload
}

func (o *GetAppPermissionsConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppPermissionsConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPermissionsConfigBadRequest creates a GetAppPermissionsConfigBadRequest with default headers values
func NewGetAppPermissionsConfigBadRequest() *GetAppPermissionsConfigBadRequest {
	return &GetAppPermissionsConfigBadRequest{}
}

/*
GetAppPermissionsConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppPermissionsConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app permissions config bad request response has a 2xx status code
func (o *GetAppPermissionsConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app permissions config bad request response has a 3xx status code
func (o *GetAppPermissionsConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app permissions config bad request response has a 4xx status code
func (o *GetAppPermissionsConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app permissions config bad request response has a 5xx status code
func (o *GetAppPermissionsConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app permissions config bad request response a status code equal to that given
func (o *GetAppPermissionsConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app permissions config bad request response
func (o *GetAppPermissionsConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppPermissionsConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppPermissionsConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppPermissionsConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPermissionsConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPermissionsConfigUnauthorized creates a GetAppPermissionsConfigUnauthorized with default headers values
func NewGetAppPermissionsConfigUnauthorized() *GetAppPermissionsConfigUnauthorized {
	return &GetAppPermissionsConfigUnauthorized{}
}

/*
GetAppPermissionsConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppPermissionsConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app permissions config unauthorized response has a 2xx status code
func (o *GetAppPermissionsConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app permissions config unauthorized response has a 3xx status code
func (o *GetAppPermissionsConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app permissions config unauthorized response has a 4xx status code
func (o *GetAppPermissionsConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app permissions config unauthorized response has a 5xx status code
func (o *GetAppPermissionsConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app permissions config unauthorized response a status code equal to that given
func (o *GetAppPermissionsConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app permissions config unauthorized response
func (o *GetAppPermissionsConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppPermissionsConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppPermissionsConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppPermissionsConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPermissionsConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPermissionsConfigForbidden creates a GetAppPermissionsConfigForbidden with default headers values
func NewGetAppPermissionsConfigForbidden() *GetAppPermissionsConfigForbidden {
	return &GetAppPermissionsConfigForbidden{}
}

/*
GetAppPermissionsConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppPermissionsConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app permissions config forbidden response has a 2xx status code
func (o *GetAppPermissionsConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app permissions config forbidden response has a 3xx status code
func (o *GetAppPermissionsConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app permissions config forbidden response has a 4xx status code
func (o *GetAppPermissionsConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app permissions config forbidden response has a 5xx status code
func (o *GetAppPermissionsConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app permissions config forbidden response a status code equal to that given
func (o *GetAppPermissionsConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app permissions config forbidden response
func (o *GetAppPermissionsConfigForbidden) Code() int {
	return 403
}

func (o *GetAppPermissionsConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppPermissionsConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppPermissionsConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPermissionsConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPermissionsConfigNotFound creates a GetAppPermissionsConfigNotFound with default headers values
func NewGetAppPermissionsConfigNotFound() *GetAppPermissionsConfigNotFound {
	return &GetAppPermissionsConfigNotFound{}
}

/*
GetAppPermissionsConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppPermissionsConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app permissions config not found response has a 2xx status code
func (o *GetAppPermissionsConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app permissions config not found response has a 3xx status code
func (o *GetAppPermissionsConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app permissions config not found response has a 4xx status code
func (o *GetAppPermissionsConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app permissions config not found response has a 5xx status code
func (o *GetAppPermissionsConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app permissions config not found response a status code equal to that given
func (o *GetAppPermissionsConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app permissions config not found response
func (o *GetAppPermissionsConfigNotFound) Code() int {
	return 404
}

func (o *GetAppPermissionsConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppPermissionsConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppPermissionsConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPermissionsConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPermissionsConfigInternalServerError creates a GetAppPermissionsConfigInternalServerError with default headers values
func NewGetAppPermissionsConfigInternalServerError() *GetAppPermissionsConfigInternalServerError {
	return &GetAppPermissionsConfigInternalServerError{}
}

/*
GetAppPermissionsConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppPermissionsConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app permissions config internal server error response has a 2xx status code
func (o *GetAppPermissionsConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app permissions config internal server error response has a 3xx status code
func (o *GetAppPermissionsConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app permissions config internal server error response has a 4xx status code
func (o *GetAppPermissionsConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app permissions config internal server error response has a 5xx status code
func (o *GetAppPermissionsConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app permissions config internal server error response a status code equal to that given
func (o *GetAppPermissionsConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app permissions config internal server error response
func (o *GetAppPermissionsConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppPermissionsConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppPermissionsConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/permissions-configs/{permissions_config_id}][%d] getAppPermissionsConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppPermissionsConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPermissionsConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
