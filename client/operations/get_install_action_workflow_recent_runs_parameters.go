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
)

// NewGetInstallActionWorkflowRecentRunsParams creates a new GetInstallActionWorkflowRecentRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallActionWorkflowRecentRunsParams() *GetInstallActionWorkflowRecentRunsParams {
	return &GetInstallActionWorkflowRecentRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallActionWorkflowRecentRunsParamsWithTimeout creates a new GetInstallActionWorkflowRecentRunsParams object
// with the ability to set a timeout on a request.
func NewGetInstallActionWorkflowRecentRunsParamsWithTimeout(timeout time.Duration) *GetInstallActionWorkflowRecentRunsParams {
	return &GetInstallActionWorkflowRecentRunsParams{
		timeout: timeout,
	}
}

// NewGetInstallActionWorkflowRecentRunsParamsWithContext creates a new GetInstallActionWorkflowRecentRunsParams object
// with the ability to set a context for a request.
func NewGetInstallActionWorkflowRecentRunsParamsWithContext(ctx context.Context) *GetInstallActionWorkflowRecentRunsParams {
	return &GetInstallActionWorkflowRecentRunsParams{
		Context: ctx,
	}
}

// NewGetInstallActionWorkflowRecentRunsParamsWithHTTPClient creates a new GetInstallActionWorkflowRecentRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallActionWorkflowRecentRunsParamsWithHTTPClient(client *http.Client) *GetInstallActionWorkflowRecentRunsParams {
	return &GetInstallActionWorkflowRecentRunsParams{
		HTTPClient: client,
	}
}

/*
GetInstallActionWorkflowRecentRunsParams contains all the parameters to send to the API endpoint

	for the get install action workflow recent runs operation.

	Typically these are written to a http.Request.
*/
type GetInstallActionWorkflowRecentRunsParams struct {

	/* ActionWorkflowID.

	   action workflow ID
	*/
	ActionWorkflowID string

	/* InstallID.

	   install ID
	*/
	InstallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install action workflow recent runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallActionWorkflowRecentRunsParams) WithDefaults() *GetInstallActionWorkflowRecentRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install action workflow recent runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallActionWorkflowRecentRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) WithTimeout(timeout time.Duration) *GetInstallActionWorkflowRecentRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) WithContext(ctx context.Context) *GetInstallActionWorkflowRecentRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) WithHTTPClient(client *http.Client) *GetInstallActionWorkflowRecentRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionWorkflowID adds the actionWorkflowID to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) WithActionWorkflowID(actionWorkflowID string) *GetInstallActionWorkflowRecentRunsParams {
	o.SetActionWorkflowID(actionWorkflowID)
	return o
}

// SetActionWorkflowID adds the actionWorkflowId to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) SetActionWorkflowID(actionWorkflowID string) {
	o.ActionWorkflowID = actionWorkflowID
}

// WithInstallID adds the installID to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) WithInstallID(installID string) *GetInstallActionWorkflowRecentRunsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get install action workflow recent runs params
func (o *GetInstallActionWorkflowRecentRunsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallActionWorkflowRecentRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_workflow_id
	if err := r.SetPathParam("action_workflow_id", o.ActionWorkflowID); err != nil {
		return err
	}

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}