// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppJobComponentConfig app job component config
//
// swagger:model app.JobComponentConfig
type AppJobComponentConfig struct {

	// cmd
	Cmd string `json:"cmd,omitempty"`

	// value
	ComponentConfigConnectionID string `json:"component_config_connection_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Image attributes, copied from a docker_buid or external_image component.
	ImageURL string `json:"image_url,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app job component config
func (m *AppJobComponentConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app job component config based on context it is used
func (m *AppJobComponentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppJobComponentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppJobComponentConfig) UnmarshalBinary(b []byte) error {
	var res AppJobComponentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
