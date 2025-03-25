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

// DeleteAppReader is a Reader for the DeleteApp structure.
type DeleteAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/apps/{app_id}] DeleteApp", response, response.Code())
	}
}

// NewDeleteAppOK creates a DeleteAppOK with default headers values
func NewDeleteAppOK() *DeleteAppOK {
	return &DeleteAppOK{}
}

/*
DeleteAppOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAppOK struct {
	Payload bool
}

// IsSuccess returns true when this delete app o k response has a 2xx status code
func (o *DeleteAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete app o k response has a 3xx status code
func (o *DeleteAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app o k response has a 4xx status code
func (o *DeleteAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete app o k response has a 5xx status code
func (o *DeleteAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app o k response a status code equal to that given
func (o *DeleteAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete app o k response
func (o *DeleteAppOK) Code() int {
	return 200
}

func (o *DeleteAppOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppOK  %+v", 200, o.Payload)
}

func (o *DeleteAppOK) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppOK  %+v", 200, o.Payload)
}

func (o *DeleteAppOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppBadRequest creates a DeleteAppBadRequest with default headers values
func NewDeleteAppBadRequest() *DeleteAppBadRequest {
	return &DeleteAppBadRequest{}
}

/*
DeleteAppBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAppBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete app bad request response has a 2xx status code
func (o *DeleteAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete app bad request response has a 3xx status code
func (o *DeleteAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app bad request response has a 4xx status code
func (o *DeleteAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete app bad request response has a 5xx status code
func (o *DeleteAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app bad request response a status code equal to that given
func (o *DeleteAppBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete app bad request response
func (o *DeleteAppBadRequest) Code() int {
	return 400
}

func (o *DeleteAppBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAppBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAppBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppUnauthorized creates a DeleteAppUnauthorized with default headers values
func NewDeleteAppUnauthorized() *DeleteAppUnauthorized {
	return &DeleteAppUnauthorized{}
}

/*
DeleteAppUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteAppUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete app unauthorized response has a 2xx status code
func (o *DeleteAppUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete app unauthorized response has a 3xx status code
func (o *DeleteAppUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app unauthorized response has a 4xx status code
func (o *DeleteAppUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete app unauthorized response has a 5xx status code
func (o *DeleteAppUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app unauthorized response a status code equal to that given
func (o *DeleteAppUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete app unauthorized response
func (o *DeleteAppUnauthorized) Code() int {
	return 401
}

func (o *DeleteAppUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAppUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAppUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppForbidden creates a DeleteAppForbidden with default headers values
func NewDeleteAppForbidden() *DeleteAppForbidden {
	return &DeleteAppForbidden{}
}

/*
DeleteAppForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAppForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete app forbidden response has a 2xx status code
func (o *DeleteAppForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete app forbidden response has a 3xx status code
func (o *DeleteAppForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app forbidden response has a 4xx status code
func (o *DeleteAppForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete app forbidden response has a 5xx status code
func (o *DeleteAppForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app forbidden response a status code equal to that given
func (o *DeleteAppForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete app forbidden response
func (o *DeleteAppForbidden) Code() int {
	return 403
}

func (o *DeleteAppForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAppForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAppForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppNotFound creates a DeleteAppNotFound with default headers values
func NewDeleteAppNotFound() *DeleteAppNotFound {
	return &DeleteAppNotFound{}
}

/*
DeleteAppNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAppNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete app not found response has a 2xx status code
func (o *DeleteAppNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete app not found response has a 3xx status code
func (o *DeleteAppNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app not found response has a 4xx status code
func (o *DeleteAppNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete app not found response has a 5xx status code
func (o *DeleteAppNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete app not found response a status code equal to that given
func (o *DeleteAppNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete app not found response
func (o *DeleteAppNotFound) Code() int {
	return 404
}

func (o *DeleteAppNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAppNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAppNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppInternalServerError creates a DeleteAppInternalServerError with default headers values
func NewDeleteAppInternalServerError() *DeleteAppInternalServerError {
	return &DeleteAppInternalServerError{}
}

/*
DeleteAppInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAppInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete app internal server error response has a 2xx status code
func (o *DeleteAppInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete app internal server error response has a 3xx status code
func (o *DeleteAppInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete app internal server error response has a 4xx status code
func (o *DeleteAppInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete app internal server error response has a 5xx status code
func (o *DeleteAppInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete app internal server error response a status code equal to that given
func (o *DeleteAppInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete app internal server error response
func (o *DeleteAppInternalServerError) Code() int {
	return 500
}

func (o *DeleteAppInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAppInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/{app_id}][%d] deleteAppInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAppInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
