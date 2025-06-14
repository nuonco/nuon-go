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

// ServiceCreateInstallConfigRequest service create install config request
//
// swagger:model service.CreateInstallConfigRequest
type ServiceCreateInstallConfigRequest struct {

	// approval option
	ApprovalOption AppInstallApprovalOption `json:"approvalOption,omitempty"`
}

// Validate validates this service create install config request
func (m *ServiceCreateInstallConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprovalOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateInstallConfigRequest) validateApprovalOption(formats strfmt.Registry) error {
	if swag.IsZero(m.ApprovalOption) { // not required
		return nil
	}

	if err := m.ApprovalOption.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("approvalOption")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("approvalOption")
		}
		return err
	}

	return nil
}

// ContextValidate validate this service create install config request based on the context it is used
func (m *ServiceCreateInstallConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApprovalOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateInstallConfigRequest) contextValidateApprovalOption(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ApprovalOption) { // not required
		return nil
	}

	if err := m.ApprovalOption.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("approvalOption")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("approvalOption")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateInstallConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateInstallConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateInstallConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
