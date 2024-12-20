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

// AppRunnerOperation app runner operation
//
// swagger:model app.RunnerOperation
type AppRunnerOperation struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// job details
	LogStream struct {
		AppLogStream
	} `json:"log_stream,omitempty"`

	// operation type
	OperationType AppRunnerOperationType `json:"operation_type,omitempty"`

	// runner id
	RunnerID string `json:"runner_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app runner operation
func (m *AppRunnerOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogStream(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRunnerOperation) validateLogStream(formats strfmt.Registry) error {
	if swag.IsZero(m.LogStream) { // not required
		return nil
	}

	return nil
}

func (m *AppRunnerOperation) validateOperationType(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationType) { // not required
		return nil
	}

	if err := m.OperationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operation_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operation_type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app runner operation based on the context it is used
func (m *AppRunnerOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogStream(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRunnerOperation) contextValidateLogStream(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *AppRunnerOperation) contextValidateOperationType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.OperationType) { // not required
		return nil
	}

	if err := m.OperationType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operation_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operation_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRunnerOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRunnerOperation) UnmarshalBinary(b []byte) error {
	var res AppRunnerOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}