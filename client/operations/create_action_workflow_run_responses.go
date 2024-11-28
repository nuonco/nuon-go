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

// CreateActionWorkflowRunReader is a Reader for the CreateActionWorkflowRun structure.
type CreateActionWorkflowRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateActionWorkflowRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateActionWorkflowRunCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateActionWorkflowRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateActionWorkflowRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateActionWorkflowRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateActionWorkflowRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateActionWorkflowRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/action-workflows/run] CreateActionWorkflowRun", response, response.Code())
	}
}

// NewCreateActionWorkflowRunCreated creates a CreateActionWorkflowRunCreated with default headers values
func NewCreateActionWorkflowRunCreated() *CreateActionWorkflowRunCreated {
	return &CreateActionWorkflowRunCreated{}
}

/*
CreateActionWorkflowRunCreated describes a response with status code 201, with default header values.

Created
*/
type CreateActionWorkflowRunCreated struct {
	Payload bool
}

// IsSuccess returns true when this create action workflow run created response has a 2xx status code
func (o *CreateActionWorkflowRunCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create action workflow run created response has a 3xx status code
func (o *CreateActionWorkflowRunCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow run created response has a 4xx status code
func (o *CreateActionWorkflowRunCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create action workflow run created response has a 5xx status code
func (o *CreateActionWorkflowRunCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow run created response a status code equal to that given
func (o *CreateActionWorkflowRunCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create action workflow run created response
func (o *CreateActionWorkflowRunCreated) Code() int {
	return 201
}

func (o *CreateActionWorkflowRunCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunCreated %s", 201, payload)
}

func (o *CreateActionWorkflowRunCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunCreated %s", 201, payload)
}

func (o *CreateActionWorkflowRunCreated) GetPayload() bool {
	return o.Payload
}

func (o *CreateActionWorkflowRunCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowRunBadRequest creates a CreateActionWorkflowRunBadRequest with default headers values
func NewCreateActionWorkflowRunBadRequest() *CreateActionWorkflowRunBadRequest {
	return &CreateActionWorkflowRunBadRequest{}
}

/*
CreateActionWorkflowRunBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateActionWorkflowRunBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow run bad request response has a 2xx status code
func (o *CreateActionWorkflowRunBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow run bad request response has a 3xx status code
func (o *CreateActionWorkflowRunBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow run bad request response has a 4xx status code
func (o *CreateActionWorkflowRunBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow run bad request response has a 5xx status code
func (o *CreateActionWorkflowRunBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow run bad request response a status code equal to that given
func (o *CreateActionWorkflowRunBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create action workflow run bad request response
func (o *CreateActionWorkflowRunBadRequest) Code() int {
	return 400
}

func (o *CreateActionWorkflowRunBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunBadRequest %s", 400, payload)
}

func (o *CreateActionWorkflowRunBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunBadRequest %s", 400, payload)
}

func (o *CreateActionWorkflowRunBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowRunUnauthorized creates a CreateActionWorkflowRunUnauthorized with default headers values
func NewCreateActionWorkflowRunUnauthorized() *CreateActionWorkflowRunUnauthorized {
	return &CreateActionWorkflowRunUnauthorized{}
}

/*
CreateActionWorkflowRunUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateActionWorkflowRunUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow run unauthorized response has a 2xx status code
func (o *CreateActionWorkflowRunUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow run unauthorized response has a 3xx status code
func (o *CreateActionWorkflowRunUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow run unauthorized response has a 4xx status code
func (o *CreateActionWorkflowRunUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow run unauthorized response has a 5xx status code
func (o *CreateActionWorkflowRunUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow run unauthorized response a status code equal to that given
func (o *CreateActionWorkflowRunUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create action workflow run unauthorized response
func (o *CreateActionWorkflowRunUnauthorized) Code() int {
	return 401
}

func (o *CreateActionWorkflowRunUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunUnauthorized %s", 401, payload)
}

func (o *CreateActionWorkflowRunUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunUnauthorized %s", 401, payload)
}

func (o *CreateActionWorkflowRunUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowRunForbidden creates a CreateActionWorkflowRunForbidden with default headers values
func NewCreateActionWorkflowRunForbidden() *CreateActionWorkflowRunForbidden {
	return &CreateActionWorkflowRunForbidden{}
}

/*
CreateActionWorkflowRunForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateActionWorkflowRunForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow run forbidden response has a 2xx status code
func (o *CreateActionWorkflowRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow run forbidden response has a 3xx status code
func (o *CreateActionWorkflowRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow run forbidden response has a 4xx status code
func (o *CreateActionWorkflowRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow run forbidden response has a 5xx status code
func (o *CreateActionWorkflowRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow run forbidden response a status code equal to that given
func (o *CreateActionWorkflowRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create action workflow run forbidden response
func (o *CreateActionWorkflowRunForbidden) Code() int {
	return 403
}

func (o *CreateActionWorkflowRunForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunForbidden %s", 403, payload)
}

func (o *CreateActionWorkflowRunForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunForbidden %s", 403, payload)
}

func (o *CreateActionWorkflowRunForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowRunNotFound creates a CreateActionWorkflowRunNotFound with default headers values
func NewCreateActionWorkflowRunNotFound() *CreateActionWorkflowRunNotFound {
	return &CreateActionWorkflowRunNotFound{}
}

/*
CreateActionWorkflowRunNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateActionWorkflowRunNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow run not found response has a 2xx status code
func (o *CreateActionWorkflowRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow run not found response has a 3xx status code
func (o *CreateActionWorkflowRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow run not found response has a 4xx status code
func (o *CreateActionWorkflowRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow run not found response has a 5xx status code
func (o *CreateActionWorkflowRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow run not found response a status code equal to that given
func (o *CreateActionWorkflowRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create action workflow run not found response
func (o *CreateActionWorkflowRunNotFound) Code() int {
	return 404
}

func (o *CreateActionWorkflowRunNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunNotFound %s", 404, payload)
}

func (o *CreateActionWorkflowRunNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunNotFound %s", 404, payload)
}

func (o *CreateActionWorkflowRunNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowRunInternalServerError creates a CreateActionWorkflowRunInternalServerError with default headers values
func NewCreateActionWorkflowRunInternalServerError() *CreateActionWorkflowRunInternalServerError {
	return &CreateActionWorkflowRunInternalServerError{}
}

/*
CreateActionWorkflowRunInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateActionWorkflowRunInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow run internal server error response has a 2xx status code
func (o *CreateActionWorkflowRunInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow run internal server error response has a 3xx status code
func (o *CreateActionWorkflowRunInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow run internal server error response has a 4xx status code
func (o *CreateActionWorkflowRunInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create action workflow run internal server error response has a 5xx status code
func (o *CreateActionWorkflowRunInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create action workflow run internal server error response a status code equal to that given
func (o *CreateActionWorkflowRunInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create action workflow run internal server error response
func (o *CreateActionWorkflowRunInternalServerError) Code() int {
	return 500
}

func (o *CreateActionWorkflowRunInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunInternalServerError %s", 500, payload)
}

func (o *CreateActionWorkflowRunInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/run][%d] createActionWorkflowRunInternalServerError %s", 500, payload)
}

func (o *CreateActionWorkflowRunInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
