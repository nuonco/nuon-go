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
