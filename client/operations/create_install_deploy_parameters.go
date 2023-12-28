// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-go/models"
)

// NewCreateInstallDeployParams creates a new CreateInstallDeployParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInstallDeployParams() *CreateInstallDeployParams {
	return &CreateInstallDeployParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstallDeployParamsWithTimeout creates a new CreateInstallDeployParams object
// with the ability to set a timeout on a request.
func NewCreateInstallDeployParamsWithTimeout(timeout time.Duration) *CreateInstallDeployParams {
	return &CreateInstallDeployParams{
		timeout: timeout,
	}
}

// NewCreateInstallDeployParamsWithContext creates a new CreateInstallDeployParams object
// with the ability to set a context for a request.
func NewCreateInstallDeployParamsWithContext(ctx context.Context) *CreateInstallDeployParams {
	return &CreateInstallDeployParams{
		Context: ctx,
	}
}

// NewCreateInstallDeployParamsWithHTTPClient creates a new CreateInstallDeployParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInstallDeployParamsWithHTTPClient(client *http.Client) *CreateInstallDeployParams {
	return &CreateInstallDeployParams{
		HTTPClient: client,
	}
}

/*
CreateInstallDeployParams contains all the parameters to send to the API endpoint

	for the create install deploy operation.

	Typically these are written to a http.Request.
*/
type CreateInstallDeployParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateInstallDeployRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create install deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstallDeployParams) WithDefaults() *CreateInstallDeployParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create install deploy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstallDeployParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create install deploy params
func (o *CreateInstallDeployParams) WithTimeout(timeout time.Duration) *CreateInstallDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create install deploy params
func (o *CreateInstallDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create install deploy params
func (o *CreateInstallDeployParams) WithContext(ctx context.Context) *CreateInstallDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create install deploy params
func (o *CreateInstallDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create install deploy params
func (o *CreateInstallDeployParams) WithHTTPClient(client *http.Client) *CreateInstallDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create install deploy params
func (o *CreateInstallDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the create install deploy params
func (o *CreateInstallDeployParams) WithInstallID(installID string) *CreateInstallDeployParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the create install deploy params
func (o *CreateInstallDeployParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithReq adds the req to the create install deploy params
func (o *CreateInstallDeployParams) WithReq(req *models.ServiceCreateInstallDeployRequest) *CreateInstallDeployParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create install deploy params
func (o *CreateInstallDeployParams) SetReq(req *models.ServiceCreateInstallDeployRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstallDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
