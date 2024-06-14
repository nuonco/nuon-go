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

// AppPolicyName app policy name
//
// swagger:model app.PolicyName
type AppPolicyName string

func NewAppPolicyName(value AppPolicyName) *AppPolicyName {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppPolicyName.
func (m AppPolicyName) Pointer() *AppPolicyName {
	return &m
}

const (

	// AppPolicyNameOrgAdmin captures enum value "org_admin"
	AppPolicyNameOrgAdmin AppPolicyName = "org_admin"

	// AppPolicyNameInstaller captures enum value "installer"
	AppPolicyNameInstaller AppPolicyName = "installer"

	// AppPolicyNameRunner captures enum value "runner"
	AppPolicyNameRunner AppPolicyName = "runner"
)

// for schema
var appPolicyNameEnum []interface{}

func init() {
	var res []AppPolicyName
	if err := json.Unmarshal([]byte(`["org_admin","installer","runner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appPolicyNameEnum = append(appPolicyNameEnum, v)
	}
}

func (m AppPolicyName) validateAppPolicyNameEnum(path, location string, value AppPolicyName) error {
	if err := validate.EnumCase(path, location, value, appPolicyNameEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app policy name
func (m AppPolicyName) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppPolicyNameEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app policy name based on context it is used
func (m AppPolicyName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
