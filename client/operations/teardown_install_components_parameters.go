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

// NewTeardownInstallComponentsParams creates a new TeardownInstallComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTeardownInstallComponentsParams() *TeardownInstallComponentsParams {
	return &TeardownInstallComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTeardownInstallComponentsParamsWithTimeout creates a new TeardownInstallComponentsParams object
// with the ability to set a timeout on a request.
func NewTeardownInstallComponentsParamsWithTimeout(timeout time.Duration) *TeardownInstallComponentsParams {
	return &TeardownInstallComponentsParams{
		timeout: timeout,
	}
}

// NewTeardownInstallComponentsParamsWithContext creates a new TeardownInstallComponentsParams object
// with the ability to set a context for a request.
func NewTeardownInstallComponentsParamsWithContext(ctx context.Context) *TeardownInstallComponentsParams {
	return &TeardownInstallComponentsParams{
		Context: ctx,
	}
}

// NewTeardownInstallComponentsParamsWithHTTPClient creates a new TeardownInstallComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTeardownInstallComponentsParamsWithHTTPClient(client *http.Client) *TeardownInstallComponentsParams {
	return &TeardownInstallComponentsParams{
		HTTPClient: client,
	}
}

/*
TeardownInstallComponentsParams contains all the parameters to send to the API endpoint

	for the teardown install components operation.

	Typically these are written to a http.Request.
*/
type TeardownInstallComponentsParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Req.

	   Input
	*/
	Req *models.ServiceTeardownInstallComponentsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the teardown install components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeardownInstallComponentsParams) WithDefaults() *TeardownInstallComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the teardown install components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TeardownInstallComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the teardown install components params
func (o *TeardownInstallComponentsParams) WithTimeout(timeout time.Duration) *TeardownInstallComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teardown install components params
func (o *TeardownInstallComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teardown install components params
func (o *TeardownInstallComponentsParams) WithContext(ctx context.Context) *TeardownInstallComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teardown install components params
func (o *TeardownInstallComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teardown install components params
func (o *TeardownInstallComponentsParams) WithHTTPClient(client *http.Client) *TeardownInstallComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teardown install components params
func (o *TeardownInstallComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the teardown install components params
func (o *TeardownInstallComponentsParams) WithInstallID(installID string) *TeardownInstallComponentsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the teardown install components params
func (o *TeardownInstallComponentsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithReq adds the req to the teardown install components params
func (o *TeardownInstallComponentsParams) WithReq(req *models.ServiceTeardownInstallComponentsRequest) *TeardownInstallComponentsParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the teardown install components params
func (o *TeardownInstallComponentsParams) SetReq(req *models.ServiceTeardownInstallComponentsRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *TeardownInstallComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
