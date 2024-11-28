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

// NewCreateActionWorkflowConfigParams creates a new CreateActionWorkflowConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateActionWorkflowConfigParams() *CreateActionWorkflowConfigParams {
	return &CreateActionWorkflowConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActionWorkflowConfigParamsWithTimeout creates a new CreateActionWorkflowConfigParams object
// with the ability to set a timeout on a request.
func NewCreateActionWorkflowConfigParamsWithTimeout(timeout time.Duration) *CreateActionWorkflowConfigParams {
	return &CreateActionWorkflowConfigParams{
		timeout: timeout,
	}
}

// NewCreateActionWorkflowConfigParamsWithContext creates a new CreateActionWorkflowConfigParams object
// with the ability to set a context for a request.
func NewCreateActionWorkflowConfigParamsWithContext(ctx context.Context) *CreateActionWorkflowConfigParams {
	return &CreateActionWorkflowConfigParams{
		Context: ctx,
	}
}

// NewCreateActionWorkflowConfigParamsWithHTTPClient creates a new CreateActionWorkflowConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateActionWorkflowConfigParamsWithHTTPClient(client *http.Client) *CreateActionWorkflowConfigParams {
	return &CreateActionWorkflowConfigParams{
		HTTPClient: client,
	}
}

/*
CreateActionWorkflowConfigParams contains all the parameters to send to the API endpoint

	for the create action workflow config operation.

	Typically these are written to a http.Request.
*/
type CreateActionWorkflowConfigParams struct {

	/* ActionWorkflowID.

	   action workflow ID
	*/
	ActionWorkflowID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateActionWorkflowConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create action workflow config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActionWorkflowConfigParams) WithDefaults() *CreateActionWorkflowConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create action workflow config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActionWorkflowConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) WithTimeout(timeout time.Duration) *CreateActionWorkflowConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) WithContext(ctx context.Context) *CreateActionWorkflowConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) WithHTTPClient(client *http.Client) *CreateActionWorkflowConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionWorkflowID adds the actionWorkflowID to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) WithActionWorkflowID(actionWorkflowID string) *CreateActionWorkflowConfigParams {
	o.SetActionWorkflowID(actionWorkflowID)
	return o
}

// SetActionWorkflowID adds the actionWorkflowId to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) SetActionWorkflowID(actionWorkflowID string) {
	o.ActionWorkflowID = actionWorkflowID
}

// WithReq adds the req to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) WithReq(req *models.ServiceCreateActionWorkflowConfigRequest) *CreateActionWorkflowConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create action workflow config params
func (o *CreateActionWorkflowConfigParams) SetReq(req *models.ServiceCreateActionWorkflowConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActionWorkflowConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_workflow_id
	if err := r.SetPathParam("action_workflow_id", o.ActionWorkflowID); err != nil {
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
