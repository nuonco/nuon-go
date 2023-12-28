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

// NewPublishMetricsParams creates a new PublishMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublishMetricsParams() *PublishMetricsParams {
	return &PublishMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublishMetricsParamsWithTimeout creates a new PublishMetricsParams object
// with the ability to set a timeout on a request.
func NewPublishMetricsParamsWithTimeout(timeout time.Duration) *PublishMetricsParams {
	return &PublishMetricsParams{
		timeout: timeout,
	}
}

// NewPublishMetricsParamsWithContext creates a new PublishMetricsParams object
// with the ability to set a context for a request.
func NewPublishMetricsParamsWithContext(ctx context.Context) *PublishMetricsParams {
	return &PublishMetricsParams{
		Context: ctx,
	}
}

// NewPublishMetricsParamsWithHTTPClient creates a new PublishMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublishMetricsParamsWithHTTPClient(client *http.Client) *PublishMetricsParams {
	return &PublishMetricsParams{
		HTTPClient: client,
	}
}

/*
PublishMetricsParams contains all the parameters to send to the API endpoint

	for the publish metrics operation.

	Typically these are written to a http.Request.
*/
type PublishMetricsParams struct {

	/* Req.

	   Input
	*/
	Req []*models.ServicePublishMetricInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the publish metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublishMetricsParams) WithDefaults() *PublishMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the publish metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublishMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the publish metrics params
func (o *PublishMetricsParams) WithTimeout(timeout time.Duration) *PublishMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publish metrics params
func (o *PublishMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publish metrics params
func (o *PublishMetricsParams) WithContext(ctx context.Context) *PublishMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publish metrics params
func (o *PublishMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publish metrics params
func (o *PublishMetricsParams) WithHTTPClient(client *http.Client) *PublishMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publish metrics params
func (o *PublishMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the publish metrics params
func (o *PublishMetricsParams) WithReq(req []*models.ServicePublishMetricInput) *PublishMetricsParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the publish metrics params
func (o *PublishMetricsParams) SetReq(req []*models.ServicePublishMetricInput) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *PublishMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
