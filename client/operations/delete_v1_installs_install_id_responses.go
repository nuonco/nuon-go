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

// DeleteV1InstallsInstallIDReader is a Reader for the DeleteV1InstallsInstallID structure.
type DeleteV1InstallsInstallIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1InstallsInstallIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1InstallsInstallIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteV1InstallsInstallIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteV1InstallsInstallIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteV1InstallsInstallIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteV1InstallsInstallIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteV1InstallsInstallIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/installs/{install_id}] DeleteV1InstallsInstallID", response, response.Code())
	}
}

// NewDeleteV1InstallsInstallIDOK creates a DeleteV1InstallsInstallIDOK with default headers values
func NewDeleteV1InstallsInstallIDOK() *DeleteV1InstallsInstallIDOK {
	return &DeleteV1InstallsInstallIDOK{}
}

/*
DeleteV1InstallsInstallIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteV1InstallsInstallIDOK struct {
	Payload bool
}

// IsSuccess returns true when this delete v1 installs install Id o k response has a 2xx status code
func (o *DeleteV1InstallsInstallIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 installs install Id o k response has a 3xx status code
func (o *DeleteV1InstallsInstallIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 installs install Id o k response has a 4xx status code
func (o *DeleteV1InstallsInstallIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 installs install Id o k response has a 5xx status code
func (o *DeleteV1InstallsInstallIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 installs install Id o k response a status code equal to that given
func (o *DeleteV1InstallsInstallIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 installs install Id o k response
func (o *DeleteV1InstallsInstallIDOK) Code() int {
	return 200
}

func (o *DeleteV1InstallsInstallIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1InstallsInstallIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1InstallsInstallIDOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteV1InstallsInstallIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV1InstallsInstallIDBadRequest creates a DeleteV1InstallsInstallIDBadRequest with default headers values
func NewDeleteV1InstallsInstallIDBadRequest() *DeleteV1InstallsInstallIDBadRequest {
	return &DeleteV1InstallsInstallIDBadRequest{}
}

/*
DeleteV1InstallsInstallIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteV1InstallsInstallIDBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete v1 installs install Id bad request response has a 2xx status code
func (o *DeleteV1InstallsInstallIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 installs install Id bad request response has a 3xx status code
func (o *DeleteV1InstallsInstallIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 installs install Id bad request response has a 4xx status code
func (o *DeleteV1InstallsInstallIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete v1 installs install Id bad request response has a 5xx status code
func (o *DeleteV1InstallsInstallIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 installs install Id bad request response a status code equal to that given
func (o *DeleteV1InstallsInstallIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete v1 installs install Id bad request response
func (o *DeleteV1InstallsInstallIDBadRequest) Code() int {
	return 400
}

func (o *DeleteV1InstallsInstallIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteV1InstallsInstallIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteV1InstallsInstallIDBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteV1InstallsInstallIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV1InstallsInstallIDUnauthorized creates a DeleteV1InstallsInstallIDUnauthorized with default headers values
func NewDeleteV1InstallsInstallIDUnauthorized() *DeleteV1InstallsInstallIDUnauthorized {
	return &DeleteV1InstallsInstallIDUnauthorized{}
}

/*
DeleteV1InstallsInstallIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteV1InstallsInstallIDUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete v1 installs install Id unauthorized response has a 2xx status code
func (o *DeleteV1InstallsInstallIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 installs install Id unauthorized response has a 3xx status code
func (o *DeleteV1InstallsInstallIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 installs install Id unauthorized response has a 4xx status code
func (o *DeleteV1InstallsInstallIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete v1 installs install Id unauthorized response has a 5xx status code
func (o *DeleteV1InstallsInstallIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 installs install Id unauthorized response a status code equal to that given
func (o *DeleteV1InstallsInstallIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete v1 installs install Id unauthorized response
func (o *DeleteV1InstallsInstallIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteV1InstallsInstallIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteV1InstallsInstallIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteV1InstallsInstallIDUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteV1InstallsInstallIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV1InstallsInstallIDForbidden creates a DeleteV1InstallsInstallIDForbidden with default headers values
func NewDeleteV1InstallsInstallIDForbidden() *DeleteV1InstallsInstallIDForbidden {
	return &DeleteV1InstallsInstallIDForbidden{}
}

/*
DeleteV1InstallsInstallIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteV1InstallsInstallIDForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete v1 installs install Id forbidden response has a 2xx status code
func (o *DeleteV1InstallsInstallIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 installs install Id forbidden response has a 3xx status code
func (o *DeleteV1InstallsInstallIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 installs install Id forbidden response has a 4xx status code
func (o *DeleteV1InstallsInstallIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete v1 installs install Id forbidden response has a 5xx status code
func (o *DeleteV1InstallsInstallIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 installs install Id forbidden response a status code equal to that given
func (o *DeleteV1InstallsInstallIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete v1 installs install Id forbidden response
func (o *DeleteV1InstallsInstallIDForbidden) Code() int {
	return 403
}

func (o *DeleteV1InstallsInstallIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteV1InstallsInstallIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteV1InstallsInstallIDForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteV1InstallsInstallIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV1InstallsInstallIDNotFound creates a DeleteV1InstallsInstallIDNotFound with default headers values
func NewDeleteV1InstallsInstallIDNotFound() *DeleteV1InstallsInstallIDNotFound {
	return &DeleteV1InstallsInstallIDNotFound{}
}

/*
DeleteV1InstallsInstallIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteV1InstallsInstallIDNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete v1 installs install Id not found response has a 2xx status code
func (o *DeleteV1InstallsInstallIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 installs install Id not found response has a 3xx status code
func (o *DeleteV1InstallsInstallIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 installs install Id not found response has a 4xx status code
func (o *DeleteV1InstallsInstallIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete v1 installs install Id not found response has a 5xx status code
func (o *DeleteV1InstallsInstallIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 installs install Id not found response a status code equal to that given
func (o *DeleteV1InstallsInstallIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete v1 installs install Id not found response
func (o *DeleteV1InstallsInstallIDNotFound) Code() int {
	return 404
}

func (o *DeleteV1InstallsInstallIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteV1InstallsInstallIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteV1InstallsInstallIDNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteV1InstallsInstallIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteV1InstallsInstallIDInternalServerError creates a DeleteV1InstallsInstallIDInternalServerError with default headers values
func NewDeleteV1InstallsInstallIDInternalServerError() *DeleteV1InstallsInstallIDInternalServerError {
	return &DeleteV1InstallsInstallIDInternalServerError{}
}

/*
DeleteV1InstallsInstallIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteV1InstallsInstallIDInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete v1 installs install Id internal server error response has a 2xx status code
func (o *DeleteV1InstallsInstallIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 installs install Id internal server error response has a 3xx status code
func (o *DeleteV1InstallsInstallIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 installs install Id internal server error response has a 4xx status code
func (o *DeleteV1InstallsInstallIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 installs install Id internal server error response has a 5xx status code
func (o *DeleteV1InstallsInstallIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete v1 installs install Id internal server error response a status code equal to that given
func (o *DeleteV1InstallsInstallIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete v1 installs install Id internal server error response
func (o *DeleteV1InstallsInstallIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteV1InstallsInstallIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteV1InstallsInstallIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteV1InstallsInstallIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteV1InstallsInstallIDInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteV1InstallsInstallIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
