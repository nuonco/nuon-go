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

// NewGetOrgAcountsParams creates a new GetOrgAcountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgAcountsParams() *GetOrgAcountsParams {
	return &GetOrgAcountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgAcountsParamsWithTimeout creates a new GetOrgAcountsParams object
// with the ability to set a timeout on a request.
func NewGetOrgAcountsParamsWithTimeout(timeout time.Duration) *GetOrgAcountsParams {
	return &GetOrgAcountsParams{
		timeout: timeout,
	}
}

// NewGetOrgAcountsParamsWithContext creates a new GetOrgAcountsParams object
// with the ability to set a context for a request.
func NewGetOrgAcountsParamsWithContext(ctx context.Context) *GetOrgAcountsParams {
	return &GetOrgAcountsParams{
		Context: ctx,
	}
}

// NewGetOrgAcountsParamsWithHTTPClient creates a new GetOrgAcountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgAcountsParamsWithHTTPClient(client *http.Client) *GetOrgAcountsParams {
	return &GetOrgAcountsParams{
		HTTPClient: client,
	}
}

/*
GetOrgAcountsParams contains all the parameters to send to the API endpoint

	for the get org acounts operation.

	Typically these are written to a http.Request.
*/
type GetOrgAcountsParams struct {

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

// WithDefaults hydrates default values in the get org acounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgAcountsParams) WithDefaults() *GetOrgAcountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get org acounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgAcountsParams) SetDefaults() {
	var (
		limitDefault = int64(10)

		offsetDefault = int64(0)
	)

	val := GetOrgAcountsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get org acounts params
func (o *GetOrgAcountsParams) WithTimeout(timeout time.Duration) *GetOrgAcountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get org acounts params
func (o *GetOrgAcountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get org acounts params
func (o *GetOrgAcountsParams) WithContext(ctx context.Context) *GetOrgAcountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get org acounts params
func (o *GetOrgAcountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get org acounts params
func (o *GetOrgAcountsParams) WithHTTPClient(client *http.Client) *GetOrgAcountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get org acounts params
func (o *GetOrgAcountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get org acounts params
func (o *GetOrgAcountsParams) WithLimit(limit *int64) *GetOrgAcountsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get org acounts params
func (o *GetOrgAcountsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get org acounts params
func (o *GetOrgAcountsParams) WithOffset(offset *int64) *GetOrgAcountsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get org acounts params
func (o *GetOrgAcountsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get org acounts params
func (o *GetOrgAcountsParams) WithXNuonPaginationEnabled(xNuonPaginationEnabled *bool) *GetOrgAcountsParams {
	o.SetXNuonPaginationEnabled(xNuonPaginationEnabled)
	return o
}

// SetXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get org acounts params
func (o *GetOrgAcountsParams) SetXNuonPaginationEnabled(xNuonPaginationEnabled *bool) {
	o.XNuonPaginationEnabled = xNuonPaginationEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgAcountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
