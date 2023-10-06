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

// PutV1AppsAppIDSandboxReader is a Reader for the PutV1AppsAppIDSandbox structure.
type PutV1AppsAppIDSandboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1AppsAppIDSandboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1AppsAppIDSandboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutV1AppsAppIDSandboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutV1AppsAppIDSandboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutV1AppsAppIDSandboxForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutV1AppsAppIDSandboxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutV1AppsAppIDSandboxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/apps/{app_id}/sandbox] PutV1AppsAppIDSandbox", response, response.Code())
	}
}

// NewPutV1AppsAppIDSandboxOK creates a PutV1AppsAppIDSandboxOK with default headers values
func NewPutV1AppsAppIDSandboxOK() *PutV1AppsAppIDSandboxOK {
	return &PutV1AppsAppIDSandboxOK{}
}

/*
PutV1AppsAppIDSandboxOK describes a response with status code 200, with default header values.

OK
*/
type PutV1AppsAppIDSandboxOK struct {
	Payload *models.AppApp
}

// IsSuccess returns true when this put v1 apps app Id sandbox o k response has a 2xx status code
func (o *PutV1AppsAppIDSandboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 apps app Id sandbox o k response has a 3xx status code
func (o *PutV1AppsAppIDSandboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 apps app Id sandbox o k response has a 4xx status code
func (o *PutV1AppsAppIDSandboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 apps app Id sandbox o k response has a 5xx status code
func (o *PutV1AppsAppIDSandboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 apps app Id sandbox o k response a status code equal to that given
func (o *PutV1AppsAppIDSandboxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put v1 apps app Id sandbox o k response
func (o *PutV1AppsAppIDSandboxOK) Code() int {
	return 200
}

func (o *PutV1AppsAppIDSandboxOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxOK  %+v", 200, o.Payload)
}

func (o *PutV1AppsAppIDSandboxOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxOK  %+v", 200, o.Payload)
}

func (o *PutV1AppsAppIDSandboxOK) GetPayload() *models.AppApp {
	return o.Payload
}

func (o *PutV1AppsAppIDSandboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AppsAppIDSandboxBadRequest creates a PutV1AppsAppIDSandboxBadRequest with default headers values
func NewPutV1AppsAppIDSandboxBadRequest() *PutV1AppsAppIDSandboxBadRequest {
	return &PutV1AppsAppIDSandboxBadRequest{}
}

/*
PutV1AppsAppIDSandboxBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutV1AppsAppIDSandboxBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this put v1 apps app Id sandbox bad request response has a 2xx status code
func (o *PutV1AppsAppIDSandboxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 apps app Id sandbox bad request response has a 3xx status code
func (o *PutV1AppsAppIDSandboxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 apps app Id sandbox bad request response has a 4xx status code
func (o *PutV1AppsAppIDSandboxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 apps app Id sandbox bad request response has a 5xx status code
func (o *PutV1AppsAppIDSandboxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 apps app Id sandbox bad request response a status code equal to that given
func (o *PutV1AppsAppIDSandboxBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put v1 apps app Id sandbox bad request response
func (o *PutV1AppsAppIDSandboxBadRequest) Code() int {
	return 400
}

func (o *PutV1AppsAppIDSandboxBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1AppsAppIDSandboxBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxBadRequest  %+v", 400, o.Payload)
}

func (o *PutV1AppsAppIDSandboxBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PutV1AppsAppIDSandboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AppsAppIDSandboxUnauthorized creates a PutV1AppsAppIDSandboxUnauthorized with default headers values
func NewPutV1AppsAppIDSandboxUnauthorized() *PutV1AppsAppIDSandboxUnauthorized {
	return &PutV1AppsAppIDSandboxUnauthorized{}
}

/*
PutV1AppsAppIDSandboxUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutV1AppsAppIDSandboxUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this put v1 apps app Id sandbox unauthorized response has a 2xx status code
func (o *PutV1AppsAppIDSandboxUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 apps app Id sandbox unauthorized response has a 3xx status code
func (o *PutV1AppsAppIDSandboxUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 apps app Id sandbox unauthorized response has a 4xx status code
func (o *PutV1AppsAppIDSandboxUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 apps app Id sandbox unauthorized response has a 5xx status code
func (o *PutV1AppsAppIDSandboxUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 apps app Id sandbox unauthorized response a status code equal to that given
func (o *PutV1AppsAppIDSandboxUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put v1 apps app Id sandbox unauthorized response
func (o *PutV1AppsAppIDSandboxUnauthorized) Code() int {
	return 401
}

func (o *PutV1AppsAppIDSandboxUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxUnauthorized  %+v", 401, o.Payload)
}

func (o *PutV1AppsAppIDSandboxUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxUnauthorized  %+v", 401, o.Payload)
}

func (o *PutV1AppsAppIDSandboxUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PutV1AppsAppIDSandboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AppsAppIDSandboxForbidden creates a PutV1AppsAppIDSandboxForbidden with default headers values
func NewPutV1AppsAppIDSandboxForbidden() *PutV1AppsAppIDSandboxForbidden {
	return &PutV1AppsAppIDSandboxForbidden{}
}

/*
PutV1AppsAppIDSandboxForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutV1AppsAppIDSandboxForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this put v1 apps app Id sandbox forbidden response has a 2xx status code
func (o *PutV1AppsAppIDSandboxForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 apps app Id sandbox forbidden response has a 3xx status code
func (o *PutV1AppsAppIDSandboxForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 apps app Id sandbox forbidden response has a 4xx status code
func (o *PutV1AppsAppIDSandboxForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 apps app Id sandbox forbidden response has a 5xx status code
func (o *PutV1AppsAppIDSandboxForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 apps app Id sandbox forbidden response a status code equal to that given
func (o *PutV1AppsAppIDSandboxForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put v1 apps app Id sandbox forbidden response
func (o *PutV1AppsAppIDSandboxForbidden) Code() int {
	return 403
}

func (o *PutV1AppsAppIDSandboxForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxForbidden  %+v", 403, o.Payload)
}

func (o *PutV1AppsAppIDSandboxForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxForbidden  %+v", 403, o.Payload)
}

func (o *PutV1AppsAppIDSandboxForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PutV1AppsAppIDSandboxForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AppsAppIDSandboxNotFound creates a PutV1AppsAppIDSandboxNotFound with default headers values
func NewPutV1AppsAppIDSandboxNotFound() *PutV1AppsAppIDSandboxNotFound {
	return &PutV1AppsAppIDSandboxNotFound{}
}

/*
PutV1AppsAppIDSandboxNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutV1AppsAppIDSandboxNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this put v1 apps app Id sandbox not found response has a 2xx status code
func (o *PutV1AppsAppIDSandboxNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 apps app Id sandbox not found response has a 3xx status code
func (o *PutV1AppsAppIDSandboxNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 apps app Id sandbox not found response has a 4xx status code
func (o *PutV1AppsAppIDSandboxNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put v1 apps app Id sandbox not found response has a 5xx status code
func (o *PutV1AppsAppIDSandboxNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 apps app Id sandbox not found response a status code equal to that given
func (o *PutV1AppsAppIDSandboxNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put v1 apps app Id sandbox not found response
func (o *PutV1AppsAppIDSandboxNotFound) Code() int {
	return 404
}

func (o *PutV1AppsAppIDSandboxNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxNotFound  %+v", 404, o.Payload)
}

func (o *PutV1AppsAppIDSandboxNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxNotFound  %+v", 404, o.Payload)
}

func (o *PutV1AppsAppIDSandboxNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PutV1AppsAppIDSandboxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutV1AppsAppIDSandboxInternalServerError creates a PutV1AppsAppIDSandboxInternalServerError with default headers values
func NewPutV1AppsAppIDSandboxInternalServerError() *PutV1AppsAppIDSandboxInternalServerError {
	return &PutV1AppsAppIDSandboxInternalServerError{}
}

/*
PutV1AppsAppIDSandboxInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutV1AppsAppIDSandboxInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this put v1 apps app Id sandbox internal server error response has a 2xx status code
func (o *PutV1AppsAppIDSandboxInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put v1 apps app Id sandbox internal server error response has a 3xx status code
func (o *PutV1AppsAppIDSandboxInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 apps app Id sandbox internal server error response has a 4xx status code
func (o *PutV1AppsAppIDSandboxInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 apps app Id sandbox internal server error response has a 5xx status code
func (o *PutV1AppsAppIDSandboxInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put v1 apps app Id sandbox internal server error response a status code equal to that given
func (o *PutV1AppsAppIDSandboxInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put v1 apps app Id sandbox internal server error response
func (o *PutV1AppsAppIDSandboxInternalServerError) Code() int {
	return 500
}

func (o *PutV1AppsAppIDSandboxInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxInternalServerError  %+v", 500, o.Payload)
}

func (o *PutV1AppsAppIDSandboxInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/{app_id}/sandbox][%d] putV1AppsAppIdSandboxInternalServerError  %+v", 500, o.Payload)
}

func (o *PutV1AppsAppIDSandboxInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PutV1AppsAppIDSandboxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
