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

// CreateVCSConnectionReader is a Reader for the CreateVCSConnection structure.
type CreateVCSConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVCSConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVCSConnectionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVCSConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateVCSConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVCSConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVCSConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVCSConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/vcs/connections] CreateVCSConnection", response, response.Code())
	}
}

// NewCreateVCSConnectionCreated creates a CreateVCSConnectionCreated with default headers values
func NewCreateVCSConnectionCreated() *CreateVCSConnectionCreated {
	return &CreateVCSConnectionCreated{}
}

/*
CreateVCSConnectionCreated describes a response with status code 201, with default header values.

Created
*/
type CreateVCSConnectionCreated struct {
	Payload *models.AppVCSConnection
}

// IsSuccess returns true when this create v c s connection created response has a 2xx status code
func (o *CreateVCSConnectionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create v c s connection created response has a 3xx status code
func (o *CreateVCSConnectionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection created response has a 4xx status code
func (o *CreateVCSConnectionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create v c s connection created response has a 5xx status code
func (o *CreateVCSConnectionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection created response a status code equal to that given
func (o *CreateVCSConnectionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create v c s connection created response
func (o *CreateVCSConnectionCreated) Code() int {
	return 201
}

func (o *CreateVCSConnectionCreated) Error() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionCreated  %+v", 201, o.Payload)
}

func (o *CreateVCSConnectionCreated) String() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionCreated  %+v", 201, o.Payload)
}

func (o *CreateVCSConnectionCreated) GetPayload() *models.AppVCSConnection {
	return o.Payload
}

func (o *CreateVCSConnectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppVCSConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionBadRequest creates a CreateVCSConnectionBadRequest with default headers values
func NewCreateVCSConnectionBadRequest() *CreateVCSConnectionBadRequest {
	return &CreateVCSConnectionBadRequest{}
}

/*
CreateVCSConnectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateVCSConnectionBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection bad request response has a 2xx status code
func (o *CreateVCSConnectionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection bad request response has a 3xx status code
func (o *CreateVCSConnectionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection bad request response has a 4xx status code
func (o *CreateVCSConnectionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection bad request response has a 5xx status code
func (o *CreateVCSConnectionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection bad request response a status code equal to that given
func (o *CreateVCSConnectionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create v c s connection bad request response
func (o *CreateVCSConnectionBadRequest) Code() int {
	return 400
}

func (o *CreateVCSConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVCSConnectionBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVCSConnectionBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionUnauthorized creates a CreateVCSConnectionUnauthorized with default headers values
func NewCreateVCSConnectionUnauthorized() *CreateVCSConnectionUnauthorized {
	return &CreateVCSConnectionUnauthorized{}
}

/*
CreateVCSConnectionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateVCSConnectionUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection unauthorized response has a 2xx status code
func (o *CreateVCSConnectionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection unauthorized response has a 3xx status code
func (o *CreateVCSConnectionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection unauthorized response has a 4xx status code
func (o *CreateVCSConnectionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection unauthorized response has a 5xx status code
func (o *CreateVCSConnectionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection unauthorized response a status code equal to that given
func (o *CreateVCSConnectionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create v c s connection unauthorized response
func (o *CreateVCSConnectionUnauthorized) Code() int {
	return 401
}

func (o *CreateVCSConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateVCSConnectionUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateVCSConnectionUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionForbidden creates a CreateVCSConnectionForbidden with default headers values
func NewCreateVCSConnectionForbidden() *CreateVCSConnectionForbidden {
	return &CreateVCSConnectionForbidden{}
}

/*
CreateVCSConnectionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateVCSConnectionForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection forbidden response has a 2xx status code
func (o *CreateVCSConnectionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection forbidden response has a 3xx status code
func (o *CreateVCSConnectionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection forbidden response has a 4xx status code
func (o *CreateVCSConnectionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection forbidden response has a 5xx status code
func (o *CreateVCSConnectionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection forbidden response a status code equal to that given
func (o *CreateVCSConnectionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create v c s connection forbidden response
func (o *CreateVCSConnectionForbidden) Code() int {
	return 403
}

func (o *CreateVCSConnectionForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionForbidden  %+v", 403, o.Payload)
}

func (o *CreateVCSConnectionForbidden) String() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionForbidden  %+v", 403, o.Payload)
}

func (o *CreateVCSConnectionForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionNotFound creates a CreateVCSConnectionNotFound with default headers values
func NewCreateVCSConnectionNotFound() *CreateVCSConnectionNotFound {
	return &CreateVCSConnectionNotFound{}
}

/*
CreateVCSConnectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateVCSConnectionNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection not found response has a 2xx status code
func (o *CreateVCSConnectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection not found response has a 3xx status code
func (o *CreateVCSConnectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection not found response has a 4xx status code
func (o *CreateVCSConnectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection not found response has a 5xx status code
func (o *CreateVCSConnectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection not found response a status code equal to that given
func (o *CreateVCSConnectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create v c s connection not found response
func (o *CreateVCSConnectionNotFound) Code() int {
	return 404
}

func (o *CreateVCSConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionNotFound  %+v", 404, o.Payload)
}

func (o *CreateVCSConnectionNotFound) String() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionNotFound  %+v", 404, o.Payload)
}

func (o *CreateVCSConnectionNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionInternalServerError creates a CreateVCSConnectionInternalServerError with default headers values
func NewCreateVCSConnectionInternalServerError() *CreateVCSConnectionInternalServerError {
	return &CreateVCSConnectionInternalServerError{}
}

/*
CreateVCSConnectionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateVCSConnectionInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection internal server error response has a 2xx status code
func (o *CreateVCSConnectionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection internal server error response has a 3xx status code
func (o *CreateVCSConnectionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection internal server error response has a 4xx status code
func (o *CreateVCSConnectionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create v c s connection internal server error response has a 5xx status code
func (o *CreateVCSConnectionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create v c s connection internal server error response a status code equal to that given
func (o *CreateVCSConnectionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create v c s connection internal server error response
func (o *CreateVCSConnectionInternalServerError) Code() int {
	return 500
}

func (o *CreateVCSConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateVCSConnectionInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/vcs/connections][%d] createVCSConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateVCSConnectionInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}