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

// GetV1InstallsInstallIDComponentComponentIDReader is a Reader for the GetV1InstallsInstallIDComponentComponentID structure.
type GetV1InstallsInstallIDComponentComponentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallsInstallIDComponentComponentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallsInstallIDComponentComponentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/component/{component_id}] GetV1InstallsInstallIDComponentComponentID", response, response.Code())
	}
}

// NewGetV1InstallsInstallIDComponentComponentIDOK creates a GetV1InstallsInstallIDComponentComponentIDOK with default headers values
func NewGetV1InstallsInstallIDComponentComponentIDOK() *GetV1InstallsInstallIDComponentComponentIDOK {
	return &GetV1InstallsInstallIDComponentComponentIDOK{}
}

/*
GetV1InstallsInstallIDComponentComponentIDOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallsInstallIDComponentComponentIDOK struct {
	Payload *models.AppInstallComponent
}

// IsSuccess returns true when this get v1 installs install Id component component Id o k response has a 2xx status code
func (o *GetV1InstallsInstallIDComponentComponentIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installs install Id component component Id o k response has a 3xx status code
func (o *GetV1InstallsInstallIDComponentComponentIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id component component Id o k response has a 4xx status code
func (o *GetV1InstallsInstallIDComponentComponentIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id component component Id o k response has a 5xx status code
func (o *GetV1InstallsInstallIDComponentComponentIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id component component Id o k response a status code equal to that given
func (o *GetV1InstallsInstallIDComponentComponentIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installs install Id component component Id o k response
func (o *GetV1InstallsInstallIDComponentComponentIDOK) Code() int {
	return 200
}

func (o *GetV1InstallsInstallIDComponentComponentIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/component/{component_id}][%d] getV1InstallsInstallIdComponentComponentIdOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDComponentComponentIDOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/component/{component_id}][%d] getV1InstallsInstallIdComponentComponentIdOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDComponentComponentIDOK) GetPayload() *models.AppInstallComponent {
	return o.Payload
}

func (o *GetV1InstallsInstallIDComponentComponentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallComponent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}