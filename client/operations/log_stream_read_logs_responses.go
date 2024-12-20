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

// LogStreamReadLogsReader is a Reader for the LogStreamReadLogs structure.
type LogStreamReadLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogStreamReadLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogStreamReadLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogStreamReadLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLogStreamReadLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLogStreamReadLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLogStreamReadLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLogStreamReadLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/log-streams/{log_stream_id}/logs] LogStreamReadLogs", response, response.Code())
	}
}

// NewLogStreamReadLogsOK creates a LogStreamReadLogsOK with default headers values
func NewLogStreamReadLogsOK() *LogStreamReadLogsOK {
	return &LogStreamReadLogsOK{}
}

/*
LogStreamReadLogsOK describes a response with status code 200, with default header values.

OK
*/
type LogStreamReadLogsOK struct {
	Payload []*models.AppOtelLogRecord
}

// IsSuccess returns true when this log stream read logs o k response has a 2xx status code
func (o *LogStreamReadLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log stream read logs o k response has a 3xx status code
func (o *LogStreamReadLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream read logs o k response has a 4xx status code
func (o *LogStreamReadLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this log stream read logs o k response has a 5xx status code
func (o *LogStreamReadLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream read logs o k response a status code equal to that given
func (o *LogStreamReadLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the log stream read logs o k response
func (o *LogStreamReadLogsOK) Code() int {
	return 200
}

func (o *LogStreamReadLogsOK) Error() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsOK  %+v", 200, o.Payload)
}

func (o *LogStreamReadLogsOK) String() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsOK  %+v", 200, o.Payload)
}

func (o *LogStreamReadLogsOK) GetPayload() []*models.AppOtelLogRecord {
	return o.Payload
}

func (o *LogStreamReadLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamReadLogsBadRequest creates a LogStreamReadLogsBadRequest with default headers values
func NewLogStreamReadLogsBadRequest() *LogStreamReadLogsBadRequest {
	return &LogStreamReadLogsBadRequest{}
}

/*
LogStreamReadLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type LogStreamReadLogsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream read logs bad request response has a 2xx status code
func (o *LogStreamReadLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream read logs bad request response has a 3xx status code
func (o *LogStreamReadLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream read logs bad request response has a 4xx status code
func (o *LogStreamReadLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream read logs bad request response has a 5xx status code
func (o *LogStreamReadLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream read logs bad request response a status code equal to that given
func (o *LogStreamReadLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the log stream read logs bad request response
func (o *LogStreamReadLogsBadRequest) Code() int {
	return 400
}

func (o *LogStreamReadLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsBadRequest  %+v", 400, o.Payload)
}

func (o *LogStreamReadLogsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsBadRequest  %+v", 400, o.Payload)
}

func (o *LogStreamReadLogsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamReadLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamReadLogsUnauthorized creates a LogStreamReadLogsUnauthorized with default headers values
func NewLogStreamReadLogsUnauthorized() *LogStreamReadLogsUnauthorized {
	return &LogStreamReadLogsUnauthorized{}
}

/*
LogStreamReadLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type LogStreamReadLogsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream read logs unauthorized response has a 2xx status code
func (o *LogStreamReadLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream read logs unauthorized response has a 3xx status code
func (o *LogStreamReadLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream read logs unauthorized response has a 4xx status code
func (o *LogStreamReadLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream read logs unauthorized response has a 5xx status code
func (o *LogStreamReadLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream read logs unauthorized response a status code equal to that given
func (o *LogStreamReadLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the log stream read logs unauthorized response
func (o *LogStreamReadLogsUnauthorized) Code() int {
	return 401
}

func (o *LogStreamReadLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *LogStreamReadLogsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *LogStreamReadLogsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamReadLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamReadLogsForbidden creates a LogStreamReadLogsForbidden with default headers values
func NewLogStreamReadLogsForbidden() *LogStreamReadLogsForbidden {
	return &LogStreamReadLogsForbidden{}
}

/*
LogStreamReadLogsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type LogStreamReadLogsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream read logs forbidden response has a 2xx status code
func (o *LogStreamReadLogsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream read logs forbidden response has a 3xx status code
func (o *LogStreamReadLogsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream read logs forbidden response has a 4xx status code
func (o *LogStreamReadLogsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream read logs forbidden response has a 5xx status code
func (o *LogStreamReadLogsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream read logs forbidden response a status code equal to that given
func (o *LogStreamReadLogsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the log stream read logs forbidden response
func (o *LogStreamReadLogsForbidden) Code() int {
	return 403
}

func (o *LogStreamReadLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsForbidden  %+v", 403, o.Payload)
}

func (o *LogStreamReadLogsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsForbidden  %+v", 403, o.Payload)
}

func (o *LogStreamReadLogsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamReadLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamReadLogsNotFound creates a LogStreamReadLogsNotFound with default headers values
func NewLogStreamReadLogsNotFound() *LogStreamReadLogsNotFound {
	return &LogStreamReadLogsNotFound{}
}

/*
LogStreamReadLogsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type LogStreamReadLogsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream read logs not found response has a 2xx status code
func (o *LogStreamReadLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream read logs not found response has a 3xx status code
func (o *LogStreamReadLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream read logs not found response has a 4xx status code
func (o *LogStreamReadLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream read logs not found response has a 5xx status code
func (o *LogStreamReadLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream read logs not found response a status code equal to that given
func (o *LogStreamReadLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the log stream read logs not found response
func (o *LogStreamReadLogsNotFound) Code() int {
	return 404
}

func (o *LogStreamReadLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsNotFound  %+v", 404, o.Payload)
}

func (o *LogStreamReadLogsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsNotFound  %+v", 404, o.Payload)
}

func (o *LogStreamReadLogsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamReadLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamReadLogsInternalServerError creates a LogStreamReadLogsInternalServerError with default headers values
func NewLogStreamReadLogsInternalServerError() *LogStreamReadLogsInternalServerError {
	return &LogStreamReadLogsInternalServerError{}
}

/*
LogStreamReadLogsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type LogStreamReadLogsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream read logs internal server error response has a 2xx status code
func (o *LogStreamReadLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream read logs internal server error response has a 3xx status code
func (o *LogStreamReadLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream read logs internal server error response has a 4xx status code
func (o *LogStreamReadLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this log stream read logs internal server error response has a 5xx status code
func (o *LogStreamReadLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this log stream read logs internal server error response a status code equal to that given
func (o *LogStreamReadLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the log stream read logs internal server error response
func (o *LogStreamReadLogsInternalServerError) Code() int {
	return 500
}

func (o *LogStreamReadLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *LogStreamReadLogsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/log-streams/{log_stream_id}/logs][%d] logStreamReadLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *LogStreamReadLogsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamReadLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}