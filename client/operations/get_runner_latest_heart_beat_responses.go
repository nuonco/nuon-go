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

// GetRunnerLatestHeartBeatReader is a Reader for the GetRunnerLatestHeartBeat structure.
type GetRunnerLatestHeartBeatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunnerLatestHeartBeatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunnerLatestHeartBeatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunnerLatestHeartBeatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRunnerLatestHeartBeatUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRunnerLatestHeartBeatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunnerLatestHeartBeatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRunnerLatestHeartBeatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/runners/{runner_id}/latest-heart-beat] GetRunnerLatestHeartBeat", response, response.Code())
	}
}

// NewGetRunnerLatestHeartBeatOK creates a GetRunnerLatestHeartBeatOK with default headers values
func NewGetRunnerLatestHeartBeatOK() *GetRunnerLatestHeartBeatOK {
	return &GetRunnerLatestHeartBeatOK{}
}

/*
GetRunnerLatestHeartBeatOK describes a response with status code 200, with default header values.

OK
*/
type GetRunnerLatestHeartBeatOK struct {
	Payload *models.AppRunnerHeartBeat
}

// IsSuccess returns true when this get runner latest heart beat o k response has a 2xx status code
func (o *GetRunnerLatestHeartBeatOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runner latest heart beat o k response has a 3xx status code
func (o *GetRunnerLatestHeartBeatOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner latest heart beat o k response has a 4xx status code
func (o *GetRunnerLatestHeartBeatOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner latest heart beat o k response has a 5xx status code
func (o *GetRunnerLatestHeartBeatOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner latest heart beat o k response a status code equal to that given
func (o *GetRunnerLatestHeartBeatOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runner latest heart beat o k response
func (o *GetRunnerLatestHeartBeatOK) Code() int {
	return 200
}

func (o *GetRunnerLatestHeartBeatOK) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatOK  %+v", 200, o.Payload)
}

func (o *GetRunnerLatestHeartBeatOK) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatOK  %+v", 200, o.Payload)
}

func (o *GetRunnerLatestHeartBeatOK) GetPayload() *models.AppRunnerHeartBeat {
	return o.Payload
}

func (o *GetRunnerLatestHeartBeatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRunnerHeartBeat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerLatestHeartBeatBadRequest creates a GetRunnerLatestHeartBeatBadRequest with default headers values
func NewGetRunnerLatestHeartBeatBadRequest() *GetRunnerLatestHeartBeatBadRequest {
	return &GetRunnerLatestHeartBeatBadRequest{}
}

/*
GetRunnerLatestHeartBeatBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRunnerLatestHeartBeatBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner latest heart beat bad request response has a 2xx status code
func (o *GetRunnerLatestHeartBeatBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner latest heart beat bad request response has a 3xx status code
func (o *GetRunnerLatestHeartBeatBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner latest heart beat bad request response has a 4xx status code
func (o *GetRunnerLatestHeartBeatBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner latest heart beat bad request response has a 5xx status code
func (o *GetRunnerLatestHeartBeatBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner latest heart beat bad request response a status code equal to that given
func (o *GetRunnerLatestHeartBeatBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get runner latest heart beat bad request response
func (o *GetRunnerLatestHeartBeatBadRequest) Code() int {
	return 400
}

func (o *GetRunnerLatestHeartBeatBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunnerLatestHeartBeatBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunnerLatestHeartBeatBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerLatestHeartBeatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerLatestHeartBeatUnauthorized creates a GetRunnerLatestHeartBeatUnauthorized with default headers values
func NewGetRunnerLatestHeartBeatUnauthorized() *GetRunnerLatestHeartBeatUnauthorized {
	return &GetRunnerLatestHeartBeatUnauthorized{}
}

/*
GetRunnerLatestHeartBeatUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRunnerLatestHeartBeatUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner latest heart beat unauthorized response has a 2xx status code
func (o *GetRunnerLatestHeartBeatUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner latest heart beat unauthorized response has a 3xx status code
func (o *GetRunnerLatestHeartBeatUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner latest heart beat unauthorized response has a 4xx status code
func (o *GetRunnerLatestHeartBeatUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner latest heart beat unauthorized response has a 5xx status code
func (o *GetRunnerLatestHeartBeatUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner latest heart beat unauthorized response a status code equal to that given
func (o *GetRunnerLatestHeartBeatUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get runner latest heart beat unauthorized response
func (o *GetRunnerLatestHeartBeatUnauthorized) Code() int {
	return 401
}

func (o *GetRunnerLatestHeartBeatUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunnerLatestHeartBeatUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunnerLatestHeartBeatUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerLatestHeartBeatUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerLatestHeartBeatForbidden creates a GetRunnerLatestHeartBeatForbidden with default headers values
func NewGetRunnerLatestHeartBeatForbidden() *GetRunnerLatestHeartBeatForbidden {
	return &GetRunnerLatestHeartBeatForbidden{}
}

/*
GetRunnerLatestHeartBeatForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRunnerLatestHeartBeatForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner latest heart beat forbidden response has a 2xx status code
func (o *GetRunnerLatestHeartBeatForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner latest heart beat forbidden response has a 3xx status code
func (o *GetRunnerLatestHeartBeatForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner latest heart beat forbidden response has a 4xx status code
func (o *GetRunnerLatestHeartBeatForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner latest heart beat forbidden response has a 5xx status code
func (o *GetRunnerLatestHeartBeatForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner latest heart beat forbidden response a status code equal to that given
func (o *GetRunnerLatestHeartBeatForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get runner latest heart beat forbidden response
func (o *GetRunnerLatestHeartBeatForbidden) Code() int {
	return 403
}

func (o *GetRunnerLatestHeartBeatForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatForbidden  %+v", 403, o.Payload)
}

func (o *GetRunnerLatestHeartBeatForbidden) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatForbidden  %+v", 403, o.Payload)
}

func (o *GetRunnerLatestHeartBeatForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerLatestHeartBeatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerLatestHeartBeatNotFound creates a GetRunnerLatestHeartBeatNotFound with default headers values
func NewGetRunnerLatestHeartBeatNotFound() *GetRunnerLatestHeartBeatNotFound {
	return &GetRunnerLatestHeartBeatNotFound{}
}

/*
GetRunnerLatestHeartBeatNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRunnerLatestHeartBeatNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner latest heart beat not found response has a 2xx status code
func (o *GetRunnerLatestHeartBeatNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner latest heart beat not found response has a 3xx status code
func (o *GetRunnerLatestHeartBeatNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner latest heart beat not found response has a 4xx status code
func (o *GetRunnerLatestHeartBeatNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner latest heart beat not found response has a 5xx status code
func (o *GetRunnerLatestHeartBeatNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner latest heart beat not found response a status code equal to that given
func (o *GetRunnerLatestHeartBeatNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get runner latest heart beat not found response
func (o *GetRunnerLatestHeartBeatNotFound) Code() int {
	return 404
}

func (o *GetRunnerLatestHeartBeatNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerLatestHeartBeatNotFound) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerLatestHeartBeatNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerLatestHeartBeatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerLatestHeartBeatInternalServerError creates a GetRunnerLatestHeartBeatInternalServerError with default headers values
func NewGetRunnerLatestHeartBeatInternalServerError() *GetRunnerLatestHeartBeatInternalServerError {
	return &GetRunnerLatestHeartBeatInternalServerError{}
}

/*
GetRunnerLatestHeartBeatInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRunnerLatestHeartBeatInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner latest heart beat internal server error response has a 2xx status code
func (o *GetRunnerLatestHeartBeatInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner latest heart beat internal server error response has a 3xx status code
func (o *GetRunnerLatestHeartBeatInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner latest heart beat internal server error response has a 4xx status code
func (o *GetRunnerLatestHeartBeatInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner latest heart beat internal server error response has a 5xx status code
func (o *GetRunnerLatestHeartBeatInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get runner latest heart beat internal server error response a status code equal to that given
func (o *GetRunnerLatestHeartBeatInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get runner latest heart beat internal server error response
func (o *GetRunnerLatestHeartBeatInternalServerError) Code() int {
	return 500
}

func (o *GetRunnerLatestHeartBeatInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunnerLatestHeartBeatInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/latest-heart-beat][%d] getRunnerLatestHeartBeatInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunnerLatestHeartBeatInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerLatestHeartBeatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
