// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceCLIConfig service c l i config
//
// swagger:model service.CLIConfig
type ServiceCLIConfig struct {

	// auth audience
	AuthAudience string `json:"auth_audience,omitempty"`

	// auth client id
	AuthClientID string `json:"auth_client_id,omitempty"`

	// auth domain
	AuthDomain string `json:"auth_domain,omitempty"`

	// dashboard url
	DashboardURL string `json:"dashboard_url,omitempty"`
}

// Validate validates this service c l i config
func (m *ServiceCLIConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service c l i config based on context it is used
func (m *ServiceCLIConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCLIConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCLIConfig) UnmarshalBinary(b []byte) error {
	var res ServiceCLIConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
