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

// GetV1ReleasesReader is a Reader for the GetV1Releases structure.
type GetV1ReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/releases] GetV1Releases", response, response.Code())
	}
}

// NewGetV1ReleasesOK creates a GetV1ReleasesOK with default headers values
func NewGetV1ReleasesOK() *GetV1ReleasesOK {
	return &GetV1ReleasesOK{}
}

/*
GetV1ReleasesOK describes a response with status code 200, with default header values.

OK
*/
type GetV1ReleasesOK struct {
	Payload []*models.AppComponentRelease
}

// IsSuccess returns true when this get v1 releases o k response has a 2xx status code
func (o *GetV1ReleasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 releases o k response has a 3xx status code
func (o *GetV1ReleasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 releases o k response has a 4xx status code
func (o *GetV1ReleasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 releases o k response has a 5xx status code
func (o *GetV1ReleasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 releases o k response a status code equal to that given
func (o *GetV1ReleasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 releases o k response
func (o *GetV1ReleasesOK) Code() int {
	return 200
}

func (o *GetV1ReleasesOK) Error() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getV1ReleasesOK  %+v", 200, o.Payload)
}

func (o *GetV1ReleasesOK) String() string {
	return fmt.Sprintf("[GET /v1/releases][%d] getV1ReleasesOK  %+v", 200, o.Payload)
}

func (o *GetV1ReleasesOK) GetPayload() []*models.AppComponentRelease {
	return o.Payload
}

func (o *GetV1ReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
