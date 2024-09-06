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

// CreateAppReader is a Reader for the CreateApp structure.
type CreateAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps] CreateApp", response, response.Code())
	}
}

// NewCreateAppCreated creates a CreateAppCreated with default headers values
func NewCreateAppCreated() *CreateAppCreated {
	return &CreateAppCreated{}
}

/*
CreateAppCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppCreated struct {
	Payload *models.AppApp
}

// IsSuccess returns true when this create app created response has a 2xx status code
func (o *CreateAppCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app created response has a 3xx status code
func (o *CreateAppCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app created response has a 4xx status code
func (o *CreateAppCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app created response has a 5xx status code
func (o *CreateAppCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app created response a status code equal to that given
func (o *CreateAppCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app created response
func (o *CreateAppCreated) Code() int {
	return 201
}

func (o *CreateAppCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppCreated %s", 201, payload)
}

func (o *CreateAppCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppCreated %s", 201, payload)
}

func (o *CreateAppCreated) GetPayload() *models.AppApp {
	return o.Payload
}

func (o *CreateAppCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppBadRequest creates a CreateAppBadRequest with default headers values
func NewCreateAppBadRequest() *CreateAppBadRequest {
	return &CreateAppBadRequest{}
}

/*
CreateAppBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app bad request response has a 2xx status code
func (o *CreateAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app bad request response has a 3xx status code
func (o *CreateAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app bad request response has a 4xx status code
func (o *CreateAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app bad request response has a 5xx status code
func (o *CreateAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app bad request response a status code equal to that given
func (o *CreateAppBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app bad request response
func (o *CreateAppBadRequest) Code() int {
	return 400
}

func (o *CreateAppBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppBadRequest %s", 400, payload)
}

func (o *CreateAppBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppBadRequest %s", 400, payload)
}

func (o *CreateAppBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppUnauthorized creates a CreateAppUnauthorized with default headers values
func NewCreateAppUnauthorized() *CreateAppUnauthorized {
	return &CreateAppUnauthorized{}
}

/*
CreateAppUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app unauthorized response has a 2xx status code
func (o *CreateAppUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app unauthorized response has a 3xx status code
func (o *CreateAppUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app unauthorized response has a 4xx status code
func (o *CreateAppUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app unauthorized response has a 5xx status code
func (o *CreateAppUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app unauthorized response a status code equal to that given
func (o *CreateAppUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app unauthorized response
func (o *CreateAppUnauthorized) Code() int {
	return 401
}

func (o *CreateAppUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppUnauthorized %s", 401, payload)
}

func (o *CreateAppUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppUnauthorized %s", 401, payload)
}

func (o *CreateAppUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppForbidden creates a CreateAppForbidden with default headers values
func NewCreateAppForbidden() *CreateAppForbidden {
	return &CreateAppForbidden{}
}

/*
CreateAppForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app forbidden response has a 2xx status code
func (o *CreateAppForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app forbidden response has a 3xx status code
func (o *CreateAppForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app forbidden response has a 4xx status code
func (o *CreateAppForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app forbidden response has a 5xx status code
func (o *CreateAppForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app forbidden response a status code equal to that given
func (o *CreateAppForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app forbidden response
func (o *CreateAppForbidden) Code() int {
	return 403
}

func (o *CreateAppForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppForbidden %s", 403, payload)
}

func (o *CreateAppForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppForbidden %s", 403, payload)
}

func (o *CreateAppForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppNotFound creates a CreateAppNotFound with default headers values
func NewCreateAppNotFound() *CreateAppNotFound {
	return &CreateAppNotFound{}
}

/*
CreateAppNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app not found response has a 2xx status code
func (o *CreateAppNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app not found response has a 3xx status code
func (o *CreateAppNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app not found response has a 4xx status code
func (o *CreateAppNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app not found response has a 5xx status code
func (o *CreateAppNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app not found response a status code equal to that given
func (o *CreateAppNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app not found response
func (o *CreateAppNotFound) Code() int {
	return 404
}

func (o *CreateAppNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppNotFound %s", 404, payload)
}

func (o *CreateAppNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppNotFound %s", 404, payload)
}

func (o *CreateAppNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppInternalServerError creates a CreateAppInternalServerError with default headers values
func NewCreateAppInternalServerError() *CreateAppInternalServerError {
	return &CreateAppInternalServerError{}
}

/*
CreateAppInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app internal server error response has a 2xx status code
func (o *CreateAppInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app internal server error response has a 3xx status code
func (o *CreateAppInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app internal server error response has a 4xx status code
func (o *CreateAppInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app internal server error response has a 5xx status code
func (o *CreateAppInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app internal server error response a status code equal to that given
func (o *CreateAppInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app internal server error response
func (o *CreateAppInternalServerError) Code() int {
	return 500
}

func (o *CreateAppInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppInternalServerError %s", 500, payload)
}

func (o *CreateAppInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps][%d] createAppInternalServerError %s", 500, payload)
}

func (o *CreateAppInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
