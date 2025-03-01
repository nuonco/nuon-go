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

// GetAppReleasesReader is a Reader for the GetAppReleases structure.
type GetAppReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppReleasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppReleasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppReleasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAppReleasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAppReleasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/releases] GetAppReleases", response, response.Code())
	}
}

// NewGetAppReleasesOK creates a GetAppReleasesOK with default headers values
func NewGetAppReleasesOK() *GetAppReleasesOK {
	return &GetAppReleasesOK{}
}

/*
GetAppReleasesOK describes a response with status code 200, with default header values.

OK
*/
type GetAppReleasesOK struct {
	Payload []*models.AppComponentRelease
}

// IsSuccess returns true when this get app releases o k response has a 2xx status code
func (o *GetAppReleasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get app releases o k response has a 3xx status code
func (o *GetAppReleasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app releases o k response has a 4xx status code
func (o *GetAppReleasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app releases o k response has a 5xx status code
func (o *GetAppReleasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get app releases o k response a status code equal to that given
func (o *GetAppReleasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get app releases o k response
func (o *GetAppReleasesOK) Code() int {
	return 200
}

func (o *GetAppReleasesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesOK %s", 200, payload)
}

func (o *GetAppReleasesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesOK %s", 200, payload)
}

func (o *GetAppReleasesOK) GetPayload() []*models.AppComponentRelease {
	return o.Payload
}

func (o *GetAppReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppReleasesBadRequest creates a GetAppReleasesBadRequest with default headers values
func NewGetAppReleasesBadRequest() *GetAppReleasesBadRequest {
	return &GetAppReleasesBadRequest{}
}

/*
GetAppReleasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAppReleasesBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app releases bad request response has a 2xx status code
func (o *GetAppReleasesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app releases bad request response has a 3xx status code
func (o *GetAppReleasesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app releases bad request response has a 4xx status code
func (o *GetAppReleasesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app releases bad request response has a 5xx status code
func (o *GetAppReleasesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get app releases bad request response a status code equal to that given
func (o *GetAppReleasesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get app releases bad request response
func (o *GetAppReleasesBadRequest) Code() int {
	return 400
}

func (o *GetAppReleasesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesBadRequest %s", 400, payload)
}

func (o *GetAppReleasesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesBadRequest %s", 400, payload)
}

func (o *GetAppReleasesBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppReleasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppReleasesUnauthorized creates a GetAppReleasesUnauthorized with default headers values
func NewGetAppReleasesUnauthorized() *GetAppReleasesUnauthorized {
	return &GetAppReleasesUnauthorized{}
}

/*
GetAppReleasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAppReleasesUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app releases unauthorized response has a 2xx status code
func (o *GetAppReleasesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app releases unauthorized response has a 3xx status code
func (o *GetAppReleasesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app releases unauthorized response has a 4xx status code
func (o *GetAppReleasesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app releases unauthorized response has a 5xx status code
func (o *GetAppReleasesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get app releases unauthorized response a status code equal to that given
func (o *GetAppReleasesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get app releases unauthorized response
func (o *GetAppReleasesUnauthorized) Code() int {
	return 401
}

func (o *GetAppReleasesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesUnauthorized %s", 401, payload)
}

func (o *GetAppReleasesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesUnauthorized %s", 401, payload)
}

func (o *GetAppReleasesUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppReleasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppReleasesForbidden creates a GetAppReleasesForbidden with default headers values
func NewGetAppReleasesForbidden() *GetAppReleasesForbidden {
	return &GetAppReleasesForbidden{}
}

/*
GetAppReleasesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAppReleasesForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app releases forbidden response has a 2xx status code
func (o *GetAppReleasesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app releases forbidden response has a 3xx status code
func (o *GetAppReleasesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app releases forbidden response has a 4xx status code
func (o *GetAppReleasesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app releases forbidden response has a 5xx status code
func (o *GetAppReleasesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get app releases forbidden response a status code equal to that given
func (o *GetAppReleasesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get app releases forbidden response
func (o *GetAppReleasesForbidden) Code() int {
	return 403
}

func (o *GetAppReleasesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesForbidden %s", 403, payload)
}

func (o *GetAppReleasesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesForbidden %s", 403, payload)
}

func (o *GetAppReleasesForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppReleasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppReleasesNotFound creates a GetAppReleasesNotFound with default headers values
func NewGetAppReleasesNotFound() *GetAppReleasesNotFound {
	return &GetAppReleasesNotFound{}
}

/*
GetAppReleasesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAppReleasesNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app releases not found response has a 2xx status code
func (o *GetAppReleasesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app releases not found response has a 3xx status code
func (o *GetAppReleasesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app releases not found response has a 4xx status code
func (o *GetAppReleasesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get app releases not found response has a 5xx status code
func (o *GetAppReleasesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get app releases not found response a status code equal to that given
func (o *GetAppReleasesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get app releases not found response
func (o *GetAppReleasesNotFound) Code() int {
	return 404
}

func (o *GetAppReleasesNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesNotFound %s", 404, payload)
}

func (o *GetAppReleasesNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesNotFound %s", 404, payload)
}

func (o *GetAppReleasesNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppReleasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppReleasesInternalServerError creates a GetAppReleasesInternalServerError with default headers values
func NewGetAppReleasesInternalServerError() *GetAppReleasesInternalServerError {
	return &GetAppReleasesInternalServerError{}
}

/*
GetAppReleasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAppReleasesInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get app releases internal server error response has a 2xx status code
func (o *GetAppReleasesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get app releases internal server error response has a 3xx status code
func (o *GetAppReleasesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get app releases internal server error response has a 4xx status code
func (o *GetAppReleasesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get app releases internal server error response has a 5xx status code
func (o *GetAppReleasesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get app releases internal server error response a status code equal to that given
func (o *GetAppReleasesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get app releases internal server error response
func (o *GetAppReleasesInternalServerError) Code() int {
	return 500
}

func (o *GetAppReleasesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesInternalServerError %s", 500, payload)
}

func (o *GetAppReleasesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getAppReleasesInternalServerError %s", 500, payload)
}

func (o *GetAppReleasesInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetAppReleasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
