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

// UpdateAppActionWorkflowReader is a Reader for the UpdateAppActionWorkflow structure.
type UpdateAppActionWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppActionWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAppActionWorkflowCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAppActionWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAppActionWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAppActionWorkflowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAppActionWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAppActionWorkflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/action-workflows/{action_workflow_id}] UpdateAppActionWorkflow", response, response.Code())
	}
}

// NewUpdateAppActionWorkflowCreated creates a UpdateAppActionWorkflowCreated with default headers values
func NewUpdateAppActionWorkflowCreated() *UpdateAppActionWorkflowCreated {
	return &UpdateAppActionWorkflowCreated{}
}

/*
UpdateAppActionWorkflowCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateAppActionWorkflowCreated struct {
	Payload *models.AppActionWorkflow
}

// IsSuccess returns true when this update app action workflow created response has a 2xx status code
func (o *UpdateAppActionWorkflowCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update app action workflow created response has a 3xx status code
func (o *UpdateAppActionWorkflowCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app action workflow created response has a 4xx status code
func (o *UpdateAppActionWorkflowCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update app action workflow created response has a 5xx status code
func (o *UpdateAppActionWorkflowCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update app action workflow created response a status code equal to that given
func (o *UpdateAppActionWorkflowCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update app action workflow created response
func (o *UpdateAppActionWorkflowCreated) Code() int {
	return 201
}

func (o *UpdateAppActionWorkflowCreated) Error() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowCreated  %+v", 201, o.Payload)
}

func (o *UpdateAppActionWorkflowCreated) String() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowCreated  %+v", 201, o.Payload)
}

func (o *UpdateAppActionWorkflowCreated) GetPayload() *models.AppActionWorkflow {
	return o.Payload
}

func (o *UpdateAppActionWorkflowCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppActionWorkflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppActionWorkflowBadRequest creates a UpdateAppActionWorkflowBadRequest with default headers values
func NewUpdateAppActionWorkflowBadRequest() *UpdateAppActionWorkflowBadRequest {
	return &UpdateAppActionWorkflowBadRequest{}
}

/*
UpdateAppActionWorkflowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAppActionWorkflowBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app action workflow bad request response has a 2xx status code
func (o *UpdateAppActionWorkflowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app action workflow bad request response has a 3xx status code
func (o *UpdateAppActionWorkflowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app action workflow bad request response has a 4xx status code
func (o *UpdateAppActionWorkflowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app action workflow bad request response has a 5xx status code
func (o *UpdateAppActionWorkflowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update app action workflow bad request response a status code equal to that given
func (o *UpdateAppActionWorkflowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update app action workflow bad request response
func (o *UpdateAppActionWorkflowBadRequest) Code() int {
	return 400
}

func (o *UpdateAppActionWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAppActionWorkflowBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAppActionWorkflowBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppActionWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppActionWorkflowUnauthorized creates a UpdateAppActionWorkflowUnauthorized with default headers values
func NewUpdateAppActionWorkflowUnauthorized() *UpdateAppActionWorkflowUnauthorized {
	return &UpdateAppActionWorkflowUnauthorized{}
}

/*
UpdateAppActionWorkflowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateAppActionWorkflowUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app action workflow unauthorized response has a 2xx status code
func (o *UpdateAppActionWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app action workflow unauthorized response has a 3xx status code
func (o *UpdateAppActionWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app action workflow unauthorized response has a 4xx status code
func (o *UpdateAppActionWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app action workflow unauthorized response has a 5xx status code
func (o *UpdateAppActionWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update app action workflow unauthorized response a status code equal to that given
func (o *UpdateAppActionWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update app action workflow unauthorized response
func (o *UpdateAppActionWorkflowUnauthorized) Code() int {
	return 401
}

func (o *UpdateAppActionWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAppActionWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAppActionWorkflowUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppActionWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppActionWorkflowForbidden creates a UpdateAppActionWorkflowForbidden with default headers values
func NewUpdateAppActionWorkflowForbidden() *UpdateAppActionWorkflowForbidden {
	return &UpdateAppActionWorkflowForbidden{}
}

/*
UpdateAppActionWorkflowForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateAppActionWorkflowForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app action workflow forbidden response has a 2xx status code
func (o *UpdateAppActionWorkflowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app action workflow forbidden response has a 3xx status code
func (o *UpdateAppActionWorkflowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app action workflow forbidden response has a 4xx status code
func (o *UpdateAppActionWorkflowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app action workflow forbidden response has a 5xx status code
func (o *UpdateAppActionWorkflowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update app action workflow forbidden response a status code equal to that given
func (o *UpdateAppActionWorkflowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update app action workflow forbidden response
func (o *UpdateAppActionWorkflowForbidden) Code() int {
	return 403
}

func (o *UpdateAppActionWorkflowForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAppActionWorkflowForbidden) String() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAppActionWorkflowForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppActionWorkflowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppActionWorkflowNotFound creates a UpdateAppActionWorkflowNotFound with default headers values
func NewUpdateAppActionWorkflowNotFound() *UpdateAppActionWorkflowNotFound {
	return &UpdateAppActionWorkflowNotFound{}
}

/*
UpdateAppActionWorkflowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAppActionWorkflowNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app action workflow not found response has a 2xx status code
func (o *UpdateAppActionWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app action workflow not found response has a 3xx status code
func (o *UpdateAppActionWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app action workflow not found response has a 4xx status code
func (o *UpdateAppActionWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app action workflow not found response has a 5xx status code
func (o *UpdateAppActionWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update app action workflow not found response a status code equal to that given
func (o *UpdateAppActionWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update app action workflow not found response
func (o *UpdateAppActionWorkflowNotFound) Code() int {
	return 404
}

func (o *UpdateAppActionWorkflowNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAppActionWorkflowNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAppActionWorkflowNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppActionWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppActionWorkflowInternalServerError creates a UpdateAppActionWorkflowInternalServerError with default headers values
func NewUpdateAppActionWorkflowInternalServerError() *UpdateAppActionWorkflowInternalServerError {
	return &UpdateAppActionWorkflowInternalServerError{}
}

/*
UpdateAppActionWorkflowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateAppActionWorkflowInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app action workflow internal server error response has a 2xx status code
func (o *UpdateAppActionWorkflowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app action workflow internal server error response has a 3xx status code
func (o *UpdateAppActionWorkflowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app action workflow internal server error response has a 4xx status code
func (o *UpdateAppActionWorkflowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update app action workflow internal server error response has a 5xx status code
func (o *UpdateAppActionWorkflowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update app action workflow internal server error response a status code equal to that given
func (o *UpdateAppActionWorkflowInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update app action workflow internal server error response
func (o *UpdateAppActionWorkflowInternalServerError) Code() int {
	return 500
}

func (o *UpdateAppActionWorkflowInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAppActionWorkflowInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/action-workflows/{action_workflow_id}][%d] updateAppActionWorkflowInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateAppActionWorkflowInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppActionWorkflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
