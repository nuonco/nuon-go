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

// GetAppActionWorkflowReader is a Reader for the GetAppActionWorkflow structure.
type GetAppActionWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppActionWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppActionWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppActionWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppActionWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppActionWorkflowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppActionWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppActionWorkflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}] GetAppActionWorkflow", response, response.Code())
	}
}

// NewGetAppActionWorkflowOK creates a GetAppActionWorkflowOK with default headers values
func NewGetAppActionWorkflowOK() *GetAppActionWorkflowOK {
	return &GetAppActionWorkflowOK{}
}

/*
GetAppActionWorkflowOK describes a response with status code 200, with default header values.

OK
*/
type GetAppActionWorkflowOK struct {
	Payload *models.AppActionWorkflow
}

// IsSuccess returns true when this get app action workflow o k response has a 2xx status code
func (o *GetAppActionWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app action workflow o k response has a 3xx status code
func (o *GetAppActionWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app action workflow o k response has a 4xx status code
func (o *GetAppActionWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app action workflow o k response has a 5xx status code
func (o *GetAppActionWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app action workflow o k response a status code equal to that given
func (o *GetAppActionWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app action workflow o k response
func (o *GetAppActionWorkflowOK) Code() int {
	return 200
}

func (o *GetAppActionWorkflowOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowOK %s", 200, payload)
}

func (o *GetAppActionWorkflowOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowOK %s", 200, payload)
}

func (o *GetAppActionWorkflowOK) GetPayload() *models.AppActionWorkflow {
	return o.Payload
}

func (o *GetAppActionWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppActionWorkflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppActionWorkflowBadRequest creates a GetAppActionWorkflowBadRequest with default headers values
func NewGetAppActionWorkflowBadRequest() *GetAppActionWorkflowBadRequest {
	return &GetAppActionWorkflowBadRequest{}
}

/*
GetAppActionWorkflowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppActionWorkflowBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app action workflow bad request response has a 2xx status code
func (o *GetAppActionWorkflowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app action workflow bad request response has a 3xx status code
func (o *GetAppActionWorkflowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app action workflow bad request response has a 4xx status code
func (o *GetAppActionWorkflowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app action workflow bad request response has a 5xx status code
func (o *GetAppActionWorkflowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app action workflow bad request response a status code equal to that given
func (o *GetAppActionWorkflowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app action workflow bad request response
func (o *GetAppActionWorkflowBadRequest) Code() int {
	return 400
}

func (o *GetAppActionWorkflowBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowBadRequest %s", 400, payload)
}

func (o *GetAppActionWorkflowBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowBadRequest %s", 400, payload)
}

func (o *GetAppActionWorkflowBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppActionWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppActionWorkflowUnauthorized creates a GetAppActionWorkflowUnauthorized with default headers values
func NewGetAppActionWorkflowUnauthorized() *GetAppActionWorkflowUnauthorized {
	return &GetAppActionWorkflowUnauthorized{}
}

/*
GetAppActionWorkflowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppActionWorkflowUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app action workflow unauthorized response has a 2xx status code
func (o *GetAppActionWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app action workflow unauthorized response has a 3xx status code
func (o *GetAppActionWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app action workflow unauthorized response has a 4xx status code
func (o *GetAppActionWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app action workflow unauthorized response has a 5xx status code
func (o *GetAppActionWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app action workflow unauthorized response a status code equal to that given
func (o *GetAppActionWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app action workflow unauthorized response
func (o *GetAppActionWorkflowUnauthorized) Code() int {
	return 401
}

func (o *GetAppActionWorkflowUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowUnauthorized %s", 401, payload)
}

func (o *GetAppActionWorkflowUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowUnauthorized %s", 401, payload)
}

func (o *GetAppActionWorkflowUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppActionWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppActionWorkflowForbidden creates a GetAppActionWorkflowForbidden with default headers values
func NewGetAppActionWorkflowForbidden() *GetAppActionWorkflowForbidden {
	return &GetAppActionWorkflowForbidden{}
}

/*
GetAppActionWorkflowForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppActionWorkflowForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app action workflow forbidden response has a 2xx status code
func (o *GetAppActionWorkflowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app action workflow forbidden response has a 3xx status code
func (o *GetAppActionWorkflowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app action workflow forbidden response has a 4xx status code
func (o *GetAppActionWorkflowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app action workflow forbidden response has a 5xx status code
func (o *GetAppActionWorkflowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app action workflow forbidden response a status code equal to that given
func (o *GetAppActionWorkflowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app action workflow forbidden response
func (o *GetAppActionWorkflowForbidden) Code() int {
	return 403
}

func (o *GetAppActionWorkflowForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowForbidden %s", 403, payload)
}

func (o *GetAppActionWorkflowForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowForbidden %s", 403, payload)
}

func (o *GetAppActionWorkflowForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppActionWorkflowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppActionWorkflowNotFound creates a GetAppActionWorkflowNotFound with default headers values
func NewGetAppActionWorkflowNotFound() *GetAppActionWorkflowNotFound {
	return &GetAppActionWorkflowNotFound{}
}

/*
GetAppActionWorkflowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppActionWorkflowNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app action workflow not found response has a 2xx status code
func (o *GetAppActionWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app action workflow not found response has a 3xx status code
func (o *GetAppActionWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app action workflow not found response has a 4xx status code
func (o *GetAppActionWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app action workflow not found response has a 5xx status code
func (o *GetAppActionWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app action workflow not found response a status code equal to that given
func (o *GetAppActionWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app action workflow not found response
func (o *GetAppActionWorkflowNotFound) Code() int {
	return 404
}

func (o *GetAppActionWorkflowNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowNotFound %s", 404, payload)
}

func (o *GetAppActionWorkflowNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowNotFound %s", 404, payload)
}

func (o *GetAppActionWorkflowNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppActionWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppActionWorkflowInternalServerError creates a GetAppActionWorkflowInternalServerError with default headers values
func NewGetAppActionWorkflowInternalServerError() *GetAppActionWorkflowInternalServerError {
	return &GetAppActionWorkflowInternalServerError{}
}

/*
GetAppActionWorkflowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppActionWorkflowInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app action workflow internal server error response has a 2xx status code
func (o *GetAppActionWorkflowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app action workflow internal server error response has a 3xx status code
func (o *GetAppActionWorkflowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app action workflow internal server error response has a 4xx status code
func (o *GetAppActionWorkflowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app action workflow internal server error response has a 5xx status code
func (o *GetAppActionWorkflowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app action workflow internal server error response a status code equal to that given
func (o *GetAppActionWorkflowInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app action workflow internal server error response
func (o *GetAppActionWorkflowInternalServerError) Code() int {
	return 500
}

func (o *GetAppActionWorkflowInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowInternalServerError %s", 500, payload)
}

func (o *GetAppActionWorkflowInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/action-workflows/{action_workflow_id}][%d] getAppActionWorkflowInternalServerError %s", 500, payload)
}

func (o *GetAppActionWorkflowInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppActionWorkflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
