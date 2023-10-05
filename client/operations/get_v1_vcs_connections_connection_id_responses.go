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

// GetV1VcsConnectionsConnectionIDReader is a Reader for the GetV1VcsConnectionsConnectionID structure.
type GetV1VcsConnectionsConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1VcsConnectionsConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1VcsConnectionsConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1VcsConnectionsConnectionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1VcsConnectionsConnectionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1VcsConnectionsConnectionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1VcsConnectionsConnectionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1VcsConnectionsConnectionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/vcs/connections/{connection_id}] GetV1VcsConnectionsConnectionID", response, response.Code())
	}
}

// NewGetV1VcsConnectionsConnectionIDOK creates a GetV1VcsConnectionsConnectionIDOK with default headers values
func NewGetV1VcsConnectionsConnectionIDOK() *GetV1VcsConnectionsConnectionIDOK {
	return &GetV1VcsConnectionsConnectionIDOK{}
}

/*
GetV1VcsConnectionsConnectionIDOK describes a response with status code 200, with default header values.

OK
*/
type GetV1VcsConnectionsConnectionIDOK struct {
	Payload *models.AppVCSConnection
}

// IsSuccess returns true when this get v1 vcs connections connection Id o k response has a 2xx status code
func (o *GetV1VcsConnectionsConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 vcs connections connection Id o k response has a 3xx status code
func (o *GetV1VcsConnectionsConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 vcs connections connection Id o k response has a 4xx status code
func (o *GetV1VcsConnectionsConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 vcs connections connection Id o k response has a 5xx status code
func (o *GetV1VcsConnectionsConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 vcs connections connection Id o k response a status code equal to that given
func (o *GetV1VcsConnectionsConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 vcs connections connection Id o k response
func (o *GetV1VcsConnectionsConnectionIDOK) Code() int {
	return 200
}

func (o *GetV1VcsConnectionsConnectionIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdOK  %+v", 200, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDOK) String() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdOK  %+v", 200, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDOK) GetPayload() *models.AppVCSConnection {
	return o.Payload
}

func (o *GetV1VcsConnectionsConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppVCSConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1VcsConnectionsConnectionIDBadRequest creates a GetV1VcsConnectionsConnectionIDBadRequest with default headers values
func NewGetV1VcsConnectionsConnectionIDBadRequest() *GetV1VcsConnectionsConnectionIDBadRequest {
	return &GetV1VcsConnectionsConnectionIDBadRequest{}
}

/*
GetV1VcsConnectionsConnectionIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1VcsConnectionsConnectionIDBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 vcs connections connection Id bad request response has a 2xx status code
func (o *GetV1VcsConnectionsConnectionIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 vcs connections connection Id bad request response has a 3xx status code
func (o *GetV1VcsConnectionsConnectionIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 vcs connections connection Id bad request response has a 4xx status code
func (o *GetV1VcsConnectionsConnectionIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 vcs connections connection Id bad request response has a 5xx status code
func (o *GetV1VcsConnectionsConnectionIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 vcs connections connection Id bad request response a status code equal to that given
func (o *GetV1VcsConnectionsConnectionIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 vcs connections connection Id bad request response
func (o *GetV1VcsConnectionsConnectionIDBadRequest) Code() int {
	return 400
}

func (o *GetV1VcsConnectionsConnectionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1VcsConnectionsConnectionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1VcsConnectionsConnectionIDUnauthorized creates a GetV1VcsConnectionsConnectionIDUnauthorized with default headers values
func NewGetV1VcsConnectionsConnectionIDUnauthorized() *GetV1VcsConnectionsConnectionIDUnauthorized {
	return &GetV1VcsConnectionsConnectionIDUnauthorized{}
}

/*
GetV1VcsConnectionsConnectionIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1VcsConnectionsConnectionIDUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 vcs connections connection Id unauthorized response has a 2xx status code
func (o *GetV1VcsConnectionsConnectionIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 vcs connections connection Id unauthorized response has a 3xx status code
func (o *GetV1VcsConnectionsConnectionIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 vcs connections connection Id unauthorized response has a 4xx status code
func (o *GetV1VcsConnectionsConnectionIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 vcs connections connection Id unauthorized response has a 5xx status code
func (o *GetV1VcsConnectionsConnectionIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 vcs connections connection Id unauthorized response a status code equal to that given
func (o *GetV1VcsConnectionsConnectionIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 vcs connections connection Id unauthorized response
func (o *GetV1VcsConnectionsConnectionIDUnauthorized) Code() int {
	return 401
}

func (o *GetV1VcsConnectionsConnectionIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1VcsConnectionsConnectionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1VcsConnectionsConnectionIDForbidden creates a GetV1VcsConnectionsConnectionIDForbidden with default headers values
func NewGetV1VcsConnectionsConnectionIDForbidden() *GetV1VcsConnectionsConnectionIDForbidden {
	return &GetV1VcsConnectionsConnectionIDForbidden{}
}

/*
GetV1VcsConnectionsConnectionIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1VcsConnectionsConnectionIDForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 vcs connections connection Id forbidden response has a 2xx status code
func (o *GetV1VcsConnectionsConnectionIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 vcs connections connection Id forbidden response has a 3xx status code
func (o *GetV1VcsConnectionsConnectionIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 vcs connections connection Id forbidden response has a 4xx status code
func (o *GetV1VcsConnectionsConnectionIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 vcs connections connection Id forbidden response has a 5xx status code
func (o *GetV1VcsConnectionsConnectionIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 vcs connections connection Id forbidden response a status code equal to that given
func (o *GetV1VcsConnectionsConnectionIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 vcs connections connection Id forbidden response
func (o *GetV1VcsConnectionsConnectionIDForbidden) Code() int {
	return 403
}

func (o *GetV1VcsConnectionsConnectionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1VcsConnectionsConnectionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1VcsConnectionsConnectionIDNotFound creates a GetV1VcsConnectionsConnectionIDNotFound with default headers values
func NewGetV1VcsConnectionsConnectionIDNotFound() *GetV1VcsConnectionsConnectionIDNotFound {
	return &GetV1VcsConnectionsConnectionIDNotFound{}
}

/*
GetV1VcsConnectionsConnectionIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1VcsConnectionsConnectionIDNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 vcs connections connection Id not found response has a 2xx status code
func (o *GetV1VcsConnectionsConnectionIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 vcs connections connection Id not found response has a 3xx status code
func (o *GetV1VcsConnectionsConnectionIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 vcs connections connection Id not found response has a 4xx status code
func (o *GetV1VcsConnectionsConnectionIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 vcs connections connection Id not found response has a 5xx status code
func (o *GetV1VcsConnectionsConnectionIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 vcs connections connection Id not found response a status code equal to that given
func (o *GetV1VcsConnectionsConnectionIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 vcs connections connection Id not found response
func (o *GetV1VcsConnectionsConnectionIDNotFound) Code() int {
	return 404
}

func (o *GetV1VcsConnectionsConnectionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1VcsConnectionsConnectionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1VcsConnectionsConnectionIDInternalServerError creates a GetV1VcsConnectionsConnectionIDInternalServerError with default headers values
func NewGetV1VcsConnectionsConnectionIDInternalServerError() *GetV1VcsConnectionsConnectionIDInternalServerError {
	return &GetV1VcsConnectionsConnectionIDInternalServerError{}
}

/*
GetV1VcsConnectionsConnectionIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1VcsConnectionsConnectionIDInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 vcs connections connection Id internal server error response has a 2xx status code
func (o *GetV1VcsConnectionsConnectionIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 vcs connections connection Id internal server error response has a 3xx status code
func (o *GetV1VcsConnectionsConnectionIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 vcs connections connection Id internal server error response has a 4xx status code
func (o *GetV1VcsConnectionsConnectionIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 vcs connections connection Id internal server error response has a 5xx status code
func (o *GetV1VcsConnectionsConnectionIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 vcs connections connection Id internal server error response a status code equal to that given
func (o *GetV1VcsConnectionsConnectionIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 vcs connections connection Id internal server error response
func (o *GetV1VcsConnectionsConnectionIDInternalServerError) Code() int {
	return 500
}

func (o *GetV1VcsConnectionsConnectionIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vcs/connections/{connection_id}][%d] getV1VcsConnectionsConnectionIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1VcsConnectionsConnectionIDInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1VcsConnectionsConnectionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
