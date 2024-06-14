// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppAccount app account
//
// swagger:model app.Account
type AppAccount struct {

	// account type
	AccountType AppAccountType `json:"account_type,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// ReadOnly Fields
	OrgIds []string `json:"org_ids"`

	// orgs
	Orgs []*AppOrg `json:"orgs"`

	// permissions
	Permissions PermissionsSet `json:"permissions,omitempty"`

	// roles
	Roles []*AppRole `json:"roles"`

	// subject
	Subject string `json:"subject,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app account
func (m *AppAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAccount) validateAccountType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("account_type")
		}
		return err
	}

	return nil
}

func (m *AppAccount) validateOrgs(formats strfmt.Registry) error {
	if swag.IsZero(m.Orgs) { // not required
		return nil
	}

	for i := 0; i < len(m.Orgs); i++ {
		if swag.IsZero(m.Orgs[i]) { // not required
			continue
		}

		if m.Orgs[i] != nil {
			if err := m.Orgs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orgs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orgs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppAccount) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *AppAccount) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app account based on the context it is used
func (m *AppAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrgs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAccount) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("account_type")
		}
		return err
	}

	return nil
}

func (m *AppAccount) contextValidateOrgs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Orgs); i++ {

		if m.Orgs[i] != nil {

			if swag.IsZero(m.Orgs[i]) { // not required
				return nil
			}

			if err := m.Orgs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orgs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orgs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppAccount) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if err := m.Permissions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permissions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permissions")
		}
		return err
	}

	return nil
}

func (m *AppAccount) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {

			if swag.IsZero(m.Roles[i]) { // not required
				return nil
			}

			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAccount) UnmarshalBinary(b []byte) error {
	var res AppAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
