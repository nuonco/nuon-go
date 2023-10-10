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

// NewGetV1InstallersInstallerIDParams creates a new GetV1InstallersInstallerIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1InstallersInstallerIDParams() *GetV1InstallersInstallerIDParams {
	return &GetV1InstallersInstallerIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1InstallersInstallerIDParamsWithTimeout creates a new GetV1InstallersInstallerIDParams object
// with the ability to set a timeout on a request.
func NewGetV1InstallersInstallerIDParamsWithTimeout(timeout time.Duration) *GetV1InstallersInstallerIDParams {
	return &GetV1InstallersInstallerIDParams{
		timeout: timeout,
	}
}

// NewGetV1InstallersInstallerIDParamsWithContext creates a new GetV1InstallersInstallerIDParams object
// with the ability to set a context for a request.
func NewGetV1InstallersInstallerIDParamsWithContext(ctx context.Context) *GetV1InstallersInstallerIDParams {
	return &GetV1InstallersInstallerIDParams{
		Context: ctx,
	}
}

// NewGetV1InstallersInstallerIDParamsWithHTTPClient creates a new GetV1InstallersInstallerIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1InstallersInstallerIDParamsWithHTTPClient(client *http.Client) *GetV1InstallersInstallerIDParams {
	return &GetV1InstallersInstallerIDParams{
		HTTPClient: client,
	}
}

/*
GetV1InstallersInstallerIDParams contains all the parameters to send to the API endpoint

	for the get v1 installers installer ID operation.

	Typically these are written to a http.Request.
*/
type GetV1InstallersInstallerIDParams struct {

	/* Authorization.

	   bearer auth token
	*/
	Authorization string

	/* XNuonOrgID.

	   org ID
	*/
	XNuonOrgID string

	/* InstallerID.

	   installer ID
	*/
	InstallerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 installers installer ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallersInstallerIDParams) WithDefaults() *GetV1InstallersInstallerIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 installers installer ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallersInstallerIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) WithTimeout(timeout time.Duration) *GetV1InstallersInstallerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) WithContext(ctx context.Context) *GetV1InstallersInstallerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) WithHTTPClient(client *http.Client) *GetV1InstallersInstallerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) WithAuthorization(authorization string) *GetV1InstallersInstallerIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) WithXNuonOrgID(xNuonOrgID string) *GetV1InstallersInstallerIDParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithInstallerID adds the installerID to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) WithInstallerID(installerID string) *GetV1InstallersInstallerIDParams {
	o.SetInstallerID(installerID)
	return o
}

// SetInstallerID adds the installerId to the get v1 installers installer ID params
func (o *GetV1InstallersInstallerIDParams) SetInstallerID(installerID string) {
	o.InstallerID = installerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1InstallersInstallerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// header param X-Nuon-Org-ID
	if err := r.SetHeaderParam("X-Nuon-Org-ID", o.XNuonOrgID); err != nil {
		return err
	}

	// path param installer_id
	if err := r.SetPathParam("installer_id", o.InstallerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
