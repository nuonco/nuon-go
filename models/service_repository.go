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

// ServiceRepository service repository
//
// swagger:model service.Repository
type ServiceRepository struct {

	// clone url
	// Required: true
	CloneURL *string `json:"clone_url"`

	// default branch
	// Required: true
	DefaultBranch *string `json:"default_branch"`

	// full name
	// Required: true
	FullName *string `json:"full_name"`

	// git url
	// Required: true
	GitURL *string `json:"git_url"`

	// github install id
	// Required: true
	GithubInstallID *string `json:"github_install_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization name
	// Required: true
	OrganizationName *string `json:"organization_name"`

	// user name
	// Required: true
	UserName *string `json:"user_name"`
}

// Validate validates this service repository
func (m *ServiceRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloneURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithubInstallID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceRepository) validateCloneURL(formats strfmt.Registry) error {

	if err := validate.Required("clone_url", "body", m.CloneURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateDefaultBranch(formats strfmt.Registry) error {

	if err := validate.Required("default_branch", "body", m.DefaultBranch); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("full_name", "body", m.FullName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateGitURL(formats strfmt.Registry) error {

	if err := validate.Required("git_url", "body", m.GitURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateGithubInstallID(formats strfmt.Registry) error {

	if err := validate.Required("github_install_id", "body", m.GithubInstallID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateOrganizationName(formats strfmt.Registry) error {

	if err := validate.Required("organization_name", "body", m.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceRepository) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("user_name", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service repository based on context it is used
func (m *ServiceRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceRepository) UnmarshalBinary(b []byte) error {
	var res ServiceRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
