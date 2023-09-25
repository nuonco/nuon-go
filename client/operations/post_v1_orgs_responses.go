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

// PostV1OrgsReader is a Reader for the PostV1Orgs structure.
type PostV1OrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1OrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1OrgsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/orgs] PostV1Orgs", response, response.Code())
	}
}

// NewPostV1OrgsCreated creates a PostV1OrgsCreated with default headers values
func NewPostV1OrgsCreated() *PostV1OrgsCreated {
	return &PostV1OrgsCreated{}
}

/*
PostV1OrgsCreated describes a response with status code 201, with default header values.

Created
*/
type PostV1OrgsCreated struct {
	Payload *models.AppOrg
}

// IsSuccess returns true when this post v1 orgs created response has a 2xx status code
func (o *PostV1OrgsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 orgs created response has a 3xx status code
func (o *PostV1OrgsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 orgs created response has a 4xx status code
func (o *PostV1OrgsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 orgs created response has a 5xx status code
func (o *PostV1OrgsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 orgs created response a status code equal to that given
func (o *PostV1OrgsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 orgs created response
func (o *PostV1OrgsCreated) Code() int {
	return 201
}

func (o *PostV1OrgsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] postV1OrgsCreated  %+v", 201, o.Payload)
}

func (o *PostV1OrgsCreated) String() string {
	return fmt.Sprintf("[POST /v1/orgs][%d] postV1OrgsCreated  %+v", 201, o.Payload)
}

func (o *PostV1OrgsCreated) GetPayload() *models.AppOrg {
	return o.Payload
}

func (o *PostV1OrgsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppOrg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}