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

// NewPatchV1ComponentsComponentIDParams creates a new PatchV1ComponentsComponentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ComponentsComponentIDParams() *PatchV1ComponentsComponentIDParams {
	return &PatchV1ComponentsComponentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ComponentsComponentIDParamsWithTimeout creates a new PatchV1ComponentsComponentIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1ComponentsComponentIDParamsWithTimeout(timeout time.Duration) *PatchV1ComponentsComponentIDParams {
	return &PatchV1ComponentsComponentIDParams{
		timeout: timeout,
	}
}

// NewPatchV1ComponentsComponentIDParamsWithContext creates a new PatchV1ComponentsComponentIDParams object
// with the ability to set a context for a request.
func NewPatchV1ComponentsComponentIDParamsWithContext(ctx context.Context) *PatchV1ComponentsComponentIDParams {
	return &PatchV1ComponentsComponentIDParams{
		Context: ctx,
	}
}

// NewPatchV1ComponentsComponentIDParamsWithHTTPClient creates a new PatchV1ComponentsComponentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ComponentsComponentIDParamsWithHTTPClient(client *http.Client) *PatchV1ComponentsComponentIDParams {
	return &PatchV1ComponentsComponentIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1ComponentsComponentIDParams contains all the parameters to send to the API endpoint

	for the patch v1 components component ID operation.

	Typically these are written to a http.Request.
*/
type PatchV1ComponentsComponentIDParams struct {

	/* Authorization.

	   bearer auth token
	*/
	Authorization string

	/* XNuonOrgID.

	   org ID
	*/
	XNuonOrgID string

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateComponentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 components component ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ComponentsComponentIDParams) WithDefaults() *PatchV1ComponentsComponentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 components component ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ComponentsComponentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithTimeout(timeout time.Duration) *PatchV1ComponentsComponentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithContext(ctx context.Context) *PatchV1ComponentsComponentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithHTTPClient(client *http.Client) *PatchV1ComponentsComponentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithAuthorization(authorization string) *PatchV1ComponentsComponentIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithXNuonOrgID(xNuonOrgID string) *PatchV1ComponentsComponentIDParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithComponentID adds the componentID to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithComponentID(componentID string) *PatchV1ComponentsComponentIDParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithReq adds the req to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) WithReq(req *models.ServiceUpdateComponentRequest) *PatchV1ComponentsComponentIDParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the patch v1 components component ID params
func (o *PatchV1ComponentsComponentIDParams) SetReq(req *models.ServiceUpdateComponentRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ComponentsComponentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
