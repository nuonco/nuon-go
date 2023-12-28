// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppSandboxRelease app sandbox release
//
// swagger:model app.SandboxRelease
type AppSandboxRelease struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// deprovision policy url
	DeprovisionPolicyURL string `json:"deprovision_policy_url,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// one click role template url
	OneClickRoleTemplateURL string `json:"one_click_role_template_url,omitempty"`

	// provision policy url
	ProvisionPolicyURL string `json:"provision_policy_url,omitempty"`

	// trust policy url
	TrustPolicyURL string `json:"trust_policy_url,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this app sandbox release
func (m *AppSandboxRelease) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app sandbox release based on context it is used
func (m *AppSandboxRelease) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppSandboxRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppSandboxRelease) UnmarshalBinary(b []byte) error {
	var res AppSandboxRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
