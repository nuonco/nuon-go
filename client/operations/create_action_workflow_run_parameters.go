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

// NewCreateActionWorkflowRunParams creates a new CreateActionWorkflowRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateActionWorkflowRunParams() *CreateActionWorkflowRunParams {
	return &CreateActionWorkflowRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActionWorkflowRunParamsWithTimeout creates a new CreateActionWorkflowRunParams object
// with the ability to set a timeout on a request.
func NewCreateActionWorkflowRunParamsWithTimeout(timeout time.Duration) *CreateActionWorkflowRunParams {
	return &CreateActionWorkflowRunParams{
		timeout: timeout,
	}
}

// NewCreateActionWorkflowRunParamsWithContext creates a new CreateActionWorkflowRunParams object
// with the ability to set a context for a request.
func NewCreateActionWorkflowRunParamsWithContext(ctx context.Context) *CreateActionWorkflowRunParams {
	return &CreateActionWorkflowRunParams{
		Context: ctx,
	}
}

// NewCreateActionWorkflowRunParamsWithHTTPClient creates a new CreateActionWorkflowRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateActionWorkflowRunParamsWithHTTPClient(client *http.Client) *CreateActionWorkflowRunParams {
	return &CreateActionWorkflowRunParams{
		HTTPClient: client,
	}
}

/*
CreateActionWorkflowRunParams contains all the parameters to send to the API endpoint

	for the create action workflow run operation.

	Typically these are written to a http.Request.
*/
type CreateActionWorkflowRunParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateActionWorkflowRunRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create action workflow run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActionWorkflowRunParams) WithDefaults() *CreateActionWorkflowRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create action workflow run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActionWorkflowRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create action workflow run params
func (o *CreateActionWorkflowRunParams) WithTimeout(timeout time.Duration) *CreateActionWorkflowRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create action workflow run params
func (o *CreateActionWorkflowRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create action workflow run params
func (o *CreateActionWorkflowRunParams) WithContext(ctx context.Context) *CreateActionWorkflowRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create action workflow run params
func (o *CreateActionWorkflowRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create action workflow run params
func (o *CreateActionWorkflowRunParams) WithHTTPClient(client *http.Client) *CreateActionWorkflowRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create action workflow run params
func (o *CreateActionWorkflowRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the create action workflow run params
func (o *CreateActionWorkflowRunParams) WithReq(req *models.ServiceCreateActionWorkflowRunRequest) *CreateActionWorkflowRunParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create action workflow run params
func (o *CreateActionWorkflowRunParams) SetReq(req *models.ServiceCreateActionWorkflowRunRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActionWorkflowRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
