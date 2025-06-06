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

// NewCreateAppCloudFormationStackConfigParams creates a new CreateAppCloudFormationStackConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppCloudFormationStackConfigParams() *CreateAppCloudFormationStackConfigParams {
	return &CreateAppCloudFormationStackConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppCloudFormationStackConfigParamsWithTimeout creates a new CreateAppCloudFormationStackConfigParams object
// with the ability to set a timeout on a request.
func NewCreateAppCloudFormationStackConfigParamsWithTimeout(timeout time.Duration) *CreateAppCloudFormationStackConfigParams {
	return &CreateAppCloudFormationStackConfigParams{
		timeout: timeout,
	}
}

// NewCreateAppCloudFormationStackConfigParamsWithContext creates a new CreateAppCloudFormationStackConfigParams object
// with the ability to set a context for a request.
func NewCreateAppCloudFormationStackConfigParamsWithContext(ctx context.Context) *CreateAppCloudFormationStackConfigParams {
	return &CreateAppCloudFormationStackConfigParams{
		Context: ctx,
	}
}

// NewCreateAppCloudFormationStackConfigParamsWithHTTPClient creates a new CreateAppCloudFormationStackConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppCloudFormationStackConfigParamsWithHTTPClient(client *http.Client) *CreateAppCloudFormationStackConfigParams {
	return &CreateAppCloudFormationStackConfigParams{
		HTTPClient: client,
	}
}

/*
CreateAppCloudFormationStackConfigParams contains all the parameters to send to the API endpoint

	for the create app cloud formation stack config operation.

	Typically these are written to a http.Request.
*/
type CreateAppCloudFormationStackConfigParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateAppCloudFormationStackConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app cloud formation stack config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppCloudFormationStackConfigParams) WithDefaults() *CreateAppCloudFormationStackConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app cloud formation stack config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppCloudFormationStackConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) WithTimeout(timeout time.Duration) *CreateAppCloudFormationStackConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) WithContext(ctx context.Context) *CreateAppCloudFormationStackConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) WithHTTPClient(client *http.Client) *CreateAppCloudFormationStackConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) WithAppID(appID string) *CreateAppCloudFormationStackConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReq adds the req to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) WithReq(req *models.ServiceCreateAppCloudFormationStackConfigRequest) *CreateAppCloudFormationStackConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create app cloud formation stack config params
func (o *CreateAppCloudFormationStackConfigParams) SetReq(req *models.ServiceCreateAppCloudFormationStackConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppCloudFormationStackConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
