// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppInstallDeploy app install deploy
//
// swagger:model app.InstallDeploy
type AppInstallDeploy struct {

	// build id
	BuildID string `json:"build_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install component id
	InstallComponentID string `json:"install_component_id,omitempty"`

	// release id
	ReleaseID string `json:"release_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install deploy
func (m *AppInstallDeploy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app install deploy based on context it is used
func (m *AppInstallDeploy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
