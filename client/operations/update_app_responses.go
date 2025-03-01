// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-go/models"
)

// UpdateAppReader is a Reader for the UpdateApp structure.
type UpdateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/apps/{app_id}] UpdateApp", response, response.Code())
	}
}

// NewUpdateAppOK creates a UpdateAppOK with default headers values
func NewUpdateAppOK() *UpdateAppOK {
	return &UpdateAppOK{}
}

/*
UpdateAppOK describes a response with status code 200, with default header values.

OK
*/
type UpdateAppOK struct {
	Payload *models.AppApp
}

// IsSuccess returns true when this update app o k response has a 2xx status code
func (o *UpdateAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update app o k response has a 3xx status code
func (o *UpdateAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app o k response has a 4xx status code
func (o *UpdateAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update app o k response has a 5xx status code
func (o *UpdateAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update app o k response a status code equal to that given
func (o *UpdateAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update app o k response
func (o *UpdateAppOK) Code() int {
	return 200
}

func (o *UpdateAppOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppOK %s", 200, payload)
}

func (o *UpdateAppOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppOK %s", 200, payload)
}

func (o *UpdateAppOK) GetPayload() *models.AppApp {
	return o.Payload
}

func (o *UpdateAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppBadRequest creates a UpdateAppBadRequest with default headers values
func NewUpdateAppBadRequest() *UpdateAppBadRequest {
	return &UpdateAppBadRequest{}
}

/*
UpdateAppBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateAppBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app bad request response has a 2xx status code
func (o *UpdateAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app bad request response has a 3xx status code
func (o *UpdateAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app bad request response has a 4xx status code
func (o *UpdateAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app bad request response has a 5xx status code
func (o *UpdateAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update app bad request response a status code equal to that given
func (o *UpdateAppBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update app bad request response
func (o *UpdateAppBadRequest) Code() int {
	return 400
}

func (o *UpdateAppBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppBadRequest %s", 400, payload)
}

func (o *UpdateAppBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppBadRequest %s", 400, payload)
}

func (o *UpdateAppBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppUnauthorized creates a UpdateAppUnauthorized with default headers values
func NewUpdateAppUnauthorized() *UpdateAppUnauthorized {
	return &UpdateAppUnauthorized{}
}

/*
UpdateAppUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateAppUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app unauthorized response has a 2xx status code
func (o *UpdateAppUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app unauthorized response has a 3xx status code
func (o *UpdateAppUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app unauthorized response has a 4xx status code
func (o *UpdateAppUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app unauthorized response has a 5xx status code
func (o *UpdateAppUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update app unauthorized response a status code equal to that given
func (o *UpdateAppUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update app unauthorized response
func (o *UpdateAppUnauthorized) Code() int {
	return 401
}

func (o *UpdateAppUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppUnauthorized %s", 401, payload)
}

func (o *UpdateAppUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppUnauthorized %s", 401, payload)
}

func (o *UpdateAppUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppForbidden creates a UpdateAppForbidden with default headers values
func NewUpdateAppForbidden() *UpdateAppForbidden {
	return &UpdateAppForbidden{}
}

/*
UpdateAppForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateAppForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app forbidden response has a 2xx status code
func (o *UpdateAppForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app forbidden response has a 3xx status code
func (o *UpdateAppForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app forbidden response has a 4xx status code
func (o *UpdateAppForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app forbidden response has a 5xx status code
func (o *UpdateAppForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update app forbidden response a status code equal to that given
func (o *UpdateAppForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update app forbidden response
func (o *UpdateAppForbidden) Code() int {
	return 403
}

func (o *UpdateAppForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppForbidden %s", 403, payload)
}

func (o *UpdateAppForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppForbidden %s", 403, payload)
}

func (o *UpdateAppForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppNotFound creates a UpdateAppNotFound with default headers values
func NewUpdateAppNotFound() *UpdateAppNotFound {
	return &UpdateAppNotFound{}
}

/*
UpdateAppNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAppNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app not found response has a 2xx status code
func (o *UpdateAppNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app not found response has a 3xx status code
func (o *UpdateAppNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app not found response has a 4xx status code
func (o *UpdateAppNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update app not found response has a 5xx status code
func (o *UpdateAppNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update app not found response a status code equal to that given
func (o *UpdateAppNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update app not found response
func (o *UpdateAppNotFound) Code() int {
	return 404
}

func (o *UpdateAppNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppNotFound %s", 404, payload)
}

func (o *UpdateAppNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppNotFound %s", 404, payload)
}

func (o *UpdateAppNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAppInternalServerError creates a UpdateAppInternalServerError with default headers values
func NewUpdateAppInternalServerError() *UpdateAppInternalServerError {
	return &UpdateAppInternalServerError{}
}

/*
UpdateAppInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateAppInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update app internal server error response has a 2xx status code
func (o *UpdateAppInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update app internal server error response has a 3xx status code
func (o *UpdateAppInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update app internal server error response has a 4xx status code
func (o *UpdateAppInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update app internal server error response has a 5xx status code
func (o *UpdateAppInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update app internal server error response a status code equal to that given
func (o *UpdateAppInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update app internal server error response
func (o *UpdateAppInternalServerError) Code() int {
	return 500
}

func (o *UpdateAppInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppInternalServerError %s", 500, payload)
}

func (o *UpdateAppInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/apps/{app_id}][%d] updateAppInternalServerError %s", 500, payload)
}

func (o *UpdateAppInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
