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

// DeleteInstallReader is a Reader for the DeleteInstall structure.
type DeleteInstallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInstallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteInstallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteInstallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteInstallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInstallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInstallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/installs/{install_id}] DeleteInstall", response, response.Code())
	}
}

// NewDeleteInstallOK creates a DeleteInstallOK with default headers values
func NewDeleteInstallOK() *DeleteInstallOK {
	return &DeleteInstallOK{}
}

/*
DeleteInstallOK describes a response with status code 200, with default header values.

OK
*/
type DeleteInstallOK struct {
	Payload bool
}

// IsSuccess returns true when this delete install o k response has a 2xx status code
func (o *DeleteInstallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete install o k response has a 3xx status code
func (o *DeleteInstallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete install o k response has a 4xx status code
func (o *DeleteInstallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete install o k response has a 5xx status code
func (o *DeleteInstallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete install o k response a status code equal to that given
func (o *DeleteInstallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete install o k response
func (o *DeleteInstallOK) Code() int {
	return 200
}

func (o *DeleteInstallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallOK %s", 200, payload)
}

func (o *DeleteInstallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallOK %s", 200, payload)
}

func (o *DeleteInstallOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteInstallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallBadRequest creates a DeleteInstallBadRequest with default headers values
func NewDeleteInstallBadRequest() *DeleteInstallBadRequest {
	return &DeleteInstallBadRequest{}
}

/*
DeleteInstallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteInstallBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete install bad request response has a 2xx status code
func (o *DeleteInstallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete install bad request response has a 3xx status code
func (o *DeleteInstallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete install bad request response has a 4xx status code
func (o *DeleteInstallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete install bad request response has a 5xx status code
func (o *DeleteInstallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete install bad request response a status code equal to that given
func (o *DeleteInstallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete install bad request response
func (o *DeleteInstallBadRequest) Code() int {
	return 400
}

func (o *DeleteInstallBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallBadRequest %s", 400, payload)
}

func (o *DeleteInstallBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallBadRequest %s", 400, payload)
}

func (o *DeleteInstallBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallUnauthorized creates a DeleteInstallUnauthorized with default headers values
func NewDeleteInstallUnauthorized() *DeleteInstallUnauthorized {
	return &DeleteInstallUnauthorized{}
}

/*
DeleteInstallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteInstallUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete install unauthorized response has a 2xx status code
func (o *DeleteInstallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete install unauthorized response has a 3xx status code
func (o *DeleteInstallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete install unauthorized response has a 4xx status code
func (o *DeleteInstallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete install unauthorized response has a 5xx status code
func (o *DeleteInstallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete install unauthorized response a status code equal to that given
func (o *DeleteInstallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete install unauthorized response
func (o *DeleteInstallUnauthorized) Code() int {
	return 401
}

func (o *DeleteInstallUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallUnauthorized %s", 401, payload)
}

func (o *DeleteInstallUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallUnauthorized %s", 401, payload)
}

func (o *DeleteInstallUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallForbidden creates a DeleteInstallForbidden with default headers values
func NewDeleteInstallForbidden() *DeleteInstallForbidden {
	return &DeleteInstallForbidden{}
}

/*
DeleteInstallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteInstallForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete install forbidden response has a 2xx status code
func (o *DeleteInstallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete install forbidden response has a 3xx status code
func (o *DeleteInstallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete install forbidden response has a 4xx status code
func (o *DeleteInstallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete install forbidden response has a 5xx status code
func (o *DeleteInstallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete install forbidden response a status code equal to that given
func (o *DeleteInstallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete install forbidden response
func (o *DeleteInstallForbidden) Code() int {
	return 403
}

func (o *DeleteInstallForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallForbidden %s", 403, payload)
}

func (o *DeleteInstallForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallForbidden %s", 403, payload)
}

func (o *DeleteInstallForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallNotFound creates a DeleteInstallNotFound with default headers values
func NewDeleteInstallNotFound() *DeleteInstallNotFound {
	return &DeleteInstallNotFound{}
}

/*
DeleteInstallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteInstallNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete install not found response has a 2xx status code
func (o *DeleteInstallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete install not found response has a 3xx status code
func (o *DeleteInstallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete install not found response has a 4xx status code
func (o *DeleteInstallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete install not found response has a 5xx status code
func (o *DeleteInstallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete install not found response a status code equal to that given
func (o *DeleteInstallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete install not found response
func (o *DeleteInstallNotFound) Code() int {
	return 404
}

func (o *DeleteInstallNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallNotFound %s", 404, payload)
}

func (o *DeleteInstallNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallNotFound %s", 404, payload)
}

func (o *DeleteInstallNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstallInternalServerError creates a DeleteInstallInternalServerError with default headers values
func NewDeleteInstallInternalServerError() *DeleteInstallInternalServerError {
	return &DeleteInstallInternalServerError{}
}

/*
DeleteInstallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteInstallInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete install internal server error response has a 2xx status code
func (o *DeleteInstallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete install internal server error response has a 3xx status code
func (o *DeleteInstallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete install internal server error response has a 4xx status code
func (o *DeleteInstallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete install internal server error response has a 5xx status code
func (o *DeleteInstallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete install internal server error response a status code equal to that given
func (o *DeleteInstallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete install internal server error response
func (o *DeleteInstallInternalServerError) Code() int {
	return 500
}

func (o *DeleteInstallInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallInternalServerError %s", 500, payload)
}

func (o *DeleteInstallInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /v1/installs/{install_id}][%d] deleteInstallInternalServerError %s", 500, payload)
}

func (o *DeleteInstallInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteInstallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
