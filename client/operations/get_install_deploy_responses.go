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

// GetInstallDeployReader is a Reader for the GetInstallDeploy structure.
type GetInstallDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallDeployBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallDeployUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallDeployForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallDeployNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallDeployInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/deploys/{deploy_id}] GetInstallDeploy", response, response.Code())
	}
}

// NewGetInstallDeployOK creates a GetInstallDeployOK with default headers values
func NewGetInstallDeployOK() *GetInstallDeployOK {
	return &GetInstallDeployOK{}
}

/*
GetInstallDeployOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallDeployOK struct {
	Payload *models.AppInstallDeploy
}

// IsSuccess returns true when this get install deploy o k response has a 2xx status code
func (o *GetInstallDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install deploy o k response has a 3xx status code
func (o *GetInstallDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploy o k response has a 4xx status code
func (o *GetInstallDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install deploy o k response has a 5xx status code
func (o *GetInstallDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploy o k response a status code equal to that given
func (o *GetInstallDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install deploy o k response
func (o *GetInstallDeployOK) Code() int {
	return 200
}

func (o *GetInstallDeployOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployOK %s", 200, payload)
}

func (o *GetInstallDeployOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployOK %s", 200, payload)
}

func (o *GetInstallDeployOK) GetPayload() *models.AppInstallDeploy {
	return o.Payload
}

func (o *GetInstallDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallDeploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeployBadRequest creates a GetInstallDeployBadRequest with default headers values
func NewGetInstallDeployBadRequest() *GetInstallDeployBadRequest {
	return &GetInstallDeployBadRequest{}
}

/*
GetInstallDeployBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallDeployBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploy bad request response has a 2xx status code
func (o *GetInstallDeployBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploy bad request response has a 3xx status code
func (o *GetInstallDeployBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploy bad request response has a 4xx status code
func (o *GetInstallDeployBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploy bad request response has a 5xx status code
func (o *GetInstallDeployBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploy bad request response a status code equal to that given
func (o *GetInstallDeployBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install deploy bad request response
func (o *GetInstallDeployBadRequest) Code() int {
	return 400
}

func (o *GetInstallDeployBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployBadRequest %s", 400, payload)
}

func (o *GetInstallDeployBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployBadRequest %s", 400, payload)
}

func (o *GetInstallDeployBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeployBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeployUnauthorized creates a GetInstallDeployUnauthorized with default headers values
func NewGetInstallDeployUnauthorized() *GetInstallDeployUnauthorized {
	return &GetInstallDeployUnauthorized{}
}

/*
GetInstallDeployUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallDeployUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploy unauthorized response has a 2xx status code
func (o *GetInstallDeployUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploy unauthorized response has a 3xx status code
func (o *GetInstallDeployUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploy unauthorized response has a 4xx status code
func (o *GetInstallDeployUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploy unauthorized response has a 5xx status code
func (o *GetInstallDeployUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploy unauthorized response a status code equal to that given
func (o *GetInstallDeployUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install deploy unauthorized response
func (o *GetInstallDeployUnauthorized) Code() int {
	return 401
}

func (o *GetInstallDeployUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployUnauthorized %s", 401, payload)
}

func (o *GetInstallDeployUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployUnauthorized %s", 401, payload)
}

func (o *GetInstallDeployUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeployUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeployForbidden creates a GetInstallDeployForbidden with default headers values
func NewGetInstallDeployForbidden() *GetInstallDeployForbidden {
	return &GetInstallDeployForbidden{}
}

/*
GetInstallDeployForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallDeployForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploy forbidden response has a 2xx status code
func (o *GetInstallDeployForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploy forbidden response has a 3xx status code
func (o *GetInstallDeployForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploy forbidden response has a 4xx status code
func (o *GetInstallDeployForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploy forbidden response has a 5xx status code
func (o *GetInstallDeployForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploy forbidden response a status code equal to that given
func (o *GetInstallDeployForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install deploy forbidden response
func (o *GetInstallDeployForbidden) Code() int {
	return 403
}

func (o *GetInstallDeployForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployForbidden %s", 403, payload)
}

func (o *GetInstallDeployForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployForbidden %s", 403, payload)
}

func (o *GetInstallDeployForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeployForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeployNotFound creates a GetInstallDeployNotFound with default headers values
func NewGetInstallDeployNotFound() *GetInstallDeployNotFound {
	return &GetInstallDeployNotFound{}
}

/*
GetInstallDeployNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallDeployNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploy not found response has a 2xx status code
func (o *GetInstallDeployNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploy not found response has a 3xx status code
func (o *GetInstallDeployNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploy not found response has a 4xx status code
func (o *GetInstallDeployNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install deploy not found response has a 5xx status code
func (o *GetInstallDeployNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install deploy not found response a status code equal to that given
func (o *GetInstallDeployNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install deploy not found response
func (o *GetInstallDeployNotFound) Code() int {
	return 404
}

func (o *GetInstallDeployNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployNotFound %s", 404, payload)
}

func (o *GetInstallDeployNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployNotFound %s", 404, payload)
}

func (o *GetInstallDeployNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeployNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallDeployInternalServerError creates a GetInstallDeployInternalServerError with default headers values
func NewGetInstallDeployInternalServerError() *GetInstallDeployInternalServerError {
	return &GetInstallDeployInternalServerError{}
}

/*
GetInstallDeployInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallDeployInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install deploy internal server error response has a 2xx status code
func (o *GetInstallDeployInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install deploy internal server error response has a 3xx status code
func (o *GetInstallDeployInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install deploy internal server error response has a 4xx status code
func (o *GetInstallDeployInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install deploy internal server error response has a 5xx status code
func (o *GetInstallDeployInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install deploy internal server error response a status code equal to that given
func (o *GetInstallDeployInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install deploy internal server error response
func (o *GetInstallDeployInternalServerError) Code() int {
	return 500
}

func (o *GetInstallDeployInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployInternalServerError %s", 500, payload)
}

func (o *GetInstallDeployInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}][%d] getInstallDeployInternalServerError %s", 500, payload)
}

func (o *GetInstallDeployInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallDeployInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
