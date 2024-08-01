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

// GetOrgHealthChecksReader is a Reader for the GetOrgHealthChecks structure.
type GetOrgHealthChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgHealthChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgHealthChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgHealthChecksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgHealthChecksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgHealthChecksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgHealthChecksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgHealthChecksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/orgs/current/health-checks] GetOrgHealthChecks", response, response.Code())
	}
}

// NewGetOrgHealthChecksOK creates a GetOrgHealthChecksOK with default headers values
func NewGetOrgHealthChecksOK() *GetOrgHealthChecksOK {
	return &GetOrgHealthChecksOK{}
}

/*
GetOrgHealthChecksOK describes a response with status code 200, with default header values.

OK
*/
type GetOrgHealthChecksOK struct {
	Payload []*models.AppOrgHealthCheck
}

// IsSuccess returns true when this get org health checks o k response has a 2xx status code
func (o *GetOrgHealthChecksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org health checks o k response has a 3xx status code
func (o *GetOrgHealthChecksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org health checks o k response has a 4xx status code
func (o *GetOrgHealthChecksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org health checks o k response has a 5xx status code
func (o *GetOrgHealthChecksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org health checks o k response a status code equal to that given
func (o *GetOrgHealthChecksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org health checks o k response
func (o *GetOrgHealthChecksOK) Code() int {
	return 200
}

func (o *GetOrgHealthChecksOK) Error() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksOK  %+v", 200, o.Payload)
}

func (o *GetOrgHealthChecksOK) String() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksOK  %+v", 200, o.Payload)
}

func (o *GetOrgHealthChecksOK) GetPayload() []*models.AppOrgHealthCheck {
	return o.Payload
}

func (o *GetOrgHealthChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgHealthChecksBadRequest creates a GetOrgHealthChecksBadRequest with default headers values
func NewGetOrgHealthChecksBadRequest() *GetOrgHealthChecksBadRequest {
	return &GetOrgHealthChecksBadRequest{}
}

/*
GetOrgHealthChecksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrgHealthChecksBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org health checks bad request response has a 2xx status code
func (o *GetOrgHealthChecksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org health checks bad request response has a 3xx status code
func (o *GetOrgHealthChecksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org health checks bad request response has a 4xx status code
func (o *GetOrgHealthChecksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org health checks bad request response has a 5xx status code
func (o *GetOrgHealthChecksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get org health checks bad request response a status code equal to that given
func (o *GetOrgHealthChecksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get org health checks bad request response
func (o *GetOrgHealthChecksBadRequest) Code() int {
	return 400
}

func (o *GetOrgHealthChecksBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgHealthChecksBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgHealthChecksBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgHealthChecksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgHealthChecksUnauthorized creates a GetOrgHealthChecksUnauthorized with default headers values
func NewGetOrgHealthChecksUnauthorized() *GetOrgHealthChecksUnauthorized {
	return &GetOrgHealthChecksUnauthorized{}
}

/*
GetOrgHealthChecksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrgHealthChecksUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org health checks unauthorized response has a 2xx status code
func (o *GetOrgHealthChecksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org health checks unauthorized response has a 3xx status code
func (o *GetOrgHealthChecksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org health checks unauthorized response has a 4xx status code
func (o *GetOrgHealthChecksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org health checks unauthorized response has a 5xx status code
func (o *GetOrgHealthChecksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org health checks unauthorized response a status code equal to that given
func (o *GetOrgHealthChecksUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org health checks unauthorized response
func (o *GetOrgHealthChecksUnauthorized) Code() int {
	return 401
}

func (o *GetOrgHealthChecksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgHealthChecksUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgHealthChecksUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgHealthChecksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgHealthChecksForbidden creates a GetOrgHealthChecksForbidden with default headers values
func NewGetOrgHealthChecksForbidden() *GetOrgHealthChecksForbidden {
	return &GetOrgHealthChecksForbidden{}
}

/*
GetOrgHealthChecksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrgHealthChecksForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org health checks forbidden response has a 2xx status code
func (o *GetOrgHealthChecksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org health checks forbidden response has a 3xx status code
func (o *GetOrgHealthChecksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org health checks forbidden response has a 4xx status code
func (o *GetOrgHealthChecksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org health checks forbidden response has a 5xx status code
func (o *GetOrgHealthChecksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org health checks forbidden response a status code equal to that given
func (o *GetOrgHealthChecksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org health checks forbidden response
func (o *GetOrgHealthChecksForbidden) Code() int {
	return 403
}

func (o *GetOrgHealthChecksForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgHealthChecksForbidden) String() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgHealthChecksForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgHealthChecksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgHealthChecksNotFound creates a GetOrgHealthChecksNotFound with default headers values
func NewGetOrgHealthChecksNotFound() *GetOrgHealthChecksNotFound {
	return &GetOrgHealthChecksNotFound{}
}

/*
GetOrgHealthChecksNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrgHealthChecksNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org health checks not found response has a 2xx status code
func (o *GetOrgHealthChecksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org health checks not found response has a 3xx status code
func (o *GetOrgHealthChecksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org health checks not found response has a 4xx status code
func (o *GetOrgHealthChecksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org health checks not found response has a 5xx status code
func (o *GetOrgHealthChecksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get org health checks not found response a status code equal to that given
func (o *GetOrgHealthChecksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get org health checks not found response
func (o *GetOrgHealthChecksNotFound) Code() int {
	return 404
}

func (o *GetOrgHealthChecksNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgHealthChecksNotFound) String() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgHealthChecksNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgHealthChecksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgHealthChecksInternalServerError creates a GetOrgHealthChecksInternalServerError with default headers values
func NewGetOrgHealthChecksInternalServerError() *GetOrgHealthChecksInternalServerError {
	return &GetOrgHealthChecksInternalServerError{}
}

/*
GetOrgHealthChecksInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrgHealthChecksInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org health checks internal server error response has a 2xx status code
func (o *GetOrgHealthChecksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org health checks internal server error response has a 3xx status code
func (o *GetOrgHealthChecksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org health checks internal server error response has a 4xx status code
func (o *GetOrgHealthChecksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org health checks internal server error response has a 5xx status code
func (o *GetOrgHealthChecksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org health checks internal server error response a status code equal to that given
func (o *GetOrgHealthChecksInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org health checks internal server error response
func (o *GetOrgHealthChecksInternalServerError) Code() int {
	return 500
}

func (o *GetOrgHealthChecksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgHealthChecksInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/orgs/current/health-checks][%d] getOrgHealthChecksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgHealthChecksInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgHealthChecksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
