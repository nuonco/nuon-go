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

// NewPatchV1OrgsCurrentParams creates a new PatchV1OrgsCurrentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1OrgsCurrentParams() *PatchV1OrgsCurrentParams {
	return &PatchV1OrgsCurrentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1OrgsCurrentParamsWithTimeout creates a new PatchV1OrgsCurrentParams object
// with the ability to set a timeout on a request.
func NewPatchV1OrgsCurrentParamsWithTimeout(timeout time.Duration) *PatchV1OrgsCurrentParams {
	return &PatchV1OrgsCurrentParams{
		timeout: timeout,
	}
}

// NewPatchV1OrgsCurrentParamsWithContext creates a new PatchV1OrgsCurrentParams object
// with the ability to set a context for a request.
func NewPatchV1OrgsCurrentParamsWithContext(ctx context.Context) *PatchV1OrgsCurrentParams {
	return &PatchV1OrgsCurrentParams{
		Context: ctx,
	}
}

// NewPatchV1OrgsCurrentParamsWithHTTPClient creates a new PatchV1OrgsCurrentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1OrgsCurrentParamsWithHTTPClient(client *http.Client) *PatchV1OrgsCurrentParams {
	return &PatchV1OrgsCurrentParams{
		HTTPClient: client,
	}
}

/*
PatchV1OrgsCurrentParams contains all the parameters to send to the API endpoint

	for the patch v1 orgs current operation.

	Typically these are written to a http.Request.
*/
type PatchV1OrgsCurrentParams struct {

	/* Authorization.

	   bearer auth token
	*/
	Authorization string

	/* XNuonOrgID.

	   org ID
	*/
	XNuonOrgID string

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateOrgRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 orgs current params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1OrgsCurrentParams) WithDefaults() *PatchV1OrgsCurrentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 orgs current params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1OrgsCurrentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) WithTimeout(timeout time.Duration) *PatchV1OrgsCurrentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) WithContext(ctx context.Context) *PatchV1OrgsCurrentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) WithHTTPClient(client *http.Client) *PatchV1OrgsCurrentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) WithAuthorization(authorization string) *PatchV1OrgsCurrentParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXNuonOrgID adds the xNuonOrgID to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) WithXNuonOrgID(xNuonOrgID string) *PatchV1OrgsCurrentParams {
	o.SetXNuonOrgID(xNuonOrgID)
	return o
}

// SetXNuonOrgID adds the xNuonOrgId to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) SetXNuonOrgID(xNuonOrgID string) {
	o.XNuonOrgID = xNuonOrgID
}

// WithReq adds the req to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) WithReq(req *models.ServiceUpdateOrgRequest) *PatchV1OrgsCurrentParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the patch v1 orgs current params
func (o *PatchV1OrgsCurrentParams) SetReq(req *models.ServiceUpdateOrgRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1OrgsCurrentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
