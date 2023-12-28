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

// NewGetCLIConfigParams creates a new GetCLIConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCLIConfigParams() *GetCLIConfigParams {
	return &GetCLIConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCLIConfigParamsWithTimeout creates a new GetCLIConfigParams object
// with the ability to set a timeout on a request.
func NewGetCLIConfigParamsWithTimeout(timeout time.Duration) *GetCLIConfigParams {
	return &GetCLIConfigParams{
		timeout: timeout,
	}
}

// NewGetCLIConfigParamsWithContext creates a new GetCLIConfigParams object
// with the ability to set a context for a request.
func NewGetCLIConfigParamsWithContext(ctx context.Context) *GetCLIConfigParams {
	return &GetCLIConfigParams{
		Context: ctx,
	}
}

// NewGetCLIConfigParamsWithHTTPClient creates a new GetCLIConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCLIConfigParamsWithHTTPClient(client *http.Client) *GetCLIConfigParams {
	return &GetCLIConfigParams{
		HTTPClient: client,
	}
}

/*
GetCLIConfigParams contains all the parameters to send to the API endpoint

	for the get c l i config operation.

	Typically these are written to a http.Request.
*/
type GetCLIConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c l i config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCLIConfigParams) WithDefaults() *GetCLIConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c l i config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCLIConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c l i config params
func (o *GetCLIConfigParams) WithTimeout(timeout time.Duration) *GetCLIConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c l i config params
func (o *GetCLIConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c l i config params
func (o *GetCLIConfigParams) WithContext(ctx context.Context) *GetCLIConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c l i config params
func (o *GetCLIConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c l i config params
func (o *GetCLIConfigParams) WithHTTPClient(client *http.Client) *GetCLIConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c l i config params
func (o *GetCLIConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCLIConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
