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

// ForgetInstallReader is a Reader for the ForgetInstall structure.
type ForgetInstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForgetInstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForgetInstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewForgetInstallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewForgetInstallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewForgetInstallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/installs/{install_id}/forget] ForgetInstall", response, response.Code())
	}
}

// NewForgetInstallOK creates a ForgetInstallOK with default headers values
func NewForgetInstallOK() *ForgetInstallOK {
	return &ForgetInstallOK{}
}

/*
ForgetInstallOK describes a response with status code 200, with default header values.

OK
*/
type ForgetInstallOK struct {
	Payload bool
}

// IsSuccess returns true when this forget install o k response has a 2xx status code
func (o *ForgetInstallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this forget install o k response has a 3xx status code
func (o *ForgetInstallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forget install o k response has a 4xx status code
func (o *ForgetInstallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this forget install o k response has a 5xx status code
func (o *ForgetInstallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this forget install o k response a status code equal to that given
func (o *ForgetInstallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the forget install o k response
func (o *ForgetInstallOK) Code() int {
	return 200
}

func (o *ForgetInstallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallOK %s", 200, payload)
}

func (o *ForgetInstallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallOK %s", 200, payload)
}

func (o *ForgetInstallOK) GetPayload() bool {
	return o.Payload
}

func (o *ForgetInstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgetInstallBadRequest creates a ForgetInstallBadRequest with default headers values
func NewForgetInstallBadRequest() *ForgetInstallBadRequest {
	return &ForgetInstallBadRequest{}
}

/*
ForgetInstallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ForgetInstallBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this forget install bad request response has a 2xx status code
func (o *ForgetInstallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this forget install bad request response has a 3xx status code
func (o *ForgetInstallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forget install bad request response has a 4xx status code
func (o *ForgetInstallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this forget install bad request response has a 5xx status code
func (o *ForgetInstallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this forget install bad request response a status code equal to that given
func (o *ForgetInstallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the forget install bad request response
func (o *ForgetInstallBadRequest) Code() int {
	return 400
}

func (o *ForgetInstallBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallBadRequest %s", 400, payload)
}

func (o *ForgetInstallBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallBadRequest %s", 400, payload)
}

func (o *ForgetInstallBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *ForgetInstallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgetInstallNotFound creates a ForgetInstallNotFound with default headers values
func NewForgetInstallNotFound() *ForgetInstallNotFound {
	return &ForgetInstallNotFound{}
}

/*
ForgetInstallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ForgetInstallNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this forget install not found response has a 2xx status code
func (o *ForgetInstallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this forget install not found response has a 3xx status code
func (o *ForgetInstallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forget install not found response has a 4xx status code
func (o *ForgetInstallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this forget install not found response has a 5xx status code
func (o *ForgetInstallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this forget install not found response a status code equal to that given
func (o *ForgetInstallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the forget install not found response
func (o *ForgetInstallNotFound) Code() int {
	return 404
}

func (o *ForgetInstallNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallNotFound %s", 404, payload)
}

func (o *ForgetInstallNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallNotFound %s", 404, payload)
}

func (o *ForgetInstallNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *ForgetInstallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForgetInstallInternalServerError creates a ForgetInstallInternalServerError with default headers values
func NewForgetInstallInternalServerError() *ForgetInstallInternalServerError {
	return &ForgetInstallInternalServerError{}
}

/*
ForgetInstallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ForgetInstallInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this forget install internal server error response has a 2xx status code
func (o *ForgetInstallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this forget install internal server error response has a 3xx status code
func (o *ForgetInstallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this forget install internal server error response has a 4xx status code
func (o *ForgetInstallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this forget install internal server error response has a 5xx status code
func (o *ForgetInstallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this forget install internal server error response a status code equal to that given
func (o *ForgetInstallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the forget install internal server error response
func (o *ForgetInstallInternalServerError) Code() int {
	return 500
}

func (o *ForgetInstallInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallInternalServerError %s", 500, payload)
}

func (o *ForgetInstallInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/forget][%d] forgetInstallInternalServerError %s", 500, payload)
}

func (o *ForgetInstallInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *ForgetInstallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
