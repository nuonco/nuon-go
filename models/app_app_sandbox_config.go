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

// AppAppSandboxConfig app app sandbox config
//
// swagger:model app.AppSandboxConfig
type AppAppSandboxConfig struct {

	// app config id
	AppConfigID string `json:"app_config_id,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// cloud specific fields
	AwsRegionType string `json:"aws_region_type,omitempty"`

	// fields set via after query
	CloudPlatform string `json:"cloud_platform,omitempty"`

	// connected github vcs config
	ConnectedGithubVcsConfig *AppConnectedGithubVCSConfig `json:"connected_github_vcs_config,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// env vars
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// public git vcs config
	PublicGitVcsConfig *AppPublicGitVCSConfig `json:"public_git_vcs_config,omitempty"`

	// terraform version
	TerraformVersion string `json:"terraform_version,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`

	// variables files
	VariablesFiles []string `json:"variables_files"`
}

// Validate validates this app app sandbox config
func (m *AppAppSandboxConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedGithubVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicGitVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppSandboxConfig) validateConnectedGithubVcsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedGithubVcsConfig) { // not required
		return nil
	}

	if m.ConnectedGithubVcsConfig != nil {
		if err := m.ConnectedGithubVcsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connected_github_vcs_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connected_github_vcs_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppAppSandboxConfig) validatePublicGitVcsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicGitVcsConfig) { // not required
		return nil
	}

	if m.PublicGitVcsConfig != nil {
		if err := m.PublicGitVcsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_git_vcs_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_git_vcs_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app app sandbox config based on the context it is used
func (m *AppAppSandboxConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectedGithubVcsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicGitVcsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppSandboxConfig) contextValidateConnectedGithubVcsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectedGithubVcsConfig != nil {

		if swag.IsZero(m.ConnectedGithubVcsConfig) { // not required
			return nil
		}

		if err := m.ConnectedGithubVcsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connected_github_vcs_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connected_github_vcs_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppAppSandboxConfig) contextValidatePublicGitVcsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicGitVcsConfig != nil {

		if swag.IsZero(m.PublicGitVcsConfig) { // not required
			return nil
		}

		if err := m.PublicGitVcsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_git_vcs_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_git_vcs_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppAppSandboxConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppSandboxConfig) UnmarshalBinary(b []byte) error {
	var res AppAppSandboxConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
