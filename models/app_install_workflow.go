// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppInstallWorkflow app install workflow
//
// swagger:model app.InstallWorkflow
type AppInstallWorkflow struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// execution time
	ExecutionTime int64 `json:"execution_time,omitempty"`

	// finished at
	FinishedAt string `json:"finished_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install id
	InstallID string `json:"install_id,omitempty"`

	// metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// started at
	StartedAt string `json:"started_at,omitempty"`

	// status
	Status *AppCompositeStatus `json:"status,omitempty"`

	// step error behavior
	StepErrorBehavior AppStepErrorBehavior `json:"step_error_behavior,omitempty"`

	// steps represent each piece of the workflow
	Steps []*AppInstallWorkflowStep `json:"steps"`

	// type
	Type AppInstallWorkflowType `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install workflow
func (m *AppInstallWorkflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepErrorBehavior(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallWorkflow) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallWorkflow) validateStepErrorBehavior(formats strfmt.Registry) error {
	if swag.IsZero(m.StepErrorBehavior) { // not required
		return nil
	}

	if err := m.StepErrorBehavior.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("step_error_behavior")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("step_error_behavior")
		}
		return err
	}

	return nil
}

func (m *AppInstallWorkflow) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstallWorkflow) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app install workflow based on the context it is used
func (m *AppInstallWorkflow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStepErrorBehavior(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallWorkflow) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallWorkflow) contextValidateStepErrorBehavior(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.StepErrorBehavior) { // not required
		return nil
	}

	if err := m.StepErrorBehavior.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("step_error_behavior")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("step_error_behavior")
		}
		return err
	}

	return nil
}

func (m *AppInstallWorkflow) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {

			if swag.IsZero(m.Steps[i]) { // not required
				return nil
			}

			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstallWorkflow) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallWorkflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallWorkflow) UnmarshalBinary(b []byte) error {
	var res AppInstallWorkflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
