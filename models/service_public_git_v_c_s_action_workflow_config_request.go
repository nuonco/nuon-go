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

// ServicePublicGitVCSActionWorkflowConfigRequest service public git v c s action workflow config request
//
// swagger:model service.PublicGitVCSActionWorkflowConfigRequest
type ServicePublicGitVCSActionWorkflowConfigRequest struct {

	// branch
	// Required: true
	Branch *string `json:"branch"`

	// directory
	// Required: true
	Directory *string `json:"directory"`

	// repo
	// Required: true
	Repo *string `json:"repo"`
}

// Validate validates this service public git v c s action workflow config request
func (m *ServicePublicGitVCSActionWorkflowConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePublicGitVCSActionWorkflowConfigRequest) validateBranch(formats strfmt.Registry) error {

	if err := validate.Required("branch", "body", m.Branch); err != nil {
		return err
	}

	return nil
}

func (m *ServicePublicGitVCSActionWorkflowConfigRequest) validateDirectory(formats strfmt.Registry) error {

	if err := validate.Required("directory", "body", m.Directory); err != nil {
		return err
	}

	return nil
}

func (m *ServicePublicGitVCSActionWorkflowConfigRequest) validateRepo(formats strfmt.Registry) error {

	if err := validate.Required("repo", "body", m.Repo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service public git v c s action workflow config request based on context it is used
func (m *ServicePublicGitVCSActionWorkflowConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServicePublicGitVCSActionWorkflowConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePublicGitVCSActionWorkflowConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServicePublicGitVCSActionWorkflowConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}