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

// NewCreateAppActionWorkflowParams creates a new CreateAppActionWorkflowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppActionWorkflowParams() *CreateAppActionWorkflowParams {
	return &CreateAppActionWorkflowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppActionWorkflowParamsWithTimeout creates a new CreateAppActionWorkflowParams object
// with the ability to set a timeout on a request.
func NewCreateAppActionWorkflowParamsWithTimeout(timeout time.Duration) *CreateAppActionWorkflowParams {
	return &CreateAppActionWorkflowParams{
		timeout: timeout,
	}
}

// NewCreateAppActionWorkflowParamsWithContext creates a new CreateAppActionWorkflowParams object
// with the ability to set a context for a request.
func NewCreateAppActionWorkflowParamsWithContext(ctx context.Context) *CreateAppActionWorkflowParams {
	return &CreateAppActionWorkflowParams{
		Context: ctx,
	}
}

// NewCreateAppActionWorkflowParamsWithHTTPClient creates a new CreateAppActionWorkflowParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppActionWorkflowParamsWithHTTPClient(client *http.Client) *CreateAppActionWorkflowParams {
	return &CreateAppActionWorkflowParams{
		HTTPClient: client,
	}
}

/*
CreateAppActionWorkflowParams contains all the parameters to send to the API endpoint

	for the create app action workflow operation.

	Typically these are written to a http.Request.
*/
type CreateAppActionWorkflowParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateAppActionWorkflowRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app action workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppActionWorkflowParams) WithDefaults() *CreateAppActionWorkflowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app action workflow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppActionWorkflowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app action workflow params
func (o *CreateAppActionWorkflowParams) WithTimeout(timeout time.Duration) *CreateAppActionWorkflowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app action workflow params
func (o *CreateAppActionWorkflowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app action workflow params
func (o *CreateAppActionWorkflowParams) WithContext(ctx context.Context) *CreateAppActionWorkflowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app action workflow params
func (o *CreateAppActionWorkflowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app action workflow params
func (o *CreateAppActionWorkflowParams) WithHTTPClient(client *http.Client) *CreateAppActionWorkflowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app action workflow params
func (o *CreateAppActionWorkflowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create app action workflow params
func (o *CreateAppActionWorkflowParams) WithAppID(appID string) *CreateAppActionWorkflowParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create app action workflow params
func (o *CreateAppActionWorkflowParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReq adds the req to the create app action workflow params
func (o *CreateAppActionWorkflowParams) WithReq(req *models.ServiceCreateAppActionWorkflowRequest) *CreateAppActionWorkflowParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create app action workflow params
func (o *CreateAppActionWorkflowParams) SetReq(req *models.ServiceCreateAppActionWorkflowRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppActionWorkflowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
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