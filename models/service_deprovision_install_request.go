// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDeprovisionInstallRequest service deprovision install request
//
// swagger:model service.DeprovisionInstallRequest
type ServiceDeprovisionInstallRequest struct {

	// error behavior
	ErrorBehavior string `json:"error_behavior,omitempty"`
}

// Validate validates this service deprovision install request
func (m *ServiceDeprovisionInstallRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service deprovision install request based on context it is used
func (m *ServiceDeprovisionInstallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDeprovisionInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDeprovisionInstallRequest) UnmarshalBinary(b []byte) error {
	var res ServiceDeprovisionInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
