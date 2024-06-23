// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceCreateAppSandboxConfigRequest service create app sandbox config request
//
// swagger:model service.CreateAppSandboxConfigRequest
type ServiceCreateAppSandboxConfigRequest struct {

	// connected github vcs config
	ConnectedGithubVcsConfig *ServiceConnectedGithubVCSSandboxConfigRequest `json:"connected_github_vcs_config,omitempty"`

	// public git vcs config
	PublicGitVcsConfig *ServicePublicGitVCSSandboxConfigRequest `json:"public_git_vcs_config,omitempty"`

	// sandbox inputs
	// Required: true
	SandboxInputs map[string]string `json:"sandbox_inputs"`

	// terraform version
	// Required: true
	TerraformVersion *string `json:"terraform_version"`
}

// Validate validates this service create app sandbox config request
func (m *ServiceCreateAppSandboxConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedGithubVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicGitVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSandboxInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateAppSandboxConfigRequest) validateConnectedGithubVcsConfig(formats strfmt.Registry) error {
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

func (m *ServiceCreateAppSandboxConfigRequest) validatePublicGitVcsConfig(formats strfmt.Registry) error {
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

func (m *ServiceCreateAppSandboxConfigRequest) validateSandboxInputs(formats strfmt.Registry) error {

	if err := validate.Required("sandbox_inputs", "body", m.SandboxInputs); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateAppSandboxConfigRequest) validateTerraformVersion(formats strfmt.Registry) error {

	if err := validate.Required("terraform_version", "body", m.TerraformVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service create app sandbox config request based on the context it is used
func (m *ServiceCreateAppSandboxConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ServiceCreateAppSandboxConfigRequest) contextValidateConnectedGithubVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceCreateAppSandboxConfigRequest) contextValidatePublicGitVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ServiceCreateAppSandboxConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateAppSandboxConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateAppSandboxConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
