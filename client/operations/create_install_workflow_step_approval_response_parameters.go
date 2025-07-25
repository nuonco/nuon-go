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

// NewCreateInstallWorkflowStepApprovalResponseParams creates a new CreateInstallWorkflowStepApprovalResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInstallWorkflowStepApprovalResponseParams() *CreateInstallWorkflowStepApprovalResponseParams {
	return &CreateInstallWorkflowStepApprovalResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstallWorkflowStepApprovalResponseParamsWithTimeout creates a new CreateInstallWorkflowStepApprovalResponseParams object
// with the ability to set a timeout on a request.
func NewCreateInstallWorkflowStepApprovalResponseParamsWithTimeout(timeout time.Duration) *CreateInstallWorkflowStepApprovalResponseParams {
	return &CreateInstallWorkflowStepApprovalResponseParams{
		timeout: timeout,
	}
}

// NewCreateInstallWorkflowStepApprovalResponseParamsWithContext creates a new CreateInstallWorkflowStepApprovalResponseParams object
// with the ability to set a context for a request.
func NewCreateInstallWorkflowStepApprovalResponseParamsWithContext(ctx context.Context) *CreateInstallWorkflowStepApprovalResponseParams {
	return &CreateInstallWorkflowStepApprovalResponseParams{
		Context: ctx,
	}
}

// NewCreateInstallWorkflowStepApprovalResponseParamsWithHTTPClient creates a new CreateInstallWorkflowStepApprovalResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInstallWorkflowStepApprovalResponseParamsWithHTTPClient(client *http.Client) *CreateInstallWorkflowStepApprovalResponseParams {
	return &CreateInstallWorkflowStepApprovalResponseParams{
		HTTPClient: client,
	}
}

/*
CreateInstallWorkflowStepApprovalResponseParams contains all the parameters to send to the API endpoint

	for the create install workflow step approval response operation.

	Typically these are written to a http.Request.
*/
type CreateInstallWorkflowStepApprovalResponseParams struct {

	/* ApprovalID.

	   approval id
	*/
	ApprovalID string

	/* InstallWorkflowID.

	   workflow id
	*/
	InstallWorkflowID string

	/* InstallWorkflowStepID.

	   step id
	*/
	InstallWorkflowStepID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateWorkflowStepApprovalResponseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create install workflow step approval response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithDefaults() *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create install workflow step approval response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithTimeout(timeout time.Duration) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithContext(ctx context.Context) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithHTTPClient(client *http.Client) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApprovalID adds the approvalID to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithApprovalID(approvalID string) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetApprovalID(approvalID)
	return o
}

// SetApprovalID adds the approvalId to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetApprovalID(approvalID string) {
	o.ApprovalID = approvalID
}

// WithInstallWorkflowID adds the installWorkflowID to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithInstallWorkflowID(installWorkflowID string) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetInstallWorkflowID(installWorkflowID)
	return o
}

// SetInstallWorkflowID adds the installWorkflowId to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetInstallWorkflowID(installWorkflowID string) {
	o.InstallWorkflowID = installWorkflowID
}

// WithInstallWorkflowStepID adds the installWorkflowStepID to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithInstallWorkflowStepID(installWorkflowStepID string) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetInstallWorkflowStepID(installWorkflowStepID)
	return o
}

// SetInstallWorkflowStepID adds the installWorkflowStepId to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetInstallWorkflowStepID(installWorkflowStepID string) {
	o.InstallWorkflowStepID = installWorkflowStepID
}

// WithReq adds the req to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) WithReq(req *models.ServiceCreateWorkflowStepApprovalResponseRequest) *CreateInstallWorkflowStepApprovalResponseParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create install workflow step approval response params
func (o *CreateInstallWorkflowStepApprovalResponseParams) SetReq(req *models.ServiceCreateWorkflowStepApprovalResponseRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstallWorkflowStepApprovalResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param approval_id
	if err := r.SetPathParam("approval_id", o.ApprovalID); err != nil {
		return err
	}

	// path param install_workflow_id
	if err := r.SetPathParam("install_workflow_id", o.InstallWorkflowID); err != nil {
		return err
	}

	// path param install_workflow_step_id
	if err := r.SetPathParam("install_workflow_step_id", o.InstallWorkflowStepID); err != nil {
		return err
	}
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
