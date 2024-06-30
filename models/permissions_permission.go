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

// PermissionsPermission permissions permission
//
// swagger:model permissions.Permission
type PermissionsPermission string

func NewPermissionsPermission(value PermissionsPermission) *PermissionsPermission {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PermissionsPermission.
func (m PermissionsPermission) Pointer() *PermissionsPermission {
	return &m
}

const (

	// PermissionsPermissionUnknown captures enum value "unknown"
	PermissionsPermissionUnknown PermissionsPermission = "unknown"

	// PermissionsPermissionAll captures enum value "all"
	PermissionsPermissionAll PermissionsPermission = "all"

	// PermissionsPermissionCreate captures enum value "create"
	PermissionsPermissionCreate PermissionsPermission = "create"

	// PermissionsPermissionRead captures enum value "read"
	PermissionsPermissionRead PermissionsPermission = "read"

	// PermissionsPermissionUpdate captures enum value "update"
	PermissionsPermissionUpdate PermissionsPermission = "update"

	// PermissionsPermissionDelete captures enum value "delete"
	PermissionsPermissionDelete PermissionsPermission = "delete"
)

// for schema
var permissionsPermissionEnum []interface{}

func init() {
	var res []PermissionsPermission
	if err := json.Unmarshal([]byte(`["unknown","all","create","read","update","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		permissionsPermissionEnum = append(permissionsPermissionEnum, v)
	}
}

func (m PermissionsPermission) validatePermissionsPermissionEnum(path, location string, value PermissionsPermission) error {
	if err := validate.EnumCase(path, location, value, permissionsPermissionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this permissions permission
func (m PermissionsPermission) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePermissionsPermissionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this permissions permission based on context it is used
func (m PermissionsPermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}