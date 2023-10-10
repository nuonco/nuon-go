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

// GetV1InstallersInstallerSlugRenderReader is a Reader for the GetV1InstallersInstallerSlugRender structure.
type GetV1InstallersInstallerSlugRenderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallersInstallerSlugRenderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallersInstallerSlugRenderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1InstallersInstallerSlugRenderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1InstallersInstallerSlugRenderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1InstallersInstallerSlugRenderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1InstallersInstallerSlugRenderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1InstallersInstallerSlugRenderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installers/{installer_slug}/render] GetV1InstallersInstallerSlugRender", response, response.Code())
	}
}

// NewGetV1InstallersInstallerSlugRenderOK creates a GetV1InstallersInstallerSlugRenderOK with default headers values
func NewGetV1InstallersInstallerSlugRenderOK() *GetV1InstallersInstallerSlugRenderOK {
	return &GetV1InstallersInstallerSlugRenderOK{}
}

/*
GetV1InstallersInstallerSlugRenderOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallersInstallerSlugRenderOK struct {
	Payload *models.ServiceAppInstaller
}

// IsSuccess returns true when this get v1 installers installer slug render o k response has a 2xx status code
func (o *GetV1InstallersInstallerSlugRenderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installers installer slug render o k response has a 3xx status code
func (o *GetV1InstallersInstallerSlugRenderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installers installer slug render o k response has a 4xx status code
func (o *GetV1InstallersInstallerSlugRenderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installers installer slug render o k response has a 5xx status code
func (o *GetV1InstallersInstallerSlugRenderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installers installer slug render o k response a status code equal to that given
func (o *GetV1InstallersInstallerSlugRenderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installers installer slug render o k response
func (o *GetV1InstallersInstallerSlugRenderOK) Code() int {
	return 200
}

func (o *GetV1InstallersInstallerSlugRenderOK) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderOK) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderOK) GetPayload() *models.ServiceAppInstaller {
	return o.Payload
}

func (o *GetV1InstallersInstallerSlugRenderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAppInstaller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallersInstallerSlugRenderBadRequest creates a GetV1InstallersInstallerSlugRenderBadRequest with default headers values
func NewGetV1InstallersInstallerSlugRenderBadRequest() *GetV1InstallersInstallerSlugRenderBadRequest {
	return &GetV1InstallersInstallerSlugRenderBadRequest{}
}

/*
GetV1InstallersInstallerSlugRenderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1InstallersInstallerSlugRenderBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installers installer slug render bad request response has a 2xx status code
func (o *GetV1InstallersInstallerSlugRenderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installers installer slug render bad request response has a 3xx status code
func (o *GetV1InstallersInstallerSlugRenderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installers installer slug render bad request response has a 4xx status code
func (o *GetV1InstallersInstallerSlugRenderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installers installer slug render bad request response has a 5xx status code
func (o *GetV1InstallersInstallerSlugRenderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installers installer slug render bad request response a status code equal to that given
func (o *GetV1InstallersInstallerSlugRenderBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 installers installer slug render bad request response
func (o *GetV1InstallersInstallerSlugRenderBadRequest) Code() int {
	return 400
}

func (o *GetV1InstallersInstallerSlugRenderBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallersInstallerSlugRenderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallersInstallerSlugRenderUnauthorized creates a GetV1InstallersInstallerSlugRenderUnauthorized with default headers values
func NewGetV1InstallersInstallerSlugRenderUnauthorized() *GetV1InstallersInstallerSlugRenderUnauthorized {
	return &GetV1InstallersInstallerSlugRenderUnauthorized{}
}

/*
GetV1InstallersInstallerSlugRenderUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1InstallersInstallerSlugRenderUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installers installer slug render unauthorized response has a 2xx status code
func (o *GetV1InstallersInstallerSlugRenderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installers installer slug render unauthorized response has a 3xx status code
func (o *GetV1InstallersInstallerSlugRenderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installers installer slug render unauthorized response has a 4xx status code
func (o *GetV1InstallersInstallerSlugRenderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installers installer slug render unauthorized response has a 5xx status code
func (o *GetV1InstallersInstallerSlugRenderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installers installer slug render unauthorized response a status code equal to that given
func (o *GetV1InstallersInstallerSlugRenderUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 installers installer slug render unauthorized response
func (o *GetV1InstallersInstallerSlugRenderUnauthorized) Code() int {
	return 401
}

func (o *GetV1InstallersInstallerSlugRenderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallersInstallerSlugRenderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallersInstallerSlugRenderForbidden creates a GetV1InstallersInstallerSlugRenderForbidden with default headers values
func NewGetV1InstallersInstallerSlugRenderForbidden() *GetV1InstallersInstallerSlugRenderForbidden {
	return &GetV1InstallersInstallerSlugRenderForbidden{}
}

/*
GetV1InstallersInstallerSlugRenderForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1InstallersInstallerSlugRenderForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installers installer slug render forbidden response has a 2xx status code
func (o *GetV1InstallersInstallerSlugRenderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installers installer slug render forbidden response has a 3xx status code
func (o *GetV1InstallersInstallerSlugRenderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installers installer slug render forbidden response has a 4xx status code
func (o *GetV1InstallersInstallerSlugRenderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installers installer slug render forbidden response has a 5xx status code
func (o *GetV1InstallersInstallerSlugRenderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installers installer slug render forbidden response a status code equal to that given
func (o *GetV1InstallersInstallerSlugRenderForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 installers installer slug render forbidden response
func (o *GetV1InstallersInstallerSlugRenderForbidden) Code() int {
	return 403
}

func (o *GetV1InstallersInstallerSlugRenderForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallersInstallerSlugRenderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallersInstallerSlugRenderNotFound creates a GetV1InstallersInstallerSlugRenderNotFound with default headers values
func NewGetV1InstallersInstallerSlugRenderNotFound() *GetV1InstallersInstallerSlugRenderNotFound {
	return &GetV1InstallersInstallerSlugRenderNotFound{}
}

/*
GetV1InstallersInstallerSlugRenderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1InstallersInstallerSlugRenderNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installers installer slug render not found response has a 2xx status code
func (o *GetV1InstallersInstallerSlugRenderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installers installer slug render not found response has a 3xx status code
func (o *GetV1InstallersInstallerSlugRenderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installers installer slug render not found response has a 4xx status code
func (o *GetV1InstallersInstallerSlugRenderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installers installer slug render not found response has a 5xx status code
func (o *GetV1InstallersInstallerSlugRenderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installers installer slug render not found response a status code equal to that given
func (o *GetV1InstallersInstallerSlugRenderNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 installers installer slug render not found response
func (o *GetV1InstallersInstallerSlugRenderNotFound) Code() int {
	return 404
}

func (o *GetV1InstallersInstallerSlugRenderNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallersInstallerSlugRenderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallersInstallerSlugRenderInternalServerError creates a GetV1InstallersInstallerSlugRenderInternalServerError with default headers values
func NewGetV1InstallersInstallerSlugRenderInternalServerError() *GetV1InstallersInstallerSlugRenderInternalServerError {
	return &GetV1InstallersInstallerSlugRenderInternalServerError{}
}

/*
GetV1InstallersInstallerSlugRenderInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1InstallersInstallerSlugRenderInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installers installer slug render internal server error response has a 2xx status code
func (o *GetV1InstallersInstallerSlugRenderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installers installer slug render internal server error response has a 3xx status code
func (o *GetV1InstallersInstallerSlugRenderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installers installer slug render internal server error response has a 4xx status code
func (o *GetV1InstallersInstallerSlugRenderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installers installer slug render internal server error response has a 5xx status code
func (o *GetV1InstallersInstallerSlugRenderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 installers installer slug render internal server error response a status code equal to that given
func (o *GetV1InstallersInstallerSlugRenderInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 installers installer slug render internal server error response
func (o *GetV1InstallersInstallerSlugRenderInternalServerError) Code() int {
	return 500
}

func (o *GetV1InstallersInstallerSlugRenderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installers/{installer_slug}/render][%d] getV1InstallersInstallerSlugRenderInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallersInstallerSlugRenderInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallersInstallerSlugRenderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}