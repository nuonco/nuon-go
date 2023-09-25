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

// NewGetV1ComponentsComponentIDReleasesParams creates a new GetV1ComponentsComponentIDReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ComponentsComponentIDReleasesParams() *GetV1ComponentsComponentIDReleasesParams {
	return &GetV1ComponentsComponentIDReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ComponentsComponentIDReleasesParamsWithTimeout creates a new GetV1ComponentsComponentIDReleasesParams object
// with the ability to set a timeout on a request.
func NewGetV1ComponentsComponentIDReleasesParamsWithTimeout(timeout time.Duration) *GetV1ComponentsComponentIDReleasesParams {
	return &GetV1ComponentsComponentIDReleasesParams{
		timeout: timeout,
	}
}

// NewGetV1ComponentsComponentIDReleasesParamsWithContext creates a new GetV1ComponentsComponentIDReleasesParams object
// with the ability to set a context for a request.
func NewGetV1ComponentsComponentIDReleasesParamsWithContext(ctx context.Context) *GetV1ComponentsComponentIDReleasesParams {
	return &GetV1ComponentsComponentIDReleasesParams{
		Context: ctx,
	}
}

// NewGetV1ComponentsComponentIDReleasesParamsWithHTTPClient creates a new GetV1ComponentsComponentIDReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ComponentsComponentIDReleasesParamsWithHTTPClient(client *http.Client) *GetV1ComponentsComponentIDReleasesParams {
	return &GetV1ComponentsComponentIDReleasesParams{
		HTTPClient: client,
	}
}

/*
GetV1ComponentsComponentIDReleasesParams contains all the parameters to send to the API endpoint

	for the get v1 components component ID releases operation.

	Typically these are written to a http.Request.
*/
type GetV1ComponentsComponentIDReleasesParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 components component ID releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ComponentsComponentIDReleasesParams) WithDefaults() *GetV1ComponentsComponentIDReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 components component ID releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ComponentsComponentIDReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) WithTimeout(timeout time.Duration) *GetV1ComponentsComponentIDReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) WithContext(ctx context.Context) *GetV1ComponentsComponentIDReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) WithHTTPClient(client *http.Client) *GetV1ComponentsComponentIDReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) WithComponentID(componentID string) *GetV1ComponentsComponentIDReleasesParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the get v1 components component ID releases params
func (o *GetV1ComponentsComponentIDReleasesParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ComponentsComponentIDReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}