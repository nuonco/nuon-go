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

// NewGetV1AppsAppIDInputConfigsParams creates a new GetV1AppsAppIDInputConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AppsAppIDInputConfigsParams() *GetV1AppsAppIDInputConfigsParams {
	return &GetV1AppsAppIDInputConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AppsAppIDInputConfigsParamsWithTimeout creates a new GetV1AppsAppIDInputConfigsParams object
// with the ability to set a timeout on a request.
func NewGetV1AppsAppIDInputConfigsParamsWithTimeout(timeout time.Duration) *GetV1AppsAppIDInputConfigsParams {
	return &GetV1AppsAppIDInputConfigsParams{
		timeout: timeout,
	}
}

// NewGetV1AppsAppIDInputConfigsParamsWithContext creates a new GetV1AppsAppIDInputConfigsParams object
// with the ability to set a context for a request.
func NewGetV1AppsAppIDInputConfigsParamsWithContext(ctx context.Context) *GetV1AppsAppIDInputConfigsParams {
	return &GetV1AppsAppIDInputConfigsParams{
		Context: ctx,
	}
}

// NewGetV1AppsAppIDInputConfigsParamsWithHTTPClient creates a new GetV1AppsAppIDInputConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AppsAppIDInputConfigsParamsWithHTTPClient(client *http.Client) *GetV1AppsAppIDInputConfigsParams {
	return &GetV1AppsAppIDInputConfigsParams{
		HTTPClient: client,
	}
}

/*
GetV1AppsAppIDInputConfigsParams contains all the parameters to send to the API endpoint

	for the get v1 apps app ID input configs operation.

	Typically these are written to a http.Request.
*/
type GetV1AppsAppIDInputConfigsParams struct {

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

// WithDefaults hydrates default values in the get v1 apps app ID input configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDInputConfigsParams) WithDefaults() *GetV1AppsAppIDInputConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 apps app ID input configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDInputConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) WithTimeout(timeout time.Duration) *GetV1AppsAppIDInputConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) WithContext(ctx context.Context) *GetV1AppsAppIDInputConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) WithHTTPClient(client *http.Client) *GetV1AppsAppIDInputConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) WithAuthorization(authorization string) *GetV1AppsAppIDInputConfigsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) WithXNuonOrgID(xNuonOrgID string) *GetV1AppsAppIDInputConfigsParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithAppID adds the appID to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) WithAppID(appID string) *GetV1AppsAppIDInputConfigsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get v1 apps app ID input configs params
func (o *GetV1AppsAppIDInputConfigsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AppsAppIDInputConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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