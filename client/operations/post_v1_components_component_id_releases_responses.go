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

// PostV1ComponentsComponentIDReleasesReader is a Reader for the PostV1ComponentsComponentIDReleases structure.
type PostV1ComponentsComponentIDReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ComponentsComponentIDReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ComponentsComponentIDReleasesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/components/{component_id}/releases] PostV1ComponentsComponentIDReleases", response, response.Code())
	}
}

// NewPostV1ComponentsComponentIDReleasesCreated creates a PostV1ComponentsComponentIDReleasesCreated with default headers values
func NewPostV1ComponentsComponentIDReleasesCreated() *PostV1ComponentsComponentIDReleasesCreated {
	return &PostV1ComponentsComponentIDReleasesCreated{}
}

/*
PostV1ComponentsComponentIDReleasesCreated describes a response with status code 201, with default header values.

Created
*/
type PostV1ComponentsComponentIDReleasesCreated struct {
	Payload *models.AppComponentRelease
}

// IsSuccess returns true when this post v1 components component Id releases created response has a 2xx status code
func (o *PostV1ComponentsComponentIDReleasesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 components component Id releases created response has a 3xx status code
func (o *PostV1ComponentsComponentIDReleasesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 components component Id releases created response has a 4xx status code
func (o *PostV1ComponentsComponentIDReleasesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 components component Id releases created response has a 5xx status code
func (o *PostV1ComponentsComponentIDReleasesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 components component Id releases created response a status code equal to that given
func (o *PostV1ComponentsComponentIDReleasesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 components component Id releases created response
func (o *PostV1ComponentsComponentIDReleasesCreated) Code() int {
	return 201
}

func (o *PostV1ComponentsComponentIDReleasesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/releases][%d] postV1ComponentsComponentIdReleasesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsComponentIDReleasesCreated) String() string {
	return fmt.Sprintf("[POST /v1/components/{component_id}/releases][%d] postV1ComponentsComponentIdReleasesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ComponentsComponentIDReleasesCreated) GetPayload() *models.AppComponentRelease {
	return o.Payload
}

func (o *PostV1ComponentsComponentIDReleasesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppComponentRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}