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
)

// NewGetV1InstallsInstallIDDeploysParams creates a new GetV1InstallsInstallIDDeploysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1InstallsInstallIDDeploysParams() *GetV1InstallsInstallIDDeploysParams {
	return &GetV1InstallsInstallIDDeploysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1InstallsInstallIDDeploysParamsWithTimeout creates a new GetV1InstallsInstallIDDeploysParams object
// with the ability to set a timeout on a request.
func NewGetV1InstallsInstallIDDeploysParamsWithTimeout(timeout time.Duration) *GetV1InstallsInstallIDDeploysParams {
	return &GetV1InstallsInstallIDDeploysParams{
		timeout: timeout,
	}
}

// NewGetV1InstallsInstallIDDeploysParamsWithContext creates a new GetV1InstallsInstallIDDeploysParams object
// with the ability to set a context for a request.
func NewGetV1InstallsInstallIDDeploysParamsWithContext(ctx context.Context) *GetV1InstallsInstallIDDeploysParams {
	return &GetV1InstallsInstallIDDeploysParams{
		Context: ctx,
	}
}

// NewGetV1InstallsInstallIDDeploysParamsWithHTTPClient creates a new GetV1InstallsInstallIDDeploysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1InstallsInstallIDDeploysParamsWithHTTPClient(client *http.Client) *GetV1InstallsInstallIDDeploysParams {
	return &GetV1InstallsInstallIDDeploysParams{
		HTTPClient: client,
	}
}

/*
GetV1InstallsInstallIDDeploysParams contains all the parameters to send to the API endpoint

	for the get v1 installs install ID deploys operation.

	Typically these are written to a http.Request.
*/
type GetV1InstallsInstallIDDeploysParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 installs install ID deploys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDDeploysParams) WithDefaults() *GetV1InstallsInstallIDDeploysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 installs install ID deploys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDDeploysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) WithTimeout(timeout time.Duration) *GetV1InstallsInstallIDDeploysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) WithContext(ctx context.Context) *GetV1InstallsInstallIDDeploysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) WithHTTPClient(client *http.Client) *GetV1InstallsInstallIDDeploysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) WithInstallID(installID string) *GetV1InstallsInstallIDDeploysParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get v1 installs install ID deploys params
func (o *GetV1InstallsInstallIDDeploysParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1InstallsInstallIDDeploysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
