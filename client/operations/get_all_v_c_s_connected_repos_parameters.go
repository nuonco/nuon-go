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

// NewGetAllVCSConnectedReposParams creates a new GetAllVCSConnectedReposParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllVCSConnectedReposParams() *GetAllVCSConnectedReposParams {
	return &GetAllVCSConnectedReposParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllVCSConnectedReposParamsWithTimeout creates a new GetAllVCSConnectedReposParams object
// with the ability to set a timeout on a request.
func NewGetAllVCSConnectedReposParamsWithTimeout(timeout time.Duration) *GetAllVCSConnectedReposParams {
	return &GetAllVCSConnectedReposParams{
		timeout: timeout,
	}
}

// NewGetAllVCSConnectedReposParamsWithContext creates a new GetAllVCSConnectedReposParams object
// with the ability to set a context for a request.
func NewGetAllVCSConnectedReposParamsWithContext(ctx context.Context) *GetAllVCSConnectedReposParams {
	return &GetAllVCSConnectedReposParams{
		Context: ctx,
	}
}

// NewGetAllVCSConnectedReposParamsWithHTTPClient creates a new GetAllVCSConnectedReposParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllVCSConnectedReposParamsWithHTTPClient(client *http.Client) *GetAllVCSConnectedReposParams {
	return &GetAllVCSConnectedReposParams{
		HTTPClient: client,
	}
}

/*
GetAllVCSConnectedReposParams contains all the parameters to send to the API endpoint

	for the get all v c s connected repos operation.

	Typically these are written to a http.Request.
*/
type GetAllVCSConnectedReposParams struct {

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

// WithDefaults hydrates default values in the get all v c s connected repos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllVCSConnectedReposParams) WithDefaults() *GetAllVCSConnectedReposParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all v c s connected repos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllVCSConnectedReposParams) SetDefaults() {
	var (
		limitDefault = int64(10)

		offsetDefault = int64(0)
	)

	val := GetAllVCSConnectedReposParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) WithTimeout(timeout time.Duration) *GetAllVCSConnectedReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) WithContext(ctx context.Context) *GetAllVCSConnectedReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) WithHTTPClient(client *http.Client) *GetAllVCSConnectedReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) WithLimit(limit *int64) *GetAllVCSConnectedReposParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) WithOffset(offset *int64) *GetAllVCSConnectedReposParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) WithXNuonPaginationEnabled(xNuonPaginationEnabled *bool) *GetAllVCSConnectedReposParams {
	o.SetXNuonPaginationEnabled(xNuonPaginationEnabled)
	return o
}

// SetXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get all v c s connected repos params
func (o *GetAllVCSConnectedReposParams) SetXNuonPaginationEnabled(xNuonPaginationEnabled *bool) {
	o.XNuonPaginationEnabled = xNuonPaginationEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllVCSConnectedReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
