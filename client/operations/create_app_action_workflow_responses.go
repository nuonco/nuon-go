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

// CreateAppActionWorkflowReader is a Reader for the CreateAppActionWorkflow structure.
type CreateAppActionWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppActionWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppActionWorkflowCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppActionWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppActionWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppActionWorkflowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppActionWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppActionWorkflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/action-workflows] CreateAppActionWorkflow", response, response.Code())
	}
}

// NewCreateAppActionWorkflowCreated creates a CreateAppActionWorkflowCreated with default headers values
func NewCreateAppActionWorkflowCreated() *CreateAppActionWorkflowCreated {
	return &CreateAppActionWorkflowCreated{}
}

/*
CreateAppActionWorkflowCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppActionWorkflowCreated struct {
	Payload *models.AppActionWorkflow
}

// IsSuccess returns true when this create app action workflow created response has a 2xx status code
func (o *CreateAppActionWorkflowCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app action workflow created response has a 3xx status code
func (o *CreateAppActionWorkflowCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app action workflow created response has a 4xx status code
func (o *CreateAppActionWorkflowCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app action workflow created response has a 5xx status code
func (o *CreateAppActionWorkflowCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app action workflow created response a status code equal to that given
func (o *CreateAppActionWorkflowCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app action workflow created response
func (o *CreateAppActionWorkflowCreated) Code() int {
	return 201
}

func (o *CreateAppActionWorkflowCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowCreated %s", 201, payload)
}

func (o *CreateAppActionWorkflowCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowCreated %s", 201, payload)
}

func (o *CreateAppActionWorkflowCreated) GetPayload() *models.AppActionWorkflow {
	return o.Payload
}

func (o *CreateAppActionWorkflowCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppActionWorkflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppActionWorkflowBadRequest creates a CreateAppActionWorkflowBadRequest with default headers values
func NewCreateAppActionWorkflowBadRequest() *CreateAppActionWorkflowBadRequest {
	return &CreateAppActionWorkflowBadRequest{}
}

/*
CreateAppActionWorkflowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppActionWorkflowBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app action workflow bad request response has a 2xx status code
func (o *CreateAppActionWorkflowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app action workflow bad request response has a 3xx status code
func (o *CreateAppActionWorkflowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app action workflow bad request response has a 4xx status code
func (o *CreateAppActionWorkflowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app action workflow bad request response has a 5xx status code
func (o *CreateAppActionWorkflowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app action workflow bad request response a status code equal to that given
func (o *CreateAppActionWorkflowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app action workflow bad request response
func (o *CreateAppActionWorkflowBadRequest) Code() int {
	return 400
}

func (o *CreateAppActionWorkflowBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowBadRequest %s", 400, payload)
}

func (o *CreateAppActionWorkflowBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowBadRequest %s", 400, payload)
}

func (o *CreateAppActionWorkflowBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppActionWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppActionWorkflowUnauthorized creates a CreateAppActionWorkflowUnauthorized with default headers values
func NewCreateAppActionWorkflowUnauthorized() *CreateAppActionWorkflowUnauthorized {
	return &CreateAppActionWorkflowUnauthorized{}
}

/*
CreateAppActionWorkflowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppActionWorkflowUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app action workflow unauthorized response has a 2xx status code
func (o *CreateAppActionWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app action workflow unauthorized response has a 3xx status code
func (o *CreateAppActionWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app action workflow unauthorized response has a 4xx status code
func (o *CreateAppActionWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app action workflow unauthorized response has a 5xx status code
func (o *CreateAppActionWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app action workflow unauthorized response a status code equal to that given
func (o *CreateAppActionWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app action workflow unauthorized response
func (o *CreateAppActionWorkflowUnauthorized) Code() int {
	return 401
}

func (o *CreateAppActionWorkflowUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowUnauthorized %s", 401, payload)
}

func (o *CreateAppActionWorkflowUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowUnauthorized %s", 401, payload)
}

func (o *CreateAppActionWorkflowUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppActionWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppActionWorkflowForbidden creates a CreateAppActionWorkflowForbidden with default headers values
func NewCreateAppActionWorkflowForbidden() *CreateAppActionWorkflowForbidden {
	return &CreateAppActionWorkflowForbidden{}
}

/*
CreateAppActionWorkflowForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppActionWorkflowForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app action workflow forbidden response has a 2xx status code
func (o *CreateAppActionWorkflowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app action workflow forbidden response has a 3xx status code
func (o *CreateAppActionWorkflowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app action workflow forbidden response has a 4xx status code
func (o *CreateAppActionWorkflowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app action workflow forbidden response has a 5xx status code
func (o *CreateAppActionWorkflowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app action workflow forbidden response a status code equal to that given
func (o *CreateAppActionWorkflowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app action workflow forbidden response
func (o *CreateAppActionWorkflowForbidden) Code() int {
	return 403
}

func (o *CreateAppActionWorkflowForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowForbidden %s", 403, payload)
}

func (o *CreateAppActionWorkflowForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowForbidden %s", 403, payload)
}

func (o *CreateAppActionWorkflowForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppActionWorkflowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppActionWorkflowNotFound creates a CreateAppActionWorkflowNotFound with default headers values
func NewCreateAppActionWorkflowNotFound() *CreateAppActionWorkflowNotFound {
	return &CreateAppActionWorkflowNotFound{}
}

/*
CreateAppActionWorkflowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppActionWorkflowNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app action workflow not found response has a 2xx status code
func (o *CreateAppActionWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app action workflow not found response has a 3xx status code
func (o *CreateAppActionWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app action workflow not found response has a 4xx status code
func (o *CreateAppActionWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app action workflow not found response has a 5xx status code
func (o *CreateAppActionWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app action workflow not found response a status code equal to that given
func (o *CreateAppActionWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app action workflow not found response
func (o *CreateAppActionWorkflowNotFound) Code() int {
	return 404
}

func (o *CreateAppActionWorkflowNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowNotFound %s", 404, payload)
}

func (o *CreateAppActionWorkflowNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowNotFound %s", 404, payload)
}

func (o *CreateAppActionWorkflowNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppActionWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppActionWorkflowInternalServerError creates a CreateAppActionWorkflowInternalServerError with default headers values
func NewCreateAppActionWorkflowInternalServerError() *CreateAppActionWorkflowInternalServerError {
	return &CreateAppActionWorkflowInternalServerError{}
}

/*
CreateAppActionWorkflowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppActionWorkflowInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app action workflow internal server error response has a 2xx status code
func (o *CreateAppActionWorkflowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app action workflow internal server error response has a 3xx status code
func (o *CreateAppActionWorkflowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app action workflow internal server error response has a 4xx status code
func (o *CreateAppActionWorkflowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app action workflow internal server error response has a 5xx status code
func (o *CreateAppActionWorkflowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app action workflow internal server error response a status code equal to that given
func (o *CreateAppActionWorkflowInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app action workflow internal server error response
func (o *CreateAppActionWorkflowInternalServerError) Code() int {
	return 500
}

func (o *CreateAppActionWorkflowInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowInternalServerError %s", 500, payload)
}

func (o *CreateAppActionWorkflowInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/action-workflows][%d] createAppActionWorkflowInternalServerError %s", 500, payload)
}

func (o *CreateAppActionWorkflowInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppActionWorkflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
