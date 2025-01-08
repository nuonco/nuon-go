// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceActionWorkflowRecentRunsResponse service action workflow recent runs response
//
// swagger:model service.ActionWorkflowRecentRunsResponse
type ServiceActionWorkflowRecentRunsResponse struct {

	// action workflow
	ActionWorkflow *AppActionWorkflow `json:"action_workflow,omitempty"`

	// recent runs
	RecentRuns []*AppInstallActionWorkflowRun `json:"recent_runs"`
}

// Validate validates this service action workflow recent runs response
func (m *ServiceActionWorkflowRecentRunsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecentRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceActionWorkflowRecentRunsResponse) validateActionWorkflow(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionWorkflow) { // not required
		return nil
	}

	if m.ActionWorkflow != nil {
		if err := m.ActionWorkflow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_workflow")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceActionWorkflowRecentRunsResponse) validateRecentRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.RecentRuns) { // not required
		return nil
	}

	for i := 0; i < len(m.RecentRuns); i++ {
		if swag.IsZero(m.RecentRuns[i]) { // not required
			continue
		}

		if m.RecentRuns[i] != nil {
			if err := m.RecentRuns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recent_runs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recent_runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service action workflow recent runs response based on the context it is used
func (m *ServiceActionWorkflowRecentRunsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionWorkflow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecentRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceActionWorkflowRecentRunsResponse) contextValidateActionWorkflow(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionWorkflow != nil {

		if swag.IsZero(m.ActionWorkflow) { // not required
			return nil
		}

		if err := m.ActionWorkflow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action_workflow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action_workflow")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceActionWorkflowRecentRunsResponse) contextValidateRecentRuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecentRuns); i++ {

		if m.RecentRuns[i] != nil {

			if swag.IsZero(m.RecentRuns[i]) { // not required
				return nil
			}

			if err := m.RecentRuns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recent_runs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recent_runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceActionWorkflowRecentRunsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceActionWorkflowRecentRunsResponse) UnmarshalBinary(b []byte) error {
	var res ServiceActionWorkflowRecentRunsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
