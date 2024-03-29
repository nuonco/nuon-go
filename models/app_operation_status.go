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

// AppOperationStatus app operation status
//
// swagger:model app.OperationStatus
type AppOperationStatus string

func NewAppOperationStatus(value AppOperationStatus) *AppOperationStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppOperationStatus.
func (m AppOperationStatus) Pointer() *AppOperationStatus {
	return &m
}

const (

	// AppOperationStatusStarted captures enum value "started"
	AppOperationStatusStarted AppOperationStatus = "started"

	// AppOperationStatusFinished captures enum value "finished"
	AppOperationStatusFinished AppOperationStatus = "finished"

	// AppOperationStatusNoop captures enum value "noop"
	AppOperationStatusNoop AppOperationStatus = "noop"

	// AppOperationStatusFailed captures enum value "failed"
	AppOperationStatusFailed AppOperationStatus = "failed"
)

// for schema
var appOperationStatusEnum []interface{}

func init() {
	var res []AppOperationStatus
	if err := json.Unmarshal([]byte(`["started","finished","noop","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appOperationStatusEnum = append(appOperationStatusEnum, v)
	}
}

func (m AppOperationStatus) validateAppOperationStatusEnum(path, location string, value AppOperationStatus) error {
	if err := validate.EnumCase(path, location, value, appOperationStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app operation status
func (m AppOperationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppOperationStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app operation status based on context it is used
func (m AppOperationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
