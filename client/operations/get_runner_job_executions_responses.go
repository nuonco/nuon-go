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

// GetRunnerJobExecutionsReader is a Reader for the GetRunnerJobExecutions structure.
type GetRunnerJobExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunnerJobExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunnerJobExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunnerJobExecutionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRunnerJobExecutionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRunnerJobExecutionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunnerJobExecutionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRunnerJobExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/runner-jobs/{runner_job_id}/executions] GetRunnerJobExecutions", response, response.Code())
	}
}

// NewGetRunnerJobExecutionsOK creates a GetRunnerJobExecutionsOK with default headers values
func NewGetRunnerJobExecutionsOK() *GetRunnerJobExecutionsOK {
	return &GetRunnerJobExecutionsOK{}
}

/*
GetRunnerJobExecutionsOK describes a response with status code 200, with default header values.

OK
*/
type GetRunnerJobExecutionsOK struct {
	Payload []*models.AppRunnerJobExecution
}

// IsSuccess returns true when this get runner job executions o k response has a 2xx status code
func (o *GetRunnerJobExecutionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runner job executions o k response has a 3xx status code
func (o *GetRunnerJobExecutionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job executions o k response has a 4xx status code
func (o *GetRunnerJobExecutionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner job executions o k response has a 5xx status code
func (o *GetRunnerJobExecutionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job executions o k response a status code equal to that given
func (o *GetRunnerJobExecutionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runner job executions o k response
func (o *GetRunnerJobExecutionsOK) Code() int {
	return 200
}

func (o *GetRunnerJobExecutionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsOK %s", 200, payload)
}

func (o *GetRunnerJobExecutionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsOK %s", 200, payload)
}

func (o *GetRunnerJobExecutionsOK) GetPayload() []*models.AppRunnerJobExecution {
	return o.Payload
}

func (o *GetRunnerJobExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobExecutionsBadRequest creates a GetRunnerJobExecutionsBadRequest with default headers values
func NewGetRunnerJobExecutionsBadRequest() *GetRunnerJobExecutionsBadRequest {
	return &GetRunnerJobExecutionsBadRequest{}
}

/*
GetRunnerJobExecutionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRunnerJobExecutionsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job executions bad request response has a 2xx status code
func (o *GetRunnerJobExecutionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job executions bad request response has a 3xx status code
func (o *GetRunnerJobExecutionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job executions bad request response has a 4xx status code
func (o *GetRunnerJobExecutionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job executions bad request response has a 5xx status code
func (o *GetRunnerJobExecutionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job executions bad request response a status code equal to that given
func (o *GetRunnerJobExecutionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get runner job executions bad request response
func (o *GetRunnerJobExecutionsBadRequest) Code() int {
	return 400
}

func (o *GetRunnerJobExecutionsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsBadRequest %s", 400, payload)
}

func (o *GetRunnerJobExecutionsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsBadRequest %s", 400, payload)
}

func (o *GetRunnerJobExecutionsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobExecutionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobExecutionsUnauthorized creates a GetRunnerJobExecutionsUnauthorized with default headers values
func NewGetRunnerJobExecutionsUnauthorized() *GetRunnerJobExecutionsUnauthorized {
	return &GetRunnerJobExecutionsUnauthorized{}
}

/*
GetRunnerJobExecutionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRunnerJobExecutionsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job executions unauthorized response has a 2xx status code
func (o *GetRunnerJobExecutionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job executions unauthorized response has a 3xx status code
func (o *GetRunnerJobExecutionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job executions unauthorized response has a 4xx status code
func (o *GetRunnerJobExecutionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job executions unauthorized response has a 5xx status code
func (o *GetRunnerJobExecutionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job executions unauthorized response a status code equal to that given
func (o *GetRunnerJobExecutionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get runner job executions unauthorized response
func (o *GetRunnerJobExecutionsUnauthorized) Code() int {
	return 401
}

func (o *GetRunnerJobExecutionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsUnauthorized %s", 401, payload)
}

func (o *GetRunnerJobExecutionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsUnauthorized %s", 401, payload)
}

func (o *GetRunnerJobExecutionsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobExecutionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobExecutionsForbidden creates a GetRunnerJobExecutionsForbidden with default headers values
func NewGetRunnerJobExecutionsForbidden() *GetRunnerJobExecutionsForbidden {
	return &GetRunnerJobExecutionsForbidden{}
}

/*
GetRunnerJobExecutionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRunnerJobExecutionsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job executions forbidden response has a 2xx status code
func (o *GetRunnerJobExecutionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job executions forbidden response has a 3xx status code
func (o *GetRunnerJobExecutionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job executions forbidden response has a 4xx status code
func (o *GetRunnerJobExecutionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job executions forbidden response has a 5xx status code
func (o *GetRunnerJobExecutionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job executions forbidden response a status code equal to that given
func (o *GetRunnerJobExecutionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get runner job executions forbidden response
func (o *GetRunnerJobExecutionsForbidden) Code() int {
	return 403
}

func (o *GetRunnerJobExecutionsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsForbidden %s", 403, payload)
}

func (o *GetRunnerJobExecutionsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsForbidden %s", 403, payload)
}

func (o *GetRunnerJobExecutionsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobExecutionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobExecutionsNotFound creates a GetRunnerJobExecutionsNotFound with default headers values
func NewGetRunnerJobExecutionsNotFound() *GetRunnerJobExecutionsNotFound {
	return &GetRunnerJobExecutionsNotFound{}
}

/*
GetRunnerJobExecutionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRunnerJobExecutionsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job executions not found response has a 2xx status code
func (o *GetRunnerJobExecutionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job executions not found response has a 3xx status code
func (o *GetRunnerJobExecutionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job executions not found response has a 4xx status code
func (o *GetRunnerJobExecutionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job executions not found response has a 5xx status code
func (o *GetRunnerJobExecutionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job executions not found response a status code equal to that given
func (o *GetRunnerJobExecutionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get runner job executions not found response
func (o *GetRunnerJobExecutionsNotFound) Code() int {
	return 404
}

func (o *GetRunnerJobExecutionsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsNotFound %s", 404, payload)
}

func (o *GetRunnerJobExecutionsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsNotFound %s", 404, payload)
}

func (o *GetRunnerJobExecutionsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobExecutionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobExecutionsInternalServerError creates a GetRunnerJobExecutionsInternalServerError with default headers values
func NewGetRunnerJobExecutionsInternalServerError() *GetRunnerJobExecutionsInternalServerError {
	return &GetRunnerJobExecutionsInternalServerError{}
}

/*
GetRunnerJobExecutionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRunnerJobExecutionsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job executions internal server error response has a 2xx status code
func (o *GetRunnerJobExecutionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job executions internal server error response has a 3xx status code
func (o *GetRunnerJobExecutionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job executions internal server error response has a 4xx status code
func (o *GetRunnerJobExecutionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner job executions internal server error response has a 5xx status code
func (o *GetRunnerJobExecutionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get runner job executions internal server error response a status code equal to that given
func (o *GetRunnerJobExecutionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get runner job executions internal server error response
func (o *GetRunnerJobExecutionsInternalServerError) Code() int {
	return 500
}

func (o *GetRunnerJobExecutionsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsInternalServerError %s", 500, payload)
}

func (o *GetRunnerJobExecutionsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}/executions][%d] getRunnerJobExecutionsInternalServerError %s", 500, payload)
}

func (o *GetRunnerJobExecutionsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
