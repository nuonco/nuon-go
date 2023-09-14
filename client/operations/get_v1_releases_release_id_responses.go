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

// GetV1ReleasesReleaseIDReader is a Reader for the GetV1ReleasesReleaseID structure.
type GetV1ReleasesReleaseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ReleasesReleaseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ReleasesReleaseIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases/{release_id}] GetV1ReleasesReleaseID", response, response.Code())
	}
}

// NewGetV1ReleasesReleaseIDOK creates a GetV1ReleasesReleaseIDOK with default headers values
func NewGetV1ReleasesReleaseIDOK() *GetV1ReleasesReleaseIDOK {
	return &GetV1ReleasesReleaseIDOK{}
}

/*
GetV1ReleasesReleaseIDOK describes a response with status code 200, with default header values.

OK
*/
type GetV1ReleasesReleaseIDOK struct {
	Payload *models.AppComponentRelease
}

// IsSuccess returns true when this get v1 releases release Id o k response has a 2xx status code
func (o *GetV1ReleasesReleaseIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 releases release Id o k response has a 3xx status code
func (o *GetV1ReleasesReleaseIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 releases release Id o k response has a 4xx status code
func (o *GetV1ReleasesReleaseIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 releases release Id o k response has a 5xx status code
func (o *GetV1ReleasesReleaseIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 releases release Id o k response a status code equal to that given
func (o *GetV1ReleasesReleaseIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 releases release Id o k response
func (o *GetV1ReleasesReleaseIDOK) Code() int {
	return 200
}

func (o *GetV1ReleasesReleaseIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/releases/{release_id}][%d] getV1ReleasesReleaseIdOK  %+v", 200, o.Payload)
}

func (o *GetV1ReleasesReleaseIDOK) String() string {
	return fmt.Sprintf("[GET /v1/releases/{release_id}][%d] getV1ReleasesReleaseIdOK  %+v", 200, o.Payload)
}

func (o *GetV1ReleasesReleaseIDOK) GetPayload() *models.AppComponentRelease {
	return o.Payload
}

func (o *GetV1ReleasesReleaseIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppComponentRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
