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

// AppRunnerJobOperationType app runner job operation type
//
// swagger:model app.RunnerJobOperationType
type AppRunnerJobOperationType string

func NewAppRunnerJobOperationType(value AppRunnerJobOperationType) *AppRunnerJobOperationType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppRunnerJobOperationType.
func (m AppRunnerJobOperationType) Pointer() *AppRunnerJobOperationType {
	return &m
}

const (

	// AppRunnerJobOperationTypeExec captures enum value "exec"
	AppRunnerJobOperationTypeExec AppRunnerJobOperationType = "exec"

	// AppRunnerJobOperationTypeApply captures enum value "apply"
	AppRunnerJobOperationTypeApply AppRunnerJobOperationType = "apply"

	// AppRunnerJobOperationTypeDestroy captures enum value "destroy"
	AppRunnerJobOperationTypeDestroy AppRunnerJobOperationType = "destroy"

	// AppRunnerJobOperationTypePlanDashOnly captures enum value "plan-only"
	AppRunnerJobOperationTypePlanDashOnly AppRunnerJobOperationType = "plan-only"

	// AppRunnerJobOperationTypeBuild captures enum value "build"
	AppRunnerJobOperationTypeBuild AppRunnerJobOperationType = "build"

	// AppRunnerJobOperationTypeUnknown captures enum value "unknown"
	AppRunnerJobOperationTypeUnknown AppRunnerJobOperationType = "unknown"
)

// for schema
var appRunnerJobOperationTypeEnum []interface{}

func init() {
	var res []AppRunnerJobOperationType
	if err := json.Unmarshal([]byte(`["exec","apply","destroy","plan-only","build","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appRunnerJobOperationTypeEnum = append(appRunnerJobOperationTypeEnum, v)
	}
}

func (m AppRunnerJobOperationType) validateAppRunnerJobOperationTypeEnum(path, location string, value AppRunnerJobOperationType) error {
	if err := validate.EnumCase(path, location, value, appRunnerJobOperationTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app runner job operation type
func (m AppRunnerJobOperationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppRunnerJobOperationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app runner job operation type based on context it is used
func (m AppRunnerJobOperationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
