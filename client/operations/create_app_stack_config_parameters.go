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

// NewCreateAppStackConfigParams creates a new CreateAppStackConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppStackConfigParams() *CreateAppStackConfigParams {
	return &CreateAppStackConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppStackConfigParamsWithTimeout creates a new CreateAppStackConfigParams object
// with the ability to set a timeout on a request.
func NewCreateAppStackConfigParamsWithTimeout(timeout time.Duration) *CreateAppStackConfigParams {
	return &CreateAppStackConfigParams{
		timeout: timeout,
	}
}

// NewCreateAppStackConfigParamsWithContext creates a new CreateAppStackConfigParams object
// with the ability to set a context for a request.
func NewCreateAppStackConfigParamsWithContext(ctx context.Context) *CreateAppStackConfigParams {
	return &CreateAppStackConfigParams{
		Context: ctx,
	}
}

// NewCreateAppStackConfigParamsWithHTTPClient creates a new CreateAppStackConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppStackConfigParamsWithHTTPClient(client *http.Client) *CreateAppStackConfigParams {
	return &CreateAppStackConfigParams{
		HTTPClient: client,
	}
}

/*
CreateAppStackConfigParams contains all the parameters to send to the API endpoint

	for the create app stack config operation.

	Typically these are written to a http.Request.
*/
type CreateAppStackConfigParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateAppStackConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app stack config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppStackConfigParams) WithDefaults() *CreateAppStackConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app stack config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppStackConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app stack config params
func (o *CreateAppStackConfigParams) WithTimeout(timeout time.Duration) *CreateAppStackConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app stack config params
func (o *CreateAppStackConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app stack config params
func (o *CreateAppStackConfigParams) WithContext(ctx context.Context) *CreateAppStackConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app stack config params
func (o *CreateAppStackConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app stack config params
func (o *CreateAppStackConfigParams) WithHTTPClient(client *http.Client) *CreateAppStackConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app stack config params
func (o *CreateAppStackConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create app stack config params
func (o *CreateAppStackConfigParams) WithAppID(appID string) *CreateAppStackConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create app stack config params
func (o *CreateAppStackConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReq adds the req to the create app stack config params
func (o *CreateAppStackConfigParams) WithReq(req *models.ServiceCreateAppStackConfigRequest) *CreateAppStackConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create app stack config params
func (o *CreateAppStackConfigParams) SetReq(req *models.ServiceCreateAppStackConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppStackConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
