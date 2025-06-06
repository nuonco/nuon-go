// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppHelmRelease app helm release
//
// swagger:model app.HelmRelease
type AppHelmRelease struct {

	// The rspb.Release body, as a base64-encoded string
	Body string `json:"body,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// helm chart
	HelmChart *AppHelmChart `json:"helmChart,omitempty"`

	// helm chart ID
	HelmChartID string `json:"helmChartID,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// labels
	Labels AppJSONMap `json:"labels,omitempty"`

	// Release "labels" that can be used as filters in the storage.Query(labels map[string]string)
	// we implemented. Note that allowing Helm users to filter against new dimensions will require a
	// new migration to be added, and the Create and/or update functions to be updated accordingly.
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// See https://github.com/helm/helm/blob/c9fe3d118caec699eb2565df9838673af379ce12/pkg/storage/driver/secrets.go#L231
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this app helm release
func (m *AppHelmRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHelmChart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppHelmRelease) validateHelmChart(formats strfmt.Registry) error {
	if swag.IsZero(m.HelmChart) { // not required
		return nil
	}

	if m.HelmChart != nil {
		if err := m.HelmChart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helmChart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helmChart")
			}
			return err
		}
	}

	return nil
}

func (m *AppHelmRelease) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if m.Labels != nil {
		if err := m.Labels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app helm release based on the context it is used
func (m *AppHelmRelease) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHelmChart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppHelmRelease) contextValidateHelmChart(ctx context.Context, formats strfmt.Registry) error {

	if m.HelmChart != nil {

		if swag.IsZero(m.HelmChart) { // not required
			return nil
		}

		if err := m.HelmChart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helmChart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helmChart")
			}
			return err
		}
	}

	return nil
}

func (m *AppHelmRelease) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppHelmRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppHelmRelease) UnmarshalBinary(b []byte) error {
	var res AppHelmRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
