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
	"github.com/go-openapi/swag"
)

// NewGetActionWorkflowConfigsParams creates a new GetActionWorkflowConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetActionWorkflowConfigsParams() *GetActionWorkflowConfigsParams {
	return &GetActionWorkflowConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionWorkflowConfigsParamsWithTimeout creates a new GetActionWorkflowConfigsParams object
// with the ability to set a timeout on a request.
func NewGetActionWorkflowConfigsParamsWithTimeout(timeout time.Duration) *GetActionWorkflowConfigsParams {
	return &GetActionWorkflowConfigsParams{
		timeout: timeout,
	}
}

// NewGetActionWorkflowConfigsParamsWithContext creates a new GetActionWorkflowConfigsParams object
// with the ability to set a context for a request.
func NewGetActionWorkflowConfigsParamsWithContext(ctx context.Context) *GetActionWorkflowConfigsParams {
	return &GetActionWorkflowConfigsParams{
		Context: ctx,
	}
}

// NewGetActionWorkflowConfigsParamsWithHTTPClient creates a new GetActionWorkflowConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetActionWorkflowConfigsParamsWithHTTPClient(client *http.Client) *GetActionWorkflowConfigsParams {
	return &GetActionWorkflowConfigsParams{
		HTTPClient: client,
	}
}

/*
GetActionWorkflowConfigsParams contains all the parameters to send to the API endpoint

	for the get action workflow configs operation.

	Typically these are written to a http.Request.
*/
type GetActionWorkflowConfigsParams struct {

	/* ActionWorkflowID.

	   action workflow ID
	*/
	ActionWorkflowID string

	/* Limit.

	   limit of results to return

	   Default: 10
	*/
	Limit *int64

	/* Offset.

	   offset of results to return
	*/
	Offset *int64

	/* XNuonPaginationEnabled.

	   Enable pagination
	*/
	XNuonPaginationEnabled *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get action workflow configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionWorkflowConfigsParams) WithDefaults() *GetActionWorkflowConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get action workflow configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionWorkflowConfigsParams) SetDefaults() {
	var (
		limitDefault = int64(10)

		offsetDefault = int64(0)
	)

	val := GetActionWorkflowConfigsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithTimeout(timeout time.Duration) *GetActionWorkflowConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithContext(ctx context.Context) *GetActionWorkflowConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithHTTPClient(client *http.Client) *GetActionWorkflowConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionWorkflowID adds the actionWorkflowID to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithActionWorkflowID(actionWorkflowID string) *GetActionWorkflowConfigsParams {
	o.SetActionWorkflowID(actionWorkflowID)
	return o
}

// SetActionWorkflowID adds the actionWorkflowId to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetActionWorkflowID(actionWorkflowID string) {
	o.ActionWorkflowID = actionWorkflowID
}

// WithLimit adds the limit to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithLimit(limit *int64) *GetActionWorkflowConfigsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithOffset(offset *int64) *GetActionWorkflowConfigsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) WithXNuonPaginationEnabled(xNuonPaginationEnabled *bool) *GetActionWorkflowConfigsParams {
	o.SetXNuonPaginationEnabled(xNuonPaginationEnabled)
	return o
}

// SetXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get action workflow configs params
func (o *GetActionWorkflowConfigsParams) SetXNuonPaginationEnabled(xNuonPaginationEnabled *bool) {
	o.XNuonPaginationEnabled = xNuonPaginationEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionWorkflowConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action_workflow_id
	if err := r.SetPathParam("action_workflow_id", o.ActionWorkflowID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.XNuonPaginationEnabled != nil {

		// header param x-nuon-pagination-enabled
		if err := r.SetHeaderParam("x-nuon-pagination-enabled", swag.FormatBool(*o.XNuonPaginationEnabled)); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
