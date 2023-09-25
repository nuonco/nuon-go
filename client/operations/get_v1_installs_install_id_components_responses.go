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

// GetV1InstallsInstallIDComponentsReader is a Reader for the GetV1InstallsInstallIDComponents structure.
type GetV1InstallsInstallIDComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallsInstallIDComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallsInstallIDComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/components] GetV1InstallsInstallIDComponents", response, response.Code())
	}
}

// NewGetV1InstallsInstallIDComponentsOK creates a GetV1InstallsInstallIDComponentsOK with default headers values
func NewGetV1InstallsInstallIDComponentsOK() *GetV1InstallsInstallIDComponentsOK {
	return &GetV1InstallsInstallIDComponentsOK{}
}

/*
GetV1InstallsInstallIDComponentsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallsInstallIDComponentsOK struct {
	Payload []*models.AppInstallComponent
}

// IsSuccess returns true when this get v1 installs install Id components o k response has a 2xx status code
func (o *GetV1InstallsInstallIDComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installs install Id components o k response has a 3xx status code
func (o *GetV1InstallsInstallIDComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id components o k response has a 4xx status code
func (o *GetV1InstallsInstallIDComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id components o k response has a 5xx status code
func (o *GetV1InstallsInstallIDComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id components o k response a status code equal to that given
func (o *GetV1InstallsInstallIDComponentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installs install Id components o k response
func (o *GetV1InstallsInstallIDComponentsOK) Code() int {
	return 200
}

func (o *GetV1InstallsInstallIDComponentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components][%d] getV1InstallsInstallIdComponentsOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDComponentsOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/components][%d] getV1InstallsInstallIdComponentsOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDComponentsOK) GetPayload() []*models.AppInstallComponent {
	return o.Payload
}

func (o *GetV1InstallsInstallIDComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}