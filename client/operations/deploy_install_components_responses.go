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

// DeployInstallComponentsReader is a Reader for the DeployInstallComponents structure.
type DeployInstallComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeployInstallComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeployInstallComponentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeployInstallComponentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeployInstallComponentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeployInstallComponentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeployInstallComponentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeployInstallComponentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/installs/{install_id}/components/deploy-all] DeployInstallComponents", response, response.Code())
	}
}

// NewDeployInstallComponentsCreated creates a DeployInstallComponentsCreated with default headers values
func NewDeployInstallComponentsCreated() *DeployInstallComponentsCreated {
	return &DeployInstallComponentsCreated{}
}

/*
DeployInstallComponentsCreated describes a response with status code 201, with default header values.

Created
*/
type DeployInstallComponentsCreated struct {
	Payload string
}

// IsSuccess returns true when this deploy install components created response has a 2xx status code
func (o *DeployInstallComponentsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this deploy install components created response has a 3xx status code
func (o *DeployInstallComponentsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy install components created response has a 4xx status code
func (o *DeployInstallComponentsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this deploy install components created response has a 5xx status code
func (o *DeployInstallComponentsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy install components created response a status code equal to that given
func (o *DeployInstallComponentsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the deploy install components created response
func (o *DeployInstallComponentsCreated) Code() int {
	return 201
}

func (o *DeployInstallComponentsCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsCreated %s", 201, payload)
}

func (o *DeployInstallComponentsCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsCreated %s", 201, payload)
}

func (o *DeployInstallComponentsCreated) GetPayload() string {
	return o.Payload
}

func (o *DeployInstallComponentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployInstallComponentsBadRequest creates a DeployInstallComponentsBadRequest with default headers values
func NewDeployInstallComponentsBadRequest() *DeployInstallComponentsBadRequest {
	return &DeployInstallComponentsBadRequest{}
}

/*
DeployInstallComponentsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeployInstallComponentsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this deploy install components bad request response has a 2xx status code
func (o *DeployInstallComponentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy install components bad request response has a 3xx status code
func (o *DeployInstallComponentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy install components bad request response has a 4xx status code
func (o *DeployInstallComponentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this deploy install components bad request response has a 5xx status code
func (o *DeployInstallComponentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy install components bad request response a status code equal to that given
func (o *DeployInstallComponentsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the deploy install components bad request response
func (o *DeployInstallComponentsBadRequest) Code() int {
	return 400
}

func (o *DeployInstallComponentsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsBadRequest %s", 400, payload)
}

func (o *DeployInstallComponentsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsBadRequest %s", 400, payload)
}

func (o *DeployInstallComponentsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeployInstallComponentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployInstallComponentsUnauthorized creates a DeployInstallComponentsUnauthorized with default headers values
func NewDeployInstallComponentsUnauthorized() *DeployInstallComponentsUnauthorized {
	return &DeployInstallComponentsUnauthorized{}
}

/*
DeployInstallComponentsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeployInstallComponentsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this deploy install components unauthorized response has a 2xx status code
func (o *DeployInstallComponentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy install components unauthorized response has a 3xx status code
func (o *DeployInstallComponentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy install components unauthorized response has a 4xx status code
func (o *DeployInstallComponentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this deploy install components unauthorized response has a 5xx status code
func (o *DeployInstallComponentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy install components unauthorized response a status code equal to that given
func (o *DeployInstallComponentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the deploy install components unauthorized response
func (o *DeployInstallComponentsUnauthorized) Code() int {
	return 401
}

func (o *DeployInstallComponentsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsUnauthorized %s", 401, payload)
}

func (o *DeployInstallComponentsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsUnauthorized %s", 401, payload)
}

func (o *DeployInstallComponentsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeployInstallComponentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployInstallComponentsForbidden creates a DeployInstallComponentsForbidden with default headers values
func NewDeployInstallComponentsForbidden() *DeployInstallComponentsForbidden {
	return &DeployInstallComponentsForbidden{}
}

/*
DeployInstallComponentsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeployInstallComponentsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this deploy install components forbidden response has a 2xx status code
func (o *DeployInstallComponentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy install components forbidden response has a 3xx status code
func (o *DeployInstallComponentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy install components forbidden response has a 4xx status code
func (o *DeployInstallComponentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this deploy install components forbidden response has a 5xx status code
func (o *DeployInstallComponentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy install components forbidden response a status code equal to that given
func (o *DeployInstallComponentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the deploy install components forbidden response
func (o *DeployInstallComponentsForbidden) Code() int {
	return 403
}

func (o *DeployInstallComponentsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsForbidden %s", 403, payload)
}

func (o *DeployInstallComponentsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsForbidden %s", 403, payload)
}

func (o *DeployInstallComponentsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeployInstallComponentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployInstallComponentsNotFound creates a DeployInstallComponentsNotFound with default headers values
func NewDeployInstallComponentsNotFound() *DeployInstallComponentsNotFound {
	return &DeployInstallComponentsNotFound{}
}

/*
DeployInstallComponentsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeployInstallComponentsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this deploy install components not found response has a 2xx status code
func (o *DeployInstallComponentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy install components not found response has a 3xx status code
func (o *DeployInstallComponentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy install components not found response has a 4xx status code
func (o *DeployInstallComponentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this deploy install components not found response has a 5xx status code
func (o *DeployInstallComponentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this deploy install components not found response a status code equal to that given
func (o *DeployInstallComponentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the deploy install components not found response
func (o *DeployInstallComponentsNotFound) Code() int {
	return 404
}

func (o *DeployInstallComponentsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsNotFound %s", 404, payload)
}

func (o *DeployInstallComponentsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsNotFound %s", 404, payload)
}

func (o *DeployInstallComponentsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeployInstallComponentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeployInstallComponentsInternalServerError creates a DeployInstallComponentsInternalServerError with default headers values
func NewDeployInstallComponentsInternalServerError() *DeployInstallComponentsInternalServerError {
	return &DeployInstallComponentsInternalServerError{}
}

/*
DeployInstallComponentsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeployInstallComponentsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this deploy install components internal server error response has a 2xx status code
func (o *DeployInstallComponentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deploy install components internal server error response has a 3xx status code
func (o *DeployInstallComponentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deploy install components internal server error response has a 4xx status code
func (o *DeployInstallComponentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this deploy install components internal server error response has a 5xx status code
func (o *DeployInstallComponentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this deploy install components internal server error response a status code equal to that given
func (o *DeployInstallComponentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the deploy install components internal server error response
func (o *DeployInstallComponentsInternalServerError) Code() int {
	return 500
}

func (o *DeployInstallComponentsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsInternalServerError %s", 500, payload)
}

func (o *DeployInstallComponentsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/installs/{install_id}/components/deploy-all][%d] deployInstallComponentsInternalServerError %s", 500, payload)
}

func (o *DeployInstallComponentsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeployInstallComponentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
