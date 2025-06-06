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

// NewGetAppCloudFormationStackConfigParams creates a new GetAppCloudFormationStackConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppCloudFormationStackConfigParams() *GetAppCloudFormationStackConfigParams {
	return &GetAppCloudFormationStackConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppCloudFormationStackConfigParamsWithTimeout creates a new GetAppCloudFormationStackConfigParams object
// with the ability to set a timeout on a request.
func NewGetAppCloudFormationStackConfigParamsWithTimeout(timeout time.Duration) *GetAppCloudFormationStackConfigParams {
	return &GetAppCloudFormationStackConfigParams{
		timeout: timeout,
	}
}

// NewGetAppCloudFormationStackConfigParamsWithContext creates a new GetAppCloudFormationStackConfigParams object
// with the ability to set a context for a request.
func NewGetAppCloudFormationStackConfigParamsWithContext(ctx context.Context) *GetAppCloudFormationStackConfigParams {
	return &GetAppCloudFormationStackConfigParams{
		Context: ctx,
	}
}

// NewGetAppCloudFormationStackConfigParamsWithHTTPClient creates a new GetAppCloudFormationStackConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppCloudFormationStackConfigParamsWithHTTPClient(client *http.Client) *GetAppCloudFormationStackConfigParams {
	return &GetAppCloudFormationStackConfigParams{
		HTTPClient: client,
	}
}

/*
GetAppCloudFormationStackConfigParams contains all the parameters to send to the API endpoint

	for the get app cloud formation stack config operation.

	Typically these are written to a http.Request.
*/
type GetAppCloudFormationStackConfigParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* ConfigID.

	   app cloudformation stack config ID
	*/
	ConfigID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app cloud formation stack config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppCloudFormationStackConfigParams) WithDefaults() *GetAppCloudFormationStackConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app cloud formation stack config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppCloudFormationStackConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) WithTimeout(timeout time.Duration) *GetAppCloudFormationStackConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) WithContext(ctx context.Context) *GetAppCloudFormationStackConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) WithHTTPClient(client *http.Client) *GetAppCloudFormationStackConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) WithAppID(appID string) *GetAppCloudFormationStackConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithConfigID adds the configID to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) WithConfigID(configID string) *GetAppCloudFormationStackConfigParams {
	o.SetConfigID(configID)
	return o
}

// SetConfigID adds the configId to the get app cloud formation stack config params
func (o *GetAppCloudFormationStackConfigParams) SetConfigID(configID string) {
	o.ConfigID = configID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppCloudFormationStackConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	// path param config_id
	if err := r.SetPathParam("config_id", o.ConfigID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
