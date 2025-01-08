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

// GetCloudPlatformRegionsReader is a Reader for the GetCloudPlatformRegions structure.
type GetCloudPlatformRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudPlatformRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudPlatformRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCloudPlatformRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCloudPlatformRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCloudPlatformRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudPlatformRegionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudPlatformRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/general/cloud-platform/{cloud_platform}/regions] GetCloudPlatformRegions", response, response.Code())
	}
}

// NewGetCloudPlatformRegionsOK creates a GetCloudPlatformRegionsOK with default headers values
func NewGetCloudPlatformRegionsOK() *GetCloudPlatformRegionsOK {
	return &GetCloudPlatformRegionsOK{}
}

/*
GetCloudPlatformRegionsOK describes a response with status code 200, with default header values.

OK
*/
type GetCloudPlatformRegionsOK struct {
	Payload []*models.AppCloudPlatformRegion
}

// IsSuccess returns true when this get cloud platform regions o k response has a 2xx status code
func (o *GetCloudPlatformRegionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud platform regions o k response has a 3xx status code
func (o *GetCloudPlatformRegionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud platform regions o k response has a 4xx status code
func (o *GetCloudPlatformRegionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud platform regions o k response has a 5xx status code
func (o *GetCloudPlatformRegionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud platform regions o k response a status code equal to that given
func (o *GetCloudPlatformRegionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cloud platform regions o k response
func (o *GetCloudPlatformRegionsOK) Code() int {
	return 200
}

func (o *GetCloudPlatformRegionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsOK %s", 200, payload)
}

func (o *GetCloudPlatformRegionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsOK %s", 200, payload)
}

func (o *GetCloudPlatformRegionsOK) GetPayload() []*models.AppCloudPlatformRegion {
	return o.Payload
}

func (o *GetCloudPlatformRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPlatformRegionsBadRequest creates a GetCloudPlatformRegionsBadRequest with default headers values
func NewGetCloudPlatformRegionsBadRequest() *GetCloudPlatformRegionsBadRequest {
	return &GetCloudPlatformRegionsBadRequest{}
}

/*
GetCloudPlatformRegionsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCloudPlatformRegionsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get cloud platform regions bad request response has a 2xx status code
func (o *GetCloudPlatformRegionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud platform regions bad request response has a 3xx status code
func (o *GetCloudPlatformRegionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud platform regions bad request response has a 4xx status code
func (o *GetCloudPlatformRegionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud platform regions bad request response has a 5xx status code
func (o *GetCloudPlatformRegionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud platform regions bad request response a status code equal to that given
func (o *GetCloudPlatformRegionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get cloud platform regions bad request response
func (o *GetCloudPlatformRegionsBadRequest) Code() int {
	return 400
}

func (o *GetCloudPlatformRegionsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsBadRequest %s", 400, payload)
}

func (o *GetCloudPlatformRegionsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsBadRequest %s", 400, payload)
}

func (o *GetCloudPlatformRegionsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCloudPlatformRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPlatformRegionsUnauthorized creates a GetCloudPlatformRegionsUnauthorized with default headers values
func NewGetCloudPlatformRegionsUnauthorized() *GetCloudPlatformRegionsUnauthorized {
	return &GetCloudPlatformRegionsUnauthorized{}
}

/*
GetCloudPlatformRegionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCloudPlatformRegionsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get cloud platform regions unauthorized response has a 2xx status code
func (o *GetCloudPlatformRegionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud platform regions unauthorized response has a 3xx status code
func (o *GetCloudPlatformRegionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud platform regions unauthorized response has a 4xx status code
func (o *GetCloudPlatformRegionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud platform regions unauthorized response has a 5xx status code
func (o *GetCloudPlatformRegionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud platform regions unauthorized response a status code equal to that given
func (o *GetCloudPlatformRegionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get cloud platform regions unauthorized response
func (o *GetCloudPlatformRegionsUnauthorized) Code() int {
	return 401
}

func (o *GetCloudPlatformRegionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsUnauthorized %s", 401, payload)
}

func (o *GetCloudPlatformRegionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsUnauthorized %s", 401, payload)
}

func (o *GetCloudPlatformRegionsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCloudPlatformRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPlatformRegionsForbidden creates a GetCloudPlatformRegionsForbidden with default headers values
func NewGetCloudPlatformRegionsForbidden() *GetCloudPlatformRegionsForbidden {
	return &GetCloudPlatformRegionsForbidden{}
}

/*
GetCloudPlatformRegionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCloudPlatformRegionsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get cloud platform regions forbidden response has a 2xx status code
func (o *GetCloudPlatformRegionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud platform regions forbidden response has a 3xx status code
func (o *GetCloudPlatformRegionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud platform regions forbidden response has a 4xx status code
func (o *GetCloudPlatformRegionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud platform regions forbidden response has a 5xx status code
func (o *GetCloudPlatformRegionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud platform regions forbidden response a status code equal to that given
func (o *GetCloudPlatformRegionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get cloud platform regions forbidden response
func (o *GetCloudPlatformRegionsForbidden) Code() int {
	return 403
}

func (o *GetCloudPlatformRegionsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsForbidden %s", 403, payload)
}

func (o *GetCloudPlatformRegionsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsForbidden %s", 403, payload)
}

func (o *GetCloudPlatformRegionsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCloudPlatformRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPlatformRegionsNotFound creates a GetCloudPlatformRegionsNotFound with default headers values
func NewGetCloudPlatformRegionsNotFound() *GetCloudPlatformRegionsNotFound {
	return &GetCloudPlatformRegionsNotFound{}
}

/*
GetCloudPlatformRegionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCloudPlatformRegionsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get cloud platform regions not found response has a 2xx status code
func (o *GetCloudPlatformRegionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud platform regions not found response has a 3xx status code
func (o *GetCloudPlatformRegionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud platform regions not found response has a 4xx status code
func (o *GetCloudPlatformRegionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cloud platform regions not found response has a 5xx status code
func (o *GetCloudPlatformRegionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud platform regions not found response a status code equal to that given
func (o *GetCloudPlatformRegionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cloud platform regions not found response
func (o *GetCloudPlatformRegionsNotFound) Code() int {
	return 404
}

func (o *GetCloudPlatformRegionsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsNotFound %s", 404, payload)
}

func (o *GetCloudPlatformRegionsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsNotFound %s", 404, payload)
}

func (o *GetCloudPlatformRegionsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCloudPlatformRegionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudPlatformRegionsInternalServerError creates a GetCloudPlatformRegionsInternalServerError with default headers values
func NewGetCloudPlatformRegionsInternalServerError() *GetCloudPlatformRegionsInternalServerError {
	return &GetCloudPlatformRegionsInternalServerError{}
}

/*
GetCloudPlatformRegionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCloudPlatformRegionsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get cloud platform regions internal server error response has a 2xx status code
func (o *GetCloudPlatformRegionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cloud platform regions internal server error response has a 3xx status code
func (o *GetCloudPlatformRegionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud platform regions internal server error response has a 4xx status code
func (o *GetCloudPlatformRegionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud platform regions internal server error response has a 5xx status code
func (o *GetCloudPlatformRegionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cloud platform regions internal server error response a status code equal to that given
func (o *GetCloudPlatformRegionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cloud platform regions internal server error response
func (o *GetCloudPlatformRegionsInternalServerError) Code() int {
	return 500
}

func (o *GetCloudPlatformRegionsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsInternalServerError %s", 500, payload)
}

func (o *GetCloudPlatformRegionsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/general/cloud-platform/{cloud_platform}/regions][%d] getCloudPlatformRegionsInternalServerError %s", 500, payload)
}

func (o *GetCloudPlatformRegionsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetCloudPlatformRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
