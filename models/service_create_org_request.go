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

// ServiceCreateOrgRequest service create org request
//
// swagger:model service.CreateOrgRequest
type ServiceCreateOrgRequest struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// These fields are used to control the behaviour of the org.
	UseSandboxMode bool `json:"use_sandbox_mode,omitempty"`
}

// Validate validates this service create org request
func (m *ServiceCreateOrgRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateOrgRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service create org request based on context it is used
func (m *ServiceCreateOrgRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateOrgRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateOrgRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateOrgRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
