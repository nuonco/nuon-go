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

// NewCreateOrgInviteParams creates a new CreateOrgInviteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateOrgInviteParams() *CreateOrgInviteParams {
	return &CreateOrgInviteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrgInviteParamsWithTimeout creates a new CreateOrgInviteParams object
// with the ability to set a timeout on a request.
func NewCreateOrgInviteParamsWithTimeout(timeout time.Duration) *CreateOrgInviteParams {
	return &CreateOrgInviteParams{
		timeout: timeout,
	}
}

// NewCreateOrgInviteParamsWithContext creates a new CreateOrgInviteParams object
// with the ability to set a context for a request.
func NewCreateOrgInviteParamsWithContext(ctx context.Context) *CreateOrgInviteParams {
	return &CreateOrgInviteParams{
		Context: ctx,
	}
}

// NewCreateOrgInviteParamsWithHTTPClient creates a new CreateOrgInviteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateOrgInviteParamsWithHTTPClient(client *http.Client) *CreateOrgInviteParams {
	return &CreateOrgInviteParams{
		HTTPClient: client,
	}
}

/*
CreateOrgInviteParams contains all the parameters to send to the API endpoint

	for the create org invite operation.

	Typically these are written to a http.Request.
*/
type CreateOrgInviteParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateOrgInviteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create org invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgInviteParams) WithDefaults() *CreateOrgInviteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create org invite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateOrgInviteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create org invite params
func (o *CreateOrgInviteParams) WithTimeout(timeout time.Duration) *CreateOrgInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create org invite params
func (o *CreateOrgInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create org invite params
func (o *CreateOrgInviteParams) WithContext(ctx context.Context) *CreateOrgInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create org invite params
func (o *CreateOrgInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create org invite params
func (o *CreateOrgInviteParams) WithHTTPClient(client *http.Client) *CreateOrgInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create org invite params
func (o *CreateOrgInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the create org invite params
func (o *CreateOrgInviteParams) WithReq(req *models.ServiceCreateOrgInviteRequest) *CreateOrgInviteParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create org invite params
func (o *CreateOrgInviteParams) SetReq(req *models.ServiceCreateOrgInviteRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrgInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
