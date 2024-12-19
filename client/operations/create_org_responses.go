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

// CreateOrgReader is a Reader for the CreateOrg structure.
type CreateOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrgCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrgNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/orgs] CreateOrg", response, response.Code())
	}
}

// NewCreateOrgCreated creates a CreateOrgCreated with default headers values
func NewCreateOrgCreated() *CreateOrgCreated {
	return &CreateOrgCreated{}
}

/*
CreateOrgCreated describes a response with status code 201, with default header values.

Created
*/
type CreateOrgCreated struct {
	Payload *models.AppOrg
}

// IsSuccess returns true when this create org created response has a 2xx status code
func (o *CreateOrgCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create org created response has a 3xx status code
func (o *CreateOrgCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org created response has a 4xx status code
func (o *CreateOrgCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create org created response has a 5xx status code
func (o *CreateOrgCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create org created response a status code equal to that given
func (o *CreateOrgCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create org created response
func (o *CreateOrgCreated) Code() int {
	return 201
}

func (o *CreateOrgCreated) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgCreated  %+v", 201, o.Payload)
}

func (o *CreateOrgCreated) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgCreated  %+v", 201, o.Payload)
}

func (o *CreateOrgCreated) GetPayload() *models.AppOrg {
	return o.Payload
}

func (o *CreateOrgCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppOrg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgBadRequest creates a CreateOrgBadRequest with default headers values
func NewCreateOrgBadRequest() *CreateOrgBadRequest {
	return &CreateOrgBadRequest{}
}

/*
CreateOrgBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOrgBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create org bad request response has a 2xx status code
func (o *CreateOrgBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org bad request response has a 3xx status code
func (o *CreateOrgBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org bad request response has a 4xx status code
func (o *CreateOrgBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create org bad request response has a 5xx status code
func (o *CreateOrgBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create org bad request response a status code equal to that given
func (o *CreateOrgBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create org bad request response
func (o *CreateOrgBadRequest) Code() int {
	return 400
}

func (o *CreateOrgBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrgBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrgBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgUnauthorized creates a CreateOrgUnauthorized with default headers values
func NewCreateOrgUnauthorized() *CreateOrgUnauthorized {
	return &CreateOrgUnauthorized{}
}

/*
CreateOrgUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateOrgUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create org unauthorized response has a 2xx status code
func (o *CreateOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org unauthorized response has a 3xx status code
func (o *CreateOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org unauthorized response has a 4xx status code
func (o *CreateOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create org unauthorized response has a 5xx status code
func (o *CreateOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create org unauthorized response a status code equal to that given
func (o *CreateOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create org unauthorized response
func (o *CreateOrgUnauthorized) Code() int {
	return 401
}

func (o *CreateOrgUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOrgUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOrgUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgForbidden creates a CreateOrgForbidden with default headers values
func NewCreateOrgForbidden() *CreateOrgForbidden {
	return &CreateOrgForbidden{}
}

/*
CreateOrgForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateOrgForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create org forbidden response has a 2xx status code
func (o *CreateOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org forbidden response has a 3xx status code
func (o *CreateOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org forbidden response has a 4xx status code
func (o *CreateOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create org forbidden response has a 5xx status code
func (o *CreateOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create org forbidden response a status code equal to that given
func (o *CreateOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create org forbidden response
func (o *CreateOrgForbidden) Code() int {
	return 403
}

func (o *CreateOrgForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrgForbidden) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrgForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgNotFound creates a CreateOrgNotFound with default headers values
func NewCreateOrgNotFound() *CreateOrgNotFound {
	return &CreateOrgNotFound{}
}

/*
CreateOrgNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateOrgNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create org not found response has a 2xx status code
func (o *CreateOrgNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org not found response has a 3xx status code
func (o *CreateOrgNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org not found response has a 4xx status code
func (o *CreateOrgNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create org not found response has a 5xx status code
func (o *CreateOrgNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create org not found response a status code equal to that given
func (o *CreateOrgNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create org not found response
func (o *CreateOrgNotFound) Code() int {
	return 404
}

func (o *CreateOrgNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrgNotFound) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrgNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateOrgNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgInternalServerError creates a CreateOrgInternalServerError with default headers values
func NewCreateOrgInternalServerError() *CreateOrgInternalServerError {
	return &CreateOrgInternalServerError{}
}

/*
CreateOrgInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateOrgInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create org internal server error response has a 2xx status code
func (o *CreateOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create org internal server error response has a 3xx status code
func (o *CreateOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create org internal server error response has a 4xx status code
func (o *CreateOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create org internal server error response has a 5xx status code
func (o *CreateOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create org internal server error response a status code equal to that given
func (o *CreateOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create org internal server error response
func (o *CreateOrgInternalServerError) Code() int {
	return 500
}

func (o *CreateOrgInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOrgInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] createOrgInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOrgInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
