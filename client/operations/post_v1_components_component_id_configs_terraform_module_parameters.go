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

// NewPostV1ComponentsComponentIDConfigsTerraformModuleParams creates a new PostV1ComponentsComponentIDConfigsTerraformModuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ComponentsComponentIDConfigsTerraformModuleParams() *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	return &PostV1ComponentsComponentIDConfigsTerraformModuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ComponentsComponentIDConfigsTerraformModuleParamsWithTimeout creates a new PostV1ComponentsComponentIDConfigsTerraformModuleParams object
// with the ability to set a timeout on a request.
func NewPostV1ComponentsComponentIDConfigsTerraformModuleParamsWithTimeout(timeout time.Duration) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	return &PostV1ComponentsComponentIDConfigsTerraformModuleParams{
		timeout: timeout,
	}
}

// NewPostV1ComponentsComponentIDConfigsTerraformModuleParamsWithContext creates a new PostV1ComponentsComponentIDConfigsTerraformModuleParams object
// with the ability to set a context for a request.
func NewPostV1ComponentsComponentIDConfigsTerraformModuleParamsWithContext(ctx context.Context) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	return &PostV1ComponentsComponentIDConfigsTerraformModuleParams{
		Context: ctx,
	}
}

// NewPostV1ComponentsComponentIDConfigsTerraformModuleParamsWithHTTPClient creates a new PostV1ComponentsComponentIDConfigsTerraformModuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ComponentsComponentIDConfigsTerraformModuleParamsWithHTTPClient(client *http.Client) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	return &PostV1ComponentsComponentIDConfigsTerraformModuleParams{
		HTTPClient: client,
	}
}

/*
PostV1ComponentsComponentIDConfigsTerraformModuleParams contains all the parameters to send to the API endpoint

	for the post v1 components component ID configs terraform module operation.

	Typically these are written to a http.Request.
*/
type PostV1ComponentsComponentIDConfigsTerraformModuleParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateTerraformModuleComponentConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 components component ID configs terraform module params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WithDefaults() *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 components component ID configs terraform module params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WithTimeout(timeout time.Duration) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WithContext(ctx context.Context) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WithHTTPClient(client *http.Client) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WithComponentID(componentID string) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithReq adds the req to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WithReq(req *models.ServiceCreateTerraformModuleComponentConfigRequest) *PostV1ComponentsComponentIDConfigsTerraformModuleParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the post v1 components component ID configs terraform module params
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) SetReq(req *models.ServiceCreateTerraformModuleComponentConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ComponentsComponentIDConfigsTerraformModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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