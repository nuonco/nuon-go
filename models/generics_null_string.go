// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenericsNullString generics null string
//
// swagger:model generics.NullString
type GenericsNullString struct {

	// string
	String string `json:"string,omitempty"`

	// Valid is true if String is not NULL
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this generics null string
func (m *GenericsNullString) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generics null string based on context it is used
func (m *GenericsNullString) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericsNullString) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericsNullString) UnmarshalBinary(b []byte) error {
	var res GenericsNullString
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
