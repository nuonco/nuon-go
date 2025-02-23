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

// NewGetInstallActionWorkflowsParams creates a new GetInstallActionWorkflowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallActionWorkflowsParams() *GetInstallActionWorkflowsParams {
	return &GetInstallActionWorkflowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallActionWorkflowsParamsWithTimeout creates a new GetInstallActionWorkflowsParams object
// with the ability to set a timeout on a request.
func NewGetInstallActionWorkflowsParamsWithTimeout(timeout time.Duration) *GetInstallActionWorkflowsParams {
	return &GetInstallActionWorkflowsParams{
		timeout: timeout,
	}
}

// NewGetInstallActionWorkflowsParamsWithContext creates a new GetInstallActionWorkflowsParams object
// with the ability to set a context for a request.
func NewGetInstallActionWorkflowsParamsWithContext(ctx context.Context) *GetInstallActionWorkflowsParams {
	return &GetInstallActionWorkflowsParams{
		Context: ctx,
	}
}

// NewGetInstallActionWorkflowsParamsWithHTTPClient creates a new GetInstallActionWorkflowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallActionWorkflowsParamsWithHTTPClient(client *http.Client) *GetInstallActionWorkflowsParams {
	return &GetInstallActionWorkflowsParams{
		HTTPClient: client,
	}
}

/*
GetInstallActionWorkflowsParams contains all the parameters to send to the API endpoint

	for the get install action workflows operation.

	Typically these are written to a http.Request.
*/
type GetInstallActionWorkflowsParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install action workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallActionWorkflowsParams) WithDefaults() *GetInstallActionWorkflowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install action workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallActionWorkflowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) WithTimeout(timeout time.Duration) *GetInstallActionWorkflowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) WithContext(ctx context.Context) *GetInstallActionWorkflowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) WithHTTPClient(client *http.Client) *GetInstallActionWorkflowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) WithInstallID(installID string) *GetInstallActionWorkflowsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get install action workflows params
func (o *GetInstallActionWorkflowsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallActionWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
