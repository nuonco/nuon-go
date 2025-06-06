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

// GetTerraformStatesReader is a Reader for the GetTerraformStates structure.
type GetTerraformStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerraformStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTerraformStatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTerraformStatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTerraformStatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTerraformStatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTerraformStatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTerraformStatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/runners/terraform-workspace/{workspace_id}/states] GetTerraformStates", response, response.Code())
	}
}

// NewGetTerraformStatesOK creates a GetTerraformStatesOK with default headers values
func NewGetTerraformStatesOK() *GetTerraformStatesOK {
	return &GetTerraformStatesOK{}
}

/*
GetTerraformStatesOK describes a response with status code 200, with default header values.

OK
*/
type GetTerraformStatesOK struct {
	Payload []*models.AppTerraformWorkspaceState
}

// IsSuccess returns true when this get terraform states o k response has a 2xx status code
func (o *GetTerraformStatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get terraform states o k response has a 3xx status code
func (o *GetTerraformStatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform states o k response has a 4xx status code
func (o *GetTerraformStatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get terraform states o k response has a 5xx status code
func (o *GetTerraformStatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform states o k response a status code equal to that given
func (o *GetTerraformStatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get terraform states o k response
func (o *GetTerraformStatesOK) Code() int {
	return 200
}

func (o *GetTerraformStatesOK) Error() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesOK  %+v", 200, o.Payload)
}

func (o *GetTerraformStatesOK) String() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesOK  %+v", 200, o.Payload)
}

func (o *GetTerraformStatesOK) GetPayload() []*models.AppTerraformWorkspaceState {
	return o.Payload
}

func (o *GetTerraformStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformStatesBadRequest creates a GetTerraformStatesBadRequest with default headers values
func NewGetTerraformStatesBadRequest() *GetTerraformStatesBadRequest {
	return &GetTerraformStatesBadRequest{}
}

/*
GetTerraformStatesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTerraformStatesBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get terraform states bad request response has a 2xx status code
func (o *GetTerraformStatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform states bad request response has a 3xx status code
func (o *GetTerraformStatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform states bad request response has a 4xx status code
func (o *GetTerraformStatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get terraform states bad request response has a 5xx status code
func (o *GetTerraformStatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform states bad request response a status code equal to that given
func (o *GetTerraformStatesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get terraform states bad request response
func (o *GetTerraformStatesBadRequest) Code() int {
	return 400
}

func (o *GetTerraformStatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTerraformStatesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTerraformStatesBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetTerraformStatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformStatesUnauthorized creates a GetTerraformStatesUnauthorized with default headers values
func NewGetTerraformStatesUnauthorized() *GetTerraformStatesUnauthorized {
	return &GetTerraformStatesUnauthorized{}
}

/*
GetTerraformStatesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTerraformStatesUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get terraform states unauthorized response has a 2xx status code
func (o *GetTerraformStatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform states unauthorized response has a 3xx status code
func (o *GetTerraformStatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform states unauthorized response has a 4xx status code
func (o *GetTerraformStatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get terraform states unauthorized response has a 5xx status code
func (o *GetTerraformStatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform states unauthorized response a status code equal to that given
func (o *GetTerraformStatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get terraform states unauthorized response
func (o *GetTerraformStatesUnauthorized) Code() int {
	return 401
}

func (o *GetTerraformStatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTerraformStatesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTerraformStatesUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetTerraformStatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformStatesForbidden creates a GetTerraformStatesForbidden with default headers values
func NewGetTerraformStatesForbidden() *GetTerraformStatesForbidden {
	return &GetTerraformStatesForbidden{}
}

/*
GetTerraformStatesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTerraformStatesForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get terraform states forbidden response has a 2xx status code
func (o *GetTerraformStatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform states forbidden response has a 3xx status code
func (o *GetTerraformStatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform states forbidden response has a 4xx status code
func (o *GetTerraformStatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get terraform states forbidden response has a 5xx status code
func (o *GetTerraformStatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform states forbidden response a status code equal to that given
func (o *GetTerraformStatesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get terraform states forbidden response
func (o *GetTerraformStatesForbidden) Code() int {
	return 403
}

func (o *GetTerraformStatesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesForbidden  %+v", 403, o.Payload)
}

func (o *GetTerraformStatesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesForbidden  %+v", 403, o.Payload)
}

func (o *GetTerraformStatesForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetTerraformStatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformStatesNotFound creates a GetTerraformStatesNotFound with default headers values
func NewGetTerraformStatesNotFound() *GetTerraformStatesNotFound {
	return &GetTerraformStatesNotFound{}
}

/*
GetTerraformStatesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTerraformStatesNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get terraform states not found response has a 2xx status code
func (o *GetTerraformStatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform states not found response has a 3xx status code
func (o *GetTerraformStatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform states not found response has a 4xx status code
func (o *GetTerraformStatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get terraform states not found response has a 5xx status code
func (o *GetTerraformStatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform states not found response a status code equal to that given
func (o *GetTerraformStatesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get terraform states not found response
func (o *GetTerraformStatesNotFound) Code() int {
	return 404
}

func (o *GetTerraformStatesNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesNotFound  %+v", 404, o.Payload)
}

func (o *GetTerraformStatesNotFound) String() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesNotFound  %+v", 404, o.Payload)
}

func (o *GetTerraformStatesNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetTerraformStatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformStatesInternalServerError creates a GetTerraformStatesInternalServerError with default headers values
func NewGetTerraformStatesInternalServerError() *GetTerraformStatesInternalServerError {
	return &GetTerraformStatesInternalServerError{}
}

/*
GetTerraformStatesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTerraformStatesInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get terraform states internal server error response has a 2xx status code
func (o *GetTerraformStatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform states internal server error response has a 3xx status code
func (o *GetTerraformStatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform states internal server error response has a 4xx status code
func (o *GetTerraformStatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get terraform states internal server error response has a 5xx status code
func (o *GetTerraformStatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get terraform states internal server error response a status code equal to that given
func (o *GetTerraformStatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get terraform states internal server error response
func (o *GetTerraformStatesInternalServerError) Code() int {
	return 500
}

func (o *GetTerraformStatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTerraformStatesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/runners/terraform-workspace/{workspace_id}/states][%d] getTerraformStatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTerraformStatesInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetTerraformStatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
