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

// NewPostV1VcsConnectionsParams creates a new PostV1VcsConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1VcsConnectionsParams() *PostV1VcsConnectionsParams {
	return &PostV1VcsConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1VcsConnectionsParamsWithTimeout creates a new PostV1VcsConnectionsParams object
// with the ability to set a timeout on a request.
func NewPostV1VcsConnectionsParamsWithTimeout(timeout time.Duration) *PostV1VcsConnectionsParams {
	return &PostV1VcsConnectionsParams{
		timeout: timeout,
	}
}

// NewPostV1VcsConnectionsParamsWithContext creates a new PostV1VcsConnectionsParams object
// with the ability to set a context for a request.
func NewPostV1VcsConnectionsParamsWithContext(ctx context.Context) *PostV1VcsConnectionsParams {
	return &PostV1VcsConnectionsParams{
		Context: ctx,
	}
}

// NewPostV1VcsConnectionsParamsWithHTTPClient creates a new PostV1VcsConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1VcsConnectionsParamsWithHTTPClient(client *http.Client) *PostV1VcsConnectionsParams {
	return &PostV1VcsConnectionsParams{
		HTTPClient: client,
	}
}

/*
PostV1VcsConnectionsParams contains all the parameters to send to the API endpoint

	for the post v1 vcs connections operation.

	Typically these are written to a http.Request.
*/
type PostV1VcsConnectionsParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateConnectionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 vcs connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1VcsConnectionsParams) WithDefaults() *PostV1VcsConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 vcs connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1VcsConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) WithTimeout(timeout time.Duration) *PostV1VcsConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) WithContext(ctx context.Context) *PostV1VcsConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) WithHTTPClient(client *http.Client) *PostV1VcsConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) WithReq(req *models.ServiceCreateConnectionRequest) *PostV1VcsConnectionsParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the post v1 vcs connections params
func (o *PostV1VcsConnectionsParams) SetReq(req *models.ServiceCreateConnectionRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1VcsConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
