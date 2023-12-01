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

// ServiceCreateDockerBuildComponentConfigRequest service create docker build component config request
//
// swagger:model service.CreateDockerBuildComponentConfigRequest
type ServiceCreateDockerBuildComponentConfigRequest struct {

	// build args
	BuildArgs []string `json:"build_args"`

	// connected github vcs config
	ConnectedGithubVcsConfig *ServiceConnectedGithubVCSConfigRequest `json:"connected_github_vcs_config,omitempty"`

	// dockerfile
	// Required: true
	Dockerfile *string `json:"dockerfile"`

	// env vars
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// public git vcs config
	PublicGitVcsConfig *ServicePublicGitVCSConfigRequest `json:"public_git_vcs_config,omitempty"`

	// target
	Target string `json:"target,omitempty"`
}

// Validate validates this service create docker build component config request
func (m *ServiceCreateDockerBuildComponentConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedGithubVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerfile(formats); err != nil {
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

func (m *ServiceCreateDockerBuildComponentConfigRequest) validateConnectedGithubVcsConfig(formats strfmt.Registry) error {
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

func (m *ServiceCreateDockerBuildComponentConfigRequest) validateDockerfile(formats strfmt.Registry) error {

	if err := validate.Required("dockerfile", "body", m.Dockerfile); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateDockerBuildComponentConfigRequest) validatePublicGitVcsConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this service create docker build component config request based on the context it is used
func (m *ServiceCreateDockerBuildComponentConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ServiceCreateDockerBuildComponentConfigRequest) contextValidateConnectedGithubVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceCreateDockerBuildComponentConfigRequest) contextValidatePublicGitVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ServiceCreateDockerBuildComponentConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateDockerBuildComponentConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateDockerBuildComponentConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
