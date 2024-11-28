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

// AppConnectedGithubVCSConfig app connected github v c s config
//
// swagger:model app.ConnectedGithubVCSConfig
type AppConnectedGithubVCSConfig struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// parent component
	ComponentConfigID string `json:"component_config_id,omitempty"`

	// component config type
	ComponentConfigType string `json:"component_config_type,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// directory
	Directory string `json:"directory,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// repo
	Repo string `json:"repo,omitempty"`

	// repo name
	RepoName string `json:"repo_name,omitempty"`

	// repo owner
	RepoOwner string `json:"repo_owner,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// vcs connection
	VcsConnection *AppVCSConnection `json:"vcs_connection,omitempty"`

	// vcs connection id
	VcsConnectionID string `json:"vcs_connection_id,omitempty"`
}

// Validate validates this app connected github v c s config
func (m *AppConnectedGithubVCSConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcsConnection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppConnectedGithubVCSConfig) validateVcsConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsConnection) { // not required
		return nil
	}

	if m.VcsConnection != nil {
		if err := m.VcsConnection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs_connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcs_connection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app connected github v c s config based on the context it is used
func (m *AppConnectedGithubVCSConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVcsConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppConnectedGithubVCSConfig) contextValidateVcsConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.VcsConnection != nil {

		if swag.IsZero(m.VcsConnection) { // not required
			return nil
		}

		if err := m.VcsConnection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs_connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vcs_connection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppConnectedGithubVCSConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppConnectedGithubVCSConfig) UnmarshalBinary(b []byte) error {
	var res AppConnectedGithubVCSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
