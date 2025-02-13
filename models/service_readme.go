// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceReadme service readme
//
// swagger:model service.Readme
type ServiceReadme struct {

	// original
	Original string `json:"original,omitempty"`

	// readme
	Readme string `json:"readme,omitempty"`

	// warnings
	Warnings []string `json:"warnings"`
}

// Validate validates this service readme
func (m *ServiceReadme) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service readme based on context it is used
func (m *ServiceReadme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceReadme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceReadme) UnmarshalBinary(b []byte) error {
	var res ServiceReadme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
