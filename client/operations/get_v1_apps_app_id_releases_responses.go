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

// GetV1AppsAppIDReleasesReader is a Reader for the GetV1AppsAppIDReleases structure.
type GetV1AppsAppIDReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AppsAppIDReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AppsAppIDReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/apps/{app_id}/releases] GetV1AppsAppIDReleases", response, response.Code())
	}
}

// NewGetV1AppsAppIDReleasesOK creates a GetV1AppsAppIDReleasesOK with default headers values
func NewGetV1AppsAppIDReleasesOK() *GetV1AppsAppIDReleasesOK {
	return &GetV1AppsAppIDReleasesOK{}
}

/*
GetV1AppsAppIDReleasesOK describes a response with status code 200, with default header values.

OK
*/
type GetV1AppsAppIDReleasesOK struct {
	Payload []*models.AppComponentRelease
}

// IsSuccess returns true when this get v1 apps app Id releases o k response has a 2xx status code
func (o *GetV1AppsAppIDReleasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 apps app Id releases o k response has a 3xx status code
func (o *GetV1AppsAppIDReleasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 apps app Id releases o k response has a 4xx status code
func (o *GetV1AppsAppIDReleasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 apps app Id releases o k response has a 5xx status code
func (o *GetV1AppsAppIDReleasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 apps app Id releases o k response a status code equal to that given
func (o *GetV1AppsAppIDReleasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 apps app Id releases o k response
func (o *GetV1AppsAppIDReleasesOK) Code() int {
	return 200
}

func (o *GetV1AppsAppIDReleasesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getV1AppsAppIdReleasesOK  %+v", 200, o.Payload)
}

func (o *GetV1AppsAppIDReleasesOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/{app_id}/releases][%d] getV1AppsAppIdReleasesOK  %+v", 200, o.Payload)
}

func (o *GetV1AppsAppIDReleasesOK) GetPayload() []*models.AppComponentRelease {
	return o.Payload
}

func (o *GetV1AppsAppIDReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
