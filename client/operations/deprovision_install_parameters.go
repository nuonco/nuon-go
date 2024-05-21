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

// NewDeprovisionInstallParams creates a new DeprovisionInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeprovisionInstallParams() *DeprovisionInstallParams {
	return &DeprovisionInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeprovisionInstallParamsWithTimeout creates a new DeprovisionInstallParams object
// with the ability to set a timeout on a request.
func NewDeprovisionInstallParamsWithTimeout(timeout time.Duration) *DeprovisionInstallParams {
	return &DeprovisionInstallParams{
		timeout: timeout,
	}
}

// NewDeprovisionInstallParamsWithContext creates a new DeprovisionInstallParams object
// with the ability to set a context for a request.
func NewDeprovisionInstallParamsWithContext(ctx context.Context) *DeprovisionInstallParams {
	return &DeprovisionInstallParams{
		Context: ctx,
	}
}

// NewDeprovisionInstallParamsWithHTTPClient creates a new DeprovisionInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeprovisionInstallParamsWithHTTPClient(client *http.Client) *DeprovisionInstallParams {
	return &DeprovisionInstallParams{
		HTTPClient: client,
	}
}

/*
DeprovisionInstallParams contains all the parameters to send to the API endpoint

	for the deprovision install operation.

	Typically these are written to a http.Request.
*/
type DeprovisionInstallParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Req.

	   Input
	*/
	Req models.ServiceDeprovisionInstallRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deprovision install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeprovisionInstallParams) WithDefaults() *DeprovisionInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deprovision install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeprovisionInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deprovision install params
func (o *DeprovisionInstallParams) WithTimeout(timeout time.Duration) *DeprovisionInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deprovision install params
func (o *DeprovisionInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deprovision install params
func (o *DeprovisionInstallParams) WithContext(ctx context.Context) *DeprovisionInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deprovision install params
func (o *DeprovisionInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deprovision install params
func (o *DeprovisionInstallParams) WithHTTPClient(client *http.Client) *DeprovisionInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deprovision install params
func (o *DeprovisionInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the deprovision install params
func (o *DeprovisionInstallParams) WithInstallID(installID string) *DeprovisionInstallParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the deprovision install params
func (o *DeprovisionInstallParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithReq adds the req to the deprovision install params
func (o *DeprovisionInstallParams) WithReq(req models.ServiceDeprovisionInstallRequest) *DeprovisionInstallParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the deprovision install params
func (o *DeprovisionInstallParams) SetReq(req models.ServiceDeprovisionInstallRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *DeprovisionInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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