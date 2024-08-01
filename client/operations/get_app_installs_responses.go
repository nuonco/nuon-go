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

// GetAppInstallsReader is a Reader for the GetAppInstalls structure.
type GetAppInstallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppInstallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppInstallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppInstallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppInstallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppInstallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppInstallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppInstallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/installs] GetAppInstalls", response, response.Code())
	}
}

// NewGetAppInstallsOK creates a GetAppInstallsOK with default headers values
func NewGetAppInstallsOK() *GetAppInstallsOK {
	return &GetAppInstallsOK{}
}

/*
GetAppInstallsOK describes a response with status code 200, with default header values.

OK
*/
type GetAppInstallsOK struct {
	Payload []*models.AppInstall
}

// IsSuccess returns true when this get app installs o k response has a 2xx status code
func (o *GetAppInstallsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app installs o k response has a 3xx status code
func (o *GetAppInstallsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app installs o k response has a 4xx status code
func (o *GetAppInstallsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app installs o k response has a 5xx status code
func (o *GetAppInstallsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app installs o k response a status code equal to that given
func (o *GetAppInstallsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app installs o k response
func (o *GetAppInstallsOK) Code() int {
	return 200
}

func (o *GetAppInstallsOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsOK  %+v", 200, o.Payload)
}

func (o *GetAppInstallsOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsOK  %+v", 200, o.Payload)
}

func (o *GetAppInstallsOK) GetPayload() []*models.AppInstall {
	return o.Payload
}

func (o *GetAppInstallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInstallsBadRequest creates a GetAppInstallsBadRequest with default headers values
func NewGetAppInstallsBadRequest() *GetAppInstallsBadRequest {
	return &GetAppInstallsBadRequest{}
}

/*
GetAppInstallsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppInstallsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app installs bad request response has a 2xx status code
func (o *GetAppInstallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app installs bad request response has a 3xx status code
func (o *GetAppInstallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app installs bad request response has a 4xx status code
func (o *GetAppInstallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app installs bad request response has a 5xx status code
func (o *GetAppInstallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app installs bad request response a status code equal to that given
func (o *GetAppInstallsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app installs bad request response
func (o *GetAppInstallsBadRequest) Code() int {
	return 400
}

func (o *GetAppInstallsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppInstallsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppInstallsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInstallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInstallsUnauthorized creates a GetAppInstallsUnauthorized with default headers values
func NewGetAppInstallsUnauthorized() *GetAppInstallsUnauthorized {
	return &GetAppInstallsUnauthorized{}
}

/*
GetAppInstallsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppInstallsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app installs unauthorized response has a 2xx status code
func (o *GetAppInstallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app installs unauthorized response has a 3xx status code
func (o *GetAppInstallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app installs unauthorized response has a 4xx status code
func (o *GetAppInstallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app installs unauthorized response has a 5xx status code
func (o *GetAppInstallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app installs unauthorized response a status code equal to that given
func (o *GetAppInstallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app installs unauthorized response
func (o *GetAppInstallsUnauthorized) Code() int {
	return 401
}

func (o *GetAppInstallsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppInstallsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAppInstallsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInstallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInstallsForbidden creates a GetAppInstallsForbidden with default headers values
func NewGetAppInstallsForbidden() *GetAppInstallsForbidden {
	return &GetAppInstallsForbidden{}
}

/*
GetAppInstallsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppInstallsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app installs forbidden response has a 2xx status code
func (o *GetAppInstallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app installs forbidden response has a 3xx status code
func (o *GetAppInstallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app installs forbidden response has a 4xx status code
func (o *GetAppInstallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app installs forbidden response has a 5xx status code
func (o *GetAppInstallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app installs forbidden response a status code equal to that given
func (o *GetAppInstallsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app installs forbidden response
func (o *GetAppInstallsForbidden) Code() int {
	return 403
}

func (o *GetAppInstallsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsForbidden  %+v", 403, o.Payload)
}

func (o *GetAppInstallsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsForbidden  %+v", 403, o.Payload)
}

func (o *GetAppInstallsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInstallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInstallsNotFound creates a GetAppInstallsNotFound with default headers values
func NewGetAppInstallsNotFound() *GetAppInstallsNotFound {
	return &GetAppInstallsNotFound{}
}

/*
GetAppInstallsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppInstallsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app installs not found response has a 2xx status code
func (o *GetAppInstallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app installs not found response has a 3xx status code
func (o *GetAppInstallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app installs not found response has a 4xx status code
func (o *GetAppInstallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app installs not found response has a 5xx status code
func (o *GetAppInstallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app installs not found response a status code equal to that given
func (o *GetAppInstallsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app installs not found response
func (o *GetAppInstallsNotFound) Code() int {
	return 404
}

func (o *GetAppInstallsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsNotFound  %+v", 404, o.Payload)
}

func (o *GetAppInstallsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsNotFound  %+v", 404, o.Payload)
}

func (o *GetAppInstallsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInstallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInstallsInternalServerError creates a GetAppInstallsInternalServerError with default headers values
func NewGetAppInstallsInternalServerError() *GetAppInstallsInternalServerError {
	return &GetAppInstallsInternalServerError{}
}

/*
GetAppInstallsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppInstallsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app installs internal server error response has a 2xx status code
func (o *GetAppInstallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app installs internal server error response has a 3xx status code
func (o *GetAppInstallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app installs internal server error response has a 4xx status code
func (o *GetAppInstallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app installs internal server error response has a 5xx status code
func (o *GetAppInstallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app installs internal server error response a status code equal to that given
func (o *GetAppInstallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app installs internal server error response
func (o *GetAppInstallsInternalServerError) Code() int {
	return 500
}

func (o *GetAppInstallsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppInstallsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getAppInstallsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppInstallsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInstallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
