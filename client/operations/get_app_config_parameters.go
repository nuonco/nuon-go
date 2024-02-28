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

// NewGetAppConfigParams creates a new GetAppConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppConfigParams() *GetAppConfigParams {
	return &GetAppConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppConfigParamsWithTimeout creates a new GetAppConfigParams object
// with the ability to set a timeout on a request.
func NewGetAppConfigParamsWithTimeout(timeout time.Duration) *GetAppConfigParams {
	return &GetAppConfigParams{
		timeout: timeout,
	}
}

// NewGetAppConfigParamsWithContext creates a new GetAppConfigParams object
// with the ability to set a context for a request.
func NewGetAppConfigParamsWithContext(ctx context.Context) *GetAppConfigParams {
	return &GetAppConfigParams{
		Context: ctx,
	}
}

// NewGetAppConfigParamsWithHTTPClient creates a new GetAppConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppConfigParamsWithHTTPClient(client *http.Client) *GetAppConfigParams {
	return &GetAppConfigParams{
		HTTPClient: client,
	}
}

/*
GetAppConfigParams contains all the parameters to send to the API endpoint

	for the get app config operation.

	Typically these are written to a http.Request.
*/
type GetAppConfigParams struct {

	/* AppConfigID.

	   app config ID
	*/
	AppConfigID string

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppConfigParams) WithDefaults() *GetAppConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app config params
func (o *GetAppConfigParams) WithTimeout(timeout time.Duration) *GetAppConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app config params
func (o *GetAppConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app config params
func (o *GetAppConfigParams) WithContext(ctx context.Context) *GetAppConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app config params
func (o *GetAppConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app config params
func (o *GetAppConfigParams) WithHTTPClient(client *http.Client) *GetAppConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app config params
func (o *GetAppConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppConfigID adds the appConfigID to the get app config params
func (o *GetAppConfigParams) WithAppConfigID(appConfigID string) *GetAppConfigParams {
	o.SetAppConfigID(appConfigID)
	return o
}

// SetAppConfigID adds the appConfigId to the get app config params
func (o *GetAppConfigParams) SetAppConfigID(appConfigID string) {
	o.AppConfigID = appConfigID
}

// WithAppID adds the appID to the get app config params
func (o *GetAppConfigParams) WithAppID(appID string) *GetAppConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app config params
func (o *GetAppConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_config_id
	if err := r.SetPathParam("app_config_id", o.AppConfigID); err != nil {
		return err
	}

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}