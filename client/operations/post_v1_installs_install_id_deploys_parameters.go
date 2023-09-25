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

// NewPostV1InstallsInstallIDDeploysParams creates a new PostV1InstallsInstallIDDeploysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1InstallsInstallIDDeploysParams() *PostV1InstallsInstallIDDeploysParams {
	return &PostV1InstallsInstallIDDeploysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1InstallsInstallIDDeploysParamsWithTimeout creates a new PostV1InstallsInstallIDDeploysParams object
// with the ability to set a timeout on a request.
func NewPostV1InstallsInstallIDDeploysParamsWithTimeout(timeout time.Duration) *PostV1InstallsInstallIDDeploysParams {
	return &PostV1InstallsInstallIDDeploysParams{
		timeout: timeout,
	}
}

// NewPostV1InstallsInstallIDDeploysParamsWithContext creates a new PostV1InstallsInstallIDDeploysParams object
// with the ability to set a context for a request.
func NewPostV1InstallsInstallIDDeploysParamsWithContext(ctx context.Context) *PostV1InstallsInstallIDDeploysParams {
	return &PostV1InstallsInstallIDDeploysParams{
		Context: ctx,
	}
}

// NewPostV1InstallsInstallIDDeploysParamsWithHTTPClient creates a new PostV1InstallsInstallIDDeploysParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1InstallsInstallIDDeploysParamsWithHTTPClient(client *http.Client) *PostV1InstallsInstallIDDeploysParams {
	return &PostV1InstallsInstallIDDeploysParams{
		HTTPClient: client,
	}
}

/*
PostV1InstallsInstallIDDeploysParams contains all the parameters to send to the API endpoint

	for the post v1 installs install ID deploys operation.

	Typically these are written to a http.Request.
*/
type PostV1InstallsInstallIDDeploysParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateInstallDeployRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 installs install ID deploys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1InstallsInstallIDDeploysParams) WithDefaults() *PostV1InstallsInstallIDDeploysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 installs install ID deploys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1InstallsInstallIDDeploysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) WithTimeout(timeout time.Duration) *PostV1InstallsInstallIDDeploysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) WithContext(ctx context.Context) *PostV1InstallsInstallIDDeploysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) WithHTTPClient(client *http.Client) *PostV1InstallsInstallIDDeploysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) WithInstallID(installID string) *PostV1InstallsInstallIDDeploysParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithReq adds the req to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) WithReq(req *models.ServiceCreateInstallDeployRequest) *PostV1InstallsInstallIDDeploysParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the post v1 installs install ID deploys params
func (o *PostV1InstallsInstallIDDeploysParams) SetReq(req *models.ServiceCreateInstallDeployRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1InstallsInstallIDDeploysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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