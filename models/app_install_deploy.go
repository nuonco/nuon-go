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

// AppInstallDeploy app install deploy
//
// swagger:model app.InstallDeploy
type AppInstallDeploy struct {

	// build id
	BuildID string `json:"build_id,omitempty"`

	// component config version
	ComponentConfigVersion int64 `json:"component_config_version,omitempty"`

	// component id
	ComponentID string `json:"component_id,omitempty"`

	// component name
	ComponentName string `json:"component_name,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install component id
	InstallComponentID string `json:"install_component_id,omitempty"`

	// install deploy type
	InstallDeployType AppInstallDeployType `json:"install_deploy_type,omitempty"`

	// Fields that are de-nested at read time using AfterQuery
	InstallID string `json:"install_id,omitempty"`

	// log stream
	LogStream *AppLogStream `json:"log_stream,omitempty"`

	// release id
	ReleaseID string `json:"release_id,omitempty"`

	// runner details
	RunnerJob struct {
		AppRunnerJob
	} `json:"runner_job,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install deploy
func (m *AppInstallDeploy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstallDeployType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogStream(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunnerJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallDeploy) validateInstallDeployType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallDeployType) { // not required
		return nil
	}

	if err := m.InstallDeployType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("install_deploy_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("install_deploy_type")
		}
		return err
	}

	return nil
}

func (m *AppInstallDeploy) validateLogStream(formats strfmt.Registry) error {
	if swag.IsZero(m.LogStream) { // not required
		return nil
	}

	if m.LogStream != nil {
		if err := m.LogStream.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_stream")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_stream")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallDeploy) validateRunnerJob(formats strfmt.Registry) error {
	if swag.IsZero(m.RunnerJob) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this app install deploy based on the context it is used
func (m *AppInstallDeploy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstallDeployType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogStream(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunnerJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallDeploy) contextValidateInstallDeployType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.InstallDeployType) { // not required
		return nil
	}

	if err := m.InstallDeployType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("install_deploy_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("install_deploy_type")
		}
		return err
	}

	return nil
}

func (m *AppInstallDeploy) contextValidateLogStream(ctx context.Context, formats strfmt.Registry) error {

	if m.LogStream != nil {

		if swag.IsZero(m.LogStream) { // not required
			return nil
		}

		if err := m.LogStream.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_stream")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_stream")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallDeploy) contextValidateRunnerJob(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallDeploy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallDeploy) UnmarshalBinary(b []byte) error {
	var res AppInstallDeploy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
