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

// NewGetInstallSandboxRunParams creates a new GetInstallSandboxRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallSandboxRunParams() *GetInstallSandboxRunParams {
	return &GetInstallSandboxRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallSandboxRunParamsWithTimeout creates a new GetInstallSandboxRunParams object
// with the ability to set a timeout on a request.
func NewGetInstallSandboxRunParamsWithTimeout(timeout time.Duration) *GetInstallSandboxRunParams {
	return &GetInstallSandboxRunParams{
		timeout: timeout,
	}
}

// NewGetInstallSandboxRunParamsWithContext creates a new GetInstallSandboxRunParams object
// with the ability to set a context for a request.
func NewGetInstallSandboxRunParamsWithContext(ctx context.Context) *GetInstallSandboxRunParams {
	return &GetInstallSandboxRunParams{
		Context: ctx,
	}
}

// NewGetInstallSandboxRunParamsWithHTTPClient creates a new GetInstallSandboxRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallSandboxRunParamsWithHTTPClient(client *http.Client) *GetInstallSandboxRunParams {
	return &GetInstallSandboxRunParams{
		HTTPClient: client,
	}
}

/*
GetInstallSandboxRunParams contains all the parameters to send to the API endpoint

	for the get install sandbox run operation.

	Typically these are written to a http.Request.
*/
type GetInstallSandboxRunParams struct {

	/* RunID.

	   run ID
	*/
	RunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install sandbox run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallSandboxRunParams) WithDefaults() *GetInstallSandboxRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install sandbox run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallSandboxRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get install sandbox run params
func (o *GetInstallSandboxRunParams) WithTimeout(timeout time.Duration) *GetInstallSandboxRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install sandbox run params
func (o *GetInstallSandboxRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install sandbox run params
func (o *GetInstallSandboxRunParams) WithContext(ctx context.Context) *GetInstallSandboxRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install sandbox run params
func (o *GetInstallSandboxRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install sandbox run params
func (o *GetInstallSandboxRunParams) WithHTTPClient(client *http.Client) *GetInstallSandboxRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install sandbox run params
func (o *GetInstallSandboxRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunID adds the runID to the get install sandbox run params
func (o *GetInstallSandboxRunParams) WithRunID(runID string) *GetInstallSandboxRunParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get install sandbox run params
func (o *GetInstallSandboxRunParams) SetRunID(runID string) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallSandboxRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param run_id
	if err := r.SetPathParam("run_id", o.RunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
