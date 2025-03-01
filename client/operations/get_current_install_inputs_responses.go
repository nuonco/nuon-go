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

// GetCurrentInstallInputsReader is a Reader for the GetCurrentInstallInputs structure.
type GetCurrentInstallInputsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentInstallInputsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentInstallInputsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCurrentInstallInputsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCurrentInstallInputsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentInstallInputsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCurrentInstallInputsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentInstallInputsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/inputs/current] GetCurrentInstallInputs", response, response.Code())
	}
}

// NewGetCurrentInstallInputsOK creates a GetCurrentInstallInputsOK with default headers values
func NewGetCurrentInstallInputsOK() *GetCurrentInstallInputsOK {
	return &GetCurrentInstallInputsOK{}
}

/*
GetCurrentInstallInputsOK describes a response with status code 200, with default header values.

OK
*/
type GetCurrentInstallInputsOK struct {
	Payload *models.AppInstallInputs
}

// IsSuccess returns true when this get current install inputs o k response has a 2xx status code
func (o *GetCurrentInstallInputsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current install inputs o k response has a 3xx status code
func (o *GetCurrentInstallInputsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current install inputs o k response has a 4xx status code
func (o *GetCurrentInstallInputsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current install inputs o k response has a 5xx status code
func (o *GetCurrentInstallInputsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current install inputs o k response a status code equal to that given
func (o *GetCurrentInstallInputsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get current install inputs o k response
func (o *GetCurrentInstallInputsOK) Code() int {
	return 200
}

func (o *GetCurrentInstallInputsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsOK %s", 200, payload)
}

func (o *GetCurrentInstallInputsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsOK %s", 200, payload)
}

func (o *GetCurrentInstallInputsOK) GetPayload() *models.AppInstallInputs {
	return o.Payload
}

func (o *GetCurrentInstallInputsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallInputs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentInstallInputsBadRequest creates a GetCurrentInstallInputsBadRequest with default headers values
func NewGetCurrentInstallInputsBadRequest() *GetCurrentInstallInputsBadRequest {
	return &GetCurrentInstallInputsBadRequest{}
}

/*
GetCurrentInstallInputsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCurrentInstallInputsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current install inputs bad request response has a 2xx status code
func (o *GetCurrentInstallInputsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current install inputs bad request response has a 3xx status code
func (o *GetCurrentInstallInputsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current install inputs bad request response has a 4xx status code
func (o *GetCurrentInstallInputsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current install inputs bad request response has a 5xx status code
func (o *GetCurrentInstallInputsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get current install inputs bad request response a status code equal to that given
func (o *GetCurrentInstallInputsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get current install inputs bad request response
func (o *GetCurrentInstallInputsBadRequest) Code() int {
	return 400
}

func (o *GetCurrentInstallInputsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsBadRequest %s", 400, payload)
}

func (o *GetCurrentInstallInputsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsBadRequest %s", 400, payload)
}

func (o *GetCurrentInstallInputsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentInstallInputsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentInstallInputsUnauthorized creates a GetCurrentInstallInputsUnauthorized with default headers values
func NewGetCurrentInstallInputsUnauthorized() *GetCurrentInstallInputsUnauthorized {
	return &GetCurrentInstallInputsUnauthorized{}
}

/*
GetCurrentInstallInputsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCurrentInstallInputsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current install inputs unauthorized response has a 2xx status code
func (o *GetCurrentInstallInputsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current install inputs unauthorized response has a 3xx status code
func (o *GetCurrentInstallInputsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current install inputs unauthorized response has a 4xx status code
func (o *GetCurrentInstallInputsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current install inputs unauthorized response has a 5xx status code
func (o *GetCurrentInstallInputsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get current install inputs unauthorized response a status code equal to that given
func (o *GetCurrentInstallInputsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get current install inputs unauthorized response
func (o *GetCurrentInstallInputsUnauthorized) Code() int {
	return 401
}

func (o *GetCurrentInstallInputsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsUnauthorized %s", 401, payload)
}

func (o *GetCurrentInstallInputsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsUnauthorized %s", 401, payload)
}

func (o *GetCurrentInstallInputsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentInstallInputsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentInstallInputsForbidden creates a GetCurrentInstallInputsForbidden with default headers values
func NewGetCurrentInstallInputsForbidden() *GetCurrentInstallInputsForbidden {
	return &GetCurrentInstallInputsForbidden{}
}

/*
GetCurrentInstallInputsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCurrentInstallInputsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current install inputs forbidden response has a 2xx status code
func (o *GetCurrentInstallInputsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current install inputs forbidden response has a 3xx status code
func (o *GetCurrentInstallInputsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current install inputs forbidden response has a 4xx status code
func (o *GetCurrentInstallInputsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current install inputs forbidden response has a 5xx status code
func (o *GetCurrentInstallInputsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get current install inputs forbidden response a status code equal to that given
func (o *GetCurrentInstallInputsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get current install inputs forbidden response
func (o *GetCurrentInstallInputsForbidden) Code() int {
	return 403
}

func (o *GetCurrentInstallInputsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsForbidden %s", 403, payload)
}

func (o *GetCurrentInstallInputsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsForbidden %s", 403, payload)
}

func (o *GetCurrentInstallInputsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentInstallInputsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentInstallInputsNotFound creates a GetCurrentInstallInputsNotFound with default headers values
func NewGetCurrentInstallInputsNotFound() *GetCurrentInstallInputsNotFound {
	return &GetCurrentInstallInputsNotFound{}
}

/*
GetCurrentInstallInputsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCurrentInstallInputsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current install inputs not found response has a 2xx status code
func (o *GetCurrentInstallInputsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current install inputs not found response has a 3xx status code
func (o *GetCurrentInstallInputsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current install inputs not found response has a 4xx status code
func (o *GetCurrentInstallInputsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current install inputs not found response has a 5xx status code
func (o *GetCurrentInstallInputsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get current install inputs not found response a status code equal to that given
func (o *GetCurrentInstallInputsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get current install inputs not found response
func (o *GetCurrentInstallInputsNotFound) Code() int {
	return 404
}

func (o *GetCurrentInstallInputsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsNotFound %s", 404, payload)
}

func (o *GetCurrentInstallInputsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsNotFound %s", 404, payload)
}

func (o *GetCurrentInstallInputsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentInstallInputsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentInstallInputsInternalServerError creates a GetCurrentInstallInputsInternalServerError with default headers values
func NewGetCurrentInstallInputsInternalServerError() *GetCurrentInstallInputsInternalServerError {
	return &GetCurrentInstallInputsInternalServerError{}
}

/*
GetCurrentInstallInputsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCurrentInstallInputsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current install inputs internal server error response has a 2xx status code
func (o *GetCurrentInstallInputsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current install inputs internal server error response has a 3xx status code
func (o *GetCurrentInstallInputsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current install inputs internal server error response has a 4xx status code
func (o *GetCurrentInstallInputsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current install inputs internal server error response has a 5xx status code
func (o *GetCurrentInstallInputsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get current install inputs internal server error response a status code equal to that given
func (o *GetCurrentInstallInputsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get current install inputs internal server error response
func (o *GetCurrentInstallInputsInternalServerError) Code() int {
	return 500
}

func (o *GetCurrentInstallInputsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsInternalServerError %s", 500, payload)
}

func (o *GetCurrentInstallInputsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getCurrentInstallInputsInternalServerError %s", 500, payload)
}

func (o *GetCurrentInstallInputsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentInstallInputsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
