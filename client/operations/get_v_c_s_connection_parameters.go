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

// NewGetVCSConnectionParams creates a new GetVCSConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVCSConnectionParams() *GetVCSConnectionParams {
	return &GetVCSConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVCSConnectionParamsWithTimeout creates a new GetVCSConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVCSConnectionParamsWithTimeout(timeout time.Duration) *GetVCSConnectionParams {
	return &GetVCSConnectionParams{
		timeout: timeout,
	}
}

// NewGetVCSConnectionParamsWithContext creates a new GetVCSConnectionParams object
// with the ability to set a context for a request.
func NewGetVCSConnectionParamsWithContext(ctx context.Context) *GetVCSConnectionParams {
	return &GetVCSConnectionParams{
		Context: ctx,
	}
}

// NewGetVCSConnectionParamsWithHTTPClient creates a new GetVCSConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVCSConnectionParamsWithHTTPClient(client *http.Client) *GetVCSConnectionParams {
	return &GetVCSConnectionParams{
		HTTPClient: client,
	}
}

/*
GetVCSConnectionParams contains all the parameters to send to the API endpoint

	for the get v c s connection operation.

	Typically these are written to a http.Request.
*/
type GetVCSConnectionParams struct {

	/* ConnectionID.

	   connection ID
	*/
	ConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v c s connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVCSConnectionParams) WithDefaults() *GetVCSConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v c s connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVCSConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v c s connection params
func (o *GetVCSConnectionParams) WithTimeout(timeout time.Duration) *GetVCSConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v c s connection params
func (o *GetVCSConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v c s connection params
func (o *GetVCSConnectionParams) WithContext(ctx context.Context) *GetVCSConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v c s connection params
func (o *GetVCSConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v c s connection params
func (o *GetVCSConnectionParams) WithHTTPClient(client *http.Client) *GetVCSConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v c s connection params
func (o *GetVCSConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the get v c s connection params
func (o *GetVCSConnectionParams) WithConnectionID(connectionID string) *GetVCSConnectionParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get v c s connection params
func (o *GetVCSConnectionParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVCSConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
