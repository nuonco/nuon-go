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

// AppApp app app
//
// swagger:model app.App
type AppApp struct {

	// cloud platform
	CloudPlatform AppCloudPlatform `json:"cloud_platform,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// fields set via after query
	InputConfig struct {
		AppAppInputConfig
	} `json:"input_config,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notifications config
	NotificationsConfig *AppNotificationsConfig `json:"notifications_config,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// runner config
	RunnerConfig *AppAppRunnerConfig `json:"runner_config,omitempty"`

	// sandbox config
	SandboxConfig *AppAppSandboxConfig `json:"sandbox_config,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app app
func (m *AppApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunnerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSandboxConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppApp) validateCloudPlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	if err := m.CloudPlatform.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloud_platform")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloud_platform")
		}
		return err
	}

	return nil
}

func (m *AppApp) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *AppApp) validateInputConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.InputConfig) { // not required
		return nil
	}

	return nil
}

func (m *AppApp) validateNotificationsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NotificationsConfig) { // not required
		return nil
	}

	if m.NotificationsConfig != nil {
		if err := m.NotificationsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppApp) validateRunnerConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.RunnerConfig) { // not required
		return nil
	}

	if m.RunnerConfig != nil {
		if err := m.RunnerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runner_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runner_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppApp) validateSandboxConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SandboxConfig) { // not required
		return nil
	}

	if m.SandboxConfig != nil {
		if err := m.SandboxConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sandbox_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sandbox_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app app based on the context it is used
func (m *AppApp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudPlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotificationsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunnerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSandboxConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppApp) contextValidateCloudPlatform(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	if err := m.CloudPlatform.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloud_platform")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloud_platform")
		}
		return err
	}

	return nil
}

func (m *AppApp) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *AppApp) contextValidateInputConfig(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *AppApp) contextValidateNotificationsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NotificationsConfig != nil {

		if swag.IsZero(m.NotificationsConfig) { // not required
			return nil
		}

		if err := m.NotificationsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppApp) contextValidateRunnerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.RunnerConfig != nil {

		if swag.IsZero(m.RunnerConfig) { // not required
			return nil
		}

		if err := m.RunnerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runner_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runner_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppApp) contextValidateSandboxConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SandboxConfig != nil {

		if swag.IsZero(m.SandboxConfig) { // not required
			return nil
		}

		if err := m.SandboxConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sandbox_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sandbox_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppApp) UnmarshalBinary(b []byte) error {
	var res AppApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
