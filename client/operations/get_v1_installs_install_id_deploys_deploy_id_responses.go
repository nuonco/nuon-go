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

// GetV1InstallsInstallIDDeploysDeployIDReader is a Reader for the GetV1InstallsInstallIDDeploysDeployID structure.
type GetV1InstallsInstallIDDeploysDeployIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallsInstallIDDeploysDeployIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallsInstallIDDeploysDeployIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1InstallsInstallIDDeploysDeployIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1InstallsInstallIDDeploysDeployIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1InstallsInstallIDDeploysDeployIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1InstallsInstallIDDeploysDeployIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1InstallsInstallIDDeploysDeployIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/deploys/{deploy_id}] GetV1InstallsInstallIDDeploysDeployID", response, response.Code())
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDOK creates a GetV1InstallsInstallIDDeploysDeployIDOK with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDOK() *GetV1InstallsInstallIDDeploysDeployIDOK {
	return &GetV1InstallsInstallIDDeploysDeployIDOK{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallsInstallIDDeploysDeployIDOK struct {
	Payload *models.AppInstallDeploy
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id o k response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id o k response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id o k response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id o k response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id o k response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id o k response
func (o *GetV1InstallsInstallIDDeploysDeployIDOK) Code() int {
	return 200
}

func (o *GetV1InstallsInstallIDDeploysDeployIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDOK) GetPayload() *models.AppInstallDeploy {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallDeploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDDeploysDeployIDBadRequest creates a GetV1InstallsInstallIDDeploysDeployIDBadRequest with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDBadRequest() *GetV1InstallsInstallIDDeploysDeployIDBadRequest {
	return &GetV1InstallsInstallIDDeploysDeployIDBadRequest{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1InstallsInstallIDDeploysDeployIDBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id bad request response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id bad request response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id bad request response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id bad request response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id bad request response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id bad request response
func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) Code() int {
	return 400
}

func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDDeploysDeployIDUnauthorized creates a GetV1InstallsInstallIDDeploysDeployIDUnauthorized with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDUnauthorized() *GetV1InstallsInstallIDDeploysDeployIDUnauthorized {
	return &GetV1InstallsInstallIDDeploysDeployIDUnauthorized{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1InstallsInstallIDDeploysDeployIDUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id unauthorized response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id unauthorized response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id unauthorized response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id unauthorized response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id unauthorized response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id unauthorized response
func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) Code() int {
	return 401
}

func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDDeploysDeployIDForbidden creates a GetV1InstallsInstallIDDeploysDeployIDForbidden with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDForbidden() *GetV1InstallsInstallIDDeploysDeployIDForbidden {
	return &GetV1InstallsInstallIDDeploysDeployIDForbidden{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1InstallsInstallIDDeploysDeployIDForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id forbidden response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id forbidden response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id forbidden response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id forbidden response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id forbidden response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id forbidden response
func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) Code() int {
	return 403
}

func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDDeploysDeployIDNotFound creates a GetV1InstallsInstallIDDeploysDeployIDNotFound with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDNotFound() *GetV1InstallsInstallIDDeploysDeployIDNotFound {
	return &GetV1InstallsInstallIDDeploysDeployIDNotFound{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1InstallsInstallIDDeploysDeployIDNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id not found response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id not found response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id not found response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id not found response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id not found response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id not found response
func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) Code() int {
	return 404
}

func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDDeploysDeployIDInternalServerError creates a GetV1InstallsInstallIDDeploysDeployIDInternalServerError with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDInternalServerError() *GetV1InstallsInstallIDDeploysDeployIDInternalServerError {
	return &GetV1InstallsInstallIDDeploysDeployIDInternalServerError{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1InstallsInstallIDDeploysDeployIDInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id internal server error response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id internal server error response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id internal server error response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id internal server error response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id internal server error response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id internal server error response
func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) Code() int {
	return 500
}

func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getV1InstallsInstallIdDeploysDeployIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
