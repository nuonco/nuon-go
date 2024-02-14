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

// ServiceRenderedInstall service rendered install
//
// swagger:model service.RenderedInstall
type ServiceRenderedInstall struct {

	// install
	Install *AppInstall `json:"install,omitempty"`

	// installer
	Installer *ServiceRenderedInstaller `json:"installer,omitempty"`

	// installer content
	InstallerContent string `json:"installer_content,omitempty"`
}

// Validate validates this service rendered install
func (m *ServiceRenderedInstall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstaller(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRenderedInstall) validateInstall(formats strfmt.Registry) error {
	if swag.IsZero(m.Install) { // not required
		return nil
	}

	if m.Install != nil {
		if err := m.Install.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("install")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("install")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceRenderedInstall) validateInstaller(formats strfmt.Registry) error {
	if swag.IsZero(m.Installer) { // not required
		return nil
	}

	if m.Installer != nil {
		if err := m.Installer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("installer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service rendered install based on the context it is used
func (m *ServiceRenderedInstall) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstaller(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRenderedInstall) contextValidateInstall(ctx context.Context, formats strfmt.Registry) error {

	if m.Install != nil {

		if swag.IsZero(m.Install) { // not required
			return nil
		}

		if err := m.Install.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("install")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("install")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceRenderedInstall) contextValidateInstaller(ctx context.Context, formats strfmt.Registry) error {

	if m.Installer != nil {

		if swag.IsZero(m.Installer) { // not required
			return nil
		}

		if err := m.Installer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("installer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("installer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceRenderedInstall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceRenderedInstall) UnmarshalBinary(b []byte) error {
	var res ServiceRenderedInstall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
