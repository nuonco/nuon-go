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

// AppActionWorkflowStepConfig app action workflow step config
//
// swagger:model app.ActionWorkflowStepConfig
type AppActionWorkflowStepConfig struct {

	// action workflow config id
	ActionWorkflowConfigID string `json:"action_workflow_config_id,omitempty"`

	// this belongs to an app config id
	AppConfigID string `json:"app_config_id,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// command
	Command string `json:"command,omitempty"`

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

	// idx
	Idx int64 `json:"idx,omitempty"`

	// metadata
	Name string `json:"name,omitempty"`

	// previous step id
	PreviousStepID string `json:"previous_step_id,omitempty"`

	// all the details needed for a step
	PublicGitVcsConfig struct {
		AppPublicGitVCSConfig
	} `json:"public_git_vcs_config,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app action workflow step config
func (m *AppActionWorkflowStepConfig) Validate(formats strfmt.Registry) error {
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

func (m *AppActionWorkflowStepConfig) validateConnectedGithubVcsConfig(formats strfmt.Registry) error {
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

func (m *AppActionWorkflowStepConfig) validatePublicGitVcsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicGitVcsConfig) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this app action workflow step config based on the context it is used
func (m *AppActionWorkflowStepConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *AppActionWorkflowStepConfig) contextValidateConnectedGithubVcsConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppActionWorkflowStepConfig) contextValidatePublicGitVcsConfig(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *AppActionWorkflowStepConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppActionWorkflowStepConfig) UnmarshalBinary(b []byte) error {
	var res AppActionWorkflowStepConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
