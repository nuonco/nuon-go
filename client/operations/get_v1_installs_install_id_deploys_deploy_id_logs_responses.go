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

// GetV1InstallsInstallIDDeploysDeployIDLogsReader is a Reader for the GetV1InstallsInstallIDDeploysDeployIDLogs structure.
type GetV1InstallsInstallIDDeploysDeployIDLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1InstallsInstallIDDeploysDeployIDLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/deploys/{deploy_id}/logs] GetV1InstallsInstallIDDeploysDeployIDLogs", response, response.Code())
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDLogsOK creates a GetV1InstallsInstallIDDeploysDeployIDLogsOK with default headers values
func NewGetV1InstallsInstallIDDeploysDeployIDLogsOK() *GetV1InstallsInstallIDDeploysDeployIDLogsOK {
	return &GetV1InstallsInstallIDDeploysDeployIDLogsOK{}
}

/*
GetV1InstallsInstallIDDeploysDeployIDLogsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1InstallsInstallIDDeploysDeployIDLogsOK struct {
	Payload []models.ServiceDeployLog
}

// IsSuccess returns true when this get v1 installs install Id deploys deploy Id logs o k response has a 2xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 installs install Id deploys deploy Id logs o k response has a 3xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 installs install Id deploys deploy Id logs o k response has a 4xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 installs install Id deploys deploy Id logs o k response has a 5xx status code
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 installs install Id deploys deploy Id logs o k response a status code equal to that given
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 installs install Id deploys deploy Id logs o k response
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) Code() int {
	return 200
}

func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}/logs][%d] getV1InstallsInstallIdDeploysDeployIdLogsOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/{deploy_id}/logs][%d] getV1InstallsInstallIdDeploysDeployIdLogsOK  %+v", 200, o.Payload)
}

func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) GetPayload() []models.ServiceDeployLog {
	return o.Payload
}

func (o *GetV1InstallsInstallIDDeploysDeployIDLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}