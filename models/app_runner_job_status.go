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

// AppRunnerJobStatus app runner job status
//
// swagger:model app.RunnerJobStatus
type AppRunnerJobStatus string

func NewAppRunnerJobStatus(value AppRunnerJobStatus) *AppRunnerJobStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppRunnerJobStatus.
func (m AppRunnerJobStatus) Pointer() *AppRunnerJobStatus {
	return &m
}

const (

	// AppRunnerJobStatusQueued captures enum value "queued"
	AppRunnerJobStatusQueued AppRunnerJobStatus = "queued"

	// AppRunnerJobStatusAvailable captures enum value "available"
	AppRunnerJobStatusAvailable AppRunnerJobStatus = "available"

	// AppRunnerJobStatusInDashProgress captures enum value "in-progress"
	AppRunnerJobStatusInDashProgress AppRunnerJobStatus = "in-progress"

	// AppRunnerJobStatusFinished captures enum value "finished"
	AppRunnerJobStatusFinished AppRunnerJobStatus = "finished"

	// AppRunnerJobStatusFailed captures enum value "failed"
	AppRunnerJobStatusFailed AppRunnerJobStatus = "failed"

	// AppRunnerJobStatusTimedDashOut captures enum value "timed-out"
	AppRunnerJobStatusTimedDashOut AppRunnerJobStatus = "timed-out"

	// AppRunnerJobStatusNotDashAttempted captures enum value "not-attempted"
	AppRunnerJobStatusNotDashAttempted AppRunnerJobStatus = "not-attempted"

	// AppRunnerJobStatusCancelled captures enum value "cancelled"
	AppRunnerJobStatusCancelled AppRunnerJobStatus = "cancelled"

	// AppRunnerJobStatusUnknown captures enum value "unknown"
	AppRunnerJobStatusUnknown AppRunnerJobStatus = "unknown"
)

// for schema
var appRunnerJobStatusEnum []interface{}

func init() {
	var res []AppRunnerJobStatus
	if err := json.Unmarshal([]byte(`["queued","available","in-progress","finished","failed","timed-out","not-attempted","cancelled","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appRunnerJobStatusEnum = append(appRunnerJobStatusEnum, v)
	}
}

func (m AppRunnerJobStatus) validateAppRunnerJobStatusEnum(path, location string, value AppRunnerJobStatus) error {
	if err := validate.EnumCase(path, location, value, appRunnerJobStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app runner job status
func (m AppRunnerJobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppRunnerJobStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app runner job status based on context it is used
func (m AppRunnerJobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
