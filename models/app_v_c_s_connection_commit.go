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

// AppVCSConnectionCommit app v c s connection commit
//
// swagger:model app.VCSConnectionCommit
type AppVCSConnectionCommit struct {

	// author email
	AuthorEmail string `json:"author_email,omitempty"`

	// author name
	AuthorName string `json:"author_name,omitempty"`

	// component config connection id
	ComponentConfigConnectionID string `json:"component_config_connection_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app v c s connection commit
func (m *AppVCSConnectionCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppVCSConnectionCommit) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app v c s connection commit based on the context it is used
func (m *AppVCSConnectionCommit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppVCSConnectionCommit) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppVCSConnectionCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppVCSConnectionCommit) UnmarshalBinary(b []byte) error {
	var res AppVCSConnectionCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
