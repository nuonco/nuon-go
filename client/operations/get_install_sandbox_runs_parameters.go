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

// NewGetInstallSandboxRunsParams creates a new GetInstallSandboxRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallSandboxRunsParams() *GetInstallSandboxRunsParams {
	return &GetInstallSandboxRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallSandboxRunsParamsWithTimeout creates a new GetInstallSandboxRunsParams object
// with the ability to set a timeout on a request.
func NewGetInstallSandboxRunsParamsWithTimeout(timeout time.Duration) *GetInstallSandboxRunsParams {
	return &GetInstallSandboxRunsParams{
		timeout: timeout,
	}
}

// NewGetInstallSandboxRunsParamsWithContext creates a new GetInstallSandboxRunsParams object
// with the ability to set a context for a request.
func NewGetInstallSandboxRunsParamsWithContext(ctx context.Context) *GetInstallSandboxRunsParams {
	return &GetInstallSandboxRunsParams{
		Context: ctx,
	}
}

// NewGetInstallSandboxRunsParamsWithHTTPClient creates a new GetInstallSandboxRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallSandboxRunsParamsWithHTTPClient(client *http.Client) *GetInstallSandboxRunsParams {
	return &GetInstallSandboxRunsParams{
		HTTPClient: client,
	}
}

/*
GetInstallSandboxRunsParams contains all the parameters to send to the API endpoint

	for the get install sandbox runs operation.

	Typically these are written to a http.Request.
*/
type GetInstallSandboxRunsParams struct {

	/* InstallID.

	   install ID
	*/
	InstallID string

	/* Limit.

	   limit of results to return

	   Default: 10
	*/
	Limit *int64

	/* Offset.

	   offset of results to return
	*/
	Offset *int64

	/* Page.

	   page number of results to return
	*/
	Page *int64

	/* XNuonPaginationEnabled.

	   Enable pagination
	*/
	XNuonPaginationEnabled *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install sandbox runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallSandboxRunsParams) WithDefaults() *GetInstallSandboxRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install sandbox runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallSandboxRunsParams) SetDefaults() {
	var (
		limitDefault = int64(10)

		offsetDefault = int64(0)

		pageDefault = int64(0)
	)

	val := GetInstallSandboxRunsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		Page:   &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithTimeout(timeout time.Duration) *GetInstallSandboxRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithContext(ctx context.Context) *GetInstallSandboxRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithHTTPClient(client *http.Client) *GetInstallSandboxRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstallID adds the installID to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithInstallID(installID string) *GetInstallSandboxRunsParams {
	o.SetInstallID(installID)
	return o
}

// SetInstallID adds the installId to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetInstallID(installID string) {
	o.InstallID = installID
}

// WithLimit adds the limit to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithLimit(limit *int64) *GetInstallSandboxRunsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithOffset(offset *int64) *GetInstallSandboxRunsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithPage(page *int64) *GetInstallSandboxRunsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetPage(page *int64) {
	o.Page = page
}

// WithXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) WithXNuonPaginationEnabled(xNuonPaginationEnabled *bool) *GetInstallSandboxRunsParams {
	o.SetXNuonPaginationEnabled(xNuonPaginationEnabled)
	return o
}

// SetXNuonPaginationEnabled adds the xNuonPaginationEnabled to the get install sandbox runs params
func (o *GetInstallSandboxRunsParams) SetXNuonPaginationEnabled(xNuonPaginationEnabled *bool) {
	o.XNuonPaginationEnabled = xNuonPaginationEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallSandboxRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param install_id
	if err := r.SetPathParam("install_id", o.InstallID); err != nil {
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

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
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
