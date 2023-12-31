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

// NewGetAppInputConfigsParams creates a new GetAppInputConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppInputConfigsParams() *GetAppInputConfigsParams {
	return &GetAppInputConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppInputConfigsParamsWithTimeout creates a new GetAppInputConfigsParams object
// with the ability to set a timeout on a request.
func NewGetAppInputConfigsParamsWithTimeout(timeout time.Duration) *GetAppInputConfigsParams {
	return &GetAppInputConfigsParams{
		timeout: timeout,
	}
}

// NewGetAppInputConfigsParamsWithContext creates a new GetAppInputConfigsParams object
// with the ability to set a context for a request.
func NewGetAppInputConfigsParamsWithContext(ctx context.Context) *GetAppInputConfigsParams {
	return &GetAppInputConfigsParams{
		Context: ctx,
	}
}

// NewGetAppInputConfigsParamsWithHTTPClient creates a new GetAppInputConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppInputConfigsParamsWithHTTPClient(client *http.Client) *GetAppInputConfigsParams {
	return &GetAppInputConfigsParams{
		HTTPClient: client,
	}
}

/*
GetAppInputConfigsParams contains all the parameters to send to the API endpoint

	for the get app input configs operation.

	Typically these are written to a http.Request.
*/
type GetAppInputConfigsParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app input configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppInputConfigsParams) WithDefaults() *GetAppInputConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app input configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppInputConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app input configs params
func (o *GetAppInputConfigsParams) WithTimeout(timeout time.Duration) *GetAppInputConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app input configs params
func (o *GetAppInputConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app input configs params
func (o *GetAppInputConfigsParams) WithContext(ctx context.Context) *GetAppInputConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app input configs params
func (o *GetAppInputConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app input configs params
func (o *GetAppInputConfigsParams) WithHTTPClient(client *http.Client) *GetAppInputConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app input configs params
func (o *GetAppInputConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app input configs params
func (o *GetAppInputConfigsParams) WithAppID(appID string) *GetAppInputConfigsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app input configs params
func (o *GetAppInputConfigsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppInputConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
