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
	"github.com/go-openapi/swag"
)

// NewGetOrgInvitesParams creates a new GetOrgInvitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgInvitesParams() *GetOrgInvitesParams {
	return &GetOrgInvitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgInvitesParamsWithTimeout creates a new GetOrgInvitesParams object
// with the ability to set a timeout on a request.
func NewGetOrgInvitesParamsWithTimeout(timeout time.Duration) *GetOrgInvitesParams {
	return &GetOrgInvitesParams{
		timeout: timeout,
	}
}

// NewGetOrgInvitesParamsWithContext creates a new GetOrgInvitesParams object
// with the ability to set a context for a request.
func NewGetOrgInvitesParamsWithContext(ctx context.Context) *GetOrgInvitesParams {
	return &GetOrgInvitesParams{
		Context: ctx,
	}
}

// NewGetOrgInvitesParamsWithHTTPClient creates a new GetOrgInvitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgInvitesParamsWithHTTPClient(client *http.Client) *GetOrgInvitesParams {
	return &GetOrgInvitesParams{
		HTTPClient: client,
	}
}

/*
GetOrgInvitesParams contains all the parameters to send to the API endpoint

	for the get org invites operation.

	Typically these are written to a http.Request.
*/
type GetOrgInvitesParams struct {

	/* Limit.

	   limit of health checks to return

	   Default: 60
	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get org invites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgInvitesParams) WithDefaults() *GetOrgInvitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get org invites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgInvitesParams) SetDefaults() {
	var (
		limitDefault = int64(60)
	)

	val := GetOrgInvitesParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get org invites params
func (o *GetOrgInvitesParams) WithTimeout(timeout time.Duration) *GetOrgInvitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org invites params
func (o *GetOrgInvitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org invites params
func (o *GetOrgInvitesParams) WithContext(ctx context.Context) *GetOrgInvitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org invites params
func (o *GetOrgInvitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org invites params
func (o *GetOrgInvitesParams) WithHTTPClient(client *http.Client) *GetOrgInvitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org invites params
func (o *GetOrgInvitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get org invites params
func (o *GetOrgInvitesParams) WithLimit(limit *int64) *GetOrgInvitesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get org invites params
func (o *GetOrgInvitesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgInvitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
