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

// GetComponentBuildReader is a Reader for the GetComponentBuild structure.
type GetComponentBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComponentBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetComponentBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetComponentBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetComponentBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComponentBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetComponentBuildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/components/{component_id}/builds/{build_id}] GetComponentBuild", response, response.Code())
	}
}

// NewGetComponentBuildOK creates a GetComponentBuildOK with default headers values
func NewGetComponentBuildOK() *GetComponentBuildOK {
	return &GetComponentBuildOK{}
}

/*
GetComponentBuildOK describes a response with status code 200, with default header values.

OK
*/
type GetComponentBuildOK struct {
	Payload *models.AppComponentBuild
}

// IsSuccess returns true when this get component build o k response has a 2xx status code
func (o *GetComponentBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get component build o k response has a 3xx status code
func (o *GetComponentBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get component build o k response has a 4xx status code
func (o *GetComponentBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get component build o k response has a 5xx status code
func (o *GetComponentBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get component build o k response a status code equal to that given
func (o *GetComponentBuildOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get component build o k response
func (o *GetComponentBuildOK) Code() int {
	return 200
}

func (o *GetComponentBuildOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildOK %s", 200, payload)
}

func (o *GetComponentBuildOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildOK %s", 200, payload)
}

func (o *GetComponentBuildOK) GetPayload() *models.AppComponentBuild {
	return o.Payload
}

func (o *GetComponentBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppComponentBuild)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentBuildBadRequest creates a GetComponentBuildBadRequest with default headers values
func NewGetComponentBuildBadRequest() *GetComponentBuildBadRequest {
	return &GetComponentBuildBadRequest{}
}

/*
GetComponentBuildBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetComponentBuildBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get component build bad request response has a 2xx status code
func (o *GetComponentBuildBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get component build bad request response has a 3xx status code
func (o *GetComponentBuildBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get component build bad request response has a 4xx status code
func (o *GetComponentBuildBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get component build bad request response has a 5xx status code
func (o *GetComponentBuildBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get component build bad request response a status code equal to that given
func (o *GetComponentBuildBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get component build bad request response
func (o *GetComponentBuildBadRequest) Code() int {
	return 400
}

func (o *GetComponentBuildBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildBadRequest %s", 400, payload)
}

func (o *GetComponentBuildBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildBadRequest %s", 400, payload)
}

func (o *GetComponentBuildBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetComponentBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentBuildUnauthorized creates a GetComponentBuildUnauthorized with default headers values
func NewGetComponentBuildUnauthorized() *GetComponentBuildUnauthorized {
	return &GetComponentBuildUnauthorized{}
}

/*
GetComponentBuildUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetComponentBuildUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get component build unauthorized response has a 2xx status code
func (o *GetComponentBuildUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get component build unauthorized response has a 3xx status code
func (o *GetComponentBuildUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get component build unauthorized response has a 4xx status code
func (o *GetComponentBuildUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get component build unauthorized response has a 5xx status code
func (o *GetComponentBuildUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get component build unauthorized response a status code equal to that given
func (o *GetComponentBuildUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get component build unauthorized response
func (o *GetComponentBuildUnauthorized) Code() int {
	return 401
}

func (o *GetComponentBuildUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildUnauthorized %s", 401, payload)
}

func (o *GetComponentBuildUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildUnauthorized %s", 401, payload)
}

func (o *GetComponentBuildUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetComponentBuildUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentBuildForbidden creates a GetComponentBuildForbidden with default headers values
func NewGetComponentBuildForbidden() *GetComponentBuildForbidden {
	return &GetComponentBuildForbidden{}
}

/*
GetComponentBuildForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetComponentBuildForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get component build forbidden response has a 2xx status code
func (o *GetComponentBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get component build forbidden response has a 3xx status code
func (o *GetComponentBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get component build forbidden response has a 4xx status code
func (o *GetComponentBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get component build forbidden response has a 5xx status code
func (o *GetComponentBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get component build forbidden response a status code equal to that given
func (o *GetComponentBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get component build forbidden response
func (o *GetComponentBuildForbidden) Code() int {
	return 403
}

func (o *GetComponentBuildForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildForbidden %s", 403, payload)
}

func (o *GetComponentBuildForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildForbidden %s", 403, payload)
}

func (o *GetComponentBuildForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetComponentBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentBuildNotFound creates a GetComponentBuildNotFound with default headers values
func NewGetComponentBuildNotFound() *GetComponentBuildNotFound {
	return &GetComponentBuildNotFound{}
}

/*
GetComponentBuildNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetComponentBuildNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get component build not found response has a 2xx status code
func (o *GetComponentBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get component build not found response has a 3xx status code
func (o *GetComponentBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get component build not found response has a 4xx status code
func (o *GetComponentBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get component build not found response has a 5xx status code
func (o *GetComponentBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get component build not found response a status code equal to that given
func (o *GetComponentBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get component build not found response
func (o *GetComponentBuildNotFound) Code() int {
	return 404
}

func (o *GetComponentBuildNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildNotFound %s", 404, payload)
}

func (o *GetComponentBuildNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildNotFound %s", 404, payload)
}

func (o *GetComponentBuildNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetComponentBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentBuildInternalServerError creates a GetComponentBuildInternalServerError with default headers values
func NewGetComponentBuildInternalServerError() *GetComponentBuildInternalServerError {
	return &GetComponentBuildInternalServerError{}
}

/*
GetComponentBuildInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetComponentBuildInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get component build internal server error response has a 2xx status code
func (o *GetComponentBuildInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get component build internal server error response has a 3xx status code
func (o *GetComponentBuildInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get component build internal server error response has a 4xx status code
func (o *GetComponentBuildInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get component build internal server error response has a 5xx status code
func (o *GetComponentBuildInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get component build internal server error response a status code equal to that given
func (o *GetComponentBuildInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get component build internal server error response
func (o *GetComponentBuildInternalServerError) Code() int {
	return 500
}

func (o *GetComponentBuildInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildInternalServerError %s", 500, payload)
}

func (o *GetComponentBuildInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds/{build_id}][%d] getComponentBuildInternalServerError %s", 500, payload)
}

func (o *GetComponentBuildInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetComponentBuildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
