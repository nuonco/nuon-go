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

// NewPostV1ReleasesParams creates a new PostV1ReleasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ReleasesParams() *PostV1ReleasesParams {
	return &PostV1ReleasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ReleasesParamsWithTimeout creates a new PostV1ReleasesParams object
// with the ability to set a timeout on a request.
func NewPostV1ReleasesParamsWithTimeout(timeout time.Duration) *PostV1ReleasesParams {
	return &PostV1ReleasesParams{
		timeout: timeout,
	}
}

// NewPostV1ReleasesParamsWithContext creates a new PostV1ReleasesParams object
// with the ability to set a context for a request.
func NewPostV1ReleasesParamsWithContext(ctx context.Context) *PostV1ReleasesParams {
	return &PostV1ReleasesParams{
		Context: ctx,
	}
}

// NewPostV1ReleasesParamsWithHTTPClient creates a new PostV1ReleasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ReleasesParamsWithHTTPClient(client *http.Client) *PostV1ReleasesParams {
	return &PostV1ReleasesParams{
		HTTPClient: client,
	}
}

/*
PostV1ReleasesParams contains all the parameters to send to the API endpoint

	for the post v1 releases operation.

	Typically these are written to a http.Request.
*/
type PostV1ReleasesParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateComponentReleaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ReleasesParams) WithDefaults() *PostV1ReleasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 releases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ReleasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 releases params
func (o *PostV1ReleasesParams) WithTimeout(timeout time.Duration) *PostV1ReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 releases params
func (o *PostV1ReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 releases params
func (o *PostV1ReleasesParams) WithContext(ctx context.Context) *PostV1ReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 releases params
func (o *PostV1ReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 releases params
func (o *PostV1ReleasesParams) WithHTTPClient(client *http.Client) *PostV1ReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 releases params
func (o *PostV1ReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the post v1 releases params
func (o *PostV1ReleasesParams) WithReq(req *models.ServiceCreateComponentReleaseRequest) *PostV1ReleasesParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the post v1 releases params
func (o *PostV1ReleasesParams) SetReq(req *models.ServiceCreateComponentReleaseRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
