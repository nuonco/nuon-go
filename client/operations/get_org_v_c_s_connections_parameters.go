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

// NewGetOrgVCSConnectionsParams creates a new GetOrgVCSConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgVCSConnectionsParams() *GetOrgVCSConnectionsParams {
	return &GetOrgVCSConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgVCSConnectionsParamsWithTimeout creates a new GetOrgVCSConnectionsParams object
// with the ability to set a timeout on a request.
func NewGetOrgVCSConnectionsParamsWithTimeout(timeout time.Duration) *GetOrgVCSConnectionsParams {
	return &GetOrgVCSConnectionsParams{
		timeout: timeout,
	}
}

// NewGetOrgVCSConnectionsParamsWithContext creates a new GetOrgVCSConnectionsParams object
// with the ability to set a context for a request.
func NewGetOrgVCSConnectionsParamsWithContext(ctx context.Context) *GetOrgVCSConnectionsParams {
	return &GetOrgVCSConnectionsParams{
		Context: ctx,
	}
}

// NewGetOrgVCSConnectionsParamsWithHTTPClient creates a new GetOrgVCSConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgVCSConnectionsParamsWithHTTPClient(client *http.Client) *GetOrgVCSConnectionsParams {
	return &GetOrgVCSConnectionsParams{
		HTTPClient: client,
	}
}

/*
GetOrgVCSConnectionsParams contains all the parameters to send to the API endpoint

	for the get org v c s connections operation.

	Typically these are written to a http.Request.
*/
type GetOrgVCSConnectionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get org v c s connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgVCSConnectionsParams) WithDefaults() *GetOrgVCSConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get org v c s connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgVCSConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get org v c s connections params
func (o *GetOrgVCSConnectionsParams) WithTimeout(timeout time.Duration) *GetOrgVCSConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org v c s connections params
func (o *GetOrgVCSConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org v c s connections params
func (o *GetOrgVCSConnectionsParams) WithContext(ctx context.Context) *GetOrgVCSConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org v c s connections params
func (o *GetOrgVCSConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org v c s connections params
func (o *GetOrgVCSConnectionsParams) WithHTTPClient(client *http.Client) *GetOrgVCSConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org v c s connections params
func (o *GetOrgVCSConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgVCSConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
