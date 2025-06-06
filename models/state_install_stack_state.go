// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StateInstallStackState state install stack state
//
// swagger:model state.InstallStackState
type StateInstallStackState struct {

	// checksum
	Checksum string `json:"checksum,omitempty"`

	// outputs
	Outputs map[string]string `json:"outputs,omitempty"`

	// populated
	Populated bool `json:"populated,omitempty"`

	// quick link url
	QuickLinkURL string `json:"quick_link_url,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// template json
	TemplateJSON string `json:"template_json,omitempty"`

	// template url
	TemplateURL string `json:"template_url,omitempty"`
}

// Validate validates this state install stack state
func (m *StateInstallStackState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this state install stack state based on context it is used
func (m *StateInstallStackState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StateInstallStackState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateInstallStackState) UnmarshalBinary(b []byte) error {
	var res StateInstallStackState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
