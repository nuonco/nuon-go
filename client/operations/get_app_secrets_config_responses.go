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

// GetAppSecretsConfigReader is a Reader for the GetAppSecretsConfig structure.
type GetAppSecretsConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppSecretsConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppSecretsConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppSecretsConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppSecretsConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppSecretsConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppSecretsConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppSecretsConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}] GetAppSecretsConfig", response, response.Code())
	}
}

// NewGetAppSecretsConfigOK creates a GetAppSecretsConfigOK with default headers values
func NewGetAppSecretsConfigOK() *GetAppSecretsConfigOK {
	return &GetAppSecretsConfigOK{}
}

/*
GetAppSecretsConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppSecretsConfigOK struct {
	Payload *models.AppAppSecretsConfig
}

// IsSuccess returns true when this get app secrets config o k response has a 2xx status code
func (o *GetAppSecretsConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app secrets config o k response has a 3xx status code
func (o *GetAppSecretsConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secrets config o k response has a 4xx status code
func (o *GetAppSecretsConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app secrets config o k response has a 5xx status code
func (o *GetAppSecretsConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secrets config o k response a status code equal to that given
func (o *GetAppSecretsConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app secrets config o k response
func (o *GetAppSecretsConfigOK) Code() int {
	return 200
}

func (o *GetAppSecretsConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppSecretsConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppSecretsConfigOK) GetPayload() *models.AppAppSecretsConfig {
	return o.Payload
}

func (o *GetAppSecretsConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppSecretsConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretsConfigBadRequest creates a GetAppSecretsConfigBadRequest with default headers values
func NewGetAppSecretsConfigBadRequest() *GetAppSecretsConfigBadRequest {
	return &GetAppSecretsConfigBadRequest{}
}

/*
GetAppSecretsConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppSecretsConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app secrets config bad request response has a 2xx status code
func (o *GetAppSecretsConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app secrets config bad request response has a 3xx status code
func (o *GetAppSecretsConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secrets config bad request response has a 4xx status code
func (o *GetAppSecretsConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app secrets config bad request response has a 5xx status code
func (o *GetAppSecretsConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secrets config bad request response a status code equal to that given
func (o *GetAppSecretsConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app secrets config bad request response
func (o *GetAppSecretsConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppSecretsConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppSecretsConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppSecretsConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSecretsConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretsConfigUnauthorized creates a GetAppSecretsConfigUnauthorized with default headers values
func NewGetAppSecretsConfigUnauthorized() *GetAppSecretsConfigUnauthorized {
	return &GetAppSecretsConfigUnauthorized{}
}

/*
GetAppSecretsConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppSecretsConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app secrets config unauthorized response has a 2xx status code
func (o *GetAppSecretsConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app secrets config unauthorized response has a 3xx status code
func (o *GetAppSecretsConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secrets config unauthorized response has a 4xx status code
func (o *GetAppSecretsConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app secrets config unauthorized response has a 5xx status code
func (o *GetAppSecretsConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secrets config unauthorized response a status code equal to that given
func (o *GetAppSecretsConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app secrets config unauthorized response
func (o *GetAppSecretsConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppSecretsConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppSecretsConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppSecretsConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSecretsConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretsConfigForbidden creates a GetAppSecretsConfigForbidden with default headers values
func NewGetAppSecretsConfigForbidden() *GetAppSecretsConfigForbidden {
	return &GetAppSecretsConfigForbidden{}
}

/*
GetAppSecretsConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppSecretsConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app secrets config forbidden response has a 2xx status code
func (o *GetAppSecretsConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app secrets config forbidden response has a 3xx status code
func (o *GetAppSecretsConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secrets config forbidden response has a 4xx status code
func (o *GetAppSecretsConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app secrets config forbidden response has a 5xx status code
func (o *GetAppSecretsConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secrets config forbidden response a status code equal to that given
func (o *GetAppSecretsConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app secrets config forbidden response
func (o *GetAppSecretsConfigForbidden) Code() int {
	return 403
}

func (o *GetAppSecretsConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppSecretsConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppSecretsConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSecretsConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretsConfigNotFound creates a GetAppSecretsConfigNotFound with default headers values
func NewGetAppSecretsConfigNotFound() *GetAppSecretsConfigNotFound {
	return &GetAppSecretsConfigNotFound{}
}

/*
GetAppSecretsConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppSecretsConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app secrets config not found response has a 2xx status code
func (o *GetAppSecretsConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app secrets config not found response has a 3xx status code
func (o *GetAppSecretsConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secrets config not found response has a 4xx status code
func (o *GetAppSecretsConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app secrets config not found response has a 5xx status code
func (o *GetAppSecretsConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app secrets config not found response a status code equal to that given
func (o *GetAppSecretsConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app secrets config not found response
func (o *GetAppSecretsConfigNotFound) Code() int {
	return 404
}

func (o *GetAppSecretsConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppSecretsConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppSecretsConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSecretsConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSecretsConfigInternalServerError creates a GetAppSecretsConfigInternalServerError with default headers values
func NewGetAppSecretsConfigInternalServerError() *GetAppSecretsConfigInternalServerError {
	return &GetAppSecretsConfigInternalServerError{}
}

/*
GetAppSecretsConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppSecretsConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app secrets config internal server error response has a 2xx status code
func (o *GetAppSecretsConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app secrets config internal server error response has a 3xx status code
func (o *GetAppSecretsConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app secrets config internal server error response has a 4xx status code
func (o *GetAppSecretsConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app secrets config internal server error response has a 5xx status code
func (o *GetAppSecretsConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app secrets config internal server error response a status code equal to that given
func (o *GetAppSecretsConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app secrets config internal server error response
func (o *GetAppSecretsConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppSecretsConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppSecretsConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/secrets-configs/{app_secrets_config_id}][%d] getAppSecretsConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppSecretsConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppSecretsConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
