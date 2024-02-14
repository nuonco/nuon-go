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

// NewRenderInstallerInstallParams creates a new RenderInstallerInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenderInstallerInstallParams() *RenderInstallerInstallParams {
	return &RenderInstallerInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenderInstallerInstallParamsWithTimeout creates a new RenderInstallerInstallParams object
// with the ability to set a timeout on a request.
func NewRenderInstallerInstallParamsWithTimeout(timeout time.Duration) *RenderInstallerInstallParams {
	return &RenderInstallerInstallParams{
		timeout: timeout,
	}
}

// NewRenderInstallerInstallParamsWithContext creates a new RenderInstallerInstallParams object
// with the ability to set a context for a request.
func NewRenderInstallerInstallParamsWithContext(ctx context.Context) *RenderInstallerInstallParams {
	return &RenderInstallerInstallParams{
		Context: ctx,
	}
}

// NewRenderInstallerInstallParamsWithHTTPClient creates a new RenderInstallerInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenderInstallerInstallParamsWithHTTPClient(client *http.Client) *RenderInstallerInstallParams {
	return &RenderInstallerInstallParams{
		HTTPClient: client,
	}
}

/*
RenderInstallerInstallParams contains all the parameters to send to the API endpoint

	for the render installer install operation.

	Typically these are written to a http.Request.
*/
type RenderInstallerInstallParams struct {

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

// WithDefaults hydrates default values in the render installer install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenderInstallerInstallParams) WithDefaults() *RenderInstallerInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the render installer install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenderInstallerInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the render installer install params
func (o *RenderInstallerInstallParams) WithTimeout(timeout time.Duration) *RenderInstallerInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the render installer install params
func (o *RenderInstallerInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the render installer install params
func (o *RenderInstallerInstallParams) WithContext(ctx context.Context) *RenderInstallerInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the render installer install params
func (o *RenderInstallerInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the render installer install params
func (o *RenderInstallerInstallParams) WithHTTPClient(client *http.Client) *RenderInstallerInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the render installer install params
func (o *RenderInstallerInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the render installer install params
func (o *RenderInstallerInstallParams) WithInstallID(installID string) *RenderInstallerInstallParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the render installer install params
func (o *RenderInstallerInstallParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithInstallerSlug adds the installerSlug to the render installer install params
func (o *RenderInstallerInstallParams) WithInstallerSlug(installerSlug string) *RenderInstallerInstallParams {
	o.SetInstallerSlug(installerSlug)
	return o
}

// SetInstallerSlug adds the installerSlug to the render installer install params
func (o *RenderInstallerInstallParams) SetInstallerSlug(installerSlug string) {
	o.InstallerSlug = installerSlug
}

// WriteToRequest writes these params to a swagger request
func (o *RenderInstallerInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
