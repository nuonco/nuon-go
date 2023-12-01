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

// GetV1InstallsInstallIDInputsCurrentReader is a Reader for the GetV1InstallsInstallIDInputsCurrent structure.
type GetV1InstallsInstallIDInputsCurrentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallsInstallIDInputsCurrentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallsInstallIDInputsCurrentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1InstallsInstallIDInputsCurrentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1InstallsInstallIDInputsCurrentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1InstallsInstallIDInputsCurrentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1InstallsInstallIDInputsCurrentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1InstallsInstallIDInputsCurrentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/inputs/current] GetV1InstallsInstallIDInputsCurrent", response, response.Code())
	}
}

// NewGetV1InstallsInstallIDInputsCurrentOK creates a GetV1InstallsInstallIDInputsCurrentOK with default headers values
func NewGetV1InstallsInstallIDInputsCurrentOK() *GetV1InstallsInstallIDInputsCurrentOK {
	return &GetV1InstallsInstallIDInputsCurrentOK{}
}

/*
GetV1InstallsInstallIDInputsCurrentOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallsInstallIDInputsCurrentOK struct {
	Payload *models.AppInstallInputs
}

// IsSuccess returns true when this get v1 installs install Id inputs current o k response has a 2xx status code
func (o *GetV1InstallsInstallIDInputsCurrentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installs install Id inputs current o k response has a 3xx status code
func (o *GetV1InstallsInstallIDInputsCurrentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id inputs current o k response has a 4xx status code
func (o *GetV1InstallsInstallIDInputsCurrentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id inputs current o k response has a 5xx status code
func (o *GetV1InstallsInstallIDInputsCurrentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id inputs current o k response a status code equal to that given
func (o *GetV1InstallsInstallIDInputsCurrentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installs install Id inputs current o k response
func (o *GetV1InstallsInstallIDInputsCurrentOK) Code() int {
	return 200
}

func (o *GetV1InstallsInstallIDInputsCurrentOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentOK) GetPayload() *models.AppInstallInputs {
	return o.Payload
}

func (o *GetV1InstallsInstallIDInputsCurrentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallInputs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDInputsCurrentBadRequest creates a GetV1InstallsInstallIDInputsCurrentBadRequest with default headers values
func NewGetV1InstallsInstallIDInputsCurrentBadRequest() *GetV1InstallsInstallIDInputsCurrentBadRequest {
	return &GetV1InstallsInstallIDInputsCurrentBadRequest{}
}

/*
GetV1InstallsInstallIDInputsCurrentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1InstallsInstallIDInputsCurrentBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id inputs current bad request response has a 2xx status code
func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id inputs current bad request response has a 3xx status code
func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id inputs current bad request response has a 4xx status code
func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id inputs current bad request response has a 5xx status code
func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id inputs current bad request response a status code equal to that given
func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 installs install Id inputs current bad request response
func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) Code() int {
	return 400
}

func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDInputsCurrentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDInputsCurrentUnauthorized creates a GetV1InstallsInstallIDInputsCurrentUnauthorized with default headers values
func NewGetV1InstallsInstallIDInputsCurrentUnauthorized() *GetV1InstallsInstallIDInputsCurrentUnauthorized {
	return &GetV1InstallsInstallIDInputsCurrentUnauthorized{}
}

/*
GetV1InstallsInstallIDInputsCurrentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1InstallsInstallIDInputsCurrentUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id inputs current unauthorized response has a 2xx status code
func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id inputs current unauthorized response has a 3xx status code
func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id inputs current unauthorized response has a 4xx status code
func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id inputs current unauthorized response has a 5xx status code
func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id inputs current unauthorized response a status code equal to that given
func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 installs install Id inputs current unauthorized response
func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) Code() int {
	return 401
}

func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDInputsCurrentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDInputsCurrentForbidden creates a GetV1InstallsInstallIDInputsCurrentForbidden with default headers values
func NewGetV1InstallsInstallIDInputsCurrentForbidden() *GetV1InstallsInstallIDInputsCurrentForbidden {
	return &GetV1InstallsInstallIDInputsCurrentForbidden{}
}

/*
GetV1InstallsInstallIDInputsCurrentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1InstallsInstallIDInputsCurrentForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id inputs current forbidden response has a 2xx status code
func (o *GetV1InstallsInstallIDInputsCurrentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id inputs current forbidden response has a 3xx status code
func (o *GetV1InstallsInstallIDInputsCurrentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id inputs current forbidden response has a 4xx status code
func (o *GetV1InstallsInstallIDInputsCurrentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id inputs current forbidden response has a 5xx status code
func (o *GetV1InstallsInstallIDInputsCurrentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id inputs current forbidden response a status code equal to that given
func (o *GetV1InstallsInstallIDInputsCurrentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 installs install Id inputs current forbidden response
func (o *GetV1InstallsInstallIDInputsCurrentForbidden) Code() int {
	return 403
}

func (o *GetV1InstallsInstallIDInputsCurrentForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDInputsCurrentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDInputsCurrentNotFound creates a GetV1InstallsInstallIDInputsCurrentNotFound with default headers values
func NewGetV1InstallsInstallIDInputsCurrentNotFound() *GetV1InstallsInstallIDInputsCurrentNotFound {
	return &GetV1InstallsInstallIDInputsCurrentNotFound{}
}

/*
GetV1InstallsInstallIDInputsCurrentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1InstallsInstallIDInputsCurrentNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id inputs current not found response has a 2xx status code
func (o *GetV1InstallsInstallIDInputsCurrentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id inputs current not found response has a 3xx status code
func (o *GetV1InstallsInstallIDInputsCurrentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id inputs current not found response has a 4xx status code
func (o *GetV1InstallsInstallIDInputsCurrentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 installs install Id inputs current not found response has a 5xx status code
func (o *GetV1InstallsInstallIDInputsCurrentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id inputs current not found response a status code equal to that given
func (o *GetV1InstallsInstallIDInputsCurrentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 installs install Id inputs current not found response
func (o *GetV1InstallsInstallIDInputsCurrentNotFound) Code() int {
	return 404
}

func (o *GetV1InstallsInstallIDInputsCurrentNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDInputsCurrentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallsInstallIDInputsCurrentInternalServerError creates a GetV1InstallsInstallIDInputsCurrentInternalServerError with default headers values
func NewGetV1InstallsInstallIDInputsCurrentInternalServerError() *GetV1InstallsInstallIDInputsCurrentInternalServerError {
	return &GetV1InstallsInstallIDInputsCurrentInternalServerError{}
}

/*
GetV1InstallsInstallIDInputsCurrentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1InstallsInstallIDInputsCurrentInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 installs install Id inputs current internal server error response has a 2xx status code
func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 installs install Id inputs current internal server error response has a 3xx status code
func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id inputs current internal server error response has a 4xx status code
func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id inputs current internal server error response has a 5xx status code
func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 installs install Id inputs current internal server error response a status code equal to that given
func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 installs install Id inputs current internal server error response
func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) Code() int {
	return 500
}

func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/inputs/current][%d] getV1InstallsInstallIdInputsCurrentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallsInstallIDInputsCurrentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
