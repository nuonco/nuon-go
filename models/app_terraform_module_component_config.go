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

// AppTerraformModuleComponentConfig app terraform module component config
//
// swagger:model app.TerraformModuleComponentConfig
type AppTerraformModuleComponentConfig struct {

	// parent reference
	ComponentConfigConnectionID string `json:"component_config_connection_id,omitempty"`

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

	// VCSConfig
	PublicGitVcsConfig struct {
		AppPublicGitVCSConfig
	} `json:"public_git_vcs_config,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// variables
	Variables map[string]string `json:"variables,omitempty"`

	// terraform configuration values
	Version string `json:"version,omitempty"`
}

// Validate validates this app terraform module component config
func (m *AppTerraformModuleComponentConfig) Validate(formats strfmt.Registry) error {
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

func (m *AppTerraformModuleComponentConfig) validateConnectedGithubVcsConfig(formats strfmt.Registry) error {
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

func (m *AppTerraformModuleComponentConfig) validatePublicGitVcsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicGitVcsConfig) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this app terraform module component config based on the context it is used
func (m *AppTerraformModuleComponentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *AppTerraformModuleComponentConfig) contextValidateConnectedGithubVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppTerraformModuleComponentConfig) contextValidatePublicGitVcsConfig(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *AppTerraformModuleComponentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppTerraformModuleComponentConfig) UnmarshalBinary(b []byte) error {
	var res AppTerraformModuleComponentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
