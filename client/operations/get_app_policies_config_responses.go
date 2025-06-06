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

// GetAppPoliciesConfigReader is a Reader for the GetAppPoliciesConfig structure.
type GetAppPoliciesConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppPoliciesConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppPoliciesConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppPoliciesConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppPoliciesConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppPoliciesConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppPoliciesConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppPoliciesConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}] GetAppPoliciesConfig", response, response.Code())
	}
}

// NewGetAppPoliciesConfigOK creates a GetAppPoliciesConfigOK with default headers values
func NewGetAppPoliciesConfigOK() *GetAppPoliciesConfigOK {
	return &GetAppPoliciesConfigOK{}
}

/*
GetAppPoliciesConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetAppPoliciesConfigOK struct {
	Payload *models.AppAppPoliciesConfig
}

// IsSuccess returns true when this get app policies config o k response has a 2xx status code
func (o *GetAppPoliciesConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app policies config o k response has a 3xx status code
func (o *GetAppPoliciesConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app policies config o k response has a 4xx status code
func (o *GetAppPoliciesConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app policies config o k response has a 5xx status code
func (o *GetAppPoliciesConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app policies config o k response a status code equal to that given
func (o *GetAppPoliciesConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app policies config o k response
func (o *GetAppPoliciesConfigOK) Code() int {
	return 200
}

func (o *GetAppPoliciesConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppPoliciesConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigOK  %+v", 200, o.Payload)
}

func (o *GetAppPoliciesConfigOK) GetPayload() *models.AppAppPoliciesConfig {
	return o.Payload
}

func (o *GetAppPoliciesConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppPoliciesConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPoliciesConfigBadRequest creates a GetAppPoliciesConfigBadRequest with default headers values
func NewGetAppPoliciesConfigBadRequest() *GetAppPoliciesConfigBadRequest {
	return &GetAppPoliciesConfigBadRequest{}
}

/*
GetAppPoliciesConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppPoliciesConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app policies config bad request response has a 2xx status code
func (o *GetAppPoliciesConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app policies config bad request response has a 3xx status code
func (o *GetAppPoliciesConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app policies config bad request response has a 4xx status code
func (o *GetAppPoliciesConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app policies config bad request response has a 5xx status code
func (o *GetAppPoliciesConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app policies config bad request response a status code equal to that given
func (o *GetAppPoliciesConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app policies config bad request response
func (o *GetAppPoliciesConfigBadRequest) Code() int {
	return 400
}

func (o *GetAppPoliciesConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppPoliciesConfigBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppPoliciesConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPoliciesConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPoliciesConfigUnauthorized creates a GetAppPoliciesConfigUnauthorized with default headers values
func NewGetAppPoliciesConfigUnauthorized() *GetAppPoliciesConfigUnauthorized {
	return &GetAppPoliciesConfigUnauthorized{}
}

/*
GetAppPoliciesConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppPoliciesConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app policies config unauthorized response has a 2xx status code
func (o *GetAppPoliciesConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app policies config unauthorized response has a 3xx status code
func (o *GetAppPoliciesConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app policies config unauthorized response has a 4xx status code
func (o *GetAppPoliciesConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app policies config unauthorized response has a 5xx status code
func (o *GetAppPoliciesConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app policies config unauthorized response a status code equal to that given
func (o *GetAppPoliciesConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app policies config unauthorized response
func (o *GetAppPoliciesConfigUnauthorized) Code() int {
	return 401
}

func (o *GetAppPoliciesConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppPoliciesConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppPoliciesConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPoliciesConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPoliciesConfigForbidden creates a GetAppPoliciesConfigForbidden with default headers values
func NewGetAppPoliciesConfigForbidden() *GetAppPoliciesConfigForbidden {
	return &GetAppPoliciesConfigForbidden{}
}

/*
GetAppPoliciesConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppPoliciesConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app policies config forbidden response has a 2xx status code
func (o *GetAppPoliciesConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app policies config forbidden response has a 3xx status code
func (o *GetAppPoliciesConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app policies config forbidden response has a 4xx status code
func (o *GetAppPoliciesConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app policies config forbidden response has a 5xx status code
func (o *GetAppPoliciesConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app policies config forbidden response a status code equal to that given
func (o *GetAppPoliciesConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app policies config forbidden response
func (o *GetAppPoliciesConfigForbidden) Code() int {
	return 403
}

func (o *GetAppPoliciesConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppPoliciesConfigForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetAppPoliciesConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPoliciesConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPoliciesConfigNotFound creates a GetAppPoliciesConfigNotFound with default headers values
func NewGetAppPoliciesConfigNotFound() *GetAppPoliciesConfigNotFound {
	return &GetAppPoliciesConfigNotFound{}
}

/*
GetAppPoliciesConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppPoliciesConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app policies config not found response has a 2xx status code
func (o *GetAppPoliciesConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app policies config not found response has a 3xx status code
func (o *GetAppPoliciesConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app policies config not found response has a 4xx status code
func (o *GetAppPoliciesConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app policies config not found response has a 5xx status code
func (o *GetAppPoliciesConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app policies config not found response a status code equal to that given
func (o *GetAppPoliciesConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app policies config not found response
func (o *GetAppPoliciesConfigNotFound) Code() int {
	return 404
}

func (o *GetAppPoliciesConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppPoliciesConfigNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetAppPoliciesConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPoliciesConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppPoliciesConfigInternalServerError creates a GetAppPoliciesConfigInternalServerError with default headers values
func NewGetAppPoliciesConfigInternalServerError() *GetAppPoliciesConfigInternalServerError {
	return &GetAppPoliciesConfigInternalServerError{}
}

/*
GetAppPoliciesConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppPoliciesConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app policies config internal server error response has a 2xx status code
func (o *GetAppPoliciesConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app policies config internal server error response has a 3xx status code
func (o *GetAppPoliciesConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app policies config internal server error response has a 4xx status code
func (o *GetAppPoliciesConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app policies config internal server error response has a 5xx status code
func (o *GetAppPoliciesConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app policies config internal server error response a status code equal to that given
func (o *GetAppPoliciesConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app policies config internal server error response
func (o *GetAppPoliciesConfigInternalServerError) Code() int {
	return 500
}

func (o *GetAppPoliciesConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppPoliciesConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/policies-configs/{policies_config_id}][%d] getAppPoliciesConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppPoliciesConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppPoliciesConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
