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

// CreateInstallInputsReader is a Reader for the CreateInstallInputs structure.
type CreateInstallInputsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstallInputsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInstallInputsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInstallInputsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInstallInputsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInstallInputsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInstallInputsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInstallInputsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/installs/{install_id}/inputs] CreateInstallInputs", response, response.Code())
	}
}

// NewCreateInstallInputsCreated creates a CreateInstallInputsCreated with default headers values
func NewCreateInstallInputsCreated() *CreateInstallInputsCreated {
	return &CreateInstallInputsCreated{}
}

/*
CreateInstallInputsCreated describes a response with status code 201, with default header values.

Created
*/
type CreateInstallInputsCreated struct {
	Payload *models.AppInstallInputs
}

// IsSuccess returns true when this create install inputs created response has a 2xx status code
func (o *CreateInstallInputsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create install inputs created response has a 3xx status code
func (o *CreateInstallInputsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create install inputs created response has a 4xx status code
func (o *CreateInstallInputsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create install inputs created response has a 5xx status code
func (o *CreateInstallInputsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create install inputs created response a status code equal to that given
func (o *CreateInstallInputsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create install inputs created response
func (o *CreateInstallInputsCreated) Code() int {
	return 201
}

func (o *CreateInstallInputsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsCreated  %+v", 201, o.Payload)
}

func (o *CreateInstallInputsCreated) String() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsCreated  %+v", 201, o.Payload)
}

func (o *CreateInstallInputsCreated) GetPayload() *models.AppInstallInputs {
	return o.Payload
}

func (o *CreateInstallInputsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallInputs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstallInputsBadRequest creates a CreateInstallInputsBadRequest with default headers values
func NewCreateInstallInputsBadRequest() *CreateInstallInputsBadRequest {
	return &CreateInstallInputsBadRequest{}
}

/*
CreateInstallInputsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateInstallInputsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create install inputs bad request response has a 2xx status code
func (o *CreateInstallInputsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create install inputs bad request response has a 3xx status code
func (o *CreateInstallInputsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create install inputs bad request response has a 4xx status code
func (o *CreateInstallInputsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create install inputs bad request response has a 5xx status code
func (o *CreateInstallInputsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create install inputs bad request response a status code equal to that given
func (o *CreateInstallInputsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create install inputs bad request response
func (o *CreateInstallInputsBadRequest) Code() int {
	return 400
}

func (o *CreateInstallInputsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstallInputsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstallInputsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateInstallInputsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstallInputsUnauthorized creates a CreateInstallInputsUnauthorized with default headers values
func NewCreateInstallInputsUnauthorized() *CreateInstallInputsUnauthorized {
	return &CreateInstallInputsUnauthorized{}
}

/*
CreateInstallInputsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateInstallInputsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create install inputs unauthorized response has a 2xx status code
func (o *CreateInstallInputsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create install inputs unauthorized response has a 3xx status code
func (o *CreateInstallInputsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create install inputs unauthorized response has a 4xx status code
func (o *CreateInstallInputsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create install inputs unauthorized response has a 5xx status code
func (o *CreateInstallInputsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create install inputs unauthorized response a status code equal to that given
func (o *CreateInstallInputsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create install inputs unauthorized response
func (o *CreateInstallInputsUnauthorized) Code() int {
	return 401
}

func (o *CreateInstallInputsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInstallInputsUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInstallInputsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateInstallInputsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstallInputsForbidden creates a CreateInstallInputsForbidden with default headers values
func NewCreateInstallInputsForbidden() *CreateInstallInputsForbidden {
	return &CreateInstallInputsForbidden{}
}

/*
CreateInstallInputsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateInstallInputsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create install inputs forbidden response has a 2xx status code
func (o *CreateInstallInputsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create install inputs forbidden response has a 3xx status code
func (o *CreateInstallInputsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create install inputs forbidden response has a 4xx status code
func (o *CreateInstallInputsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create install inputs forbidden response has a 5xx status code
func (o *CreateInstallInputsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create install inputs forbidden response a status code equal to that given
func (o *CreateInstallInputsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create install inputs forbidden response
func (o *CreateInstallInputsForbidden) Code() int {
	return 403
}

func (o *CreateInstallInputsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsForbidden  %+v", 403, o.Payload)
}

func (o *CreateInstallInputsForbidden) String() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsForbidden  %+v", 403, o.Payload)
}

func (o *CreateInstallInputsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateInstallInputsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstallInputsNotFound creates a CreateInstallInputsNotFound with default headers values
func NewCreateInstallInputsNotFound() *CreateInstallInputsNotFound {
	return &CreateInstallInputsNotFound{}
}

/*
CreateInstallInputsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateInstallInputsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create install inputs not found response has a 2xx status code
func (o *CreateInstallInputsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create install inputs not found response has a 3xx status code
func (o *CreateInstallInputsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create install inputs not found response has a 4xx status code
func (o *CreateInstallInputsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create install inputs not found response has a 5xx status code
func (o *CreateInstallInputsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create install inputs not found response a status code equal to that given
func (o *CreateInstallInputsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create install inputs not found response
func (o *CreateInstallInputsNotFound) Code() int {
	return 404
}

func (o *CreateInstallInputsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsNotFound  %+v", 404, o.Payload)
}

func (o *CreateInstallInputsNotFound) String() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsNotFound  %+v", 404, o.Payload)
}

func (o *CreateInstallInputsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateInstallInputsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstallInputsInternalServerError creates a CreateInstallInputsInternalServerError with default headers values
func NewCreateInstallInputsInternalServerError() *CreateInstallInputsInternalServerError {
	return &CreateInstallInputsInternalServerError{}
}

/*
CreateInstallInputsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateInstallInputsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create install inputs internal server error response has a 2xx status code
func (o *CreateInstallInputsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create install inputs internal server error response has a 3xx status code
func (o *CreateInstallInputsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create install inputs internal server error response has a 4xx status code
func (o *CreateInstallInputsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create install inputs internal server error response has a 5xx status code
func (o *CreateInstallInputsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create install inputs internal server error response a status code equal to that given
func (o *CreateInstallInputsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create install inputs internal server error response
func (o *CreateInstallInputsInternalServerError) Code() int {
	return 500
}

func (o *CreateInstallInputsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInstallInputsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/installs/{install_id}/inputs][%d] createInstallInputsInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInstallInputsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateInstallInputsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}