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

// DeleteInstallerReader is a Reader for the DeleteInstaller structure.
type DeleteInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInstallerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteInstallerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteInstallerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteInstallerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInstallerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInstallerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/installers/{installer_id}] DeleteInstaller", response, response.Code())
	}
}

// NewDeleteInstallerOK creates a DeleteInstallerOK with default headers values
func NewDeleteInstallerOK() *DeleteInstallerOK {
	return &DeleteInstallerOK{}
}

/*
DeleteInstallerOK describes a response with status code 200, with default header values.

OK
*/
type DeleteInstallerOK struct {
	Payload bool
}

// IsSuccess returns true when this delete installer o k response has a 2xx status code
func (o *DeleteInstallerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete installer o k response has a 3xx status code
func (o *DeleteInstallerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete installer o k response has a 4xx status code
func (o *DeleteInstallerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete installer o k response has a 5xx status code
func (o *DeleteInstallerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete installer o k response a status code equal to that given
func (o *DeleteInstallerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete installer o k response
func (o *DeleteInstallerOK) Code() int {
	return 200
}

func (o *DeleteInstallerOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerOK  %+v", 200, o.Payload)
}

func (o *DeleteInstallerOK) String() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerOK  %+v", 200, o.Payload)
}

func (o *DeleteInstallerOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteInstallerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallerBadRequest creates a DeleteInstallerBadRequest with default headers values
func NewDeleteInstallerBadRequest() *DeleteInstallerBadRequest {
	return &DeleteInstallerBadRequest{}
}

/*
DeleteInstallerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteInstallerBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete installer bad request response has a 2xx status code
func (o *DeleteInstallerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete installer bad request response has a 3xx status code
func (o *DeleteInstallerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete installer bad request response has a 4xx status code
func (o *DeleteInstallerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete installer bad request response has a 5xx status code
func (o *DeleteInstallerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete installer bad request response a status code equal to that given
func (o *DeleteInstallerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete installer bad request response
func (o *DeleteInstallerBadRequest) Code() int {
	return 400
}

func (o *DeleteInstallerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteInstallerBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteInstallerBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallerUnauthorized creates a DeleteInstallerUnauthorized with default headers values
func NewDeleteInstallerUnauthorized() *DeleteInstallerUnauthorized {
	return &DeleteInstallerUnauthorized{}
}

/*
DeleteInstallerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteInstallerUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete installer unauthorized response has a 2xx status code
func (o *DeleteInstallerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete installer unauthorized response has a 3xx status code
func (o *DeleteInstallerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete installer unauthorized response has a 4xx status code
func (o *DeleteInstallerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete installer unauthorized response has a 5xx status code
func (o *DeleteInstallerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete installer unauthorized response a status code equal to that given
func (o *DeleteInstallerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete installer unauthorized response
func (o *DeleteInstallerUnauthorized) Code() int {
	return 401
}

func (o *DeleteInstallerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteInstallerUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteInstallerUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallerForbidden creates a DeleteInstallerForbidden with default headers values
func NewDeleteInstallerForbidden() *DeleteInstallerForbidden {
	return &DeleteInstallerForbidden{}
}

/*
DeleteInstallerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteInstallerForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete installer forbidden response has a 2xx status code
func (o *DeleteInstallerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete installer forbidden response has a 3xx status code
func (o *DeleteInstallerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete installer forbidden response has a 4xx status code
func (o *DeleteInstallerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete installer forbidden response has a 5xx status code
func (o *DeleteInstallerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete installer forbidden response a status code equal to that given
func (o *DeleteInstallerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete installer forbidden response
func (o *DeleteInstallerForbidden) Code() int {
	return 403
}

func (o *DeleteInstallerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInstallerForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInstallerForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallerNotFound creates a DeleteInstallerNotFound with default headers values
func NewDeleteInstallerNotFound() *DeleteInstallerNotFound {
	return &DeleteInstallerNotFound{}
}

/*
DeleteInstallerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteInstallerNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete installer not found response has a 2xx status code
func (o *DeleteInstallerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete installer not found response has a 3xx status code
func (o *DeleteInstallerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete installer not found response has a 4xx status code
func (o *DeleteInstallerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete installer not found response has a 5xx status code
func (o *DeleteInstallerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete installer not found response a status code equal to that given
func (o *DeleteInstallerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete installer not found response
func (o *DeleteInstallerNotFound) Code() int {
	return 404
}

func (o *DeleteInstallerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstallerNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstallerNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallerInternalServerError creates a DeleteInstallerInternalServerError with default headers values
func NewDeleteInstallerInternalServerError() *DeleteInstallerInternalServerError {
	return &DeleteInstallerInternalServerError{}
}

/*
DeleteInstallerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteInstallerInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete installer internal server error response has a 2xx status code
func (o *DeleteInstallerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete installer internal server error response has a 3xx status code
func (o *DeleteInstallerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete installer internal server error response has a 4xx status code
func (o *DeleteInstallerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete installer internal server error response has a 5xx status code
func (o *DeleteInstallerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete installer internal server error response a status code equal to that given
func (o *DeleteInstallerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete installer internal server error response
func (o *DeleteInstallerInternalServerError) Code() int {
	return 500
}

func (o *DeleteInstallerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteInstallerInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/installers/{installer_id}][%d] deleteInstallerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteInstallerInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
