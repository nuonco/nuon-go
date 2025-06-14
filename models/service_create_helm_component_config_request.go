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

// ServiceCreateHelmComponentConfigRequest service create helm component config request
//
// swagger:model service.CreateHelmComponentConfigRequest
type ServiceCreateHelmComponentConfigRequest struct {

	// app config id
	AppConfigID string `json:"app_config_id,omitempty"`

	// chart name
	// Required: true
	// Max Length: 62
	// Min Length: 5
	ChartName *string `json:"chart_name"`

	// checksum
	Checksum string `json:"checksum,omitempty"`

	// connected github vcs config
	ConnectedGithubVcsConfig *ServiceConnectedGithubVCSConfigRequest `json:"connected_github_vcs_config,omitempty"`

	// dependencies
	Dependencies []string `json:"dependencies"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// public git vcs config
	PublicGitVcsConfig *ServicePublicGitVCSConfigRequest `json:"public_git_vcs_config,omitempty"`

	// references
	References []string `json:"references"`

	// storage driver
	StorageDriver string `json:"storage_driver,omitempty"`

	// take ownership
	TakeOwnership bool `json:"take_ownership,omitempty"`

	// values
	// Required: true
	Values map[string]string `json:"values"`

	// values files
	ValuesFiles []string `json:"values_files"`
}

// Validate validates this service create helm component config request
func (m *ServiceCreateHelmComponentConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedGithubVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicGitVcsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateHelmComponentConfigRequest) validateChartName(formats strfmt.Registry) error {

	if err := validate.Required("chart_name", "body", m.ChartName); err != nil {
		return err
	}

	if err := validate.MinLength("chart_name", "body", *m.ChartName, 5); err != nil {
		return err
	}

	if err := validate.MaxLength("chart_name", "body", *m.ChartName, 62); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateHelmComponentConfigRequest) validateConnectedGithubVcsConfig(formats strfmt.Registry) error {
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

func (m *ServiceCreateHelmComponentConfigRequest) validatePublicGitVcsConfig(formats strfmt.Registry) error {
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

func (m *ServiceCreateHelmComponentConfigRequest) validateValues(formats strfmt.Registry) error {

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service create helm component config request based on the context it is used
func (m *ServiceCreateHelmComponentConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ServiceCreateHelmComponentConfigRequest) contextValidateConnectedGithubVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceCreateHelmComponentConfigRequest) contextValidatePublicGitVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ServiceCreateHelmComponentConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateHelmComponentConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateHelmComponentConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
