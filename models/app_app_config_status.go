// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AppAppConfigStatus app app config status
//
// swagger:model app.AppConfigStatus
type AppAppConfigStatus string

func NewAppAppConfigStatus(value AppAppConfigStatus) *AppAppConfigStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppAppConfigStatus.
func (m AppAppConfigStatus) Pointer() *AppAppConfigStatus {
	return &m
}

const (

	// AppAppConfigStatusApplied captures enum value "applied"
	AppAppConfigStatusApplied AppAppConfigStatus = "applied"

	// AppAppConfigStatusPending captures enum value "pending"
	AppAppConfigStatusPending AppAppConfigStatus = "pending"

	// AppAppConfigStatusError captures enum value "error"
	AppAppConfigStatusError AppAppConfigStatus = "error"

	// AppAppConfigStatusOutdated captures enum value "outdated"
	AppAppConfigStatusOutdated AppAppConfigStatus = "outdated"
)

// for schema
var appAppConfigStatusEnum []interface{}

func init() {
	var res []AppAppConfigStatus
	if err := json.Unmarshal([]byte(`["applied","pending","error","outdated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appAppConfigStatusEnum = append(appAppConfigStatusEnum, v)
	}
}

func (m AppAppConfigStatus) validateAppAppConfigStatusEnum(path, location string, value AppAppConfigStatus) error {
	if err := validate.EnumCase(path, location, value, appAppConfigStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app app config status
func (m AppAppConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppAppConfigStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app app config status based on context it is used
func (m AppAppConfigStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
