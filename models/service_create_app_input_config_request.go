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
	Inputs map[string]ServiceAppInputRequest `json:"inputs"`
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

	for k := range m.Inputs {

		if err := validate.Required("inputs"+"."+k, "body", m.Inputs[k]); err != nil {
			return err
		}
		if val, ok := m.Inputs[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service create app input config request based on the context it is used
func (m *ServiceCreateAppInputConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateAppInputConfigRequest) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("inputs", "body", m.Inputs); err != nil {
		return err
	}

	for k := range m.Inputs {

		if val, ok := m.Inputs[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

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
