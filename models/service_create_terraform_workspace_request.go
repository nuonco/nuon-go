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

// ServiceCreateTerraformWorkspaceRequest service create terraform workspace request
//
// swagger:model service.CreateTerraformWorkspaceRequest
type ServiceCreateTerraformWorkspaceRequest struct {

	// owner id
	// Required: true
	OwnerID *string `json:"owner_id"`

	// owner type
	// Required: true
	OwnerType *string `json:"owner_type"`
}

// Validate validates this service create terraform workspace request
func (m *ServiceCreateTerraformWorkspaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateTerraformWorkspaceRequest) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("owner_id", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateTerraformWorkspaceRequest) validateOwnerType(formats strfmt.Registry) error {

	if err := validate.Required("owner_type", "body", m.OwnerType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service create terraform workspace request based on context it is used
func (m *ServiceCreateTerraformWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateTerraformWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateTerraformWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateTerraformWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
