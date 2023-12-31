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

// ServiceUpdateAppInstallerRequest service update app installer request
//
// swagger:model service.UpdateAppInstallerRequest
type ServiceUpdateAppInstallerRequest struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// links
	Links *ServiceUpdateAppInstallerRequestLinks `json:"links,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this service update app installer request
func (m *ServiceUpdateAppInstallerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceUpdateAppInstallerRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequest) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service update app installer request based on the context it is used
func (m *ServiceUpdateAppInstallerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceUpdateAppInstallerRequest) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {

		if swag.IsZero(m.Links) { // not required
			return nil
		}

		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceUpdateAppInstallerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceUpdateAppInstallerRequest) UnmarshalBinary(b []byte) error {
	var res ServiceUpdateAppInstallerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceUpdateAppInstallerRequestLinks service update app installer request links
//
// swagger:model ServiceUpdateAppInstallerRequestLinks
type ServiceUpdateAppInstallerRequestLinks struct {

	// community
	// Required: true
	Community *string `json:"community"`

	// demo
	// Required: true
	Demo *string `json:"demo"`

	// documentation
	// Required: true
	Documentation *string `json:"documentation"`

	// github
	// Required: true
	Github *string `json:"github"`

	// homepage
	// Required: true
	Homepage *string `json:"homepage"`

	// logo
	// Required: true
	Logo *string `json:"logo"`
}

// Validate validates this service update app installer request links
func (m *ServiceUpdateAppInstallerRequestLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithub(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomepage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceUpdateAppInstallerRequestLinks) validateCommunity(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"community", "body", m.Community); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequestLinks) validateDemo(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"demo", "body", m.Demo); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequestLinks) validateDocumentation(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"documentation", "body", m.Documentation); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequestLinks) validateGithub(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"github", "body", m.Github); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequestLinks) validateHomepage(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"homepage", "body", m.Homepage); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUpdateAppInstallerRequestLinks) validateLogo(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"logo", "body", m.Logo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service update app installer request links based on context it is used
func (m *ServiceUpdateAppInstallerRequestLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceUpdateAppInstallerRequestLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceUpdateAppInstallerRequestLinks) UnmarshalBinary(b []byte) error {
	var res ServiceUpdateAppInstallerRequestLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
