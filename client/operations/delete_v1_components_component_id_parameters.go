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

// NewDeleteV1ComponentsComponentIDParams creates a new DeleteV1ComponentsComponentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ComponentsComponentIDParams() *DeleteV1ComponentsComponentIDParams {
	return &DeleteV1ComponentsComponentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ComponentsComponentIDParamsWithTimeout creates a new DeleteV1ComponentsComponentIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ComponentsComponentIDParamsWithTimeout(timeout time.Duration) *DeleteV1ComponentsComponentIDParams {
	return &DeleteV1ComponentsComponentIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ComponentsComponentIDParamsWithContext creates a new DeleteV1ComponentsComponentIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ComponentsComponentIDParamsWithContext(ctx context.Context) *DeleteV1ComponentsComponentIDParams {
	return &DeleteV1ComponentsComponentIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ComponentsComponentIDParamsWithHTTPClient creates a new DeleteV1ComponentsComponentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ComponentsComponentIDParamsWithHTTPClient(client *http.Client) *DeleteV1ComponentsComponentIDParams {
	return &DeleteV1ComponentsComponentIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ComponentsComponentIDParams contains all the parameters to send to the API endpoint

	for the delete v1 components component ID operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ComponentsComponentIDParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 components component ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ComponentsComponentIDParams) WithDefaults() *DeleteV1ComponentsComponentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 components component ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ComponentsComponentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) WithTimeout(timeout time.Duration) *DeleteV1ComponentsComponentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) WithContext(ctx context.Context) *DeleteV1ComponentsComponentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) WithHTTPClient(client *http.Client) *DeleteV1ComponentsComponentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) WithComponentID(componentID string) *DeleteV1ComponentsComponentIDParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the delete v1 components component ID params
func (o *DeleteV1ComponentsComponentIDParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ComponentsComponentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
