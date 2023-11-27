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

// ServiceCreateAppInputConfigRequest service create app input config request
//
// swagger:model service.CreateAppInputConfigRequest
type ServiceCreateAppInputConfigRequest struct {

	// inputs
	// Required: true
	Inputs map[string]string `json:"inputs"`
}

// Validate validates this service create app input config request
func (m *ServiceCreateAppInputConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateAppInputConfigRequest) validateInputs(formats strfmt.Registry) error {

	if err := validate.Required("inputs", "body", m.Inputs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service create app input config request based on context it is used
func (m *ServiceCreateAppInputConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateAppInputConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateAppInputConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateAppInputConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
