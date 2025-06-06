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

// GetInstallComponentOutputsReader is a Reader for the GetInstallComponentOutputs structure.
type GetInstallComponentOutputsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallComponentOutputsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallComponentOutputsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallComponentOutputsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallComponentOutputsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallComponentOutputsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallComponentOutputsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallComponentOutputsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/components/{component_id}/outputs] GetInstallComponentOutputs", response, response.Code())
	}
}

// NewGetInstallComponentOutputsOK creates a GetInstallComponentOutputsOK with default headers values
func NewGetInstallComponentOutputsOK() *GetInstallComponentOutputsOK {
	return &GetInstallComponentOutputsOK{}
}

/*
GetInstallComponentOutputsOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallComponentOutputsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get install component outputs o k response has a 2xx status code
func (o *GetInstallComponentOutputsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install component outputs o k response has a 3xx status code
func (o *GetInstallComponentOutputsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component outputs o k response has a 4xx status code
func (o *GetInstallComponentOutputsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install component outputs o k response has a 5xx status code
func (o *GetInstallComponentOutputsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component outputs o k response a status code equal to that given
func (o *GetInstallComponentOutputsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install component outputs o k response
func (o *GetInstallComponentOutputsOK) Code() int {
	return 200
}

func (o *GetInstallComponentOutputsOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsOK  %+v", 200, o.Payload)
}

func (o *GetInstallComponentOutputsOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsOK  %+v", 200, o.Payload)
}

func (o *GetInstallComponentOutputsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetInstallComponentOutputsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentOutputsBadRequest creates a GetInstallComponentOutputsBadRequest with default headers values
func NewGetInstallComponentOutputsBadRequest() *GetInstallComponentOutputsBadRequest {
	return &GetInstallComponentOutputsBadRequest{}
}

/*
GetInstallComponentOutputsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallComponentOutputsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component outputs bad request response has a 2xx status code
func (o *GetInstallComponentOutputsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component outputs bad request response has a 3xx status code
func (o *GetInstallComponentOutputsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component outputs bad request response has a 4xx status code
func (o *GetInstallComponentOutputsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component outputs bad request response has a 5xx status code
func (o *GetInstallComponentOutputsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component outputs bad request response a status code equal to that given
func (o *GetInstallComponentOutputsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install component outputs bad request response
func (o *GetInstallComponentOutputsBadRequest) Code() int {
	return 400
}

func (o *GetInstallComponentOutputsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallComponentOutputsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallComponentOutputsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentOutputsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentOutputsUnauthorized creates a GetInstallComponentOutputsUnauthorized with default headers values
func NewGetInstallComponentOutputsUnauthorized() *GetInstallComponentOutputsUnauthorized {
	return &GetInstallComponentOutputsUnauthorized{}
}

/*
GetInstallComponentOutputsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallComponentOutputsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component outputs unauthorized response has a 2xx status code
func (o *GetInstallComponentOutputsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component outputs unauthorized response has a 3xx status code
func (o *GetInstallComponentOutputsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component outputs unauthorized response has a 4xx status code
func (o *GetInstallComponentOutputsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component outputs unauthorized response has a 5xx status code
func (o *GetInstallComponentOutputsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component outputs unauthorized response a status code equal to that given
func (o *GetInstallComponentOutputsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install component outputs unauthorized response
func (o *GetInstallComponentOutputsUnauthorized) Code() int {
	return 401
}

func (o *GetInstallComponentOutputsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallComponentOutputsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallComponentOutputsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentOutputsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentOutputsForbidden creates a GetInstallComponentOutputsForbidden with default headers values
func NewGetInstallComponentOutputsForbidden() *GetInstallComponentOutputsForbidden {
	return &GetInstallComponentOutputsForbidden{}
}

/*
GetInstallComponentOutputsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallComponentOutputsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component outputs forbidden response has a 2xx status code
func (o *GetInstallComponentOutputsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component outputs forbidden response has a 3xx status code
func (o *GetInstallComponentOutputsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component outputs forbidden response has a 4xx status code
func (o *GetInstallComponentOutputsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component outputs forbidden response has a 5xx status code
func (o *GetInstallComponentOutputsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component outputs forbidden response a status code equal to that given
func (o *GetInstallComponentOutputsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install component outputs forbidden response
func (o *GetInstallComponentOutputsForbidden) Code() int {
	return 403
}

func (o *GetInstallComponentOutputsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallComponentOutputsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallComponentOutputsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentOutputsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentOutputsNotFound creates a GetInstallComponentOutputsNotFound with default headers values
func NewGetInstallComponentOutputsNotFound() *GetInstallComponentOutputsNotFound {
	return &GetInstallComponentOutputsNotFound{}
}

/*
GetInstallComponentOutputsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallComponentOutputsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component outputs not found response has a 2xx status code
func (o *GetInstallComponentOutputsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component outputs not found response has a 3xx status code
func (o *GetInstallComponentOutputsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component outputs not found response has a 4xx status code
func (o *GetInstallComponentOutputsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install component outputs not found response has a 5xx status code
func (o *GetInstallComponentOutputsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install component outputs not found response a status code equal to that given
func (o *GetInstallComponentOutputsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install component outputs not found response
func (o *GetInstallComponentOutputsNotFound) Code() int {
	return 404
}

func (o *GetInstallComponentOutputsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallComponentOutputsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallComponentOutputsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentOutputsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallComponentOutputsInternalServerError creates a GetInstallComponentOutputsInternalServerError with default headers values
func NewGetInstallComponentOutputsInternalServerError() *GetInstallComponentOutputsInternalServerError {
	return &GetInstallComponentOutputsInternalServerError{}
}

/*
GetInstallComponentOutputsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallComponentOutputsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install component outputs internal server error response has a 2xx status code
func (o *GetInstallComponentOutputsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install component outputs internal server error response has a 3xx status code
func (o *GetInstallComponentOutputsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install component outputs internal server error response has a 4xx status code
func (o *GetInstallComponentOutputsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install component outputs internal server error response has a 5xx status code
func (o *GetInstallComponentOutputsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install component outputs internal server error response a status code equal to that given
func (o *GetInstallComponentOutputsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install component outputs internal server error response
func (o *GetInstallComponentOutputsInternalServerError) Code() int {
	return 500
}

func (o *GetInstallComponentOutputsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallComponentOutputsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components/{component_id}/outputs][%d] getInstallComponentOutputsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallComponentOutputsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallComponentOutputsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
