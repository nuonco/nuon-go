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

// AppAccountType app account type
//
// swagger:model app.AccountType
type AppAccountType string

func NewAppAccountType(value AppAccountType) *AppAccountType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppAccountType.
func (m AppAccountType) Pointer() *AppAccountType {
	return &m
}

const (

	// AppAccountTypeAuth0 captures enum value "auth0"
	AppAccountTypeAuth0 AppAccountType = "auth0"

	// AppAccountTypeService captures enum value "service"
	AppAccountTypeService AppAccountType = "service"

	// AppAccountTypeCanary captures enum value "canary"
	AppAccountTypeCanary AppAccountType = "canary"

	// AppAccountTypeIntegration captures enum value "integration"
	AppAccountTypeIntegration AppAccountType = "integration"
)

// for schema
var appAccountTypeEnum []interface{}

func init() {
	var res []AppAccountType
	if err := json.Unmarshal([]byte(`["auth0","service","canary","integration"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appAccountTypeEnum = append(appAccountTypeEnum, v)
	}
}

func (m AppAccountType) validateAppAccountTypeEnum(path, location string, value AppAccountType) error {
	if err := validate.EnumCase(path, location, value, appAccountTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app account type
func (m AppAccountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppAccountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app account type based on context it is used
func (m AppAccountType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
