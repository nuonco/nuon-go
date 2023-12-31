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

// NewGetInstallComponentsParams creates a new GetInstallComponentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallComponentsParams() *GetInstallComponentsParams {
	return &GetInstallComponentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallComponentsParamsWithTimeout creates a new GetInstallComponentsParams object
// with the ability to set a timeout on a request.
func NewGetInstallComponentsParamsWithTimeout(timeout time.Duration) *GetInstallComponentsParams {
	return &GetInstallComponentsParams{
		timeout: timeout,
	}
}

// NewGetInstallComponentsParamsWithContext creates a new GetInstallComponentsParams object
// with the ability to set a context for a request.
func NewGetInstallComponentsParamsWithContext(ctx context.Context) *GetInstallComponentsParams {
	return &GetInstallComponentsParams{
		Context: ctx,
	}
}

// NewGetInstallComponentsParamsWithHTTPClient creates a new GetInstallComponentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallComponentsParamsWithHTTPClient(client *http.Client) *GetInstallComponentsParams {
	return &GetInstallComponentsParams{
		HTTPClient: client,
	}
}

/*
GetInstallComponentsParams contains all the parameters to send to the API endpoint

	for the get install components operation.

	Typically these are written to a http.Request.
*/
type GetInstallComponentsParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallComponentsParams) WithDefaults() *GetInstallComponentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install components params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallComponentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get install components params
func (o *GetInstallComponentsParams) WithTimeout(timeout time.Duration) *GetInstallComponentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install components params
func (o *GetInstallComponentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install components params
func (o *GetInstallComponentsParams) WithContext(ctx context.Context) *GetInstallComponentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install components params
func (o *GetInstallComponentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install components params
func (o *GetInstallComponentsParams) WithHTTPClient(client *http.Client) *GetInstallComponentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install components params
func (o *GetInstallComponentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the get install components params
func (o *GetInstallComponentsParams) WithInstallID(installID string) *GetInstallComponentsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get install components params
func (o *GetInstallComponentsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallComponentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
