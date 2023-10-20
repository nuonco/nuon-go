// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsIncr metrics incr
//
// swagger:model metrics.Incr
type MetricsIncr struct {

	// name
	Name string `json:"name,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this metrics incr
func (m *MetricsIncr) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics incr based on context it is used
func (m *MetricsIncr) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsIncr) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsIncr) UnmarshalBinary(b []byte) error {
	var res MetricsIncr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
