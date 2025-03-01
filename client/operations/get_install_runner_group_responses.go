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

// GetInstallRunnerGroupReader is a Reader for the GetInstallRunnerGroup structure.
type GetInstallRunnerGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallRunnerGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallRunnerGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallRunnerGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallRunnerGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallRunnerGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallRunnerGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallRunnerGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/runner-group] GetInstallRunnerGroup", response, response.Code())
	}
}

// NewGetInstallRunnerGroupOK creates a GetInstallRunnerGroupOK with default headers values
func NewGetInstallRunnerGroupOK() *GetInstallRunnerGroupOK {
	return &GetInstallRunnerGroupOK{}
}

/*
GetInstallRunnerGroupOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallRunnerGroupOK struct {
	Payload *models.AppRunnerGroup
}

// IsSuccess returns true when this get install runner group o k response has a 2xx status code
func (o *GetInstallRunnerGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install runner group o k response has a 3xx status code
func (o *GetInstallRunnerGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install runner group o k response has a 4xx status code
func (o *GetInstallRunnerGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install runner group o k response has a 5xx status code
func (o *GetInstallRunnerGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install runner group o k response a status code equal to that given
func (o *GetInstallRunnerGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install runner group o k response
func (o *GetInstallRunnerGroupOK) Code() int {
	return 200
}

func (o *GetInstallRunnerGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupOK %s", 200, payload)
}

func (o *GetInstallRunnerGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupOK %s", 200, payload)
}

func (o *GetInstallRunnerGroupOK) GetPayload() *models.AppRunnerGroup {
	return o.Payload
}

func (o *GetInstallRunnerGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRunnerGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallRunnerGroupBadRequest creates a GetInstallRunnerGroupBadRequest with default headers values
func NewGetInstallRunnerGroupBadRequest() *GetInstallRunnerGroupBadRequest {
	return &GetInstallRunnerGroupBadRequest{}
}

/*
GetInstallRunnerGroupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallRunnerGroupBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install runner group bad request response has a 2xx status code
func (o *GetInstallRunnerGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install runner group bad request response has a 3xx status code
func (o *GetInstallRunnerGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install runner group bad request response has a 4xx status code
func (o *GetInstallRunnerGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install runner group bad request response has a 5xx status code
func (o *GetInstallRunnerGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install runner group bad request response a status code equal to that given
func (o *GetInstallRunnerGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install runner group bad request response
func (o *GetInstallRunnerGroupBadRequest) Code() int {
	return 400
}

func (o *GetInstallRunnerGroupBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupBadRequest %s", 400, payload)
}

func (o *GetInstallRunnerGroupBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupBadRequest %s", 400, payload)
}

func (o *GetInstallRunnerGroupBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallRunnerGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallRunnerGroupUnauthorized creates a GetInstallRunnerGroupUnauthorized with default headers values
func NewGetInstallRunnerGroupUnauthorized() *GetInstallRunnerGroupUnauthorized {
	return &GetInstallRunnerGroupUnauthorized{}
}

/*
GetInstallRunnerGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallRunnerGroupUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install runner group unauthorized response has a 2xx status code
func (o *GetInstallRunnerGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install runner group unauthorized response has a 3xx status code
func (o *GetInstallRunnerGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install runner group unauthorized response has a 4xx status code
func (o *GetInstallRunnerGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install runner group unauthorized response has a 5xx status code
func (o *GetInstallRunnerGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install runner group unauthorized response a status code equal to that given
func (o *GetInstallRunnerGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install runner group unauthorized response
func (o *GetInstallRunnerGroupUnauthorized) Code() int {
	return 401
}

func (o *GetInstallRunnerGroupUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupUnauthorized %s", 401, payload)
}

func (o *GetInstallRunnerGroupUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupUnauthorized %s", 401, payload)
}

func (o *GetInstallRunnerGroupUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallRunnerGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallRunnerGroupForbidden creates a GetInstallRunnerGroupForbidden with default headers values
func NewGetInstallRunnerGroupForbidden() *GetInstallRunnerGroupForbidden {
	return &GetInstallRunnerGroupForbidden{}
}

/*
GetInstallRunnerGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallRunnerGroupForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install runner group forbidden response has a 2xx status code
func (o *GetInstallRunnerGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install runner group forbidden response has a 3xx status code
func (o *GetInstallRunnerGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install runner group forbidden response has a 4xx status code
func (o *GetInstallRunnerGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install runner group forbidden response has a 5xx status code
func (o *GetInstallRunnerGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install runner group forbidden response a status code equal to that given
func (o *GetInstallRunnerGroupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install runner group forbidden response
func (o *GetInstallRunnerGroupForbidden) Code() int {
	return 403
}

func (o *GetInstallRunnerGroupForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupForbidden %s", 403, payload)
}

func (o *GetInstallRunnerGroupForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupForbidden %s", 403, payload)
}

func (o *GetInstallRunnerGroupForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallRunnerGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallRunnerGroupNotFound creates a GetInstallRunnerGroupNotFound with default headers values
func NewGetInstallRunnerGroupNotFound() *GetInstallRunnerGroupNotFound {
	return &GetInstallRunnerGroupNotFound{}
}

/*
GetInstallRunnerGroupNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallRunnerGroupNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install runner group not found response has a 2xx status code
func (o *GetInstallRunnerGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install runner group not found response has a 3xx status code
func (o *GetInstallRunnerGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install runner group not found response has a 4xx status code
func (o *GetInstallRunnerGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install runner group not found response has a 5xx status code
func (o *GetInstallRunnerGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install runner group not found response a status code equal to that given
func (o *GetInstallRunnerGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install runner group not found response
func (o *GetInstallRunnerGroupNotFound) Code() int {
	return 404
}

func (o *GetInstallRunnerGroupNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupNotFound %s", 404, payload)
}

func (o *GetInstallRunnerGroupNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupNotFound %s", 404, payload)
}

func (o *GetInstallRunnerGroupNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallRunnerGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallRunnerGroupInternalServerError creates a GetInstallRunnerGroupInternalServerError with default headers values
func NewGetInstallRunnerGroupInternalServerError() *GetInstallRunnerGroupInternalServerError {
	return &GetInstallRunnerGroupInternalServerError{}
}

/*
GetInstallRunnerGroupInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallRunnerGroupInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install runner group internal server error response has a 2xx status code
func (o *GetInstallRunnerGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install runner group internal server error response has a 3xx status code
func (o *GetInstallRunnerGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install runner group internal server error response has a 4xx status code
func (o *GetInstallRunnerGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install runner group internal server error response has a 5xx status code
func (o *GetInstallRunnerGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install runner group internal server error response a status code equal to that given
func (o *GetInstallRunnerGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install runner group internal server error response
func (o *GetInstallRunnerGroupInternalServerError) Code() int {
	return 500
}

func (o *GetInstallRunnerGroupInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupInternalServerError %s", 500, payload)
}

func (o *GetInstallRunnerGroupInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/runner-group][%d] getInstallRunnerGroupInternalServerError %s", 500, payload)
}

func (o *GetInstallRunnerGroupInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallRunnerGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
