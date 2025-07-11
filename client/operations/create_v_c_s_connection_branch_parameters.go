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

// NewCreateVCSConnectionBranchParams creates a new CreateVCSConnectionBranchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVCSConnectionBranchParams() *CreateVCSConnectionBranchParams {
	return &CreateVCSConnectionBranchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVCSConnectionBranchParamsWithTimeout creates a new CreateVCSConnectionBranchParams object
// with the ability to set a timeout on a request.
func NewCreateVCSConnectionBranchParamsWithTimeout(timeout time.Duration) *CreateVCSConnectionBranchParams {
	return &CreateVCSConnectionBranchParams{
		timeout: timeout,
	}
}

// NewCreateVCSConnectionBranchParamsWithContext creates a new CreateVCSConnectionBranchParams object
// with the ability to set a context for a request.
func NewCreateVCSConnectionBranchParamsWithContext(ctx context.Context) *CreateVCSConnectionBranchParams {
	return &CreateVCSConnectionBranchParams{
		Context: ctx,
	}
}

// NewCreateVCSConnectionBranchParamsWithHTTPClient creates a new CreateVCSConnectionBranchParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVCSConnectionBranchParamsWithHTTPClient(client *http.Client) *CreateVCSConnectionBranchParams {
	return &CreateVCSConnectionBranchParams{
		HTTPClient: client,
	}
}

/*
CreateVCSConnectionBranchParams contains all the parameters to send to the API endpoint

	for the create v c s connection branch operation.

	Typically these are written to a http.Request.
*/
type CreateVCSConnectionBranchParams struct {

	/* ConnectionID.

	   connection ID
	*/
	ConnectionID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateConnectionBranchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create v c s connection branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVCSConnectionBranchParams) WithDefaults() *CreateVCSConnectionBranchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create v c s connection branch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVCSConnectionBranchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) WithTimeout(timeout time.Duration) *CreateVCSConnectionBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) WithContext(ctx context.Context) *CreateVCSConnectionBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) WithHTTPClient(client *http.Client) *CreateVCSConnectionBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) WithConnectionID(connectionID string) *CreateVCSConnectionBranchParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WithReq adds the req to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) WithReq(req *models.ServiceCreateConnectionBranchRequest) *CreateVCSConnectionBranchParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create v c s connection branch params
func (o *CreateVCSConnectionBranchParams) SetReq(req *models.ServiceCreateConnectionBranchRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVCSConnectionBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
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
