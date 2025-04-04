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

// ServiceActionWorkflowLatestRunResponse service action workflow latest run response
//
// swagger:model service.ActionWorkflowLatestRunResponse
type ServiceActionWorkflowLatestRunResponse struct {

	// action workflow
	ActionWorkflow *AppActionWorkflow `json:"action_workflow,omitempty"`

	// install action workflow run
	InstallActionWorkflowRun *AppInstallActionWorkflowRun `json:"install_action_workflow_run,omitempty"`
}

// Validate validates this service action workflow latest run response
func (m *ServiceActionWorkflowLatestRunResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallActionWorkflowRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceActionWorkflowLatestRunResponse) validateActionWorkflow(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionWorkflow) { // not required
		return nil
	}

	if m.ActionWorkflow != nil {
		if err := m.ActionWorkflow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_workflow")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceActionWorkflowLatestRunResponse) validateInstallActionWorkflowRun(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallActionWorkflowRun) { // not required
		return nil
	}

	if m.InstallActionWorkflowRun != nil {
		if err := m.InstallActionWorkflowRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("install_action_workflow_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("install_action_workflow_run")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service action workflow latest run response based on the context it is used
func (m *ServiceActionWorkflowLatestRunResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionWorkflow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallActionWorkflowRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceActionWorkflowLatestRunResponse) contextValidateActionWorkflow(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionWorkflow != nil {

		if swag.IsZero(m.ActionWorkflow) { // not required
			return nil
		}

		if err := m.ActionWorkflow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_workflow")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceActionWorkflowLatestRunResponse) contextValidateInstallActionWorkflowRun(ctx context.Context, formats strfmt.Registry) error {

	if m.InstallActionWorkflowRun != nil {

		if swag.IsZero(m.InstallActionWorkflowRun) { // not required
			return nil
		}

		if err := m.InstallActionWorkflowRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("install_action_workflow_run")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("install_action_workflow_run")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceActionWorkflowLatestRunResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceActionWorkflowLatestRunResponse) UnmarshalBinary(b []byte) error {
	var res ServiceActionWorkflowLatestRunResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
