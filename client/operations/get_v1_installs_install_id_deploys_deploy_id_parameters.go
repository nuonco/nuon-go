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

// NewGetV1InstallsInstallIDDeploysDeployIDParams creates a new GetV1InstallsInstallIDDeploysDeployIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1InstallsInstallIDDeploysDeployIDParams() *GetV1InstallsInstallIDDeploysDeployIDParams {
	return &GetV1InstallsInstallIDDeploysDeployIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDParamsWithTimeout creates a new GetV1InstallsInstallIDDeploysDeployIDParams object
// with the ability to set a timeout on a request.
func NewGetV1InstallsInstallIDDeploysDeployIDParamsWithTimeout(timeout time.Duration) *GetV1InstallsInstallIDDeploysDeployIDParams {
	return &GetV1InstallsInstallIDDeploysDeployIDParams{
		timeout: timeout,
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDParamsWithContext creates a new GetV1InstallsInstallIDDeploysDeployIDParams object
// with the ability to set a context for a request.
func NewGetV1InstallsInstallIDDeploysDeployIDParamsWithContext(ctx context.Context) *GetV1InstallsInstallIDDeploysDeployIDParams {
	return &GetV1InstallsInstallIDDeploysDeployIDParams{
		Context: ctx,
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDParamsWithHTTPClient creates a new GetV1InstallsInstallIDDeploysDeployIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1InstallsInstallIDDeploysDeployIDParamsWithHTTPClient(client *http.Client) *GetV1InstallsInstallIDDeploysDeployIDParams {
	return &GetV1InstallsInstallIDDeploysDeployIDParams{
		HTTPClient: client,
	}
}

/*
GetV1InstallsInstallIDDeploysDeployIDParams contains all the parameters to send to the API endpoint

	for the get v1 installs install ID deploys deploy ID operation.

	Typically these are written to a http.Request.
*/
type GetV1InstallsInstallIDDeploysDeployIDParams struct {

	/* Authorization.

	   bearer auth token
	*/
	Authorization string

	/* XNuonOrgID.

	   org ID
	*/
	XNuonOrgID string

	/* DeployID.

	   deploy ID
	*/
	DeployID string

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 installs install ID deploys deploy ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithDefaults() *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 installs install ID deploys deploy ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithTimeout(timeout time.Duration) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithContext(ctx context.Context) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithHTTPClient(client *http.Client) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithAuthorization(authorization string) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithXNuonOrgID(xNuonOrgID string) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithDeployID adds the deployID to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithDeployID(deployID string) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WithInstallID adds the installID to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WithInstallID(installID string) *GetV1InstallsInstallIDDeploysDeployIDParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get v1 installs install ID deploys deploy ID params
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1InstallsInstallIDDeploysDeployIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
