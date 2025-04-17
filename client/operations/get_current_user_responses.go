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

// GetCurrentUserReader is a Reader for the GetCurrentUser structure.
type GetCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCurrentUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCurrentUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCurrentUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/general/current-user] GetCurrentUser", response, response.Code())
	}
}

// NewGetCurrentUserOK creates a GetCurrentUserOK with default headers values
func NewGetCurrentUserOK() *GetCurrentUserOK {
	return &GetCurrentUserOK{}
}

/*
GetCurrentUserOK describes a response with status code 200, with default header values.

OK
*/
type GetCurrentUserOK struct {
	Payload *models.AppAccount
}

// IsSuccess returns true when this get current user o k response has a 2xx status code
func (o *GetCurrentUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current user o k response has a 3xx status code
func (o *GetCurrentUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user o k response has a 4xx status code
func (o *GetCurrentUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current user o k response has a 5xx status code
func (o *GetCurrentUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user o k response a status code equal to that given
func (o *GetCurrentUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get current user o k response
func (o *GetCurrentUserOK) Code() int {
	return 200
}

func (o *GetCurrentUserOK) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserOK) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserOK) GetPayload() *models.AppAccount {
	return o.Payload
}

func (o *GetCurrentUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserBadRequest creates a GetCurrentUserBadRequest with default headers values
func NewGetCurrentUserBadRequest() *GetCurrentUserBadRequest {
	return &GetCurrentUserBadRequest{}
}

/*
GetCurrentUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCurrentUserBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current user bad request response has a 2xx status code
func (o *GetCurrentUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user bad request response has a 3xx status code
func (o *GetCurrentUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user bad request response has a 4xx status code
func (o *GetCurrentUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current user bad request response has a 5xx status code
func (o *GetCurrentUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user bad request response a status code equal to that given
func (o *GetCurrentUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get current user bad request response
func (o *GetCurrentUserBadRequest) Code() int {
	return 400
}

func (o *GetCurrentUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetCurrentUserBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetCurrentUserBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserUnauthorized creates a GetCurrentUserUnauthorized with default headers values
func NewGetCurrentUserUnauthorized() *GetCurrentUserUnauthorized {
	return &GetCurrentUserUnauthorized{}
}

/*
GetCurrentUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCurrentUserUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current user unauthorized response has a 2xx status code
func (o *GetCurrentUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user unauthorized response has a 3xx status code
func (o *GetCurrentUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user unauthorized response has a 4xx status code
func (o *GetCurrentUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current user unauthorized response has a 5xx status code
func (o *GetCurrentUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user unauthorized response a status code equal to that given
func (o *GetCurrentUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get current user unauthorized response
func (o *GetCurrentUserUnauthorized) Code() int {
	return 401
}

func (o *GetCurrentUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCurrentUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCurrentUserUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserForbidden creates a GetCurrentUserForbidden with default headers values
func NewGetCurrentUserForbidden() *GetCurrentUserForbidden {
	return &GetCurrentUserForbidden{}
}

/*
GetCurrentUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCurrentUserForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current user forbidden response has a 2xx status code
func (o *GetCurrentUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user forbidden response has a 3xx status code
func (o *GetCurrentUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user forbidden response has a 4xx status code
func (o *GetCurrentUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current user forbidden response has a 5xx status code
func (o *GetCurrentUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user forbidden response a status code equal to that given
func (o *GetCurrentUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get current user forbidden response
func (o *GetCurrentUserForbidden) Code() int {
	return 403
}

func (o *GetCurrentUserForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserForbidden  %+v", 403, o.Payload)
}

func (o *GetCurrentUserForbidden) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserForbidden  %+v", 403, o.Payload)
}

func (o *GetCurrentUserForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserNotFound creates a GetCurrentUserNotFound with default headers values
func NewGetCurrentUserNotFound() *GetCurrentUserNotFound {
	return &GetCurrentUserNotFound{}
}

/*
GetCurrentUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCurrentUserNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current user not found response has a 2xx status code
func (o *GetCurrentUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user not found response has a 3xx status code
func (o *GetCurrentUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user not found response has a 4xx status code
func (o *GetCurrentUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get current user not found response has a 5xx status code
func (o *GetCurrentUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user not found response a status code equal to that given
func (o *GetCurrentUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get current user not found response
func (o *GetCurrentUserNotFound) Code() int {
	return 404
}

func (o *GetCurrentUserNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrentUserNotFound) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrentUserNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserInternalServerError creates a GetCurrentUserInternalServerError with default headers values
func NewGetCurrentUserInternalServerError() *GetCurrentUserInternalServerError {
	return &GetCurrentUserInternalServerError{}
}

/*
GetCurrentUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCurrentUserInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get current user internal server error response has a 2xx status code
func (o *GetCurrentUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get current user internal server error response has a 3xx status code
func (o *GetCurrentUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user internal server error response has a 4xx status code
func (o *GetCurrentUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current user internal server error response has a 5xx status code
func (o *GetCurrentUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get current user internal server error response a status code equal to that given
func (o *GetCurrentUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get current user internal server error response
func (o *GetCurrentUserInternalServerError) Code() int {
	return 500
}

func (o *GetCurrentUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCurrentUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getCurrentUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCurrentUserInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCurrentUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
