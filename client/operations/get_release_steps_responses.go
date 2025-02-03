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

// GetReleaseStepsReader is a Reader for the GetReleaseSteps structure.
type GetReleaseStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReleaseStepsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReleaseStepsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReleaseStepsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReleaseStepsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReleaseStepsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases/{release_id}/steps] GetReleaseSteps", response, response.Code())
	}
}

// NewGetReleaseStepsOK creates a GetReleaseStepsOK with default headers values
func NewGetReleaseStepsOK() *GetReleaseStepsOK {
	return &GetReleaseStepsOK{}
}

/*
GetReleaseStepsOK describes a response with status code 200, with default header values.

OK
*/
type GetReleaseStepsOK struct {
	Payload []*models.AppComponentReleaseStep
}

// IsSuccess returns true when this get release steps o k response has a 2xx status code
func (o *GetReleaseStepsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get release steps o k response has a 3xx status code
func (o *GetReleaseStepsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release steps o k response has a 4xx status code
func (o *GetReleaseStepsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get release steps o k response has a 5xx status code
func (o *GetReleaseStepsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get release steps o k response a status code equal to that given
func (o *GetReleaseStepsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get release steps o k response
func (o *GetReleaseStepsOK) Code() int {
	return 200
}

func (o *GetReleaseStepsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsOK %s", 200, payload)
}

func (o *GetReleaseStepsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsOK %s", 200, payload)
}

func (o *GetReleaseStepsOK) GetPayload() []*models.AppComponentReleaseStep {
	return o.Payload
}

func (o *GetReleaseStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseStepsBadRequest creates a GetReleaseStepsBadRequest with default headers values
func NewGetReleaseStepsBadRequest() *GetReleaseStepsBadRequest {
	return &GetReleaseStepsBadRequest{}
}

/*
GetReleaseStepsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetReleaseStepsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get release steps bad request response has a 2xx status code
func (o *GetReleaseStepsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release steps bad request response has a 3xx status code
func (o *GetReleaseStepsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release steps bad request response has a 4xx status code
func (o *GetReleaseStepsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release steps bad request response has a 5xx status code
func (o *GetReleaseStepsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get release steps bad request response a status code equal to that given
func (o *GetReleaseStepsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get release steps bad request response
func (o *GetReleaseStepsBadRequest) Code() int {
	return 400
}

func (o *GetReleaseStepsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsBadRequest %s", 400, payload)
}

func (o *GetReleaseStepsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsBadRequest %s", 400, payload)
}

func (o *GetReleaseStepsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetReleaseStepsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseStepsUnauthorized creates a GetReleaseStepsUnauthorized with default headers values
func NewGetReleaseStepsUnauthorized() *GetReleaseStepsUnauthorized {
	return &GetReleaseStepsUnauthorized{}
}

/*
GetReleaseStepsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReleaseStepsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get release steps unauthorized response has a 2xx status code
func (o *GetReleaseStepsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release steps unauthorized response has a 3xx status code
func (o *GetReleaseStepsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release steps unauthorized response has a 4xx status code
func (o *GetReleaseStepsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release steps unauthorized response has a 5xx status code
func (o *GetReleaseStepsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get release steps unauthorized response a status code equal to that given
func (o *GetReleaseStepsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get release steps unauthorized response
func (o *GetReleaseStepsUnauthorized) Code() int {
	return 401
}

func (o *GetReleaseStepsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsUnauthorized %s", 401, payload)
}

func (o *GetReleaseStepsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsUnauthorized %s", 401, payload)
}

func (o *GetReleaseStepsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetReleaseStepsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseStepsForbidden creates a GetReleaseStepsForbidden with default headers values
func NewGetReleaseStepsForbidden() *GetReleaseStepsForbidden {
	return &GetReleaseStepsForbidden{}
}

/*
GetReleaseStepsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReleaseStepsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get release steps forbidden response has a 2xx status code
func (o *GetReleaseStepsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release steps forbidden response has a 3xx status code
func (o *GetReleaseStepsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release steps forbidden response has a 4xx status code
func (o *GetReleaseStepsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release steps forbidden response has a 5xx status code
func (o *GetReleaseStepsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get release steps forbidden response a status code equal to that given
func (o *GetReleaseStepsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get release steps forbidden response
func (o *GetReleaseStepsForbidden) Code() int {
	return 403
}

func (o *GetReleaseStepsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsForbidden %s", 403, payload)
}

func (o *GetReleaseStepsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsForbidden %s", 403, payload)
}

func (o *GetReleaseStepsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetReleaseStepsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseStepsNotFound creates a GetReleaseStepsNotFound with default headers values
func NewGetReleaseStepsNotFound() *GetReleaseStepsNotFound {
	return &GetReleaseStepsNotFound{}
}

/*
GetReleaseStepsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetReleaseStepsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get release steps not found response has a 2xx status code
func (o *GetReleaseStepsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release steps not found response has a 3xx status code
func (o *GetReleaseStepsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release steps not found response has a 4xx status code
func (o *GetReleaseStepsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get release steps not found response has a 5xx status code
func (o *GetReleaseStepsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get release steps not found response a status code equal to that given
func (o *GetReleaseStepsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get release steps not found response
func (o *GetReleaseStepsNotFound) Code() int {
	return 404
}

func (o *GetReleaseStepsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsNotFound %s", 404, payload)
}

func (o *GetReleaseStepsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsNotFound %s", 404, payload)
}

func (o *GetReleaseStepsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetReleaseStepsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseStepsInternalServerError creates a GetReleaseStepsInternalServerError with default headers values
func NewGetReleaseStepsInternalServerError() *GetReleaseStepsInternalServerError {
	return &GetReleaseStepsInternalServerError{}
}

/*
GetReleaseStepsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReleaseStepsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get release steps internal server error response has a 2xx status code
func (o *GetReleaseStepsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get release steps internal server error response has a 3xx status code
func (o *GetReleaseStepsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get release steps internal server error response has a 4xx status code
func (o *GetReleaseStepsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get release steps internal server error response has a 5xx status code
func (o *GetReleaseStepsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get release steps internal server error response a status code equal to that given
func (o *GetReleaseStepsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get release steps internal server error response
func (o *GetReleaseStepsInternalServerError) Code() int {
	return 500
}

func (o *GetReleaseStepsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsInternalServerError %s", 500, payload)
}

func (o *GetReleaseStepsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/releases/{release_id}/steps][%d] getReleaseStepsInternalServerError %s", 500, payload)
}

func (o *GetReleaseStepsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetReleaseStepsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
