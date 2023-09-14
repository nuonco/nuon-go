// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceCreateOrgUserRequest service create org user request
//
// swagger:model service.CreateOrgUserRequest
type ServiceCreateOrgUserRequest struct {

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this service create org user request
func (m *ServiceCreateOrgUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service create org user request based on context it is used
func (m *ServiceCreateOrgUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateOrgUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateOrgUserRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateOrgUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
