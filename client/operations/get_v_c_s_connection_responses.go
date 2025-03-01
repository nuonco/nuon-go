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

// GetVCSConnectionReader is a Reader for the GetVCSConnection structure.
type GetVCSConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVCSConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVCSConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVCSConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVCSConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVCSConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVCSConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVCSConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/vcs/connections/{connection_id}] GetVCSConnection", response, response.Code())
	}
}

// NewGetVCSConnectionOK creates a GetVCSConnectionOK with default headers values
func NewGetVCSConnectionOK() *GetVCSConnectionOK {
	return &GetVCSConnectionOK{}
}

/*
GetVCSConnectionOK describes a response with status code 200, with default header values.

OK
*/
type GetVCSConnectionOK struct {
	Payload *models.AppVCSConnection
}

// IsSuccess returns true when this get v c s connection o k response has a 2xx status code
func (o *GetVCSConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v c s connection o k response has a 3xx status code
func (o *GetVCSConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s connection o k response has a 4xx status code
func (o *GetVCSConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v c s connection o k response has a 5xx status code
func (o *GetVCSConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s connection o k response a status code equal to that given
func (o *GetVCSConnectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v c s connection o k response
func (o *GetVCSConnectionOK) Code() int {
	return 200
}

func (o *GetVCSConnectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionOK %s", 200, payload)
}

func (o *GetVCSConnectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionOK %s", 200, payload)
}

func (o *GetVCSConnectionOK) GetPayload() *models.AppVCSConnection {
	return o.Payload
}

func (o *GetVCSConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppVCSConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCSConnectionBadRequest creates a GetVCSConnectionBadRequest with default headers values
func NewGetVCSConnectionBadRequest() *GetVCSConnectionBadRequest {
	return &GetVCSConnectionBadRequest{}
}

/*
GetVCSConnectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetVCSConnectionBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v c s connection bad request response has a 2xx status code
func (o *GetVCSConnectionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s connection bad request response has a 3xx status code
func (o *GetVCSConnectionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s connection bad request response has a 4xx status code
func (o *GetVCSConnectionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s connection bad request response has a 5xx status code
func (o *GetVCSConnectionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s connection bad request response a status code equal to that given
func (o *GetVCSConnectionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v c s connection bad request response
func (o *GetVCSConnectionBadRequest) Code() int {
	return 400
}

func (o *GetVCSConnectionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionBadRequest %s", 400, payload)
}

func (o *GetVCSConnectionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionBadRequest %s", 400, payload)
}

func (o *GetVCSConnectionBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetVCSConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCSConnectionUnauthorized creates a GetVCSConnectionUnauthorized with default headers values
func NewGetVCSConnectionUnauthorized() *GetVCSConnectionUnauthorized {
	return &GetVCSConnectionUnauthorized{}
}

/*
GetVCSConnectionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVCSConnectionUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v c s connection unauthorized response has a 2xx status code
func (o *GetVCSConnectionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s connection unauthorized response has a 3xx status code
func (o *GetVCSConnectionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s connection unauthorized response has a 4xx status code
func (o *GetVCSConnectionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s connection unauthorized response has a 5xx status code
func (o *GetVCSConnectionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s connection unauthorized response a status code equal to that given
func (o *GetVCSConnectionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v c s connection unauthorized response
func (o *GetVCSConnectionUnauthorized) Code() int {
	return 401
}

func (o *GetVCSConnectionUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionUnauthorized %s", 401, payload)
}

func (o *GetVCSConnectionUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionUnauthorized %s", 401, payload)
}

func (o *GetVCSConnectionUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetVCSConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCSConnectionForbidden creates a GetVCSConnectionForbidden with default headers values
func NewGetVCSConnectionForbidden() *GetVCSConnectionForbidden {
	return &GetVCSConnectionForbidden{}
}

/*
GetVCSConnectionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVCSConnectionForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v c s connection forbidden response has a 2xx status code
func (o *GetVCSConnectionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s connection forbidden response has a 3xx status code
func (o *GetVCSConnectionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s connection forbidden response has a 4xx status code
func (o *GetVCSConnectionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s connection forbidden response has a 5xx status code
func (o *GetVCSConnectionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s connection forbidden response a status code equal to that given
func (o *GetVCSConnectionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v c s connection forbidden response
func (o *GetVCSConnectionForbidden) Code() int {
	return 403
}

func (o *GetVCSConnectionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionForbidden %s", 403, payload)
}

func (o *GetVCSConnectionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionForbidden %s", 403, payload)
}

func (o *GetVCSConnectionForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetVCSConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCSConnectionNotFound creates a GetVCSConnectionNotFound with default headers values
func NewGetVCSConnectionNotFound() *GetVCSConnectionNotFound {
	return &GetVCSConnectionNotFound{}
}

/*
GetVCSConnectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVCSConnectionNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v c s connection not found response has a 2xx status code
func (o *GetVCSConnectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s connection not found response has a 3xx status code
func (o *GetVCSConnectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s connection not found response has a 4xx status code
func (o *GetVCSConnectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s connection not found response has a 5xx status code
func (o *GetVCSConnectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s connection not found response a status code equal to that given
func (o *GetVCSConnectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v c s connection not found response
func (o *GetVCSConnectionNotFound) Code() int {
	return 404
}

func (o *GetVCSConnectionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionNotFound %s", 404, payload)
}

func (o *GetVCSConnectionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionNotFound %s", 404, payload)
}

func (o *GetVCSConnectionNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetVCSConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCSConnectionInternalServerError creates a GetVCSConnectionInternalServerError with default headers values
func NewGetVCSConnectionInternalServerError() *GetVCSConnectionInternalServerError {
	return &GetVCSConnectionInternalServerError{}
}

/*
GetVCSConnectionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetVCSConnectionInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v c s connection internal server error response has a 2xx status code
func (o *GetVCSConnectionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s connection internal server error response has a 3xx status code
func (o *GetVCSConnectionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s connection internal server error response has a 4xx status code
func (o *GetVCSConnectionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v c s connection internal server error response has a 5xx status code
func (o *GetVCSConnectionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v c s connection internal server error response a status code equal to that given
func (o *GetVCSConnectionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v c s connection internal server error response
func (o *GetVCSConnectionInternalServerError) Code() int {
	return 500
}

func (o *GetVCSConnectionInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionInternalServerError %s", 500, payload)
}

func (o *GetVCSConnectionInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getVCSConnectionInternalServerError %s", 500, payload)
}

func (o *GetVCSConnectionInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetVCSConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
