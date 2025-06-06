// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceAppAWSIAMPolicyConfig service app a w s i a m policy config
//
// swagger:model service.AppAWSIAMPolicyConfig
type ServiceAppAWSIAMPolicyConfig struct {

	// contents
	Contents string `json:"contents,omitempty"`

	// managed policy name
	ManagedPolicyName string `json:"managed_policy_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this service app a w s i a m policy config
func (m *ServiceAppAWSIAMPolicyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service app a w s i a m policy config based on context it is used
func (m *ServiceAppAWSIAMPolicyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceAppAWSIAMPolicyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceAppAWSIAMPolicyConfig) UnmarshalBinary(b []byte) error {
	var res ServiceAppAWSIAMPolicyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
