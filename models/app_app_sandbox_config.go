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

	// artifacts
	Artifacts *AppAppSandboxConfigArtifacts `json:"artifacts,omitempty"`

	// cloud specific fields
	AwsDelegationConfig struct {
		AppAppAWSDelegationConfig
	} `json:"aws_delegation_config,omitempty"`

	// aws region type
	AwsRegionType string `json:"aws_region_type,omitempty"`

	// fields set via after query
	CloudPlatform string `json:"cloud_platform,omitempty"`

	// connected github vcs config
	ConnectedGithubVcsConfig *AppConnectedGithubVCSConfig `json:"connected_github_vcs_config,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

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
}

// Validate validates this app app sandbox config
func (m *AppAppSandboxConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsDelegationConfig(formats); err != nil {
		res = append(res, err)
	}

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

func (m *AppAppSandboxConfig) validateArtifacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Artifacts) { // not required
		return nil
	}

	if m.Artifacts != nil {
		if err := m.Artifacts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifacts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifacts")
			}
			return err
		}
	}

	return nil
}

func (m *AppAppSandboxConfig) validateAwsDelegationConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsDelegationConfig) { // not required
		return nil
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

	if err := m.contextValidateArtifacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsDelegationConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *AppAppSandboxConfig) contextValidateArtifacts(ctx context.Context, formats strfmt.Registry) error {

	if m.Artifacts != nil {

		if swag.IsZero(m.Artifacts) { // not required
			return nil
		}

		if err := m.Artifacts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifacts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("artifacts")
			}
			return err
		}
	}

	return nil
}

func (m *AppAppSandboxConfig) contextValidateAwsDelegationConfig(ctx context.Context, formats strfmt.Registry) error {

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

// AppAppSandboxConfigArtifacts Links are dynamically loaded using an after query
//
// swagger:model AppAppSandboxConfigArtifacts
type AppAppSandboxConfigArtifacts struct {

	// cloudformation stack template
	CloudformationStackTemplate string `json:"cloudformation_stack_template,omitempty"`

	// deprovision policy
	DeprovisionPolicy string `json:"deprovision_policy,omitempty"`

	// provision policy
	ProvisionPolicy string `json:"provision_policy,omitempty"`

	// trust policy
	TrustPolicy string `json:"trust_policy,omitempty"`
}

// Validate validates this app app sandbox config artifacts
func (m *AppAppSandboxConfigArtifacts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app app sandbox config artifacts based on context it is used
func (m *AppAppSandboxConfigArtifacts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppAppSandboxConfigArtifacts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppSandboxConfigArtifacts) UnmarshalBinary(b []byte) error {
	var res AppAppSandboxConfigArtifacts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
