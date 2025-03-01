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

// GetAppReader is a Reader for the GetApp structure.
type GetAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}] GetApp", response, response.Code())
	}
}

// NewGetAppOK creates a GetAppOK with default headers values
func NewGetAppOK() *GetAppOK {
	return &GetAppOK{}
}

/*
GetAppOK describes a response with status code 200, with default header values.

OK
*/
type GetAppOK struct {
	Payload *models.AppApp
}

// IsSuccess returns true when this get app o k response has a 2xx status code
func (o *GetAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app o k response has a 3xx status code
func (o *GetAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app o k response has a 4xx status code
func (o *GetAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app o k response has a 5xx status code
func (o *GetAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app o k response a status code equal to that given
func (o *GetAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app o k response
func (o *GetAppOK) Code() int {
	return 200
}

func (o *GetAppOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppOK %s", 200, payload)
}

func (o *GetAppOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppOK %s", 200, payload)
}

func (o *GetAppOK) GetPayload() *models.AppApp {
	return o.Payload
}

func (o *GetAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppBadRequest creates a GetAppBadRequest with default headers values
func NewGetAppBadRequest() *GetAppBadRequest {
	return &GetAppBadRequest{}
}

/*
GetAppBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app bad request response has a 2xx status code
func (o *GetAppBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app bad request response has a 3xx status code
func (o *GetAppBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app bad request response has a 4xx status code
func (o *GetAppBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app bad request response has a 5xx status code
func (o *GetAppBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app bad request response a status code equal to that given
func (o *GetAppBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app bad request response
func (o *GetAppBadRequest) Code() int {
	return 400
}

func (o *GetAppBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppBadRequest %s", 400, payload)
}

func (o *GetAppBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppBadRequest %s", 400, payload)
}

func (o *GetAppBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppUnauthorized creates a GetAppUnauthorized with default headers values
func NewGetAppUnauthorized() *GetAppUnauthorized {
	return &GetAppUnauthorized{}
}

/*
GetAppUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app unauthorized response has a 2xx status code
func (o *GetAppUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app unauthorized response has a 3xx status code
func (o *GetAppUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app unauthorized response has a 4xx status code
func (o *GetAppUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app unauthorized response has a 5xx status code
func (o *GetAppUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app unauthorized response a status code equal to that given
func (o *GetAppUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app unauthorized response
func (o *GetAppUnauthorized) Code() int {
	return 401
}

func (o *GetAppUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppUnauthorized %s", 401, payload)
}

func (o *GetAppUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppUnauthorized %s", 401, payload)
}

func (o *GetAppUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppForbidden creates a GetAppForbidden with default headers values
func NewGetAppForbidden() *GetAppForbidden {
	return &GetAppForbidden{}
}

/*
GetAppForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app forbidden response has a 2xx status code
func (o *GetAppForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app forbidden response has a 3xx status code
func (o *GetAppForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app forbidden response has a 4xx status code
func (o *GetAppForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app forbidden response has a 5xx status code
func (o *GetAppForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app forbidden response a status code equal to that given
func (o *GetAppForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app forbidden response
func (o *GetAppForbidden) Code() int {
	return 403
}

func (o *GetAppForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppForbidden %s", 403, payload)
}

func (o *GetAppForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppForbidden %s", 403, payload)
}

func (o *GetAppForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppNotFound creates a GetAppNotFound with default headers values
func NewGetAppNotFound() *GetAppNotFound {
	return &GetAppNotFound{}
}

/*
GetAppNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app not found response has a 2xx status code
func (o *GetAppNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app not found response has a 3xx status code
func (o *GetAppNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app not found response has a 4xx status code
func (o *GetAppNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app not found response has a 5xx status code
func (o *GetAppNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app not found response a status code equal to that given
func (o *GetAppNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app not found response
func (o *GetAppNotFound) Code() int {
	return 404
}

func (o *GetAppNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppNotFound %s", 404, payload)
}

func (o *GetAppNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppNotFound %s", 404, payload)
}

func (o *GetAppNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppInternalServerError creates a GetAppInternalServerError with default headers values
func NewGetAppInternalServerError() *GetAppInternalServerError {
	return &GetAppInternalServerError{}
}

/*
GetAppInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app internal server error response has a 2xx status code
func (o *GetAppInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app internal server error response has a 3xx status code
func (o *GetAppInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app internal server error response has a 4xx status code
func (o *GetAppInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app internal server error response has a 5xx status code
func (o *GetAppInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app internal server error response a status code equal to that given
func (o *GetAppInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app internal server error response
func (o *GetAppInternalServerError) Code() int {
	return 500
}

func (o *GetAppInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppInternalServerError %s", 500, payload)
}

func (o *GetAppInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}][%d] getAppInternalServerError %s", 500, payload)
}

func (o *GetAppInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
