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

// NewGetV1AppsAppIDParams creates a new GetV1AppsAppIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AppsAppIDParams() *GetV1AppsAppIDParams {
	return &GetV1AppsAppIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AppsAppIDParamsWithTimeout creates a new GetV1AppsAppIDParams object
// with the ability to set a timeout on a request.
func NewGetV1AppsAppIDParamsWithTimeout(timeout time.Duration) *GetV1AppsAppIDParams {
	return &GetV1AppsAppIDParams{
		timeout: timeout,
	}
}

// NewGetV1AppsAppIDParamsWithContext creates a new GetV1AppsAppIDParams object
// with the ability to set a context for a request.
func NewGetV1AppsAppIDParamsWithContext(ctx context.Context) *GetV1AppsAppIDParams {
	return &GetV1AppsAppIDParams{
		Context: ctx,
	}
}

// NewGetV1AppsAppIDParamsWithHTTPClient creates a new GetV1AppsAppIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AppsAppIDParamsWithHTTPClient(client *http.Client) *GetV1AppsAppIDParams {
	return &GetV1AppsAppIDParams{
		HTTPClient: client,
	}
}

/*
GetV1AppsAppIDParams contains all the parameters to send to the API endpoint

	for the get v1 apps app ID operation.

	Typically these are written to a http.Request.
*/
type GetV1AppsAppIDParams struct {

	/* Authorization.

	   bearer auth token
	*/
	Authorization string

	/* XNuonOrgID.

	   org ID
	*/
	XNuonOrgID string

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 apps app ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDParams) WithDefaults() *GetV1AppsAppIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 apps app ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) WithTimeout(timeout time.Duration) *GetV1AppsAppIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) WithContext(ctx context.Context) *GetV1AppsAppIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) WithHTTPClient(client *http.Client) *GetV1AppsAppIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) WithAuthorization(authorization string) *GetV1AppsAppIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) WithXNuonOrgID(xNuonOrgID string) *GetV1AppsAppIDParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithAppID adds the appID to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) WithAppID(appID string) *GetV1AppsAppIDParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get v1 apps app ID params
func (o *GetV1AppsAppIDParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AppsAppIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// header param X-Nuon-Org-ID
	if err := r.SetHeaderParam("X-Nuon-Org-ID", o.XNuonOrgID); err != nil {
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
