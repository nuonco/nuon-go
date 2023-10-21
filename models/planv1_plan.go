// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Planv1Plan planv1 plan
//
// swagger:model planv1.Plan
type Planv1Plan struct {

	// Types that are assignable to Actual:
	//
	// 	*Plan_WaypointPlan
	// 	*Plan_TerraformPlan
	// 	*Plan_NoopPlan
	Actual interface{} `json:"actual,omitempty"`
}

// Validate validates this planv1 plan
func (m *Planv1Plan) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this planv1 plan based on context it is used
func (m *Planv1Plan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Planv1Plan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Planv1Plan) UnmarshalBinary(b []byte) error {
	var res Planv1Plan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
