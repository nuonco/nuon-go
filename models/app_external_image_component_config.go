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

// AppExternalImageComponentConfig app external image component config
//
// swagger:model app.ExternalImageComponentConfig
type AppExternalImageComponentConfig struct {

	// aws ecr image config
	AwsEcrImageConfig *AppAWSECRImageConfig `json:"aws_ecr_image_config,omitempty"`

	// value
	ComponentConfigConnectionID string `json:"component_config_connection_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image url
	ImageURL string `json:"image_url,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app external image component config
func (m *AppExternalImageComponentConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsEcrImageConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppExternalImageComponentConfig) validateAwsEcrImageConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsEcrImageConfig) { // not required
		return nil
	}

	if m.AwsEcrImageConfig != nil {
		if err := m.AwsEcrImageConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_ecr_image_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_ecr_image_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppExternalImageComponentConfig) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app external image component config based on the context it is used
func (m *AppExternalImageComponentConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsEcrImageConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppExternalImageComponentConfig) contextValidateAwsEcrImageConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsEcrImageConfig != nil {

		if swag.IsZero(m.AwsEcrImageConfig) { // not required
			return nil
		}

		if err := m.AwsEcrImageConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_ecr_image_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_ecr_image_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppExternalImageComponentConfig) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppExternalImageComponentConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppExternalImageComponentConfig) UnmarshalBinary(b []byte) error {
	var res AppExternalImageComponentConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
