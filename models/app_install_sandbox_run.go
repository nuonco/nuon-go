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

// AppInstallSandboxRun app install sandbox run
//
// swagger:model app.InstallSandboxRun
type AppInstallSandboxRun struct {

	// app sandbox config
	AppSandboxConfig *AppAppSandboxConfig `json:"app_sandbox_config,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install id
	InstallID string `json:"install_id,omitempty"`

	// run type
	RunType AppSandboxRunType `json:"run_type,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install sandbox run
func (m *AppInstallSandboxRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppSandboxConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallSandboxRun) validateAppSandboxConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AppSandboxConfig) { // not required
		return nil
	}

	if m.AppSandboxConfig != nil {
		if err := m.AppSandboxConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_sandbox_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app_sandbox_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallSandboxRun) validateRunType(formats strfmt.Registry) error {
	if swag.IsZero(m.RunType) { // not required
		return nil
	}

	if err := m.RunType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("run_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("run_type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app install sandbox run based on the context it is used
func (m *AppInstallSandboxRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppSandboxConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallSandboxRun) contextValidateAppSandboxConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AppSandboxConfig != nil {

		if swag.IsZero(m.AppSandboxConfig) { // not required
			return nil
		}

		if err := m.AppSandboxConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_sandbox_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app_sandbox_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallSandboxRun) contextValidateRunType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.RunType) { // not required
		return nil
	}

	if err := m.RunType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("run_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("run_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallSandboxRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallSandboxRun) UnmarshalBinary(b []byte) error {
	var res AppInstallSandboxRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
