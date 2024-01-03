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

// NewTeardownInstallComponentParams creates a new TeardownInstallComponentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeardownInstallComponentParams() *TeardownInstallComponentParams {
	return &TeardownInstallComponentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeardownInstallComponentParamsWithTimeout creates a new TeardownInstallComponentParams object
// with the ability to set a timeout on a request.
func NewTeardownInstallComponentParamsWithTimeout(timeout time.Duration) *TeardownInstallComponentParams {
	return &TeardownInstallComponentParams{
		timeout: timeout,
	}
}

// NewTeardownInstallComponentParamsWithContext creates a new TeardownInstallComponentParams object
// with the ability to set a context for a request.
func NewTeardownInstallComponentParamsWithContext(ctx context.Context) *TeardownInstallComponentParams {
	return &TeardownInstallComponentParams{
		Context: ctx,
	}
}

// NewTeardownInstallComponentParamsWithHTTPClient creates a new TeardownInstallComponentParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeardownInstallComponentParamsWithHTTPClient(client *http.Client) *TeardownInstallComponentParams {
	return &TeardownInstallComponentParams{
		HTTPClient: client,
	}
}

/*
TeardownInstallComponentParams contains all the parameters to send to the API endpoint

	for the teardown install component operation.

	Typically these are written to a http.Request.
*/
type TeardownInstallComponentParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Req.

	   Input
	*/
	Req models.ServiceTeardownInstallComponentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the teardown install component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeardownInstallComponentParams) WithDefaults() *TeardownInstallComponentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the teardown install component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeardownInstallComponentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the teardown install component params
func (o *TeardownInstallComponentParams) WithTimeout(timeout time.Duration) *TeardownInstallComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teardown install component params
func (o *TeardownInstallComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teardown install component params
func (o *TeardownInstallComponentParams) WithContext(ctx context.Context) *TeardownInstallComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teardown install component params
func (o *TeardownInstallComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teardown install component params
func (o *TeardownInstallComponentParams) WithHTTPClient(client *http.Client) *TeardownInstallComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teardown install component params
func (o *TeardownInstallComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the teardown install component params
func (o *TeardownInstallComponentParams) WithComponentID(componentID string) *TeardownInstallComponentParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the teardown install component params
func (o *TeardownInstallComponentParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithInstallID adds the installID to the teardown install component params
func (o *TeardownInstallComponentParams) WithInstallID(installID string) *TeardownInstallComponentParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the teardown install component params
func (o *TeardownInstallComponentParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithReq adds the req to the teardown install component params
func (o *TeardownInstallComponentParams) WithReq(req models.ServiceTeardownInstallComponentRequest) *TeardownInstallComponentParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the teardown install component params
func (o *TeardownInstallComponentParams) SetReq(req models.ServiceTeardownInstallComponentRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *TeardownInstallComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}

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