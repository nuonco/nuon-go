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

// GetInstallWorkflowStepsReader is a Reader for the GetInstallWorkflowSteps structure.
type GetInstallWorkflowStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallWorkflowStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallWorkflowStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallWorkflowStepsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallWorkflowStepsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallWorkflowStepsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallWorkflowStepsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallWorkflowStepsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/install-workflows/{install_workflow_id}/steps] GetInstallWorkflowSteps", response, response.Code())
	}
}

// NewGetInstallWorkflowStepsOK creates a GetInstallWorkflowStepsOK with default headers values
func NewGetInstallWorkflowStepsOK() *GetInstallWorkflowStepsOK {
	return &GetInstallWorkflowStepsOK{}
}

/*
GetInstallWorkflowStepsOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallWorkflowStepsOK struct {
	Payload []*models.AppInstallWorkflowStep
}

// IsSuccess returns true when this get install workflow steps o k response has a 2xx status code
func (o *GetInstallWorkflowStepsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install workflow steps o k response has a 3xx status code
func (o *GetInstallWorkflowStepsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow steps o k response has a 4xx status code
func (o *GetInstallWorkflowStepsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install workflow steps o k response has a 5xx status code
func (o *GetInstallWorkflowStepsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow steps o k response a status code equal to that given
func (o *GetInstallWorkflowStepsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install workflow steps o k response
func (o *GetInstallWorkflowStepsOK) Code() int {
	return 200
}

func (o *GetInstallWorkflowStepsOK) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsOK  %+v", 200, o.Payload)
}

func (o *GetInstallWorkflowStepsOK) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsOK  %+v", 200, o.Payload)
}

func (o *GetInstallWorkflowStepsOK) GetPayload() []*models.AppInstallWorkflowStep {
	return o.Payload
}

func (o *GetInstallWorkflowStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowStepsBadRequest creates a GetInstallWorkflowStepsBadRequest with default headers values
func NewGetInstallWorkflowStepsBadRequest() *GetInstallWorkflowStepsBadRequest {
	return &GetInstallWorkflowStepsBadRequest{}
}

/*
GetInstallWorkflowStepsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallWorkflowStepsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow steps bad request response has a 2xx status code
func (o *GetInstallWorkflowStepsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow steps bad request response has a 3xx status code
func (o *GetInstallWorkflowStepsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow steps bad request response has a 4xx status code
func (o *GetInstallWorkflowStepsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow steps bad request response has a 5xx status code
func (o *GetInstallWorkflowStepsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow steps bad request response a status code equal to that given
func (o *GetInstallWorkflowStepsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install workflow steps bad request response
func (o *GetInstallWorkflowStepsBadRequest) Code() int {
	return 400
}

func (o *GetInstallWorkflowStepsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallWorkflowStepsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallWorkflowStepsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowStepsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowStepsUnauthorized creates a GetInstallWorkflowStepsUnauthorized with default headers values
func NewGetInstallWorkflowStepsUnauthorized() *GetInstallWorkflowStepsUnauthorized {
	return &GetInstallWorkflowStepsUnauthorized{}
}

/*
GetInstallWorkflowStepsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallWorkflowStepsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow steps unauthorized response has a 2xx status code
func (o *GetInstallWorkflowStepsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow steps unauthorized response has a 3xx status code
func (o *GetInstallWorkflowStepsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow steps unauthorized response has a 4xx status code
func (o *GetInstallWorkflowStepsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow steps unauthorized response has a 5xx status code
func (o *GetInstallWorkflowStepsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow steps unauthorized response a status code equal to that given
func (o *GetInstallWorkflowStepsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install workflow steps unauthorized response
func (o *GetInstallWorkflowStepsUnauthorized) Code() int {
	return 401
}

func (o *GetInstallWorkflowStepsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallWorkflowStepsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallWorkflowStepsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowStepsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowStepsForbidden creates a GetInstallWorkflowStepsForbidden with default headers values
func NewGetInstallWorkflowStepsForbidden() *GetInstallWorkflowStepsForbidden {
	return &GetInstallWorkflowStepsForbidden{}
}

/*
GetInstallWorkflowStepsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallWorkflowStepsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow steps forbidden response has a 2xx status code
func (o *GetInstallWorkflowStepsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow steps forbidden response has a 3xx status code
func (o *GetInstallWorkflowStepsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow steps forbidden response has a 4xx status code
func (o *GetInstallWorkflowStepsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow steps forbidden response has a 5xx status code
func (o *GetInstallWorkflowStepsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow steps forbidden response a status code equal to that given
func (o *GetInstallWorkflowStepsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install workflow steps forbidden response
func (o *GetInstallWorkflowStepsForbidden) Code() int {
	return 403
}

func (o *GetInstallWorkflowStepsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallWorkflowStepsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallWorkflowStepsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowStepsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowStepsNotFound creates a GetInstallWorkflowStepsNotFound with default headers values
func NewGetInstallWorkflowStepsNotFound() *GetInstallWorkflowStepsNotFound {
	return &GetInstallWorkflowStepsNotFound{}
}

/*
GetInstallWorkflowStepsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallWorkflowStepsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow steps not found response has a 2xx status code
func (o *GetInstallWorkflowStepsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow steps not found response has a 3xx status code
func (o *GetInstallWorkflowStepsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow steps not found response has a 4xx status code
func (o *GetInstallWorkflowStepsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow steps not found response has a 5xx status code
func (o *GetInstallWorkflowStepsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow steps not found response a status code equal to that given
func (o *GetInstallWorkflowStepsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install workflow steps not found response
func (o *GetInstallWorkflowStepsNotFound) Code() int {
	return 404
}

func (o *GetInstallWorkflowStepsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallWorkflowStepsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallWorkflowStepsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowStepsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowStepsInternalServerError creates a GetInstallWorkflowStepsInternalServerError with default headers values
func NewGetInstallWorkflowStepsInternalServerError() *GetInstallWorkflowStepsInternalServerError {
	return &GetInstallWorkflowStepsInternalServerError{}
}

/*
GetInstallWorkflowStepsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallWorkflowStepsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow steps internal server error response has a 2xx status code
func (o *GetInstallWorkflowStepsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow steps internal server error response has a 3xx status code
func (o *GetInstallWorkflowStepsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow steps internal server error response has a 4xx status code
func (o *GetInstallWorkflowStepsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install workflow steps internal server error response has a 5xx status code
func (o *GetInstallWorkflowStepsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install workflow steps internal server error response a status code equal to that given
func (o *GetInstallWorkflowStepsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install workflow steps internal server error response
func (o *GetInstallWorkflowStepsInternalServerError) Code() int {
	return 500
}

func (o *GetInstallWorkflowStepsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallWorkflowStepsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}/steps][%d] getInstallWorkflowStepsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallWorkflowStepsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowStepsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
