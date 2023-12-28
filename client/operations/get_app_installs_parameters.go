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

// NewGetAppInstallsParams creates a new GetAppInstallsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppInstallsParams() *GetAppInstallsParams {
	return &GetAppInstallsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppInstallsParamsWithTimeout creates a new GetAppInstallsParams object
// with the ability to set a timeout on a request.
func NewGetAppInstallsParamsWithTimeout(timeout time.Duration) *GetAppInstallsParams {
	return &GetAppInstallsParams{
		timeout: timeout,
	}
}

// NewGetAppInstallsParamsWithContext creates a new GetAppInstallsParams object
// with the ability to set a context for a request.
func NewGetAppInstallsParamsWithContext(ctx context.Context) *GetAppInstallsParams {
	return &GetAppInstallsParams{
		Context: ctx,
	}
}

// NewGetAppInstallsParamsWithHTTPClient creates a new GetAppInstallsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppInstallsParamsWithHTTPClient(client *http.Client) *GetAppInstallsParams {
	return &GetAppInstallsParams{
		HTTPClient: client,
	}
}

/*
GetAppInstallsParams contains all the parameters to send to the API endpoint

	for the get app installs operation.

	Typically these are written to a http.Request.
*/
type GetAppInstallsParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app installs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppInstallsParams) WithDefaults() *GetAppInstallsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app installs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppInstallsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app installs params
func (o *GetAppInstallsParams) WithTimeout(timeout time.Duration) *GetAppInstallsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app installs params
func (o *GetAppInstallsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app installs params
func (o *GetAppInstallsParams) WithContext(ctx context.Context) *GetAppInstallsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app installs params
func (o *GetAppInstallsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app installs params
func (o *GetAppInstallsParams) WithHTTPClient(client *http.Client) *GetAppInstallsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app installs params
func (o *GetAppInstallsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app installs params
func (o *GetAppInstallsParams) WithAppID(appID string) *GetAppInstallsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app installs params
func (o *GetAppInstallsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppInstallsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
