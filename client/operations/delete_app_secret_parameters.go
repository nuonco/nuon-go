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

// NewDeleteAppSecretParams creates a new DeleteAppSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAppSecretParams() *DeleteAppSecretParams {
	return &DeleteAppSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppSecretParamsWithTimeout creates a new DeleteAppSecretParams object
// with the ability to set a timeout on a request.
func NewDeleteAppSecretParamsWithTimeout(timeout time.Duration) *DeleteAppSecretParams {
	return &DeleteAppSecretParams{
		timeout: timeout,
	}
}

// NewDeleteAppSecretParamsWithContext creates a new DeleteAppSecretParams object
// with the ability to set a context for a request.
func NewDeleteAppSecretParamsWithContext(ctx context.Context) *DeleteAppSecretParams {
	return &DeleteAppSecretParams{
		Context: ctx,
	}
}

// NewDeleteAppSecretParamsWithHTTPClient creates a new DeleteAppSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAppSecretParamsWithHTTPClient(client *http.Client) *DeleteAppSecretParams {
	return &DeleteAppSecretParams{
		HTTPClient: client,
	}
}

/*
DeleteAppSecretParams contains all the parameters to send to the API endpoint

	for the delete app secret operation.

	Typically these are written to a http.Request.
*/
type DeleteAppSecretParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* SecretID.

	   secret ID
	*/
	SecretID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete app secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppSecretParams) WithDefaults() *DeleteAppSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete app secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete app secret params
func (o *DeleteAppSecretParams) WithTimeout(timeout time.Duration) *DeleteAppSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete app secret params
func (o *DeleteAppSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete app secret params
func (o *DeleteAppSecretParams) WithContext(ctx context.Context) *DeleteAppSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete app secret params
func (o *DeleteAppSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete app secret params
func (o *DeleteAppSecretParams) WithHTTPClient(client *http.Client) *DeleteAppSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete app secret params
func (o *DeleteAppSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete app secret params
func (o *DeleteAppSecretParams) WithAppID(appID string) *DeleteAppSecretParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete app secret params
func (o *DeleteAppSecretParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithSecretID adds the secretID to the delete app secret params
func (o *DeleteAppSecretParams) WithSecretID(secretID string) *DeleteAppSecretParams {
	o.SetSecretID(secretID)
	return o
}

// SetSecretID adds the secretId to the delete app secret params
func (o *DeleteAppSecretParams) SetSecretID(secretID string) {
	o.SecretID = secretID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	// path param secret_id
	if err := r.SetPathParam("secret_id", o.SecretID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
