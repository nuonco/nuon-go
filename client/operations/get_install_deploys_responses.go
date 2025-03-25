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

// GetInstallDeploysReader is a Reader for the GetInstallDeploys structure.
type GetInstallDeploysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallDeploysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallDeploysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallDeploysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallDeploysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallDeploysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallDeploysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallDeploysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/deploys] GetInstallDeploys", response, response.Code())
	}
}

// NewGetInstallDeploysOK creates a GetInstallDeploysOK with default headers values
func NewGetInstallDeploysOK() *GetInstallDeploysOK {
	return &GetInstallDeploysOK{}
}

/*
GetInstallDeploysOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallDeploysOK struct {
	Payload []*models.AppInstallDeploy
}

// IsSuccess returns true when this get install deploys o k response has a 2xx status code
func (o *GetInstallDeploysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install deploys o k response has a 3xx status code
func (o *GetInstallDeploysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploys o k response has a 4xx status code
func (o *GetInstallDeploysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install deploys o k response has a 5xx status code
func (o *GetInstallDeploysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploys o k response a status code equal to that given
func (o *GetInstallDeploysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install deploys o k response
func (o *GetInstallDeploysOK) Code() int {
	return 200
}

func (o *GetInstallDeploysOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysOK  %+v", 200, o.Payload)
}

func (o *GetInstallDeploysOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysOK  %+v", 200, o.Payload)
}

func (o *GetInstallDeploysOK) GetPayload() []*models.AppInstallDeploy {
	return o.Payload
}

func (o *GetInstallDeploysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeploysBadRequest creates a GetInstallDeploysBadRequest with default headers values
func NewGetInstallDeploysBadRequest() *GetInstallDeploysBadRequest {
	return &GetInstallDeploysBadRequest{}
}

/*
GetInstallDeploysBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallDeploysBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploys bad request response has a 2xx status code
func (o *GetInstallDeploysBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploys bad request response has a 3xx status code
func (o *GetInstallDeploysBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploys bad request response has a 4xx status code
func (o *GetInstallDeploysBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploys bad request response has a 5xx status code
func (o *GetInstallDeploysBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploys bad request response a status code equal to that given
func (o *GetInstallDeploysBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install deploys bad request response
func (o *GetInstallDeploysBadRequest) Code() int {
	return 400
}

func (o *GetInstallDeploysBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallDeploysBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallDeploysBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeploysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeploysUnauthorized creates a GetInstallDeploysUnauthorized with default headers values
func NewGetInstallDeploysUnauthorized() *GetInstallDeploysUnauthorized {
	return &GetInstallDeploysUnauthorized{}
}

/*
GetInstallDeploysUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallDeploysUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploys unauthorized response has a 2xx status code
func (o *GetInstallDeploysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploys unauthorized response has a 3xx status code
func (o *GetInstallDeploysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploys unauthorized response has a 4xx status code
func (o *GetInstallDeploysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploys unauthorized response has a 5xx status code
func (o *GetInstallDeploysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploys unauthorized response a status code equal to that given
func (o *GetInstallDeploysUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install deploys unauthorized response
func (o *GetInstallDeploysUnauthorized) Code() int {
	return 401
}

func (o *GetInstallDeploysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallDeploysUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallDeploysUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeploysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeploysForbidden creates a GetInstallDeploysForbidden with default headers values
func NewGetInstallDeploysForbidden() *GetInstallDeploysForbidden {
	return &GetInstallDeploysForbidden{}
}

/*
GetInstallDeploysForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallDeploysForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploys forbidden response has a 2xx status code
func (o *GetInstallDeploysForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploys forbidden response has a 3xx status code
func (o *GetInstallDeploysForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploys forbidden response has a 4xx status code
func (o *GetInstallDeploysForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploys forbidden response has a 5xx status code
func (o *GetInstallDeploysForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploys forbidden response a status code equal to that given
func (o *GetInstallDeploysForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install deploys forbidden response
func (o *GetInstallDeploysForbidden) Code() int {
	return 403
}

func (o *GetInstallDeploysForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallDeploysForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallDeploysForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeploysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeploysNotFound creates a GetInstallDeploysNotFound with default headers values
func NewGetInstallDeploysNotFound() *GetInstallDeploysNotFound {
	return &GetInstallDeploysNotFound{}
}

/*
GetInstallDeploysNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallDeploysNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploys not found response has a 2xx status code
func (o *GetInstallDeploysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploys not found response has a 3xx status code
func (o *GetInstallDeploysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploys not found response has a 4xx status code
func (o *GetInstallDeploysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploys not found response has a 5xx status code
func (o *GetInstallDeploysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploys not found response a status code equal to that given
func (o *GetInstallDeploysNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install deploys not found response
func (o *GetInstallDeploysNotFound) Code() int {
	return 404
}

func (o *GetInstallDeploysNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallDeploysNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallDeploysNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeploysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeploysInternalServerError creates a GetInstallDeploysInternalServerError with default headers values
func NewGetInstallDeploysInternalServerError() *GetInstallDeploysInternalServerError {
	return &GetInstallDeploysInternalServerError{}
}

/*
GetInstallDeploysInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallDeploysInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploys internal server error response has a 2xx status code
func (o *GetInstallDeploysInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploys internal server error response has a 3xx status code
func (o *GetInstallDeploysInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploys internal server error response has a 4xx status code
func (o *GetInstallDeploysInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install deploys internal server error response has a 5xx status code
func (o *GetInstallDeploysInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install deploys internal server error response a status code equal to that given
func (o *GetInstallDeploysInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install deploys internal server error response
func (o *GetInstallDeploysInternalServerError) Code() int {
	return 500
}

func (o *GetInstallDeploysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallDeploysInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys][%d] getInstallDeploysInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallDeploysInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeploysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
