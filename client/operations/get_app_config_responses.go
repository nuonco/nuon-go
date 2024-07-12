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

// GetAppConfigReader is a Reader for the GetAppConfig structure.
type GetAppConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/config/{app_config_id}] GetAppConfig", response, response.Code())
	}
}

// NewGetAppConfigOK creates a GetAppConfigOK with default headers values
func NewGetAppConfigOK() *GetAppConfigOK {
	return &GetAppConfigOK{}
}

/*
GetAppConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppConfigOK struct {
	Payload *models.AppAppConfig
}

// IsSuccess returns true when this get app config o k response has a 2xx status code
func (o *GetAppConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app config o k response has a 3xx status code
func (o *GetAppConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app config o k response has a 4xx status code
func (o *GetAppConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app config o k response has a 5xx status code
func (o *GetAppConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app config o k response a status code equal to that given
func (o *GetAppConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app config o k response
func (o *GetAppConfigOK) Code() int {
	return 200
}

func (o *GetAppConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppConfigOK) GetPayload() *models.AppAppConfig {
	return o.Payload
}

func (o *GetAppConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppConfigBadRequest creates a GetAppConfigBadRequest with default headers values
func NewGetAppConfigBadRequest() *GetAppConfigBadRequest {
	return &GetAppConfigBadRequest{}
}

/*
GetAppConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app config bad request response has a 2xx status code
func (o *GetAppConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app config bad request response has a 3xx status code
func (o *GetAppConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app config bad request response has a 4xx status code
func (o *GetAppConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app config bad request response has a 5xx status code
func (o *GetAppConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app config bad request response a status code equal to that given
func (o *GetAppConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app config bad request response
func (o *GetAppConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppConfigUnauthorized creates a GetAppConfigUnauthorized with default headers values
func NewGetAppConfigUnauthorized() *GetAppConfigUnauthorized {
	return &GetAppConfigUnauthorized{}
}

/*
GetAppConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app config unauthorized response has a 2xx status code
func (o *GetAppConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app config unauthorized response has a 3xx status code
func (o *GetAppConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app config unauthorized response has a 4xx status code
func (o *GetAppConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app config unauthorized response has a 5xx status code
func (o *GetAppConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app config unauthorized response a status code equal to that given
func (o *GetAppConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app config unauthorized response
func (o *GetAppConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppConfigForbidden creates a GetAppConfigForbidden with default headers values
func NewGetAppConfigForbidden() *GetAppConfigForbidden {
	return &GetAppConfigForbidden{}
}

/*
GetAppConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app config forbidden response has a 2xx status code
func (o *GetAppConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app config forbidden response has a 3xx status code
func (o *GetAppConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app config forbidden response has a 4xx status code
func (o *GetAppConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app config forbidden response has a 5xx status code
func (o *GetAppConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app config forbidden response a status code equal to that given
func (o *GetAppConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app config forbidden response
func (o *GetAppConfigForbidden) Code() int {
	return 403
}

func (o *GetAppConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppConfigNotFound creates a GetAppConfigNotFound with default headers values
func NewGetAppConfigNotFound() *GetAppConfigNotFound {
	return &GetAppConfigNotFound{}
}

/*
GetAppConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app config not found response has a 2xx status code
func (o *GetAppConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app config not found response has a 3xx status code
func (o *GetAppConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app config not found response has a 4xx status code
func (o *GetAppConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app config not found response has a 5xx status code
func (o *GetAppConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app config not found response a status code equal to that given
func (o *GetAppConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app config not found response
func (o *GetAppConfigNotFound) Code() int {
	return 404
}

func (o *GetAppConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppConfigInternalServerError creates a GetAppConfigInternalServerError with default headers values
func NewGetAppConfigInternalServerError() *GetAppConfigInternalServerError {
	return &GetAppConfigInternalServerError{}
}

/*
GetAppConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app config internal server error response has a 2xx status code
func (o *GetAppConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app config internal server error response has a 3xx status code
func (o *GetAppConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app config internal server error response has a 4xx status code
func (o *GetAppConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app config internal server error response has a 5xx status code
func (o *GetAppConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app config internal server error response a status code equal to that given
func (o *GetAppConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app config internal server error response
func (o *GetAppConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/config/{app_config_id}][%d] getAppConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
