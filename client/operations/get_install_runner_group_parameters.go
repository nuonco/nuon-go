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

// NewGetInstallRunnerGroupParams creates a new GetInstallRunnerGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallRunnerGroupParams() *GetInstallRunnerGroupParams {
	return &GetInstallRunnerGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallRunnerGroupParamsWithTimeout creates a new GetInstallRunnerGroupParams object
// with the ability to set a timeout on a request.
func NewGetInstallRunnerGroupParamsWithTimeout(timeout time.Duration) *GetInstallRunnerGroupParams {
	return &GetInstallRunnerGroupParams{
		timeout: timeout,
	}
}

// NewGetInstallRunnerGroupParamsWithContext creates a new GetInstallRunnerGroupParams object
// with the ability to set a context for a request.
func NewGetInstallRunnerGroupParamsWithContext(ctx context.Context) *GetInstallRunnerGroupParams {
	return &GetInstallRunnerGroupParams{
		Context: ctx,
	}
}

// NewGetInstallRunnerGroupParamsWithHTTPClient creates a new GetInstallRunnerGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallRunnerGroupParamsWithHTTPClient(client *http.Client) *GetInstallRunnerGroupParams {
	return &GetInstallRunnerGroupParams{
		HTTPClient: client,
	}
}

/*
GetInstallRunnerGroupParams contains all the parameters to send to the API endpoint

	for the get install runner group operation.

	Typically these are written to a http.Request.
*/
type GetInstallRunnerGroupParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install runner group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallRunnerGroupParams) WithDefaults() *GetInstallRunnerGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install runner group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallRunnerGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get install runner group params
func (o *GetInstallRunnerGroupParams) WithTimeout(timeout time.Duration) *GetInstallRunnerGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install runner group params
func (o *GetInstallRunnerGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install runner group params
func (o *GetInstallRunnerGroupParams) WithContext(ctx context.Context) *GetInstallRunnerGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install runner group params
func (o *GetInstallRunnerGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install runner group params
func (o *GetInstallRunnerGroupParams) WithHTTPClient(client *http.Client) *GetInstallRunnerGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install runner group params
func (o *GetInstallRunnerGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the get install runner group params
func (o *GetInstallRunnerGroupParams) WithInstallID(installID string) *GetInstallRunnerGroupParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get install runner group params
func (o *GetInstallRunnerGroupParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallRunnerGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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