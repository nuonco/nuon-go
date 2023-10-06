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

// PostV1ComponentsComponentIDConfigsExternalImageReader is a Reader for the PostV1ComponentsComponentIDConfigsExternalImage structure.
type PostV1ComponentsComponentIDConfigsExternalImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ComponentsComponentIDConfigsExternalImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ComponentsComponentIDConfigsExternalImageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ComponentsComponentIDConfigsExternalImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostV1ComponentsComponentIDConfigsExternalImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostV1ComponentsComponentIDConfigsExternalImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostV1ComponentsComponentIDConfigsExternalImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1ComponentsComponentIDConfigsExternalImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/components/{component_id}/configs/external-image] PostV1ComponentsComponentIDConfigsExternalImage", response, response.Code())
	}
}

// NewPostV1ComponentsComponentIDConfigsExternalImageCreated creates a PostV1ComponentsComponentIDConfigsExternalImageCreated with default headers values
func NewPostV1ComponentsComponentIDConfigsExternalImageCreated() *PostV1ComponentsComponentIDConfigsExternalImageCreated {
	return &PostV1ComponentsComponentIDConfigsExternalImageCreated{}
}

/*
PostV1ComponentsComponentIDConfigsExternalImageCreated describes a response with status code 201, with default header values.

Created
*/
type PostV1ComponentsComponentIDConfigsExternalImageCreated struct {
	Payload *models.AppExternalImageComponentConfig
}

// IsSuccess returns true when this post v1 components component Id configs external image created response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 components component Id configs external image created response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs external image created response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 components component Id configs external image created response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs external image created response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 components component Id configs external image created response
func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) Code() int {
	return 201
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) GetPayload() *models.AppExternalImageComponentConfig {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppExternalImageComponentConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsExternalImageBadRequest creates a PostV1ComponentsComponentIDConfigsExternalImageBadRequest with default headers values
func NewPostV1ComponentsComponentIDConfigsExternalImageBadRequest() *PostV1ComponentsComponentIDConfigsExternalImageBadRequest {
	return &PostV1ComponentsComponentIDConfigsExternalImageBadRequest{}
}

/*
PostV1ComponentsComponentIDConfigsExternalImageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ComponentsComponentIDConfigsExternalImageBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs external image bad request response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs external image bad request response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs external image bad request response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs external image bad request response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs external image bad request response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 components component Id configs external image bad request response
func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) Code() int {
	return 400
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsExternalImageUnauthorized creates a PostV1ComponentsComponentIDConfigsExternalImageUnauthorized with default headers values
func NewPostV1ComponentsComponentIDConfigsExternalImageUnauthorized() *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized {
	return &PostV1ComponentsComponentIDConfigsExternalImageUnauthorized{}
}

/*
PostV1ComponentsComponentIDConfigsExternalImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostV1ComponentsComponentIDConfigsExternalImageUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs external image unauthorized response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs external image unauthorized response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs external image unauthorized response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs external image unauthorized response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs external image unauthorized response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post v1 components component Id configs external image unauthorized response
func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) Code() int {
	return 401
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsExternalImageForbidden creates a PostV1ComponentsComponentIDConfigsExternalImageForbidden with default headers values
func NewPostV1ComponentsComponentIDConfigsExternalImageForbidden() *PostV1ComponentsComponentIDConfigsExternalImageForbidden {
	return &PostV1ComponentsComponentIDConfigsExternalImageForbidden{}
}

/*
PostV1ComponentsComponentIDConfigsExternalImageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostV1ComponentsComponentIDConfigsExternalImageForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs external image forbidden response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs external image forbidden response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs external image forbidden response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs external image forbidden response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs external image forbidden response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post v1 components component Id configs external image forbidden response
func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) Code() int {
	return 403
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageForbidden  %+v", 403, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageForbidden  %+v", 403, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsExternalImageNotFound creates a PostV1ComponentsComponentIDConfigsExternalImageNotFound with default headers values
func NewPostV1ComponentsComponentIDConfigsExternalImageNotFound() *PostV1ComponentsComponentIDConfigsExternalImageNotFound {
	return &PostV1ComponentsComponentIDConfigsExternalImageNotFound{}
}

/*
PostV1ComponentsComponentIDConfigsExternalImageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostV1ComponentsComponentIDConfigsExternalImageNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs external image not found response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs external image not found response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs external image not found response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 components component Id configs external image not found response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id configs external image not found response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post v1 components component Id configs external image not found response
func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) Code() int {
	return 404
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageNotFound  %+v", 404, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageNotFound  %+v", 404, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ComponentsComponentIDConfigsExternalImageInternalServerError creates a PostV1ComponentsComponentIDConfigsExternalImageInternalServerError with default headers values
func NewPostV1ComponentsComponentIDConfigsExternalImageInternalServerError() *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError {
	return &PostV1ComponentsComponentIDConfigsExternalImageInternalServerError{}
}

/*
PostV1ComponentsComponentIDConfigsExternalImageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1ComponentsComponentIDConfigsExternalImageInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this post v1 components component Id configs external image internal server error response has a 2xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 components component Id configs external image internal server error response has a 3xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id configs external image internal server error response has a 4xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 components component Id configs external image internal server error response has a 5xx status code
func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 components component Id configs external image internal server error response a status code equal to that given
func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 components component Id configs external image internal server error response
func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) Code() int {
	return 500
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageInternalServerError  %+v", 500, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/configs/external-image][%d] postV1ComponentsComponentIdConfigsExternalImageInternalServerError  %+v", 500, o.Payload)
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDConfigsExternalImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
