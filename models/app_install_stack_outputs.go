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

// AppInstallStackOutputs app install stack outputs
//
// swagger:model app.InstallStackOutputs
type AppInstallStackOutputs struct {

	// aws
	Aws *AppAWSStackOutputs `json:"aws,omitempty"`

	// azure
	Azure *AppAzureStackOutputs `json:"azure,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// data
	Data map[string]string `json:"data,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install stack id
	InstallStackID string `json:"install_stack_id,omitempty"`

	// install version run id
	InstallVersionRunID string `json:"install_version_run_id,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install stack outputs
func (m *AppInstallStackOutputs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallStackOutputs) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallStackOutputs) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app install stack outputs based on the context it is used
func (m *AppInstallStackOutputs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallStackOutputs) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {

		if swag.IsZero(m.Aws) { // not required
			return nil
		}

		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallStackOutputs) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {

		if swag.IsZero(m.Azure) { // not required
			return nil
		}

		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallStackOutputs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallStackOutputs) UnmarshalBinary(b []byte) error {
	var res AppInstallStackOutputs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
