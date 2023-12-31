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

	"github.com/nuonco/nuon-go/models"
)

// NewCreateHelmComponentConfigParams creates a new CreateHelmComponentConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHelmComponentConfigParams() *CreateHelmComponentConfigParams {
	return &CreateHelmComponentConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHelmComponentConfigParamsWithTimeout creates a new CreateHelmComponentConfigParams object
// with the ability to set a timeout on a request.
func NewCreateHelmComponentConfigParamsWithTimeout(timeout time.Duration) *CreateHelmComponentConfigParams {
	return &CreateHelmComponentConfigParams{
		timeout: timeout,
	}
}

// NewCreateHelmComponentConfigParamsWithContext creates a new CreateHelmComponentConfigParams object
// with the ability to set a context for a request.
func NewCreateHelmComponentConfigParamsWithContext(ctx context.Context) *CreateHelmComponentConfigParams {
	return &CreateHelmComponentConfigParams{
		Context: ctx,
	}
}

// NewCreateHelmComponentConfigParamsWithHTTPClient creates a new CreateHelmComponentConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHelmComponentConfigParamsWithHTTPClient(client *http.Client) *CreateHelmComponentConfigParams {
	return &CreateHelmComponentConfigParams{
		HTTPClient: client,
	}
}

/*
CreateHelmComponentConfigParams contains all the parameters to send to the API endpoint

	for the create helm component config operation.

	Typically these are written to a http.Request.
*/
type CreateHelmComponentConfigParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateHelmComponentConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create helm component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHelmComponentConfigParams) WithDefaults() *CreateHelmComponentConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create helm component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHelmComponentConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create helm component config params
func (o *CreateHelmComponentConfigParams) WithTimeout(timeout time.Duration) *CreateHelmComponentConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create helm component config params
func (o *CreateHelmComponentConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create helm component config params
func (o *CreateHelmComponentConfigParams) WithContext(ctx context.Context) *CreateHelmComponentConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create helm component config params
func (o *CreateHelmComponentConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create helm component config params
func (o *CreateHelmComponentConfigParams) WithHTTPClient(client *http.Client) *CreateHelmComponentConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create helm component config params
func (o *CreateHelmComponentConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the create helm component config params
func (o *CreateHelmComponentConfigParams) WithComponentID(componentID string) *CreateHelmComponentConfigParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the create helm component config params
func (o *CreateHelmComponentConfigParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithReq adds the req to the create helm component config params
func (o *CreateHelmComponentConfigParams) WithReq(req *models.ServiceCreateHelmComponentConfigRequest) *CreateHelmComponentConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create helm component config params
func (o *CreateHelmComponentConfigParams) SetReq(req *models.ServiceCreateHelmComponentConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHelmComponentConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
