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

// NewGetV1InstallsInstallIDDeploysDeployIDLogsParams creates a new GetV1InstallsInstallIDDeploysDeployIDLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1InstallsInstallIDDeploysDeployIDLogsParams() *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	return &GetV1InstallsInstallIDDeploysDeployIDLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDLogsParamsWithTimeout creates a new GetV1InstallsInstallIDDeploysDeployIDLogsParams object
// with the ability to set a timeout on a request.
func NewGetV1InstallsInstallIDDeploysDeployIDLogsParamsWithTimeout(timeout time.Duration) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	return &GetV1InstallsInstallIDDeploysDeployIDLogsParams{
		timeout: timeout,
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDLogsParamsWithContext creates a new GetV1InstallsInstallIDDeploysDeployIDLogsParams object
// with the ability to set a context for a request.
func NewGetV1InstallsInstallIDDeploysDeployIDLogsParamsWithContext(ctx context.Context) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	return &GetV1InstallsInstallIDDeploysDeployIDLogsParams{
		Context: ctx,
	}
}

// NewGetV1InstallsInstallIDDeploysDeployIDLogsParamsWithHTTPClient creates a new GetV1InstallsInstallIDDeploysDeployIDLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1InstallsInstallIDDeploysDeployIDLogsParamsWithHTTPClient(client *http.Client) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	return &GetV1InstallsInstallIDDeploysDeployIDLogsParams{
		HTTPClient: client,
	}
}

/*
GetV1InstallsInstallIDDeploysDeployIDLogsParams contains all the parameters to send to the API endpoint

	for the get v1 installs install ID deploys deploy ID logs operation.

	Typically these are written to a http.Request.
*/
type GetV1InstallsInstallIDDeploysDeployIDLogsParams struct {

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

// WithDefaults hydrates default values in the get v1 installs install ID deploys deploy ID logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithDefaults() *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 installs install ID deploys deploy ID logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithTimeout(timeout time.Duration) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithContext(ctx context.Context) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithHTTPClient(client *http.Client) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithAuthorization(authorization string) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithXNuonOrgID(xNuonOrgID string) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithDeployID adds the deployID to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithDeployID(deployID string) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WithInstallID adds the installID to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WithInstallID(installID string) *GetV1InstallsInstallIDDeploysDeployIDLogsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get v1 installs install ID deploys deploy ID logs params
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1InstallsInstallIDDeploysDeployIDLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
