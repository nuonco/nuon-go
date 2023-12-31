// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppAWSAccount app a w s account
//
// swagger:model app.AWSAccount
type AppAWSAccount struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// iam role arn
	IamRoleArn string `json:"iam_role_arn,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app a w s account
func (m *AppAWSAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app a w s account based on context it is used
func (m *AppAWSAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppAWSAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAWSAccount) UnmarshalBinary(b []byte) error {
	var res AppAWSAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
