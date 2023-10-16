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

// GetV1InstallerInstallerSlugInstallInstallIDReader is a Reader for the GetV1InstallerInstallerSlugInstallInstallID structure.
type GetV1InstallerInstallerSlugInstallInstallIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallerInstallerSlugInstallInstallIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallerInstallerSlugInstallInstallIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1InstallerInstallerSlugInstallInstallIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1InstallerInstallerSlugInstallInstallIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1InstallerInstallerSlugInstallInstallIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1InstallerInstallerSlugInstallInstallIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1InstallerInstallerSlugInstallInstallIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installer/{installer_slug}/install/{install_id}] GetV1InstallerInstallerSlugInstallInstallID", response, response.Code())
	}
}

// NewGetV1InstallerInstallerSlugInstallInstallIDOK creates a GetV1InstallerInstallerSlugInstallInstallIDOK with default headers values
func NewGetV1InstallerInstallerSlugInstallInstallIDOK() *GetV1InstallerInstallerSlugInstallInstallIDOK {
	return &GetV1InstallerInstallerSlugInstallInstallIDOK{}
}

/*
GetV1InstallerInstallerSlugInstallInstallIDOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallerInstallerSlugInstallInstallIDOK struct {
	Payload *models.AppInstall
}

// IsSuccess returns true when this get v1 installer installer slug install install Id o k response has a 2xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installer installer slug install install Id o k response has a 3xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installer installer slug install install Id o k response has a 4xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installer installer slug install install Id o k response has a 5xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installer installer slug install install Id o k response a status code equal to that given
func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installer installer slug install install Id o k response
func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) Code() int {
	return 200
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) String() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) GetPayload() *models.AppInstall {
	return o.Payload
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstall)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallerInstallerSlugInstallInstallIDBadRequest creates a GetV1InstallerInstallerSlugInstallInstallIDBadRequest with default headers values
func NewGetV1InstallerInstallerSlugInstallInstallIDBadRequest() *GetV1InstallerInstallerSlugInstallInstallIDBadRequest {
	return &GetV1InstallerInstallerSlugInstallInstallIDBadRequest{}
}

/*
GetV1InstallerInstallerSlugInstallInstallIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1InstallerInstallerSlugInstallInstallIDBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installer installer slug install install Id bad request response has a 2xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installer installer slug install install Id bad request response has a 3xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installer installer slug install install Id bad request response has a 4xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installer installer slug install install Id bad request response has a 5xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installer installer slug install install Id bad request response a status code equal to that given
func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 installer installer slug install install Id bad request response
func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) Code() int {
	return 400
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallerInstallerSlugInstallInstallIDUnauthorized creates a GetV1InstallerInstallerSlugInstallInstallIDUnauthorized with default headers values
func NewGetV1InstallerInstallerSlugInstallInstallIDUnauthorized() *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized {
	return &GetV1InstallerInstallerSlugInstallInstallIDUnauthorized{}
}

/*
GetV1InstallerInstallerSlugInstallInstallIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1InstallerInstallerSlugInstallInstallIDUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installer installer slug install install Id unauthorized response has a 2xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installer installer slug install install Id unauthorized response has a 3xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installer installer slug install install Id unauthorized response has a 4xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installer installer slug install install Id unauthorized response has a 5xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installer installer slug install install Id unauthorized response a status code equal to that given
func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 installer installer slug install install Id unauthorized response
func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) Code() int {
	return 401
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallerInstallerSlugInstallInstallIDForbidden creates a GetV1InstallerInstallerSlugInstallInstallIDForbidden with default headers values
func NewGetV1InstallerInstallerSlugInstallInstallIDForbidden() *GetV1InstallerInstallerSlugInstallInstallIDForbidden {
	return &GetV1InstallerInstallerSlugInstallInstallIDForbidden{}
}

/*
GetV1InstallerInstallerSlugInstallInstallIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1InstallerInstallerSlugInstallInstallIDForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installer installer slug install install Id forbidden response has a 2xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installer installer slug install install Id forbidden response has a 3xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installer installer slug install install Id forbidden response has a 4xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installer installer slug install install Id forbidden response has a 5xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installer installer slug install install Id forbidden response a status code equal to that given
func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 installer installer slug install install Id forbidden response
func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) Code() int {
	return 403
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallerInstallerSlugInstallInstallIDNotFound creates a GetV1InstallerInstallerSlugInstallInstallIDNotFound with default headers values
func NewGetV1InstallerInstallerSlugInstallInstallIDNotFound() *GetV1InstallerInstallerSlugInstallInstallIDNotFound {
	return &GetV1InstallerInstallerSlugInstallInstallIDNotFound{}
}

/*
GetV1InstallerInstallerSlugInstallInstallIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1InstallerInstallerSlugInstallInstallIDNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installer installer slug install install Id not found response has a 2xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installer installer slug install install Id not found response has a 3xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installer installer slug install install Id not found response has a 4xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installer installer slug install install Id not found response has a 5xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installer installer slug install install Id not found response a status code equal to that given
func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 installer installer slug install install Id not found response
func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) Code() int {
	return 404
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallerInstallerSlugInstallInstallIDInternalServerError creates a GetV1InstallerInstallerSlugInstallInstallIDInternalServerError with default headers values
func NewGetV1InstallerInstallerSlugInstallInstallIDInternalServerError() *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError {
	return &GetV1InstallerInstallerSlugInstallInstallIDInternalServerError{}
}

/*
GetV1InstallerInstallerSlugInstallInstallIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1InstallerInstallerSlugInstallInstallIDInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installer installer slug install install Id internal server error response has a 2xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installer installer slug install install Id internal server error response has a 3xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installer installer slug install install Id internal server error response has a 4xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installer installer slug install install Id internal server error response has a 5xx status code
func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 installer installer slug install install Id internal server error response a status code equal to that given
func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 installer installer slug install install Id internal server error response
func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) Code() int {
	return 500
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installer/{installer_slug}/install/{install_id}][%d] getV1InstallerInstallerSlugInstallInstallIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallerInstallerSlugInstallInstallIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}