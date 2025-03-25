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

// NewGetActionWorkflowsParams creates a new GetActionWorkflowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetActionWorkflowsParams() *GetActionWorkflowsParams {
	return &GetActionWorkflowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetActionWorkflowsParamsWithTimeout creates a new GetActionWorkflowsParams object
// with the ability to set a timeout on a request.
func NewGetActionWorkflowsParamsWithTimeout(timeout time.Duration) *GetActionWorkflowsParams {
	return &GetActionWorkflowsParams{
		timeout: timeout,
	}
}

// NewGetActionWorkflowsParamsWithContext creates a new GetActionWorkflowsParams object
// with the ability to set a context for a request.
func NewGetActionWorkflowsParamsWithContext(ctx context.Context) *GetActionWorkflowsParams {
	return &GetActionWorkflowsParams{
		Context: ctx,
	}
}

// NewGetActionWorkflowsParamsWithHTTPClient creates a new GetActionWorkflowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetActionWorkflowsParamsWithHTTPClient(client *http.Client) *GetActionWorkflowsParams {
	return &GetActionWorkflowsParams{
		HTTPClient: client,
	}
}

/*
GetActionWorkflowsParams contains all the parameters to send to the API endpoint

	for the get action workflows operation.

	Typically these are written to a http.Request.
*/
type GetActionWorkflowsParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

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

// WithDefaults hydrates default values in the get action workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionWorkflowsParams) WithDefaults() *GetActionWorkflowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get action workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetActionWorkflowsParams) SetDefaults() {
	var (
		limitDefault = int64(10)

		offsetDefault = int64(0)
	)

	val := GetActionWorkflowsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get action workflows params
func (o *GetActionWorkflowsParams) WithTimeout(timeout time.Duration) *GetActionWorkflowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get action workflows params
func (o *GetActionWorkflowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get action workflows params
func (o *GetActionWorkflowsParams) WithContext(ctx context.Context) *GetActionWorkflowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get action workflows params
func (o *GetActionWorkflowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get action workflows params
func (o *GetActionWorkflowsParams) WithHTTPClient(client *http.Client) *GetActionWorkflowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get action workflows params
func (o *GetActionWorkflowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get action workflows params
func (o *GetActionWorkflowsParams) WithAppID(appID string) *GetActionWorkflowsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get action workflows params
func (o *GetActionWorkflowsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithLimit adds the limit to the get action workflows params
func (o *GetActionWorkflowsParams) WithLimit(limit *int64) *GetActionWorkflowsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get action workflows params
func (o *GetActionWorkflowsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get action workflows params
func (o *GetActionWorkflowsParams) WithOffset(offset *int64) *GetActionWorkflowsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get action workflows params
func (o *GetActionWorkflowsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get action workflows params
func (o *GetActionWorkflowsParams) WithXNuonPaginationEnabled(xNuonPaginationEnabled *bool) *GetActionWorkflowsParams {
	o.SetXNuonPaginationEnabled(xNuonPaginationEnabled)
	return o
}

// SetXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get action workflows params
func (o *GetActionWorkflowsParams) SetXNuonPaginationEnabled(xNuonPaginationEnabled *bool) {
	o.XNuonPaginationEnabled = xNuonPaginationEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetActionWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
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
