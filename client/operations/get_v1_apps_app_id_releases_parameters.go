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

// NewGetV1AppsAppIDReleasesParams creates a new GetV1AppsAppIDReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AppsAppIDReleasesParams() *GetV1AppsAppIDReleasesParams {
	return &GetV1AppsAppIDReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AppsAppIDReleasesParamsWithTimeout creates a new GetV1AppsAppIDReleasesParams object
// with the ability to set a timeout on a request.
func NewGetV1AppsAppIDReleasesParamsWithTimeout(timeout time.Duration) *GetV1AppsAppIDReleasesParams {
	return &GetV1AppsAppIDReleasesParams{
		timeout: timeout,
	}
}

// NewGetV1AppsAppIDReleasesParamsWithContext creates a new GetV1AppsAppIDReleasesParams object
// with the ability to set a context for a request.
func NewGetV1AppsAppIDReleasesParamsWithContext(ctx context.Context) *GetV1AppsAppIDReleasesParams {
	return &GetV1AppsAppIDReleasesParams{
		Context: ctx,
	}
}

// NewGetV1AppsAppIDReleasesParamsWithHTTPClient creates a new GetV1AppsAppIDReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AppsAppIDReleasesParamsWithHTTPClient(client *http.Client) *GetV1AppsAppIDReleasesParams {
	return &GetV1AppsAppIDReleasesParams{
		HTTPClient: client,
	}
}

/*
GetV1AppsAppIDReleasesParams contains all the parameters to send to the API endpoint

	for the get v1 apps app ID releases operation.

	Typically these are written to a http.Request.
*/
type GetV1AppsAppIDReleasesParams struct {

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

// WithDefaults hydrates default values in the get v1 apps app ID releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDReleasesParams) WithDefaults() *GetV1AppsAppIDReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 apps app ID releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) WithTimeout(timeout time.Duration) *GetV1AppsAppIDReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) WithContext(ctx context.Context) *GetV1AppsAppIDReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) WithHTTPClient(client *http.Client) *GetV1AppsAppIDReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) WithAuthorization(authorization string) *GetV1AppsAppIDReleasesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) WithXNuonOrgID(xNuonOrgID string) *GetV1AppsAppIDReleasesParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithAppID adds the appID to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) WithAppID(appID string) *GetV1AppsAppIDReleasesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get v1 apps app ID releases params
func (o *GetV1AppsAppIDReleasesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AppsAppIDReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
