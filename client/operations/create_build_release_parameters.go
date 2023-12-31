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

// NewCreateBuildReleaseParams creates a new CreateBuildReleaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBuildReleaseParams() *CreateBuildReleaseParams {
	return &CreateBuildReleaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBuildReleaseParamsWithTimeout creates a new CreateBuildReleaseParams object
// with the ability to set a timeout on a request.
func NewCreateBuildReleaseParamsWithTimeout(timeout time.Duration) *CreateBuildReleaseParams {
	return &CreateBuildReleaseParams{
		timeout: timeout,
	}
}

// NewCreateBuildReleaseParamsWithContext creates a new CreateBuildReleaseParams object
// with the ability to set a context for a request.
func NewCreateBuildReleaseParamsWithContext(ctx context.Context) *CreateBuildReleaseParams {
	return &CreateBuildReleaseParams{
		Context: ctx,
	}
}

// NewCreateBuildReleaseParamsWithHTTPClient creates a new CreateBuildReleaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBuildReleaseParamsWithHTTPClient(client *http.Client) *CreateBuildReleaseParams {
	return &CreateBuildReleaseParams{
		HTTPClient: client,
	}
}

/*
CreateBuildReleaseParams contains all the parameters to send to the API endpoint

	for the create build release operation.

	Typically these are written to a http.Request.
*/
type CreateBuildReleaseParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateComponentReleaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create build release params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBuildReleaseParams) WithDefaults() *CreateBuildReleaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create build release params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBuildReleaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create build release params
func (o *CreateBuildReleaseParams) WithTimeout(timeout time.Duration) *CreateBuildReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create build release params
func (o *CreateBuildReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create build release params
func (o *CreateBuildReleaseParams) WithContext(ctx context.Context) *CreateBuildReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create build release params
func (o *CreateBuildReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create build release params
func (o *CreateBuildReleaseParams) WithHTTPClient(client *http.Client) *CreateBuildReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create build release params
func (o *CreateBuildReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the create build release params
func (o *CreateBuildReleaseParams) WithReq(req *models.ServiceCreateComponentReleaseRequest) *CreateBuildReleaseParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create build release params
func (o *CreateBuildReleaseParams) SetReq(req *models.ServiceCreateComponentReleaseRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBuildReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
