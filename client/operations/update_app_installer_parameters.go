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

// NewUpdateAppInstallerParams creates a new UpdateAppInstallerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAppInstallerParams() *UpdateAppInstallerParams {
	return &UpdateAppInstallerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppInstallerParamsWithTimeout creates a new UpdateAppInstallerParams object
// with the ability to set a timeout on a request.
func NewUpdateAppInstallerParamsWithTimeout(timeout time.Duration) *UpdateAppInstallerParams {
	return &UpdateAppInstallerParams{
		timeout: timeout,
	}
}

// NewUpdateAppInstallerParamsWithContext creates a new UpdateAppInstallerParams object
// with the ability to set a context for a request.
func NewUpdateAppInstallerParamsWithContext(ctx context.Context) *UpdateAppInstallerParams {
	return &UpdateAppInstallerParams{
		Context: ctx,
	}
}

// NewUpdateAppInstallerParamsWithHTTPClient creates a new UpdateAppInstallerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAppInstallerParamsWithHTTPClient(client *http.Client) *UpdateAppInstallerParams {
	return &UpdateAppInstallerParams{
		HTTPClient: client,
	}
}

/*
UpdateAppInstallerParams contains all the parameters to send to the API endpoint

	for the update app installer operation.

	Typically these are written to a http.Request.
*/
type UpdateAppInstallerParams struct {

	/* InstallerID.

	   installer ID
	*/
	InstallerID string

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateAppInstallerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update app installer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppInstallerParams) WithDefaults() *UpdateAppInstallerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update app installer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppInstallerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update app installer params
func (o *UpdateAppInstallerParams) WithTimeout(timeout time.Duration) *UpdateAppInstallerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app installer params
func (o *UpdateAppInstallerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app installer params
func (o *UpdateAppInstallerParams) WithContext(ctx context.Context) *UpdateAppInstallerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app installer params
func (o *UpdateAppInstallerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app installer params
func (o *UpdateAppInstallerParams) WithHTTPClient(client *http.Client) *UpdateAppInstallerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app installer params
func (o *UpdateAppInstallerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallerID adds the installerID to the update app installer params
func (o *UpdateAppInstallerParams) WithInstallerID(installerID string) *UpdateAppInstallerParams {
	o.SetInstallerID(installerID)
	return o
}

// SetInstallerID adds the installerId to the update app installer params
func (o *UpdateAppInstallerParams) SetInstallerID(installerID string) {
	o.InstallerID = installerID
}

// WithReq adds the req to the update app installer params
func (o *UpdateAppInstallerParams) WithReq(req *models.ServiceUpdateAppInstallerRequest) *UpdateAppInstallerParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the update app installer params
func (o *UpdateAppInstallerParams) SetReq(req *models.ServiceUpdateAppInstallerRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppInstallerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
