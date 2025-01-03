// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppAppAWSDelegationConfig app app a w s delegation config
//
// swagger:model app.AppAWSDelegationConfig
type AppAppAWSDelegationConfig struct {

	// app sandbox config id
	AppSandboxConfigID string `json:"app_sandbox_config_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// iam role arn
	IamRoleArn string `json:"iam_role_arn,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app app a w s delegation config
func (m *AppAppAWSDelegationConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app app a w s delegation config based on context it is used
func (m *AppAppAWSDelegationConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppAppAWSDelegationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppAWSDelegationConfig) UnmarshalBinary(b []byte) error {
	var res AppAppAWSDelegationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
