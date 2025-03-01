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

// GetInstallComponentLatestDeployReader is a Reader for the GetInstallComponentLatestDeploy structure.
type GetInstallComponentLatestDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallComponentLatestDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallComponentLatestDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallComponentLatestDeployBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallComponentLatestDeployUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallComponentLatestDeployForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallComponentLatestDeployNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallComponentLatestDeployInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest] GetInstallComponentLatestDeploy", response, response.Code())
	}
}

// NewGetInstallComponentLatestDeployOK creates a GetInstallComponentLatestDeployOK with default headers values
func NewGetInstallComponentLatestDeployOK() *GetInstallComponentLatestDeployOK {
	return &GetInstallComponentLatestDeployOK{}
}

/*
GetInstallComponentLatestDeployOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallComponentLatestDeployOK struct {
	Payload *models.AppInstallDeploy
}

// IsSuccess returns true when this get install component latest deploy o k response has a 2xx status code
func (o *GetInstallComponentLatestDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install component latest deploy o k response has a 3xx status code
func (o *GetInstallComponentLatestDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component latest deploy o k response has a 4xx status code
func (o *GetInstallComponentLatestDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install component latest deploy o k response has a 5xx status code
func (o *GetInstallComponentLatestDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component latest deploy o k response a status code equal to that given
func (o *GetInstallComponentLatestDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install component latest deploy o k response
func (o *GetInstallComponentLatestDeployOK) Code() int {
	return 200
}

func (o *GetInstallComponentLatestDeployOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployOK %s", 200, payload)
}

func (o *GetInstallComponentLatestDeployOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployOK %s", 200, payload)
}

func (o *GetInstallComponentLatestDeployOK) GetPayload() *models.AppInstallDeploy {
	return o.Payload
}

func (o *GetInstallComponentLatestDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallDeploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentLatestDeployBadRequest creates a GetInstallComponentLatestDeployBadRequest with default headers values
func NewGetInstallComponentLatestDeployBadRequest() *GetInstallComponentLatestDeployBadRequest {
	return &GetInstallComponentLatestDeployBadRequest{}
}

/*
GetInstallComponentLatestDeployBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallComponentLatestDeployBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component latest deploy bad request response has a 2xx status code
func (o *GetInstallComponentLatestDeployBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component latest deploy bad request response has a 3xx status code
func (o *GetInstallComponentLatestDeployBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component latest deploy bad request response has a 4xx status code
func (o *GetInstallComponentLatestDeployBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component latest deploy bad request response has a 5xx status code
func (o *GetInstallComponentLatestDeployBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component latest deploy bad request response a status code equal to that given
func (o *GetInstallComponentLatestDeployBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install component latest deploy bad request response
func (o *GetInstallComponentLatestDeployBadRequest) Code() int {
	return 400
}

func (o *GetInstallComponentLatestDeployBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployBadRequest %s", 400, payload)
}

func (o *GetInstallComponentLatestDeployBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployBadRequest %s", 400, payload)
}

func (o *GetInstallComponentLatestDeployBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentLatestDeployBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentLatestDeployUnauthorized creates a GetInstallComponentLatestDeployUnauthorized with default headers values
func NewGetInstallComponentLatestDeployUnauthorized() *GetInstallComponentLatestDeployUnauthorized {
	return &GetInstallComponentLatestDeployUnauthorized{}
}

/*
GetInstallComponentLatestDeployUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallComponentLatestDeployUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component latest deploy unauthorized response has a 2xx status code
func (o *GetInstallComponentLatestDeployUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component latest deploy unauthorized response has a 3xx status code
func (o *GetInstallComponentLatestDeployUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component latest deploy unauthorized response has a 4xx status code
func (o *GetInstallComponentLatestDeployUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component latest deploy unauthorized response has a 5xx status code
func (o *GetInstallComponentLatestDeployUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component latest deploy unauthorized response a status code equal to that given
func (o *GetInstallComponentLatestDeployUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install component latest deploy unauthorized response
func (o *GetInstallComponentLatestDeployUnauthorized) Code() int {
	return 401
}

func (o *GetInstallComponentLatestDeployUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployUnauthorized %s", 401, payload)
}

func (o *GetInstallComponentLatestDeployUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployUnauthorized %s", 401, payload)
}

func (o *GetInstallComponentLatestDeployUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentLatestDeployUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentLatestDeployForbidden creates a GetInstallComponentLatestDeployForbidden with default headers values
func NewGetInstallComponentLatestDeployForbidden() *GetInstallComponentLatestDeployForbidden {
	return &GetInstallComponentLatestDeployForbidden{}
}

/*
GetInstallComponentLatestDeployForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallComponentLatestDeployForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component latest deploy forbidden response has a 2xx status code
func (o *GetInstallComponentLatestDeployForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component latest deploy forbidden response has a 3xx status code
func (o *GetInstallComponentLatestDeployForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component latest deploy forbidden response has a 4xx status code
func (o *GetInstallComponentLatestDeployForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component latest deploy forbidden response has a 5xx status code
func (o *GetInstallComponentLatestDeployForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component latest deploy forbidden response a status code equal to that given
func (o *GetInstallComponentLatestDeployForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install component latest deploy forbidden response
func (o *GetInstallComponentLatestDeployForbidden) Code() int {
	return 403
}

func (o *GetInstallComponentLatestDeployForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployForbidden %s", 403, payload)
}

func (o *GetInstallComponentLatestDeployForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployForbidden %s", 403, payload)
}

func (o *GetInstallComponentLatestDeployForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentLatestDeployForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentLatestDeployNotFound creates a GetInstallComponentLatestDeployNotFound with default headers values
func NewGetInstallComponentLatestDeployNotFound() *GetInstallComponentLatestDeployNotFound {
	return &GetInstallComponentLatestDeployNotFound{}
}

/*
GetInstallComponentLatestDeployNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallComponentLatestDeployNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component latest deploy not found response has a 2xx status code
func (o *GetInstallComponentLatestDeployNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component latest deploy not found response has a 3xx status code
func (o *GetInstallComponentLatestDeployNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component latest deploy not found response has a 4xx status code
func (o *GetInstallComponentLatestDeployNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component latest deploy not found response has a 5xx status code
func (o *GetInstallComponentLatestDeployNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component latest deploy not found response a status code equal to that given
func (o *GetInstallComponentLatestDeployNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install component latest deploy not found response
func (o *GetInstallComponentLatestDeployNotFound) Code() int {
	return 404
}

func (o *GetInstallComponentLatestDeployNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployNotFound %s", 404, payload)
}

func (o *GetInstallComponentLatestDeployNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployNotFound %s", 404, payload)
}

func (o *GetInstallComponentLatestDeployNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentLatestDeployNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentLatestDeployInternalServerError creates a GetInstallComponentLatestDeployInternalServerError with default headers values
func NewGetInstallComponentLatestDeployInternalServerError() *GetInstallComponentLatestDeployInternalServerError {
	return &GetInstallComponentLatestDeployInternalServerError{}
}

/*
GetInstallComponentLatestDeployInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallComponentLatestDeployInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component latest deploy internal server error response has a 2xx status code
func (o *GetInstallComponentLatestDeployInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component latest deploy internal server error response has a 3xx status code
func (o *GetInstallComponentLatestDeployInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component latest deploy internal server error response has a 4xx status code
func (o *GetInstallComponentLatestDeployInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install component latest deploy internal server error response has a 5xx status code
func (o *GetInstallComponentLatestDeployInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install component latest deploy internal server error response a status code equal to that given
func (o *GetInstallComponentLatestDeployInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install component latest deploy internal server error response
func (o *GetInstallComponentLatestDeployInternalServerError) Code() int {
	return 500
}

func (o *GetInstallComponentLatestDeployInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployInternalServerError %s", 500, payload)
}

func (o *GetInstallComponentLatestDeployInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/deploys/latest][%d] getInstallComponentLatestDeployInternalServerError %s", 500, payload)
}

func (o *GetInstallComponentLatestDeployInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentLatestDeployInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
