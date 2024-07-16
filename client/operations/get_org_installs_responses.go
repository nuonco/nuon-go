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

// GetOrgInstallsReader is a Reader for the GetOrgInstalls structure.
type GetOrgInstallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgInstallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgInstallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgInstallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgInstallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgInstallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgInstallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgInstallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs] GetOrgInstalls", response, response.Code())
	}
}

// NewGetOrgInstallsOK creates a GetOrgInstallsOK with default headers values
func NewGetOrgInstallsOK() *GetOrgInstallsOK {
	return &GetOrgInstallsOK{}
}

/*
GetOrgInstallsOK describes a response with status code 200, with default header values.

OK
*/
type GetOrgInstallsOK struct {
	Payload []*models.AppInstall
}

// IsSuccess returns true when this get org installs o k response has a 2xx status code
func (o *GetOrgInstallsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org installs o k response has a 3xx status code
func (o *GetOrgInstallsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org installs o k response has a 4xx status code
func (o *GetOrgInstallsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org installs o k response has a 5xx status code
func (o *GetOrgInstallsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org installs o k response a status code equal to that given
func (o *GetOrgInstallsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org installs o k response
func (o *GetOrgInstallsOK) Code() int {
	return 200
}

func (o *GetOrgInstallsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsOK %s", 200, payload)
}

func (o *GetOrgInstallsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsOK %s", 200, payload)
}

func (o *GetOrgInstallsOK) GetPayload() []*models.AppInstall {
	return o.Payload
}

func (o *GetOrgInstallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgInstallsBadRequest creates a GetOrgInstallsBadRequest with default headers values
func NewGetOrgInstallsBadRequest() *GetOrgInstallsBadRequest {
	return &GetOrgInstallsBadRequest{}
}

/*
GetOrgInstallsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrgInstallsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org installs bad request response has a 2xx status code
func (o *GetOrgInstallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org installs bad request response has a 3xx status code
func (o *GetOrgInstallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org installs bad request response has a 4xx status code
func (o *GetOrgInstallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org installs bad request response has a 5xx status code
func (o *GetOrgInstallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get org installs bad request response a status code equal to that given
func (o *GetOrgInstallsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get org installs bad request response
func (o *GetOrgInstallsBadRequest) Code() int {
	return 400
}

func (o *GetOrgInstallsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsBadRequest %s", 400, payload)
}

func (o *GetOrgInstallsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsBadRequest %s", 400, payload)
}

func (o *GetOrgInstallsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgInstallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgInstallsUnauthorized creates a GetOrgInstallsUnauthorized with default headers values
func NewGetOrgInstallsUnauthorized() *GetOrgInstallsUnauthorized {
	return &GetOrgInstallsUnauthorized{}
}

/*
GetOrgInstallsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrgInstallsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org installs unauthorized response has a 2xx status code
func (o *GetOrgInstallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org installs unauthorized response has a 3xx status code
func (o *GetOrgInstallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org installs unauthorized response has a 4xx status code
func (o *GetOrgInstallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org installs unauthorized response has a 5xx status code
func (o *GetOrgInstallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org installs unauthorized response a status code equal to that given
func (o *GetOrgInstallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org installs unauthorized response
func (o *GetOrgInstallsUnauthorized) Code() int {
	return 401
}

func (o *GetOrgInstallsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsUnauthorized %s", 401, payload)
}

func (o *GetOrgInstallsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsUnauthorized %s", 401, payload)
}

func (o *GetOrgInstallsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgInstallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgInstallsForbidden creates a GetOrgInstallsForbidden with default headers values
func NewGetOrgInstallsForbidden() *GetOrgInstallsForbidden {
	return &GetOrgInstallsForbidden{}
}

/*
GetOrgInstallsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrgInstallsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org installs forbidden response has a 2xx status code
func (o *GetOrgInstallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org installs forbidden response has a 3xx status code
func (o *GetOrgInstallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org installs forbidden response has a 4xx status code
func (o *GetOrgInstallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org installs forbidden response has a 5xx status code
func (o *GetOrgInstallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org installs forbidden response a status code equal to that given
func (o *GetOrgInstallsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org installs forbidden response
func (o *GetOrgInstallsForbidden) Code() int {
	return 403
}

func (o *GetOrgInstallsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsForbidden %s", 403, payload)
}

func (o *GetOrgInstallsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsForbidden %s", 403, payload)
}

func (o *GetOrgInstallsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgInstallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgInstallsNotFound creates a GetOrgInstallsNotFound with default headers values
func NewGetOrgInstallsNotFound() *GetOrgInstallsNotFound {
	return &GetOrgInstallsNotFound{}
}

/*
GetOrgInstallsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrgInstallsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org installs not found response has a 2xx status code
func (o *GetOrgInstallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org installs not found response has a 3xx status code
func (o *GetOrgInstallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org installs not found response has a 4xx status code
func (o *GetOrgInstallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org installs not found response has a 5xx status code
func (o *GetOrgInstallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get org installs not found response a status code equal to that given
func (o *GetOrgInstallsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get org installs not found response
func (o *GetOrgInstallsNotFound) Code() int {
	return 404
}

func (o *GetOrgInstallsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsNotFound %s", 404, payload)
}

func (o *GetOrgInstallsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsNotFound %s", 404, payload)
}

func (o *GetOrgInstallsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgInstallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgInstallsInternalServerError creates a GetOrgInstallsInternalServerError with default headers values
func NewGetOrgInstallsInternalServerError() *GetOrgInstallsInternalServerError {
	return &GetOrgInstallsInternalServerError{}
}

/*
GetOrgInstallsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrgInstallsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org installs internal server error response has a 2xx status code
func (o *GetOrgInstallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org installs internal server error response has a 3xx status code
func (o *GetOrgInstallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org installs internal server error response has a 4xx status code
func (o *GetOrgInstallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org installs internal server error response has a 5xx status code
func (o *GetOrgInstallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org installs internal server error response a status code equal to that given
func (o *GetOrgInstallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org installs internal server error response
func (o *GetOrgInstallsInternalServerError) Code() int {
	return 500
}

func (o *GetOrgInstallsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsInternalServerError %s", 500, payload)
}

func (o *GetOrgInstallsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs][%d] getOrgInstallsInternalServerError %s", 500, payload)
}

func (o *GetOrgInstallsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgInstallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
