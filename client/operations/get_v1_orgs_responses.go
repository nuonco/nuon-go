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

// GetV1OrgsReader is a Reader for the GetV1Orgs structure.
type GetV1OrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1OrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1OrgsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1OrgsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1OrgsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1OrgsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1OrgsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1OrgsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/orgs] GetV1Orgs", response, response.Code())
	}
}

// NewGetV1OrgsOK creates a GetV1OrgsOK with default headers values
func NewGetV1OrgsOK() *GetV1OrgsOK {
	return &GetV1OrgsOK{}
}

/*
GetV1OrgsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1OrgsOK struct {
	Payload []*models.AppOrg
}

// IsSuccess returns true when this get v1 orgs o k response has a 2xx status code
func (o *GetV1OrgsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 orgs o k response has a 3xx status code
func (o *GetV1OrgsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 orgs o k response has a 4xx status code
func (o *GetV1OrgsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 orgs o k response has a 5xx status code
func (o *GetV1OrgsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 orgs o k response a status code equal to that given
func (o *GetV1OrgsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 orgs o k response
func (o *GetV1OrgsOK) Code() int {
	return 200
}

func (o *GetV1OrgsOK) Error() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsOK  %+v", 200, o.Payload)
}

func (o *GetV1OrgsOK) String() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsOK  %+v", 200, o.Payload)
}

func (o *GetV1OrgsOK) GetPayload() []*models.AppOrg {
	return o.Payload
}

func (o *GetV1OrgsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1OrgsBadRequest creates a GetV1OrgsBadRequest with default headers values
func NewGetV1OrgsBadRequest() *GetV1OrgsBadRequest {
	return &GetV1OrgsBadRequest{}
}

/*
GetV1OrgsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1OrgsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 orgs bad request response has a 2xx status code
func (o *GetV1OrgsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 orgs bad request response has a 3xx status code
func (o *GetV1OrgsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 orgs bad request response has a 4xx status code
func (o *GetV1OrgsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 orgs bad request response has a 5xx status code
func (o *GetV1OrgsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 orgs bad request response a status code equal to that given
func (o *GetV1OrgsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 orgs bad request response
func (o *GetV1OrgsBadRequest) Code() int {
	return 400
}

func (o *GetV1OrgsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1OrgsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1OrgsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1OrgsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1OrgsUnauthorized creates a GetV1OrgsUnauthorized with default headers values
func NewGetV1OrgsUnauthorized() *GetV1OrgsUnauthorized {
	return &GetV1OrgsUnauthorized{}
}

/*
GetV1OrgsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1OrgsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 orgs unauthorized response has a 2xx status code
func (o *GetV1OrgsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 orgs unauthorized response has a 3xx status code
func (o *GetV1OrgsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 orgs unauthorized response has a 4xx status code
func (o *GetV1OrgsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 orgs unauthorized response has a 5xx status code
func (o *GetV1OrgsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 orgs unauthorized response a status code equal to that given
func (o *GetV1OrgsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 orgs unauthorized response
func (o *GetV1OrgsUnauthorized) Code() int {
	return 401
}

func (o *GetV1OrgsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1OrgsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1OrgsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1OrgsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1OrgsForbidden creates a GetV1OrgsForbidden with default headers values
func NewGetV1OrgsForbidden() *GetV1OrgsForbidden {
	return &GetV1OrgsForbidden{}
}

/*
GetV1OrgsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1OrgsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 orgs forbidden response has a 2xx status code
func (o *GetV1OrgsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 orgs forbidden response has a 3xx status code
func (o *GetV1OrgsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 orgs forbidden response has a 4xx status code
func (o *GetV1OrgsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 orgs forbidden response has a 5xx status code
func (o *GetV1OrgsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 orgs forbidden response a status code equal to that given
func (o *GetV1OrgsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 orgs forbidden response
func (o *GetV1OrgsForbidden) Code() int {
	return 403
}

func (o *GetV1OrgsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsForbidden  %+v", 403, o.Payload)
}

func (o *GetV1OrgsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsForbidden  %+v", 403, o.Payload)
}

func (o *GetV1OrgsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1OrgsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1OrgsNotFound creates a GetV1OrgsNotFound with default headers values
func NewGetV1OrgsNotFound() *GetV1OrgsNotFound {
	return &GetV1OrgsNotFound{}
}

/*
GetV1OrgsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1OrgsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 orgs not found response has a 2xx status code
func (o *GetV1OrgsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 orgs not found response has a 3xx status code
func (o *GetV1OrgsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 orgs not found response has a 4xx status code
func (o *GetV1OrgsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 orgs not found response has a 5xx status code
func (o *GetV1OrgsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 orgs not found response a status code equal to that given
func (o *GetV1OrgsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 orgs not found response
func (o *GetV1OrgsNotFound) Code() int {
	return 404
}

func (o *GetV1OrgsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsNotFound  %+v", 404, o.Payload)
}

func (o *GetV1OrgsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsNotFound  %+v", 404, o.Payload)
}

func (o *GetV1OrgsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1OrgsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1OrgsInternalServerError creates a GetV1OrgsInternalServerError with default headers values
func NewGetV1OrgsInternalServerError() *GetV1OrgsInternalServerError {
	return &GetV1OrgsInternalServerError{}
}

/*
GetV1OrgsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1OrgsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 orgs internal server error response has a 2xx status code
func (o *GetV1OrgsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 orgs internal server error response has a 3xx status code
func (o *GetV1OrgsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 orgs internal server error response has a 4xx status code
func (o *GetV1OrgsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 orgs internal server error response has a 5xx status code
func (o *GetV1OrgsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 orgs internal server error response a status code equal to that given
func (o *GetV1OrgsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 orgs internal server error response
func (o *GetV1OrgsInternalServerError) Code() int {
	return 500
}

func (o *GetV1OrgsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1OrgsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/orgs][%d] getV1OrgsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1OrgsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1OrgsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
