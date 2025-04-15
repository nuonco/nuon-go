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

// NewCreateAppPoliciesConfigParams creates a new CreateAppPoliciesConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppPoliciesConfigParams() *CreateAppPoliciesConfigParams {
	return &CreateAppPoliciesConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppPoliciesConfigParamsWithTimeout creates a new CreateAppPoliciesConfigParams object
// with the ability to set a timeout on a request.
func NewCreateAppPoliciesConfigParamsWithTimeout(timeout time.Duration) *CreateAppPoliciesConfigParams {
	return &CreateAppPoliciesConfigParams{
		timeout: timeout,
	}
}

// NewCreateAppPoliciesConfigParamsWithContext creates a new CreateAppPoliciesConfigParams object
// with the ability to set a context for a request.
func NewCreateAppPoliciesConfigParamsWithContext(ctx context.Context) *CreateAppPoliciesConfigParams {
	return &CreateAppPoliciesConfigParams{
		Context: ctx,
	}
}

// NewCreateAppPoliciesConfigParamsWithHTTPClient creates a new CreateAppPoliciesConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppPoliciesConfigParamsWithHTTPClient(client *http.Client) *CreateAppPoliciesConfigParams {
	return &CreateAppPoliciesConfigParams{
		HTTPClient: client,
	}
}

/*
CreateAppPoliciesConfigParams contains all the parameters to send to the API endpoint

	for the create app policies config operation.

	Typically these are written to a http.Request.
*/
type CreateAppPoliciesConfigParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateAppPoliciesConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app policies config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppPoliciesConfigParams) WithDefaults() *CreateAppPoliciesConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app policies config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppPoliciesConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app policies config params
func (o *CreateAppPoliciesConfigParams) WithTimeout(timeout time.Duration) *CreateAppPoliciesConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app policies config params
func (o *CreateAppPoliciesConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app policies config params
func (o *CreateAppPoliciesConfigParams) WithContext(ctx context.Context) *CreateAppPoliciesConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app policies config params
func (o *CreateAppPoliciesConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app policies config params
func (o *CreateAppPoliciesConfigParams) WithHTTPClient(client *http.Client) *CreateAppPoliciesConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app policies config params
func (o *CreateAppPoliciesConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create app policies config params
func (o *CreateAppPoliciesConfigParams) WithAppID(appID string) *CreateAppPoliciesConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create app policies config params
func (o *CreateAppPoliciesConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReq adds the req to the create app policies config params
func (o *CreateAppPoliciesConfigParams) WithReq(req *models.ServiceCreateAppPoliciesConfigRequest) *CreateAppPoliciesConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create app policies config params
func (o *CreateAppPoliciesConfigParams) SetReq(req *models.ServiceCreateAppPoliciesConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppPoliciesConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
