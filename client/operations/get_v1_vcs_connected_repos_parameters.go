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

// NewGetV1VcsConnectedReposParams creates a new GetV1VcsConnectedReposParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1VcsConnectedReposParams() *GetV1VcsConnectedReposParams {
	return &GetV1VcsConnectedReposParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1VcsConnectedReposParamsWithTimeout creates a new GetV1VcsConnectedReposParams object
// with the ability to set a timeout on a request.
func NewGetV1VcsConnectedReposParamsWithTimeout(timeout time.Duration) *GetV1VcsConnectedReposParams {
	return &GetV1VcsConnectedReposParams{
		timeout: timeout,
	}
}

// NewGetV1VcsConnectedReposParamsWithContext creates a new GetV1VcsConnectedReposParams object
// with the ability to set a context for a request.
func NewGetV1VcsConnectedReposParamsWithContext(ctx context.Context) *GetV1VcsConnectedReposParams {
	return &GetV1VcsConnectedReposParams{
		Context: ctx,
	}
}

// NewGetV1VcsConnectedReposParamsWithHTTPClient creates a new GetV1VcsConnectedReposParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1VcsConnectedReposParamsWithHTTPClient(client *http.Client) *GetV1VcsConnectedReposParams {
	return &GetV1VcsConnectedReposParams{
		HTTPClient: client,
	}
}

/*
GetV1VcsConnectedReposParams contains all the parameters to send to the API endpoint

	for the get v1 vcs connected repos operation.

	Typically these are written to a http.Request.
*/
type GetV1VcsConnectedReposParams struct {

	/* Authorization.

	   bearer auth token
	*/
	Authorization string

	/* XNuonOrgID.

	   org ID
	*/
	XNuonOrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 vcs connected repos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1VcsConnectedReposParams) WithDefaults() *GetV1VcsConnectedReposParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 vcs connected repos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1VcsConnectedReposParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) WithTimeout(timeout time.Duration) *GetV1VcsConnectedReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) WithContext(ctx context.Context) *GetV1VcsConnectedReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) WithHTTPClient(client *http.Client) *GetV1VcsConnectedReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) WithAuthorization(authorization string) *GetV1VcsConnectedReposParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) WithXNuonOrgID(xNuonOrgID string) *GetV1VcsConnectedReposParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the get v1 vcs connected repos params
func (o *GetV1VcsConnectedReposParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1VcsConnectedReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
