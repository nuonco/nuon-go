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

// GetV1ComponentsComponentIDBuildsReader is a Reader for the GetV1ComponentsComponentIDBuilds structure.
type GetV1ComponentsComponentIDBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ComponentsComponentIDBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ComponentsComponentIDBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1ComponentsComponentIDBuildsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1ComponentsComponentIDBuildsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1ComponentsComponentIDBuildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1ComponentsComponentIDBuildsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1ComponentsComponentIDBuildsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/components/{component_id}/builds] GetV1ComponentsComponentIDBuilds", response, response.Code())
	}
}

// NewGetV1ComponentsComponentIDBuildsOK creates a GetV1ComponentsComponentIDBuildsOK with default headers values
func NewGetV1ComponentsComponentIDBuildsOK() *GetV1ComponentsComponentIDBuildsOK {
	return &GetV1ComponentsComponentIDBuildsOK{}
}

/*
GetV1ComponentsComponentIDBuildsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1ComponentsComponentIDBuildsOK struct {
	Payload []*models.AppComponentBuild
}

// IsSuccess returns true when this get v1 components component Id builds o k response has a 2xx status code
func (o *GetV1ComponentsComponentIDBuildsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 components component Id builds o k response has a 3xx status code
func (o *GetV1ComponentsComponentIDBuildsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 components component Id builds o k response has a 4xx status code
func (o *GetV1ComponentsComponentIDBuildsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 components component Id builds o k response has a 5xx status code
func (o *GetV1ComponentsComponentIDBuildsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 components component Id builds o k response a status code equal to that given
func (o *GetV1ComponentsComponentIDBuildsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 components component Id builds o k response
func (o *GetV1ComponentsComponentIDBuildsOK) Code() int {
	return 200
}

func (o *GetV1ComponentsComponentIDBuildsOK) Error() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsOK  %+v", 200, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsOK) String() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsOK  %+v", 200, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsOK) GetPayload() []*models.AppComponentBuild {
	return o.Payload
}

func (o *GetV1ComponentsComponentIDBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ComponentsComponentIDBuildsBadRequest creates a GetV1ComponentsComponentIDBuildsBadRequest with default headers values
func NewGetV1ComponentsComponentIDBuildsBadRequest() *GetV1ComponentsComponentIDBuildsBadRequest {
	return &GetV1ComponentsComponentIDBuildsBadRequest{}
}

/*
GetV1ComponentsComponentIDBuildsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1ComponentsComponentIDBuildsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 components component Id builds bad request response has a 2xx status code
func (o *GetV1ComponentsComponentIDBuildsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 components component Id builds bad request response has a 3xx status code
func (o *GetV1ComponentsComponentIDBuildsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 components component Id builds bad request response has a 4xx status code
func (o *GetV1ComponentsComponentIDBuildsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 components component Id builds bad request response has a 5xx status code
func (o *GetV1ComponentsComponentIDBuildsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 components component Id builds bad request response a status code equal to that given
func (o *GetV1ComponentsComponentIDBuildsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 components component Id builds bad request response
func (o *GetV1ComponentsComponentIDBuildsBadRequest) Code() int {
	return 400
}

func (o *GetV1ComponentsComponentIDBuildsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1ComponentsComponentIDBuildsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ComponentsComponentIDBuildsUnauthorized creates a GetV1ComponentsComponentIDBuildsUnauthorized with default headers values
func NewGetV1ComponentsComponentIDBuildsUnauthorized() *GetV1ComponentsComponentIDBuildsUnauthorized {
	return &GetV1ComponentsComponentIDBuildsUnauthorized{}
}

/*
GetV1ComponentsComponentIDBuildsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1ComponentsComponentIDBuildsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 components component Id builds unauthorized response has a 2xx status code
func (o *GetV1ComponentsComponentIDBuildsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 components component Id builds unauthorized response has a 3xx status code
func (o *GetV1ComponentsComponentIDBuildsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 components component Id builds unauthorized response has a 4xx status code
func (o *GetV1ComponentsComponentIDBuildsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 components component Id builds unauthorized response has a 5xx status code
func (o *GetV1ComponentsComponentIDBuildsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 components component Id builds unauthorized response a status code equal to that given
func (o *GetV1ComponentsComponentIDBuildsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 components component Id builds unauthorized response
func (o *GetV1ComponentsComponentIDBuildsUnauthorized) Code() int {
	return 401
}

func (o *GetV1ComponentsComponentIDBuildsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1ComponentsComponentIDBuildsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ComponentsComponentIDBuildsForbidden creates a GetV1ComponentsComponentIDBuildsForbidden with default headers values
func NewGetV1ComponentsComponentIDBuildsForbidden() *GetV1ComponentsComponentIDBuildsForbidden {
	return &GetV1ComponentsComponentIDBuildsForbidden{}
}

/*
GetV1ComponentsComponentIDBuildsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1ComponentsComponentIDBuildsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 components component Id builds forbidden response has a 2xx status code
func (o *GetV1ComponentsComponentIDBuildsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 components component Id builds forbidden response has a 3xx status code
func (o *GetV1ComponentsComponentIDBuildsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 components component Id builds forbidden response has a 4xx status code
func (o *GetV1ComponentsComponentIDBuildsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 components component Id builds forbidden response has a 5xx status code
func (o *GetV1ComponentsComponentIDBuildsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 components component Id builds forbidden response a status code equal to that given
func (o *GetV1ComponentsComponentIDBuildsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 components component Id builds forbidden response
func (o *GetV1ComponentsComponentIDBuildsForbidden) Code() int {
	return 403
}

func (o *GetV1ComponentsComponentIDBuildsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsForbidden  %+v", 403, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsForbidden  %+v", 403, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1ComponentsComponentIDBuildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ComponentsComponentIDBuildsNotFound creates a GetV1ComponentsComponentIDBuildsNotFound with default headers values
func NewGetV1ComponentsComponentIDBuildsNotFound() *GetV1ComponentsComponentIDBuildsNotFound {
	return &GetV1ComponentsComponentIDBuildsNotFound{}
}

/*
GetV1ComponentsComponentIDBuildsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1ComponentsComponentIDBuildsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 components component Id builds not found response has a 2xx status code
func (o *GetV1ComponentsComponentIDBuildsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 components component Id builds not found response has a 3xx status code
func (o *GetV1ComponentsComponentIDBuildsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 components component Id builds not found response has a 4xx status code
func (o *GetV1ComponentsComponentIDBuildsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 components component Id builds not found response has a 5xx status code
func (o *GetV1ComponentsComponentIDBuildsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 components component Id builds not found response a status code equal to that given
func (o *GetV1ComponentsComponentIDBuildsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 components component Id builds not found response
func (o *GetV1ComponentsComponentIDBuildsNotFound) Code() int {
	return 404
}

func (o *GetV1ComponentsComponentIDBuildsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsNotFound  %+v", 404, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsNotFound  %+v", 404, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1ComponentsComponentIDBuildsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1ComponentsComponentIDBuildsInternalServerError creates a GetV1ComponentsComponentIDBuildsInternalServerError with default headers values
func NewGetV1ComponentsComponentIDBuildsInternalServerError() *GetV1ComponentsComponentIDBuildsInternalServerError {
	return &GetV1ComponentsComponentIDBuildsInternalServerError{}
}

/*
GetV1ComponentsComponentIDBuildsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1ComponentsComponentIDBuildsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 components component Id builds internal server error response has a 2xx status code
func (o *GetV1ComponentsComponentIDBuildsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 components component Id builds internal server error response has a 3xx status code
func (o *GetV1ComponentsComponentIDBuildsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 components component Id builds internal server error response has a 4xx status code
func (o *GetV1ComponentsComponentIDBuildsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 components component Id builds internal server error response has a 5xx status code
func (o *GetV1ComponentsComponentIDBuildsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 components component Id builds internal server error response a status code equal to that given
func (o *GetV1ComponentsComponentIDBuildsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 components component Id builds internal server error response
func (o *GetV1ComponentsComponentIDBuildsInternalServerError) Code() int {
	return 500
}

func (o *GetV1ComponentsComponentIDBuildsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/components/{component_id}/builds][%d] getV1ComponentsComponentIdBuildsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1ComponentsComponentIDBuildsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1ComponentsComponentIDBuildsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
