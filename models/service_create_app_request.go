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

// ServiceCreateAppRequest service create app request
//
// swagger:model service.CreateAppRequest
type ServiceCreateAppRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// slack webhook url
	SlackWebhookURL string `json:"slack_webhook_url,omitempty"`
}

// Validate validates this service create app request
func (m *ServiceCreateAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateAppRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service create app request based on context it is used
func (m *ServiceCreateAppRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateAppRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
