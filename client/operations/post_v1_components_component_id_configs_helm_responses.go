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

// PostV1ComponentsComponentIDConfigsHelmReader is a Reader for the PostV1ComponentsComponentIDConfigsHelm structure.
type PostV1ComponentsComponentIDConfigsHelmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ComponentsComponentIDConfigsHelmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ComponentsComponentIDConfigsHelmCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ComponentsComponentIDConfigsHelmBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostV1ComponentsComponentIDConfigsHelmUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostV1ComponentsComponentIDConfigsHelmForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostV1ComponentsComponentIDConfigsHelmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1ComponentsComponentIDConfigsHelmInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/components/{component_id}/configs/helm] PostV1ComponentsComponentIDConfigsHelm", response, response.Code())
	}
}

// NewPostV1ComponentsComponentIDConfigsHelmCreated creates a PostV1ComponentsComponentIDConfigsHelmCreated with default headers values
func NewPostV1ComponentsComponentIDConfigsHelmCreated() *PostV1ComponentsComponentIDConfigsHelmCreated {
	return &PostV1ComponentsComponentIDConfigsHelmCreated{}
}

/*
PostV1ComponentsComponentIDConfigsHelmCreated describes a response with status code 201, with default header values.

Created
*/
type PostV1ComponentsComponentIDConfigsHelmCreated struct {
	Payload *models.AppHelmComponentConfig
}

// IsSuccess returns true when this post v1 components component Id configs helm created response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 components component Id configs helm created response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs helm created response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 components component Id configs helm created response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs helm created response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsHelmCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 components component Id configs helm created response
func (o *PostV1ComponentsComponentIDConfigsHelmCreated) Code() int {
	return 201
}

func (o *PostV1ComponentsComponentIDConfigsHelmCreated) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmCreated) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmCreated) GetPayload() *models.AppHelmComponentConfig {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsHelmCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppHelmComponentConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsHelmBadRequest creates a PostV1ComponentsComponentIDConfigsHelmBadRequest with default headers values
func NewPostV1ComponentsComponentIDConfigsHelmBadRequest() *PostV1ComponentsComponentIDConfigsHelmBadRequest {
	return &PostV1ComponentsComponentIDConfigsHelmBadRequest{}
}

/*
PostV1ComponentsComponentIDConfigsHelmBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ComponentsComponentIDConfigsHelmBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs helm bad request response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs helm bad request response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs helm bad request response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs helm bad request response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs helm bad request response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 components component Id configs helm bad request response
func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) Code() int {
	return 400
}

func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsHelmBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsHelmUnauthorized creates a PostV1ComponentsComponentIDConfigsHelmUnauthorized with default headers values
func NewPostV1ComponentsComponentIDConfigsHelmUnauthorized() *PostV1ComponentsComponentIDConfigsHelmUnauthorized {
	return &PostV1ComponentsComponentIDConfigsHelmUnauthorized{}
}

/*
PostV1ComponentsComponentIDConfigsHelmUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostV1ComponentsComponentIDConfigsHelmUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs helm unauthorized response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs helm unauthorized response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs helm unauthorized response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs helm unauthorized response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs helm unauthorized response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post v1 components component Id configs helm unauthorized response
func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) Code() int {
	return 401
}

func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsHelmUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsHelmForbidden creates a PostV1ComponentsComponentIDConfigsHelmForbidden with default headers values
func NewPostV1ComponentsComponentIDConfigsHelmForbidden() *PostV1ComponentsComponentIDConfigsHelmForbidden {
	return &PostV1ComponentsComponentIDConfigsHelmForbidden{}
}

/*
PostV1ComponentsComponentIDConfigsHelmForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostV1ComponentsComponentIDConfigsHelmForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs helm forbidden response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs helm forbidden response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs helm forbidden response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs helm forbidden response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs helm forbidden response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post v1 components component Id configs helm forbidden response
func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) Code() int {
	return 403
}

func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmForbidden  %+v", 403, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmForbidden  %+v", 403, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsHelmForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsHelmNotFound creates a PostV1ComponentsComponentIDConfigsHelmNotFound with default headers values
func NewPostV1ComponentsComponentIDConfigsHelmNotFound() *PostV1ComponentsComponentIDConfigsHelmNotFound {
	return &PostV1ComponentsComponentIDConfigsHelmNotFound{}
}

/*
PostV1ComponentsComponentIDConfigsHelmNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostV1ComponentsComponentIDConfigsHelmNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs helm not found response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs helm not found response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs helm not found response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs helm not found response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs helm not found response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post v1 components component Id configs helm not found response
func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) Code() int {
	return 404
}

func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmNotFound  %+v", 404, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmNotFound  %+v", 404, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsHelmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsHelmInternalServerError creates a PostV1ComponentsComponentIDConfigsHelmInternalServerError with default headers values
func NewPostV1ComponentsComponentIDConfigsHelmInternalServerError() *PostV1ComponentsComponentIDConfigsHelmInternalServerError {
	return &PostV1ComponentsComponentIDConfigsHelmInternalServerError{}
}

/*
PostV1ComponentsComponentIDConfigsHelmInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1ComponentsComponentIDConfigsHelmInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs helm internal server error response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs helm internal server error response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs helm internal server error response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 components component Id configs helm internal server error response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 components component Id configs helm internal server error response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 components component Id configs helm internal server error response
func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) Code() int {
	return 500
}

func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmInternalServerError  %+v", 500, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/helm][%d] postV1ComponentsComponentIdConfigsHelmInternalServerError  %+v", 500, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsHelmInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
