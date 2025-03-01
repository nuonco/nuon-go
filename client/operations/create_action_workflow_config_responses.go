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

// CreateActionWorkflowConfigReader is a Reader for the CreateActionWorkflowConfig structure.
type CreateActionWorkflowConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateActionWorkflowConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateActionWorkflowConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateActionWorkflowConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateActionWorkflowConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateActionWorkflowConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateActionWorkflowConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateActionWorkflowConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/action-workflows/{action_workflow_id}/configs] CreateActionWorkflowConfig", response, response.Code())
	}
}

// NewCreateActionWorkflowConfigCreated creates a CreateActionWorkflowConfigCreated with default headers values
func NewCreateActionWorkflowConfigCreated() *CreateActionWorkflowConfigCreated {
	return &CreateActionWorkflowConfigCreated{}
}

/*
CreateActionWorkflowConfigCreated describes a response with status code 201, with default header values.

Created
*/
type CreateActionWorkflowConfigCreated struct {
	Payload *models.AppActionWorkflowConfig
}

// IsSuccess returns true when this create action workflow config created response has a 2xx status code
func (o *CreateActionWorkflowConfigCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create action workflow config created response has a 3xx status code
func (o *CreateActionWorkflowConfigCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow config created response has a 4xx status code
func (o *CreateActionWorkflowConfigCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create action workflow config created response has a 5xx status code
func (o *CreateActionWorkflowConfigCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow config created response a status code equal to that given
func (o *CreateActionWorkflowConfigCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create action workflow config created response
func (o *CreateActionWorkflowConfigCreated) Code() int {
	return 201
}

func (o *CreateActionWorkflowConfigCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigCreated %s", 201, payload)
}

func (o *CreateActionWorkflowConfigCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigCreated %s", 201, payload)
}

func (o *CreateActionWorkflowConfigCreated) GetPayload() *models.AppActionWorkflowConfig {
	return o.Payload
}

func (o *CreateActionWorkflowConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppActionWorkflowConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowConfigBadRequest creates a CreateActionWorkflowConfigBadRequest with default headers values
func NewCreateActionWorkflowConfigBadRequest() *CreateActionWorkflowConfigBadRequest {
	return &CreateActionWorkflowConfigBadRequest{}
}

/*
CreateActionWorkflowConfigBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateActionWorkflowConfigBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow config bad request response has a 2xx status code
func (o *CreateActionWorkflowConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow config bad request response has a 3xx status code
func (o *CreateActionWorkflowConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow config bad request response has a 4xx status code
func (o *CreateActionWorkflowConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow config bad request response has a 5xx status code
func (o *CreateActionWorkflowConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow config bad request response a status code equal to that given
func (o *CreateActionWorkflowConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create action workflow config bad request response
func (o *CreateActionWorkflowConfigBadRequest) Code() int {
	return 400
}

func (o *CreateActionWorkflowConfigBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigBadRequest %s", 400, payload)
}

func (o *CreateActionWorkflowConfigBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigBadRequest %s", 400, payload)
}

func (o *CreateActionWorkflowConfigBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowConfigUnauthorized creates a CreateActionWorkflowConfigUnauthorized with default headers values
func NewCreateActionWorkflowConfigUnauthorized() *CreateActionWorkflowConfigUnauthorized {
	return &CreateActionWorkflowConfigUnauthorized{}
}

/*
CreateActionWorkflowConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateActionWorkflowConfigUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow config unauthorized response has a 2xx status code
func (o *CreateActionWorkflowConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow config unauthorized response has a 3xx status code
func (o *CreateActionWorkflowConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow config unauthorized response has a 4xx status code
func (o *CreateActionWorkflowConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow config unauthorized response has a 5xx status code
func (o *CreateActionWorkflowConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow config unauthorized response a status code equal to that given
func (o *CreateActionWorkflowConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create action workflow config unauthorized response
func (o *CreateActionWorkflowConfigUnauthorized) Code() int {
	return 401
}

func (o *CreateActionWorkflowConfigUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigUnauthorized %s", 401, payload)
}

func (o *CreateActionWorkflowConfigUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigUnauthorized %s", 401, payload)
}

func (o *CreateActionWorkflowConfigUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowConfigForbidden creates a CreateActionWorkflowConfigForbidden with default headers values
func NewCreateActionWorkflowConfigForbidden() *CreateActionWorkflowConfigForbidden {
	return &CreateActionWorkflowConfigForbidden{}
}

/*
CreateActionWorkflowConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateActionWorkflowConfigForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow config forbidden response has a 2xx status code
func (o *CreateActionWorkflowConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow config forbidden response has a 3xx status code
func (o *CreateActionWorkflowConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow config forbidden response has a 4xx status code
func (o *CreateActionWorkflowConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow config forbidden response has a 5xx status code
func (o *CreateActionWorkflowConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow config forbidden response a status code equal to that given
func (o *CreateActionWorkflowConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create action workflow config forbidden response
func (o *CreateActionWorkflowConfigForbidden) Code() int {
	return 403
}

func (o *CreateActionWorkflowConfigForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigForbidden %s", 403, payload)
}

func (o *CreateActionWorkflowConfigForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigForbidden %s", 403, payload)
}

func (o *CreateActionWorkflowConfigForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowConfigNotFound creates a CreateActionWorkflowConfigNotFound with default headers values
func NewCreateActionWorkflowConfigNotFound() *CreateActionWorkflowConfigNotFound {
	return &CreateActionWorkflowConfigNotFound{}
}

/*
CreateActionWorkflowConfigNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateActionWorkflowConfigNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow config not found response has a 2xx status code
func (o *CreateActionWorkflowConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow config not found response has a 3xx status code
func (o *CreateActionWorkflowConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow config not found response has a 4xx status code
func (o *CreateActionWorkflowConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create action workflow config not found response has a 5xx status code
func (o *CreateActionWorkflowConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create action workflow config not found response a status code equal to that given
func (o *CreateActionWorkflowConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create action workflow config not found response
func (o *CreateActionWorkflowConfigNotFound) Code() int {
	return 404
}

func (o *CreateActionWorkflowConfigNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigNotFound %s", 404, payload)
}

func (o *CreateActionWorkflowConfigNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigNotFound %s", 404, payload)
}

func (o *CreateActionWorkflowConfigNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActionWorkflowConfigInternalServerError creates a CreateActionWorkflowConfigInternalServerError with default headers values
func NewCreateActionWorkflowConfigInternalServerError() *CreateActionWorkflowConfigInternalServerError {
	return &CreateActionWorkflowConfigInternalServerError{}
}

/*
CreateActionWorkflowConfigInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateActionWorkflowConfigInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create action workflow config internal server error response has a 2xx status code
func (o *CreateActionWorkflowConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create action workflow config internal server error response has a 3xx status code
func (o *CreateActionWorkflowConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create action workflow config internal server error response has a 4xx status code
func (o *CreateActionWorkflowConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create action workflow config internal server error response has a 5xx status code
func (o *CreateActionWorkflowConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create action workflow config internal server error response a status code equal to that given
func (o *CreateActionWorkflowConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create action workflow config internal server error response
func (o *CreateActionWorkflowConfigInternalServerError) Code() int {
	return 500
}

func (o *CreateActionWorkflowConfigInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigInternalServerError %s", 500, payload)
}

func (o *CreateActionWorkflowConfigInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/action-workflows/{action_workflow_id}/configs][%d] createActionWorkflowConfigInternalServerError %s", 500, payload)
}

func (o *CreateActionWorkflowConfigInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateActionWorkflowConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
