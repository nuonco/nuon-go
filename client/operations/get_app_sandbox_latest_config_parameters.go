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

// NewGetAppSandboxLatestConfigParams creates a new GetAppSandboxLatestConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppSandboxLatestConfigParams() *GetAppSandboxLatestConfigParams {
	return &GetAppSandboxLatestConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppSandboxLatestConfigParamsWithTimeout creates a new GetAppSandboxLatestConfigParams object
// with the ability to set a timeout on a request.
func NewGetAppSandboxLatestConfigParamsWithTimeout(timeout time.Duration) *GetAppSandboxLatestConfigParams {
	return &GetAppSandboxLatestConfigParams{
		timeout: timeout,
	}
}

// NewGetAppSandboxLatestConfigParamsWithContext creates a new GetAppSandboxLatestConfigParams object
// with the ability to set a context for a request.
func NewGetAppSandboxLatestConfigParamsWithContext(ctx context.Context) *GetAppSandboxLatestConfigParams {
	return &GetAppSandboxLatestConfigParams{
		Context: ctx,
	}
}

// NewGetAppSandboxLatestConfigParamsWithHTTPClient creates a new GetAppSandboxLatestConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppSandboxLatestConfigParamsWithHTTPClient(client *http.Client) *GetAppSandboxLatestConfigParams {
	return &GetAppSandboxLatestConfigParams{
		HTTPClient: client,
	}
}

/*
GetAppSandboxLatestConfigParams contains all the parameters to send to the API endpoint

	for the get app sandbox latest config operation.

	Typically these are written to a http.Request.
*/
type GetAppSandboxLatestConfigParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app sandbox latest config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppSandboxLatestConfigParams) WithDefaults() *GetAppSandboxLatestConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app sandbox latest config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppSandboxLatestConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) WithTimeout(timeout time.Duration) *GetAppSandboxLatestConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) WithContext(ctx context.Context) *GetAppSandboxLatestConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) WithHTTPClient(client *http.Client) *GetAppSandboxLatestConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) WithAppID(appID string) *GetAppSandboxLatestConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app sandbox latest config params
func (o *GetAppSandboxLatestConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppSandboxLatestConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}