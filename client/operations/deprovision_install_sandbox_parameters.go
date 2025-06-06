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

// NewDeprovisionInstallSandboxParams creates a new DeprovisionInstallSandboxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeprovisionInstallSandboxParams() *DeprovisionInstallSandboxParams {
	return &DeprovisionInstallSandboxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeprovisionInstallSandboxParamsWithTimeout creates a new DeprovisionInstallSandboxParams object
// with the ability to set a timeout on a request.
func NewDeprovisionInstallSandboxParamsWithTimeout(timeout time.Duration) *DeprovisionInstallSandboxParams {
	return &DeprovisionInstallSandboxParams{
		timeout: timeout,
	}
}

// NewDeprovisionInstallSandboxParamsWithContext creates a new DeprovisionInstallSandboxParams object
// with the ability to set a context for a request.
func NewDeprovisionInstallSandboxParamsWithContext(ctx context.Context) *DeprovisionInstallSandboxParams {
	return &DeprovisionInstallSandboxParams{
		Context: ctx,
	}
}

// NewDeprovisionInstallSandboxParamsWithHTTPClient creates a new DeprovisionInstallSandboxParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeprovisionInstallSandboxParamsWithHTTPClient(client *http.Client) *DeprovisionInstallSandboxParams {
	return &DeprovisionInstallSandboxParams{
		HTTPClient: client,
	}
}

/*
DeprovisionInstallSandboxParams contains all the parameters to send to the API endpoint

	for the deprovision install sandbox operation.

	Typically these are written to a http.Request.
*/
type DeprovisionInstallSandboxParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Req.

	   Input
	*/
	Req *models.ServiceDeprovisionInstallSandboxRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deprovision install sandbox params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeprovisionInstallSandboxParams) WithDefaults() *DeprovisionInstallSandboxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deprovision install sandbox params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeprovisionInstallSandboxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) WithTimeout(timeout time.Duration) *DeprovisionInstallSandboxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) WithContext(ctx context.Context) *DeprovisionInstallSandboxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) WithHTTPClient(client *http.Client) *DeprovisionInstallSandboxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) WithInstallID(installID string) *DeprovisionInstallSandboxParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithReq adds the req to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) WithReq(req *models.ServiceDeprovisionInstallSandboxRequest) *DeprovisionInstallSandboxParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the deprovision install sandbox params
func (o *DeprovisionInstallSandboxParams) SetReq(req *models.ServiceDeprovisionInstallSandboxRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *DeprovisionInstallSandboxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
