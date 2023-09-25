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

// GetV1SandboxesSandboxIDReader is a Reader for the GetV1SandboxesSandboxID structure.
type GetV1SandboxesSandboxIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1SandboxesSandboxIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1SandboxesSandboxIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/sandboxes/{sandbox_id}] GetV1SandboxesSandboxID", response, response.Code())
	}
}

// NewGetV1SandboxesSandboxIDOK creates a GetV1SandboxesSandboxIDOK with default headers values
func NewGetV1SandboxesSandboxIDOK() *GetV1SandboxesSandboxIDOK {
	return &GetV1SandboxesSandboxIDOK{}
}

/*
GetV1SandboxesSandboxIDOK describes a response with status code 200, with default header values.

OK
*/
type GetV1SandboxesSandboxIDOK struct {
	Payload *models.AppSandbox
}

// IsSuccess returns true when this get v1 sandboxes sandbox Id o k response has a 2xx status code
func (o *GetV1SandboxesSandboxIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 sandboxes sandbox Id o k response has a 3xx status code
func (o *GetV1SandboxesSandboxIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 sandboxes sandbox Id o k response has a 4xx status code
func (o *GetV1SandboxesSandboxIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 sandboxes sandbox Id o k response has a 5xx status code
func (o *GetV1SandboxesSandboxIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 sandboxes sandbox Id o k response a status code equal to that given
func (o *GetV1SandboxesSandboxIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 sandboxes sandbox Id o k response
func (o *GetV1SandboxesSandboxIDOK) Code() int {
	return 200
}

func (o *GetV1SandboxesSandboxIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getV1SandboxesSandboxIdOK  %+v", 200, o.Payload)
}

func (o *GetV1SandboxesSandboxIDOK) String() string {
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getV1SandboxesSandboxIdOK  %+v", 200, o.Payload)
}

func (o *GetV1SandboxesSandboxIDOK) GetPayload() *models.AppSandbox {
	return o.Payload
}

func (o *GetV1SandboxesSandboxIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppSandbox)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}