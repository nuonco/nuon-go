// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppInstallWorkflowStepPolicyValidation app install workflow step policy validation
//
// swagger:model app.InstallWorkflowStepPolicyValidation
type AppInstallWorkflowStepPolicyValidation struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install workflow step is the install step that this was performed within
	InstallWorkflowStepID string `json:"installWorkflowStepID,omitempty"`

	// response is the kyverno response
	Response string `json:"response,omitempty"`

	// runnerJobID is the runner job that this was performed within
	RunnerJobID string `json:"runnerJobID,omitempty"`

	// status denotes whether this passed, or whether it failed the job.
	Status struct {
		AppCompositeStatus
	} `json:"status,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install workflow step policy validation
func (m *AppInstallWorkflowStepPolicyValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallWorkflowStepPolicyValidation) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this app install workflow step policy validation based on the context it is used
func (m *AppInstallWorkflowStepPolicyValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallWorkflowStepPolicyValidation) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallWorkflowStepPolicyValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallWorkflowStepPolicyValidation) UnmarshalBinary(b []byte) error {
	var res AppInstallWorkflowStepPolicyValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
