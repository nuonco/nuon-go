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

// NewGetV1GeneralCliConfigParams creates a new GetV1GeneralCliConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1GeneralCliConfigParams() *GetV1GeneralCliConfigParams {
	return &GetV1GeneralCliConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1GeneralCliConfigParamsWithTimeout creates a new GetV1GeneralCliConfigParams object
// with the ability to set a timeout on a request.
func NewGetV1GeneralCliConfigParamsWithTimeout(timeout time.Duration) *GetV1GeneralCliConfigParams {
	return &GetV1GeneralCliConfigParams{
		timeout: timeout,
	}
}

// NewGetV1GeneralCliConfigParamsWithContext creates a new GetV1GeneralCliConfigParams object
// with the ability to set a context for a request.
func NewGetV1GeneralCliConfigParamsWithContext(ctx context.Context) *GetV1GeneralCliConfigParams {
	return &GetV1GeneralCliConfigParams{
		Context: ctx,
	}
}

// NewGetV1GeneralCliConfigParamsWithHTTPClient creates a new GetV1GeneralCliConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1GeneralCliConfigParamsWithHTTPClient(client *http.Client) *GetV1GeneralCliConfigParams {
	return &GetV1GeneralCliConfigParams{
		HTTPClient: client,
	}
}

/*
GetV1GeneralCliConfigParams contains all the parameters to send to the API endpoint

	for the get v1 general cli config operation.

	Typically these are written to a http.Request.
*/
type GetV1GeneralCliConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 general cli config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1GeneralCliConfigParams) WithDefaults() *GetV1GeneralCliConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 general cli config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1GeneralCliConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 general cli config params
func (o *GetV1GeneralCliConfigParams) WithTimeout(timeout time.Duration) *GetV1GeneralCliConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 general cli config params
func (o *GetV1GeneralCliConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 general cli config params
func (o *GetV1GeneralCliConfigParams) WithContext(ctx context.Context) *GetV1GeneralCliConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 general cli config params
func (o *GetV1GeneralCliConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 general cli config params
func (o *GetV1GeneralCliConfigParams) WithHTTPClient(client *http.Client) *GetV1GeneralCliConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 general cli config params
func (o *GetV1GeneralCliConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1GeneralCliConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
