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

// GetInstallerReader is a Reader for the GetInstaller structure.
type GetInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installers/{installer_id}] GetInstaller", response, response.Code())
	}
}

// NewGetInstallerOK creates a GetInstallerOK with default headers values
func NewGetInstallerOK() *GetInstallerOK {
	return &GetInstallerOK{}
}

/*
GetInstallerOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallerOK struct {
	Payload *models.AppAppInstaller
}

// IsSuccess returns true when this get installer o k response has a 2xx status code
func (o *GetInstallerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get installer o k response has a 3xx status code
func (o *GetInstallerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get installer o k response has a 4xx status code
func (o *GetInstallerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get installer o k response has a 5xx status code
func (o *GetInstallerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get installer o k response a status code equal to that given
func (o *GetInstallerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get installer o k response
func (o *GetInstallerOK) Code() int {
	return 200
}

func (o *GetInstallerOK) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerOK  %+v", 200, o.Payload)
}

func (o *GetInstallerOK) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerOK  %+v", 200, o.Payload)
}

func (o *GetInstallerOK) GetPayload() *models.AppAppInstaller {
	return o.Payload
}

func (o *GetInstallerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppInstaller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallerBadRequest creates a GetInstallerBadRequest with default headers values
func NewGetInstallerBadRequest() *GetInstallerBadRequest {
	return &GetInstallerBadRequest{}
}

/*
GetInstallerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallerBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get installer bad request response has a 2xx status code
func (o *GetInstallerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get installer bad request response has a 3xx status code
func (o *GetInstallerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get installer bad request response has a 4xx status code
func (o *GetInstallerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get installer bad request response has a 5xx status code
func (o *GetInstallerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get installer bad request response a status code equal to that given
func (o *GetInstallerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get installer bad request response
func (o *GetInstallerBadRequest) Code() int {
	return 400
}

func (o *GetInstallerBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallerBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallerBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallerUnauthorized creates a GetInstallerUnauthorized with default headers values
func NewGetInstallerUnauthorized() *GetInstallerUnauthorized {
	return &GetInstallerUnauthorized{}
}

/*
GetInstallerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallerUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get installer unauthorized response has a 2xx status code
func (o *GetInstallerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get installer unauthorized response has a 3xx status code
func (o *GetInstallerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get installer unauthorized response has a 4xx status code
func (o *GetInstallerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get installer unauthorized response has a 5xx status code
func (o *GetInstallerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get installer unauthorized response a status code equal to that given
func (o *GetInstallerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get installer unauthorized response
func (o *GetInstallerUnauthorized) Code() int {
	return 401
}

func (o *GetInstallerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallerUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallerUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallerForbidden creates a GetInstallerForbidden with default headers values
func NewGetInstallerForbidden() *GetInstallerForbidden {
	return &GetInstallerForbidden{}
}

/*
GetInstallerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallerForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get installer forbidden response has a 2xx status code
func (o *GetInstallerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get installer forbidden response has a 3xx status code
func (o *GetInstallerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get installer forbidden response has a 4xx status code
func (o *GetInstallerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get installer forbidden response has a 5xx status code
func (o *GetInstallerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get installer forbidden response a status code equal to that given
func (o *GetInstallerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get installer forbidden response
func (o *GetInstallerForbidden) Code() int {
	return 403
}

func (o *GetInstallerForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallerForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallerForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallerNotFound creates a GetInstallerNotFound with default headers values
func NewGetInstallerNotFound() *GetInstallerNotFound {
	return &GetInstallerNotFound{}
}

/*
GetInstallerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallerNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get installer not found response has a 2xx status code
func (o *GetInstallerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get installer not found response has a 3xx status code
func (o *GetInstallerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get installer not found response has a 4xx status code
func (o *GetInstallerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get installer not found response has a 5xx status code
func (o *GetInstallerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get installer not found response a status code equal to that given
func (o *GetInstallerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get installer not found response
func (o *GetInstallerNotFound) Code() int {
	return 404
}

func (o *GetInstallerNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallerNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallerNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallerInternalServerError creates a GetInstallerInternalServerError with default headers values
func NewGetInstallerInternalServerError() *GetInstallerInternalServerError {
	return &GetInstallerInternalServerError{}
}

/*
GetInstallerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallerInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get installer internal server error response has a 2xx status code
func (o *GetInstallerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get installer internal server error response has a 3xx status code
func (o *GetInstallerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get installer internal server error response has a 4xx status code
func (o *GetInstallerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get installer internal server error response has a 5xx status code
func (o *GetInstallerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get installer internal server error response a status code equal to that given
func (o *GetInstallerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get installer internal server error response
func (o *GetInstallerInternalServerError) Code() int {
	return 500
}

func (o *GetInstallerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallerInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_id}][%d] getInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallerInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}