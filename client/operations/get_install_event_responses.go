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

// GetInstallEventReader is a Reader for the GetInstallEvent structure.
type GetInstallEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallEventInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/events/{event_id}] GetInstallEvent", response, response.Code())
	}
}

// NewGetInstallEventOK creates a GetInstallEventOK with default headers values
func NewGetInstallEventOK() *GetInstallEventOK {
	return &GetInstallEventOK{}
}

/*
GetInstallEventOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallEventOK struct {
	Payload *models.AppInstallEvent
}

// IsSuccess returns true when this get install event o k response has a 2xx status code
func (o *GetInstallEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install event o k response has a 3xx status code
func (o *GetInstallEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install event o k response has a 4xx status code
func (o *GetInstallEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install event o k response has a 5xx status code
func (o *GetInstallEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install event o k response a status code equal to that given
func (o *GetInstallEventOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install event o k response
func (o *GetInstallEventOK) Code() int {
	return 200
}

func (o *GetInstallEventOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventOK  %+v", 200, o.Payload)
}

func (o *GetInstallEventOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventOK  %+v", 200, o.Payload)
}

func (o *GetInstallEventOK) GetPayload() *models.AppInstallEvent {
	return o.Payload
}

func (o *GetInstallEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallEventBadRequest creates a GetInstallEventBadRequest with default headers values
func NewGetInstallEventBadRequest() *GetInstallEventBadRequest {
	return &GetInstallEventBadRequest{}
}

/*
GetInstallEventBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallEventBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install event bad request response has a 2xx status code
func (o *GetInstallEventBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install event bad request response has a 3xx status code
func (o *GetInstallEventBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install event bad request response has a 4xx status code
func (o *GetInstallEventBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install event bad request response has a 5xx status code
func (o *GetInstallEventBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install event bad request response a status code equal to that given
func (o *GetInstallEventBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install event bad request response
func (o *GetInstallEventBadRequest) Code() int {
	return 400
}

func (o *GetInstallEventBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallEventBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallEventBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallEventUnauthorized creates a GetInstallEventUnauthorized with default headers values
func NewGetInstallEventUnauthorized() *GetInstallEventUnauthorized {
	return &GetInstallEventUnauthorized{}
}

/*
GetInstallEventUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallEventUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install event unauthorized response has a 2xx status code
func (o *GetInstallEventUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install event unauthorized response has a 3xx status code
func (o *GetInstallEventUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install event unauthorized response has a 4xx status code
func (o *GetInstallEventUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install event unauthorized response has a 5xx status code
func (o *GetInstallEventUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install event unauthorized response a status code equal to that given
func (o *GetInstallEventUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install event unauthorized response
func (o *GetInstallEventUnauthorized) Code() int {
	return 401
}

func (o *GetInstallEventUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallEventUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallEventUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallEventForbidden creates a GetInstallEventForbidden with default headers values
func NewGetInstallEventForbidden() *GetInstallEventForbidden {
	return &GetInstallEventForbidden{}
}

/*
GetInstallEventForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallEventForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install event forbidden response has a 2xx status code
func (o *GetInstallEventForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install event forbidden response has a 3xx status code
func (o *GetInstallEventForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install event forbidden response has a 4xx status code
func (o *GetInstallEventForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install event forbidden response has a 5xx status code
func (o *GetInstallEventForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install event forbidden response a status code equal to that given
func (o *GetInstallEventForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install event forbidden response
func (o *GetInstallEventForbidden) Code() int {
	return 403
}

func (o *GetInstallEventForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallEventForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallEventForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallEventNotFound creates a GetInstallEventNotFound with default headers values
func NewGetInstallEventNotFound() *GetInstallEventNotFound {
	return &GetInstallEventNotFound{}
}

/*
GetInstallEventNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallEventNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install event not found response has a 2xx status code
func (o *GetInstallEventNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install event not found response has a 3xx status code
func (o *GetInstallEventNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install event not found response has a 4xx status code
func (o *GetInstallEventNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install event not found response has a 5xx status code
func (o *GetInstallEventNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install event not found response a status code equal to that given
func (o *GetInstallEventNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install event not found response
func (o *GetInstallEventNotFound) Code() int {
	return 404
}

func (o *GetInstallEventNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallEventNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallEventNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallEventInternalServerError creates a GetInstallEventInternalServerError with default headers values
func NewGetInstallEventInternalServerError() *GetInstallEventInternalServerError {
	return &GetInstallEventInternalServerError{}
}

/*
GetInstallEventInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallEventInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install event internal server error response has a 2xx status code
func (o *GetInstallEventInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install event internal server error response has a 3xx status code
func (o *GetInstallEventInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install event internal server error response has a 4xx status code
func (o *GetInstallEventInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install event internal server error response has a 5xx status code
func (o *GetInstallEventInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install event internal server error response a status code equal to that given
func (o *GetInstallEventInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install event internal server error response
func (o *GetInstallEventInternalServerError) Code() int {
	return 500
}

func (o *GetInstallEventInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallEventInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/events/{event_id}][%d] getInstallEventInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallEventInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallEventInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}