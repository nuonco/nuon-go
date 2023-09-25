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

// GetV1GeneralCurrentUserReader is a Reader for the GetV1GeneralCurrentUser structure.
type GetV1GeneralCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1GeneralCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1GeneralCurrentUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/general/current-user] GetV1GeneralCurrentUser", response, response.Code())
	}
}

// NewGetV1GeneralCurrentUserOK creates a GetV1GeneralCurrentUserOK with default headers values
func NewGetV1GeneralCurrentUserOK() *GetV1GeneralCurrentUserOK {
	return &GetV1GeneralCurrentUserOK{}
}

/*
GetV1GeneralCurrentUserOK describes a response with status code 200, with default header values.

OK
*/
type GetV1GeneralCurrentUserOK struct {
	Payload *models.AppUserToken
}

// IsSuccess returns true when this get v1 general current user o k response has a 2xx status code
func (o *GetV1GeneralCurrentUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 general current user o k response has a 3xx status code
func (o *GetV1GeneralCurrentUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 general current user o k response has a 4xx status code
func (o *GetV1GeneralCurrentUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 general current user o k response has a 5xx status code
func (o *GetV1GeneralCurrentUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 general current user o k response a status code equal to that given
func (o *GetV1GeneralCurrentUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 general current user o k response
func (o *GetV1GeneralCurrentUserOK) Code() int {
	return 200
}

func (o *GetV1GeneralCurrentUserOK) Error() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getV1GeneralCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetV1GeneralCurrentUserOK) String() string {
	return fmt.Sprintf("[GET /v1/general/current-user][%d] getV1GeneralCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetV1GeneralCurrentUserOK) GetPayload() *models.AppUserToken {
	return o.Payload
}

func (o *GetV1GeneralCurrentUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppUserToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}