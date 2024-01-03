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

// NewInstallerGetInstallParams creates a new InstallerGetInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstallerGetInstallParams() *InstallerGetInstallParams {
	return &InstallerGetInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstallerGetInstallParamsWithTimeout creates a new InstallerGetInstallParams object
// with the ability to set a timeout on a request.
func NewInstallerGetInstallParamsWithTimeout(timeout time.Duration) *InstallerGetInstallParams {
	return &InstallerGetInstallParams{
		timeout: timeout,
	}
}

// NewInstallerGetInstallParamsWithContext creates a new InstallerGetInstallParams object
// with the ability to set a context for a request.
func NewInstallerGetInstallParamsWithContext(ctx context.Context) *InstallerGetInstallParams {
	return &InstallerGetInstallParams{
		Context: ctx,
	}
}

// NewInstallerGetInstallParamsWithHTTPClient creates a new InstallerGetInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstallerGetInstallParamsWithHTTPClient(client *http.Client) *InstallerGetInstallParams {
	return &InstallerGetInstallParams{
		HTTPClient: client,
	}
}

/*
InstallerGetInstallParams contains all the parameters to send to the API endpoint

	for the installer get install operation.

	Typically these are written to a http.Request.
*/
type InstallerGetInstallParams struct {

	/* InstallID.

	   install id
	*/
	InstallID string

	/* InstallerSlug.

	   installer slug or ID
	*/
	InstallerSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the installer get install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallerGetInstallParams) WithDefaults() *InstallerGetInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the installer get install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallerGetInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the installer get install params
func (o *InstallerGetInstallParams) WithTimeout(timeout time.Duration) *InstallerGetInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the installer get install params
func (o *InstallerGetInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the installer get install params
func (o *InstallerGetInstallParams) WithContext(ctx context.Context) *InstallerGetInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the installer get install params
func (o *InstallerGetInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the installer get install params
func (o *InstallerGetInstallParams) WithHTTPClient(client *http.Client) *InstallerGetInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the installer get install params
func (o *InstallerGetInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the installer get install params
func (o *InstallerGetInstallParams) WithInstallID(installID string) *InstallerGetInstallParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the installer get install params
func (o *InstallerGetInstallParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithInstallerSlug adds the installerSlug to the installer get install params
func (o *InstallerGetInstallParams) WithInstallerSlug(installerSlug string) *InstallerGetInstallParams {
	o.SetInstallerSlug(installerSlug)
	return o
}

// SetInstallerSlug adds the installerSlug to the installer get install params
func (o *InstallerGetInstallParams) SetInstallerSlug(installerSlug string) {
	o.InstallerSlug = installerSlug
}

// WriteToRequest writes these params to a swagger request
func (o *InstallerGetInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}

	// path param installer_slug
	if err := r.SetPathParam("installer_slug", o.InstallerSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}