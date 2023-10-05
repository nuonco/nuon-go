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

// GetV1AppsAppIDInstallsReader is a Reader for the GetV1AppsAppIDInstalls structure.
type GetV1AppsAppIDInstallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AppsAppIDInstallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AppsAppIDInstallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1AppsAppIDInstallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1AppsAppIDInstallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1AppsAppIDInstallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1AppsAppIDInstallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1AppsAppIDInstallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/installs] GetV1AppsAppIDInstalls", response, response.Code())
	}
}

// NewGetV1AppsAppIDInstallsOK creates a GetV1AppsAppIDInstallsOK with default headers values
func NewGetV1AppsAppIDInstallsOK() *GetV1AppsAppIDInstallsOK {
	return &GetV1AppsAppIDInstallsOK{}
}

/*
GetV1AppsAppIDInstallsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1AppsAppIDInstallsOK struct {
	Payload []*models.AppInstall
}

// IsSuccess returns true when this get v1 apps app Id installs o k response has a 2xx status code
func (o *GetV1AppsAppIDInstallsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 apps app Id installs o k response has a 3xx status code
func (o *GetV1AppsAppIDInstallsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id installs o k response has a 4xx status code
func (o *GetV1AppsAppIDInstallsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 apps app Id installs o k response has a 5xx status code
func (o *GetV1AppsAppIDInstallsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id installs o k response a status code equal to that given
func (o *GetV1AppsAppIDInstallsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 apps app Id installs o k response
func (o *GetV1AppsAppIDInstallsOK) Code() int {
	return 200
}

func (o *GetV1AppsAppIDInstallsOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsOK  %+v", 200, o.Payload)
}

func (o *GetV1AppsAppIDInstallsOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsOK  %+v", 200, o.Payload)
}

func (o *GetV1AppsAppIDInstallsOK) GetPayload() []*models.AppInstall {
	return o.Payload
}

func (o *GetV1AppsAppIDInstallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDInstallsBadRequest creates a GetV1AppsAppIDInstallsBadRequest with default headers values
func NewGetV1AppsAppIDInstallsBadRequest() *GetV1AppsAppIDInstallsBadRequest {
	return &GetV1AppsAppIDInstallsBadRequest{}
}

/*
GetV1AppsAppIDInstallsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1AppsAppIDInstallsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id installs bad request response has a 2xx status code
func (o *GetV1AppsAppIDInstallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id installs bad request response has a 3xx status code
func (o *GetV1AppsAppIDInstallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id installs bad request response has a 4xx status code
func (o *GetV1AppsAppIDInstallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id installs bad request response has a 5xx status code
func (o *GetV1AppsAppIDInstallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id installs bad request response a status code equal to that given
func (o *GetV1AppsAppIDInstallsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 apps app Id installs bad request response
func (o *GetV1AppsAppIDInstallsBadRequest) Code() int {
	return 400
}

func (o *GetV1AppsAppIDInstallsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1AppsAppIDInstallsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1AppsAppIDInstallsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDInstallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDInstallsUnauthorized creates a GetV1AppsAppIDInstallsUnauthorized with default headers values
func NewGetV1AppsAppIDInstallsUnauthorized() *GetV1AppsAppIDInstallsUnauthorized {
	return &GetV1AppsAppIDInstallsUnauthorized{}
}

/*
GetV1AppsAppIDInstallsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1AppsAppIDInstallsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id installs unauthorized response has a 2xx status code
func (o *GetV1AppsAppIDInstallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id installs unauthorized response has a 3xx status code
func (o *GetV1AppsAppIDInstallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id installs unauthorized response has a 4xx status code
func (o *GetV1AppsAppIDInstallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id installs unauthorized response has a 5xx status code
func (o *GetV1AppsAppIDInstallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id installs unauthorized response a status code equal to that given
func (o *GetV1AppsAppIDInstallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 apps app Id installs unauthorized response
func (o *GetV1AppsAppIDInstallsUnauthorized) Code() int {
	return 401
}

func (o *GetV1AppsAppIDInstallsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1AppsAppIDInstallsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1AppsAppIDInstallsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDInstallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDInstallsForbidden creates a GetV1AppsAppIDInstallsForbidden with default headers values
func NewGetV1AppsAppIDInstallsForbidden() *GetV1AppsAppIDInstallsForbidden {
	return &GetV1AppsAppIDInstallsForbidden{}
}

/*
GetV1AppsAppIDInstallsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1AppsAppIDInstallsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id installs forbidden response has a 2xx status code
func (o *GetV1AppsAppIDInstallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id installs forbidden response has a 3xx status code
func (o *GetV1AppsAppIDInstallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id installs forbidden response has a 4xx status code
func (o *GetV1AppsAppIDInstallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id installs forbidden response has a 5xx status code
func (o *GetV1AppsAppIDInstallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id installs forbidden response a status code equal to that given
func (o *GetV1AppsAppIDInstallsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 apps app Id installs forbidden response
func (o *GetV1AppsAppIDInstallsForbidden) Code() int {
	return 403
}

func (o *GetV1AppsAppIDInstallsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsForbidden  %+v", 403, o.Payload)
}

func (o *GetV1AppsAppIDInstallsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsForbidden  %+v", 403, o.Payload)
}

func (o *GetV1AppsAppIDInstallsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDInstallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDInstallsNotFound creates a GetV1AppsAppIDInstallsNotFound with default headers values
func NewGetV1AppsAppIDInstallsNotFound() *GetV1AppsAppIDInstallsNotFound {
	return &GetV1AppsAppIDInstallsNotFound{}
}

/*
GetV1AppsAppIDInstallsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1AppsAppIDInstallsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id installs not found response has a 2xx status code
func (o *GetV1AppsAppIDInstallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id installs not found response has a 3xx status code
func (o *GetV1AppsAppIDInstallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id installs not found response has a 4xx status code
func (o *GetV1AppsAppIDInstallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 apps app Id installs not found response has a 5xx status code
func (o *GetV1AppsAppIDInstallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id installs not found response a status code equal to that given
func (o *GetV1AppsAppIDInstallsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 apps app Id installs not found response
func (o *GetV1AppsAppIDInstallsNotFound) Code() int {
	return 404
}

func (o *GetV1AppsAppIDInstallsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsNotFound  %+v", 404, o.Payload)
}

func (o *GetV1AppsAppIDInstallsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsNotFound  %+v", 404, o.Payload)
}

func (o *GetV1AppsAppIDInstallsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDInstallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1AppsAppIDInstallsInternalServerError creates a GetV1AppsAppIDInstallsInternalServerError with default headers values
func NewGetV1AppsAppIDInstallsInternalServerError() *GetV1AppsAppIDInstallsInternalServerError {
	return &GetV1AppsAppIDInstallsInternalServerError{}
}

/*
GetV1AppsAppIDInstallsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1AppsAppIDInstallsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 apps app Id installs internal server error response has a 2xx status code
func (o *GetV1AppsAppIDInstallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 apps app Id installs internal server error response has a 3xx status code
func (o *GetV1AppsAppIDInstallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id installs internal server error response has a 4xx status code
func (o *GetV1AppsAppIDInstallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 apps app Id installs internal server error response has a 5xx status code
func (o *GetV1AppsAppIDInstallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 apps app Id installs internal server error response a status code equal to that given
func (o *GetV1AppsAppIDInstallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 apps app Id installs internal server error response
func (o *GetV1AppsAppIDInstallsInternalServerError) Code() int {
	return 500
}

func (o *GetV1AppsAppIDInstallsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1AppsAppIDInstallsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/installs][%d] getV1AppsAppIdInstallsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1AppsAppIDInstallsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1AppsAppIDInstallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
