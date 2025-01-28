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

// CreateAppSecretReader is a Reader for the CreateAppSecret structure.
type CreateAppSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAppSecretCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAppSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAppSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAppSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAppSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAppSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/apps/{app_id}/secret] CreateAppSecret", response, response.Code())
	}
}

// NewCreateAppSecretCreated creates a CreateAppSecretCreated with default headers values
func NewCreateAppSecretCreated() *CreateAppSecretCreated {
	return &CreateAppSecretCreated{}
}

/*
CreateAppSecretCreated describes a response with status code 201, with default header values.

Created
*/
type CreateAppSecretCreated struct {
	Payload *models.AppAppSecret
}

// IsSuccess returns true when this create app secret created response has a 2xx status code
func (o *CreateAppSecretCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app secret created response has a 3xx status code
func (o *CreateAppSecretCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app secret created response has a 4xx status code
func (o *CreateAppSecretCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app secret created response has a 5xx status code
func (o *CreateAppSecretCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create app secret created response a status code equal to that given
func (o *CreateAppSecretCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create app secret created response
func (o *CreateAppSecretCreated) Code() int {
	return 201
}

func (o *CreateAppSecretCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretCreated %s", 201, payload)
}

func (o *CreateAppSecretCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretCreated %s", 201, payload)
}

func (o *CreateAppSecretCreated) GetPayload() *models.AppAppSecret {
	return o.Payload
}

func (o *CreateAppSecretCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppAppSecret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSecretBadRequest creates a CreateAppSecretBadRequest with default headers values
func NewCreateAppSecretBadRequest() *CreateAppSecretBadRequest {
	return &CreateAppSecretBadRequest{}
}

/*
CreateAppSecretBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateAppSecretBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app secret bad request response has a 2xx status code
func (o *CreateAppSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app secret bad request response has a 3xx status code
func (o *CreateAppSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app secret bad request response has a 4xx status code
func (o *CreateAppSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app secret bad request response has a 5xx status code
func (o *CreateAppSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create app secret bad request response a status code equal to that given
func (o *CreateAppSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create app secret bad request response
func (o *CreateAppSecretBadRequest) Code() int {
	return 400
}

func (o *CreateAppSecretBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretBadRequest %s", 400, payload)
}

func (o *CreateAppSecretBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretBadRequest %s", 400, payload)
}

func (o *CreateAppSecretBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSecretUnauthorized creates a CreateAppSecretUnauthorized with default headers values
func NewCreateAppSecretUnauthorized() *CreateAppSecretUnauthorized {
	return &CreateAppSecretUnauthorized{}
}

/*
CreateAppSecretUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateAppSecretUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app secret unauthorized response has a 2xx status code
func (o *CreateAppSecretUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app secret unauthorized response has a 3xx status code
func (o *CreateAppSecretUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app secret unauthorized response has a 4xx status code
func (o *CreateAppSecretUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app secret unauthorized response has a 5xx status code
func (o *CreateAppSecretUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create app secret unauthorized response a status code equal to that given
func (o *CreateAppSecretUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create app secret unauthorized response
func (o *CreateAppSecretUnauthorized) Code() int {
	return 401
}

func (o *CreateAppSecretUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretUnauthorized %s", 401, payload)
}

func (o *CreateAppSecretUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretUnauthorized %s", 401, payload)
}

func (o *CreateAppSecretUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSecretForbidden creates a CreateAppSecretForbidden with default headers values
func NewCreateAppSecretForbidden() *CreateAppSecretForbidden {
	return &CreateAppSecretForbidden{}
}

/*
CreateAppSecretForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateAppSecretForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app secret forbidden response has a 2xx status code
func (o *CreateAppSecretForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app secret forbidden response has a 3xx status code
func (o *CreateAppSecretForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app secret forbidden response has a 4xx status code
func (o *CreateAppSecretForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app secret forbidden response has a 5xx status code
func (o *CreateAppSecretForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create app secret forbidden response a status code equal to that given
func (o *CreateAppSecretForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create app secret forbidden response
func (o *CreateAppSecretForbidden) Code() int {
	return 403
}

func (o *CreateAppSecretForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretForbidden %s", 403, payload)
}

func (o *CreateAppSecretForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretForbidden %s", 403, payload)
}

func (o *CreateAppSecretForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSecretNotFound creates a CreateAppSecretNotFound with default headers values
func NewCreateAppSecretNotFound() *CreateAppSecretNotFound {
	return &CreateAppSecretNotFound{}
}

/*
CreateAppSecretNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateAppSecretNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app secret not found response has a 2xx status code
func (o *CreateAppSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app secret not found response has a 3xx status code
func (o *CreateAppSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app secret not found response has a 4xx status code
func (o *CreateAppSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create app secret not found response has a 5xx status code
func (o *CreateAppSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create app secret not found response a status code equal to that given
func (o *CreateAppSecretNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create app secret not found response
func (o *CreateAppSecretNotFound) Code() int {
	return 404
}

func (o *CreateAppSecretNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretNotFound %s", 404, payload)
}

func (o *CreateAppSecretNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretNotFound %s", 404, payload)
}

func (o *CreateAppSecretNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppSecretInternalServerError creates a CreateAppSecretInternalServerError with default headers values
func NewCreateAppSecretInternalServerError() *CreateAppSecretInternalServerError {
	return &CreateAppSecretInternalServerError{}
}

/*
CreateAppSecretInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateAppSecretInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create app secret internal server error response has a 2xx status code
func (o *CreateAppSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create app secret internal server error response has a 3xx status code
func (o *CreateAppSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app secret internal server error response has a 4xx status code
func (o *CreateAppSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app secret internal server error response has a 5xx status code
func (o *CreateAppSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create app secret internal server error response a status code equal to that given
func (o *CreateAppSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create app secret internal server error response
func (o *CreateAppSecretInternalServerError) Code() int {
	return 500
}

func (o *CreateAppSecretInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretInternalServerError %s", 500, payload)
}

func (o *CreateAppSecretInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/apps/{app_id}/secret][%d] createAppSecretInternalServerError %s", 500, payload)
}

func (o *CreateAppSecretInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateAppSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
