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

// NewUpdateInstallerParams creates a new UpdateInstallerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateInstallerParams() *UpdateInstallerParams {
	return &UpdateInstallerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInstallerParamsWithTimeout creates a new UpdateInstallerParams object
// with the ability to set a timeout on a request.
func NewUpdateInstallerParamsWithTimeout(timeout time.Duration) *UpdateInstallerParams {
	return &UpdateInstallerParams{
		timeout: timeout,
	}
}

// NewUpdateInstallerParamsWithContext creates a new UpdateInstallerParams object
// with the ability to set a context for a request.
func NewUpdateInstallerParamsWithContext(ctx context.Context) *UpdateInstallerParams {
	return &UpdateInstallerParams{
		Context: ctx,
	}
}

// NewUpdateInstallerParamsWithHTTPClient creates a new UpdateInstallerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateInstallerParamsWithHTTPClient(client *http.Client) *UpdateInstallerParams {
	return &UpdateInstallerParams{
		HTTPClient: client,
	}
}

/*
UpdateInstallerParams contains all the parameters to send to the API endpoint

	for the update installer operation.

	Typically these are written to a http.Request.
*/
type UpdateInstallerParams struct {

	/* InstallerID.

	   installer ID
	*/
	InstallerID string

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateInstallerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update installer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInstallerParams) WithDefaults() *UpdateInstallerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update installer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInstallerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update installer params
func (o *UpdateInstallerParams) WithTimeout(timeout time.Duration) *UpdateInstallerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update installer params
func (o *UpdateInstallerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update installer params
func (o *UpdateInstallerParams) WithContext(ctx context.Context) *UpdateInstallerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update installer params
func (o *UpdateInstallerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update installer params
func (o *UpdateInstallerParams) WithHTTPClient(client *http.Client) *UpdateInstallerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update installer params
func (o *UpdateInstallerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallerID adds the installerID to the update installer params
func (o *UpdateInstallerParams) WithInstallerID(installerID string) *UpdateInstallerParams {
	o.SetInstallerID(installerID)
	return o
}

// SetInstallerID adds the installerId to the update installer params
func (o *UpdateInstallerParams) SetInstallerID(installerID string) {
	o.InstallerID = installerID
}

// WithReq adds the req to the update installer params
func (o *UpdateInstallerParams) WithReq(req *models.ServiceUpdateInstallerRequest) *UpdateInstallerParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the update installer params
func (o *UpdateInstallerParams) SetReq(req *models.ServiceUpdateInstallerRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInstallerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param installer_id
	if err := r.SetPathParam("installer_id", o.InstallerID); err != nil {
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