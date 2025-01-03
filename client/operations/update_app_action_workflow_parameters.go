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

// NewUpdateAppActionWorkflowParams creates a new UpdateAppActionWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAppActionWorkflowParams() *UpdateAppActionWorkflowParams {
	return &UpdateAppActionWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppActionWorkflowParamsWithTimeout creates a new UpdateAppActionWorkflowParams object
// with the ability to set a timeout on a request.
func NewUpdateAppActionWorkflowParamsWithTimeout(timeout time.Duration) *UpdateAppActionWorkflowParams {
	return &UpdateAppActionWorkflowParams{
		timeout: timeout,
	}
}

// NewUpdateAppActionWorkflowParamsWithContext creates a new UpdateAppActionWorkflowParams object
// with the ability to set a context for a request.
func NewUpdateAppActionWorkflowParamsWithContext(ctx context.Context) *UpdateAppActionWorkflowParams {
	return &UpdateAppActionWorkflowParams{
		Context: ctx,
	}
}

// NewUpdateAppActionWorkflowParamsWithHTTPClient creates a new UpdateAppActionWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAppActionWorkflowParamsWithHTTPClient(client *http.Client) *UpdateAppActionWorkflowParams {
	return &UpdateAppActionWorkflowParams{
		HTTPClient: client,
	}
}

/*
UpdateAppActionWorkflowParams contains all the parameters to send to the API endpoint

	for the update app action workflow operation.

	Typically these are written to a http.Request.
*/
type UpdateAppActionWorkflowParams struct {

	/* ActionWorkflowID.

	   action workflow ID
	*/
	ActionWorkflowID string

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateActionWorkflowRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update app action workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppActionWorkflowParams) WithDefaults() *UpdateAppActionWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update app action workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppActionWorkflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) WithTimeout(timeout time.Duration) *UpdateAppActionWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) WithContext(ctx context.Context) *UpdateAppActionWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) WithHTTPClient(client *http.Client) *UpdateAppActionWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionWorkflowID adds the actionWorkflowID to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) WithActionWorkflowID(actionWorkflowID string) *UpdateAppActionWorkflowParams {
	o.SetActionWorkflowID(actionWorkflowID)
	return o
}

// SetActionWorkflowID adds the actionWorkflowId to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) SetActionWorkflowID(actionWorkflowID string) {
	o.ActionWorkflowID = actionWorkflowID
}

// WithReq adds the req to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) WithReq(req *models.ServiceUpdateActionWorkflowRequest) *UpdateAppActionWorkflowParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the update app action workflow params
func (o *UpdateAppActionWorkflowParams) SetReq(req *models.ServiceUpdateActionWorkflowRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppActionWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
