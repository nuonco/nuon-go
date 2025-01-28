// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StateComponentImage state component image
//
// swagger:model state.componentImage
type StateComponentImage struct {

	// registry
	Registry *StateComponentImageRegistry `json:"registry,omitempty"`

	// repository
	Repository *StateComponentImageRepository `json:"repository,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`
}

// Validate validates this state component image
func (m *StateComponentImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateComponentImage) validateRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

func (m *StateComponentImage) validateRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this state component image based on the context it is used
func (m *StateComponentImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateComponentImage) contextValidateRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.Registry != nil {

		if swag.IsZero(m.Registry) { // not required
			return nil
		}

		if err := m.Registry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

func (m *StateComponentImage) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {

		if swag.IsZero(m.Repository) { // not required
			return nil
		}

		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateComponentImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateComponentImage) UnmarshalBinary(b []byte) error {
	var res StateComponentImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
