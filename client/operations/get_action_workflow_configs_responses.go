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

// GetActionWorkflowConfigsReader is a Reader for the GetActionWorkflowConfigs structure.
type GetActionWorkflowConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionWorkflowConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionWorkflowConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetActionWorkflowConfigsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetActionWorkflowConfigsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetActionWorkflowConfigsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetActionWorkflowConfigsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetActionWorkflowConfigsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/action-workflows/{action_workflow_id}/configs] GetActionWorkflowConfigs", response, response.Code())
	}
}

// NewGetActionWorkflowConfigsOK creates a GetActionWorkflowConfigsOK with default headers values
func NewGetActionWorkflowConfigsOK() *GetActionWorkflowConfigsOK {
	return &GetActionWorkflowConfigsOK{}
}

/*
GetActionWorkflowConfigsOK describes a response with status code 200, with default header values.

OK
*/
type GetActionWorkflowConfigsOK struct {
	Payload []*models.AppActionWorkflowConfig
}

// IsSuccess returns true when this get action workflow configs o k response has a 2xx status code
func (o *GetActionWorkflowConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get action workflow configs o k response has a 3xx status code
func (o *GetActionWorkflowConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action workflow configs o k response has a 4xx status code
func (o *GetActionWorkflowConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get action workflow configs o k response has a 5xx status code
func (o *GetActionWorkflowConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get action workflow configs o k response a status code equal to that given
func (o *GetActionWorkflowConfigsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get action workflow configs o k response
func (o *GetActionWorkflowConfigsOK) Code() int {
	return 200
}

func (o *GetActionWorkflowConfigsOK) Error() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsOK  %+v", 200, o.Payload)
}

func (o *GetActionWorkflowConfigsOK) String() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsOK  %+v", 200, o.Payload)
}

func (o *GetActionWorkflowConfigsOK) GetPayload() []*models.AppActionWorkflowConfig {
	return o.Payload
}

func (o *GetActionWorkflowConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionWorkflowConfigsBadRequest creates a GetActionWorkflowConfigsBadRequest with default headers values
func NewGetActionWorkflowConfigsBadRequest() *GetActionWorkflowConfigsBadRequest {
	return &GetActionWorkflowConfigsBadRequest{}
}

/*
GetActionWorkflowConfigsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetActionWorkflowConfigsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get action workflow configs bad request response has a 2xx status code
func (o *GetActionWorkflowConfigsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action workflow configs bad request response has a 3xx status code
func (o *GetActionWorkflowConfigsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action workflow configs bad request response has a 4xx status code
func (o *GetActionWorkflowConfigsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get action workflow configs bad request response has a 5xx status code
func (o *GetActionWorkflowConfigsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get action workflow configs bad request response a status code equal to that given
func (o *GetActionWorkflowConfigsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get action workflow configs bad request response
func (o *GetActionWorkflowConfigsBadRequest) Code() int {
	return 400
}

func (o *GetActionWorkflowConfigsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsBadRequest  %+v", 400, o.Payload)
}

func (o *GetActionWorkflowConfigsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsBadRequest  %+v", 400, o.Payload)
}

func (o *GetActionWorkflowConfigsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetActionWorkflowConfigsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionWorkflowConfigsUnauthorized creates a GetActionWorkflowConfigsUnauthorized with default headers values
func NewGetActionWorkflowConfigsUnauthorized() *GetActionWorkflowConfigsUnauthorized {
	return &GetActionWorkflowConfigsUnauthorized{}
}

/*
GetActionWorkflowConfigsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetActionWorkflowConfigsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get action workflow configs unauthorized response has a 2xx status code
func (o *GetActionWorkflowConfigsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action workflow configs unauthorized response has a 3xx status code
func (o *GetActionWorkflowConfigsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action workflow configs unauthorized response has a 4xx status code
func (o *GetActionWorkflowConfigsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get action workflow configs unauthorized response has a 5xx status code
func (o *GetActionWorkflowConfigsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get action workflow configs unauthorized response a status code equal to that given
func (o *GetActionWorkflowConfigsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get action workflow configs unauthorized response
func (o *GetActionWorkflowConfigsUnauthorized) Code() int {
	return 401
}

func (o *GetActionWorkflowConfigsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetActionWorkflowConfigsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetActionWorkflowConfigsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetActionWorkflowConfigsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionWorkflowConfigsForbidden creates a GetActionWorkflowConfigsForbidden with default headers values
func NewGetActionWorkflowConfigsForbidden() *GetActionWorkflowConfigsForbidden {
	return &GetActionWorkflowConfigsForbidden{}
}

/*
GetActionWorkflowConfigsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetActionWorkflowConfigsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get action workflow configs forbidden response has a 2xx status code
func (o *GetActionWorkflowConfigsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action workflow configs forbidden response has a 3xx status code
func (o *GetActionWorkflowConfigsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action workflow configs forbidden response has a 4xx status code
func (o *GetActionWorkflowConfigsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get action workflow configs forbidden response has a 5xx status code
func (o *GetActionWorkflowConfigsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get action workflow configs forbidden response a status code equal to that given
func (o *GetActionWorkflowConfigsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get action workflow configs forbidden response
func (o *GetActionWorkflowConfigsForbidden) Code() int {
	return 403
}

func (o *GetActionWorkflowConfigsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsForbidden  %+v", 403, o.Payload)
}

func (o *GetActionWorkflowConfigsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsForbidden  %+v", 403, o.Payload)
}

func (o *GetActionWorkflowConfigsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetActionWorkflowConfigsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionWorkflowConfigsNotFound creates a GetActionWorkflowConfigsNotFound with default headers values
func NewGetActionWorkflowConfigsNotFound() *GetActionWorkflowConfigsNotFound {
	return &GetActionWorkflowConfigsNotFound{}
}

/*
GetActionWorkflowConfigsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetActionWorkflowConfigsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get action workflow configs not found response has a 2xx status code
func (o *GetActionWorkflowConfigsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action workflow configs not found response has a 3xx status code
func (o *GetActionWorkflowConfigsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action workflow configs not found response has a 4xx status code
func (o *GetActionWorkflowConfigsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get action workflow configs not found response has a 5xx status code
func (o *GetActionWorkflowConfigsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get action workflow configs not found response a status code equal to that given
func (o *GetActionWorkflowConfigsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get action workflow configs not found response
func (o *GetActionWorkflowConfigsNotFound) Code() int {
	return 404
}

func (o *GetActionWorkflowConfigsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsNotFound  %+v", 404, o.Payload)
}

func (o *GetActionWorkflowConfigsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsNotFound  %+v", 404, o.Payload)
}

func (o *GetActionWorkflowConfigsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetActionWorkflowConfigsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionWorkflowConfigsInternalServerError creates a GetActionWorkflowConfigsInternalServerError with default headers values
func NewGetActionWorkflowConfigsInternalServerError() *GetActionWorkflowConfigsInternalServerError {
	return &GetActionWorkflowConfigsInternalServerError{}
}

/*
GetActionWorkflowConfigsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetActionWorkflowConfigsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get action workflow configs internal server error response has a 2xx status code
func (o *GetActionWorkflowConfigsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get action workflow configs internal server error response has a 3xx status code
func (o *GetActionWorkflowConfigsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get action workflow configs internal server error response has a 4xx status code
func (o *GetActionWorkflowConfigsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get action workflow configs internal server error response has a 5xx status code
func (o *GetActionWorkflowConfigsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get action workflow configs internal server error response a status code equal to that given
func (o *GetActionWorkflowConfigsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get action workflow configs internal server error response
func (o *GetActionWorkflowConfigsInternalServerError) Code() int {
	return 500
}

func (o *GetActionWorkflowConfigsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetActionWorkflowConfigsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/action-workflows/{action_workflow_id}/configs][%d] getActionWorkflowConfigsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetActionWorkflowConfigsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetActionWorkflowConfigsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
