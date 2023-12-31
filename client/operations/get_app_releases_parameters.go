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

// NewGetAppReleasesParams creates a new GetAppReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppReleasesParams() *GetAppReleasesParams {
	return &GetAppReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppReleasesParamsWithTimeout creates a new GetAppReleasesParams object
// with the ability to set a timeout on a request.
func NewGetAppReleasesParamsWithTimeout(timeout time.Duration) *GetAppReleasesParams {
	return &GetAppReleasesParams{
		timeout: timeout,
	}
}

// NewGetAppReleasesParamsWithContext creates a new GetAppReleasesParams object
// with the ability to set a context for a request.
func NewGetAppReleasesParamsWithContext(ctx context.Context) *GetAppReleasesParams {
	return &GetAppReleasesParams{
		Context: ctx,
	}
}

// NewGetAppReleasesParamsWithHTTPClient creates a new GetAppReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppReleasesParamsWithHTTPClient(client *http.Client) *GetAppReleasesParams {
	return &GetAppReleasesParams{
		HTTPClient: client,
	}
}

/*
GetAppReleasesParams contains all the parameters to send to the API endpoint

	for the get app releases operation.

	Typically these are written to a http.Request.
*/
type GetAppReleasesParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppReleasesParams) WithDefaults() *GetAppReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app releases params
func (o *GetAppReleasesParams) WithTimeout(timeout time.Duration) *GetAppReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app releases params
func (o *GetAppReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app releases params
func (o *GetAppReleasesParams) WithContext(ctx context.Context) *GetAppReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app releases params
func (o *GetAppReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app releases params
func (o *GetAppReleasesParams) WithHTTPClient(client *http.Client) *GetAppReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app releases params
func (o *GetAppReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app releases params
func (o *GetAppReleasesParams) WithAppID(appID string) *GetAppReleasesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app releases params
func (o *GetAppReleasesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
