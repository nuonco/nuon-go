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

// NewGetV1AppsAppIDSandboxConfigsParams creates a new GetV1AppsAppIDSandboxConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AppsAppIDSandboxConfigsParams() *GetV1AppsAppIDSandboxConfigsParams {
	return &GetV1AppsAppIDSandboxConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AppsAppIDSandboxConfigsParamsWithTimeout creates a new GetV1AppsAppIDSandboxConfigsParams object
// with the ability to set a timeout on a request.
func NewGetV1AppsAppIDSandboxConfigsParamsWithTimeout(timeout time.Duration) *GetV1AppsAppIDSandboxConfigsParams {
	return &GetV1AppsAppIDSandboxConfigsParams{
		timeout: timeout,
	}
}

// NewGetV1AppsAppIDSandboxConfigsParamsWithContext creates a new GetV1AppsAppIDSandboxConfigsParams object
// with the ability to set a context for a request.
func NewGetV1AppsAppIDSandboxConfigsParamsWithContext(ctx context.Context) *GetV1AppsAppIDSandboxConfigsParams {
	return &GetV1AppsAppIDSandboxConfigsParams{
		Context: ctx,
	}
}

// NewGetV1AppsAppIDSandboxConfigsParamsWithHTTPClient creates a new GetV1AppsAppIDSandboxConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AppsAppIDSandboxConfigsParamsWithHTTPClient(client *http.Client) *GetV1AppsAppIDSandboxConfigsParams {
	return &GetV1AppsAppIDSandboxConfigsParams{
		HTTPClient: client,
	}
}

/*
GetV1AppsAppIDSandboxConfigsParams contains all the parameters to send to the API endpoint

	for the get v1 apps app ID sandbox configs operation.

	Typically these are written to a http.Request.
*/
type GetV1AppsAppIDSandboxConfigsParams struct {

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

// WithDefaults hydrates default values in the get v1 apps app ID sandbox configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDSandboxConfigsParams) WithDefaults() *GetV1AppsAppIDSandboxConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 apps app ID sandbox configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AppsAppIDSandboxConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) WithTimeout(timeout time.Duration) *GetV1AppsAppIDSandboxConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) WithContext(ctx context.Context) *GetV1AppsAppIDSandboxConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) WithHTTPClient(client *http.Client) *GetV1AppsAppIDSandboxConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) WithAuthorization(authorization string) *GetV1AppsAppIDSandboxConfigsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) WithXNuonOrgID(xNuonOrgID string) *GetV1AppsAppIDSandboxConfigsParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithAppID adds the appID to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) WithAppID(appID string) *GetV1AppsAppIDSandboxConfigsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get v1 apps app ID sandbox configs params
func (o *GetV1AppsAppIDSandboxConfigsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AppsAppIDSandboxConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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