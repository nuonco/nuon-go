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

// AppInstallComponent app install component
//
// swagger:model app.InstallComponent
type AppInstallComponent struct {

	// component
	Component *AppComponent `json:"component,omitempty"`

	// component id
	ComponentID string `json:"component_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// helm chart
	HelmChart *AppHelmChart `json:"helm_chart,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install deploys
	InstallDeploys []*AppInstallDeploy `json:"install_deploys"`

	// install id
	InstallID string `json:"install_id,omitempty"`

	// links
	Links interface{} `json:"links,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// status v2
	StatusV2 *AppCompositeStatus `json:"status_v2,omitempty"`

	// terraform workspace
	TerraformWorkspace *AppTerraformWorkspace `json:"terraform_workspace,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install component
func (m *AppInstallComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelmChart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallDeploys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusV2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallComponent) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponent) validateHelmChart(formats strfmt.Registry) error {
	if swag.IsZero(m.HelmChart) { // not required
		return nil
	}

	if m.HelmChart != nil {
		if err := m.HelmChart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm_chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm_chart")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponent) validateInstallDeploys(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallDeploys) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallDeploys); i++ {
		if swag.IsZero(m.InstallDeploys[i]) { // not required
			continue
		}

		if m.InstallDeploys[i] != nil {
			if err := m.InstallDeploys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_deploys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_deploys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstallComponent) validateStatusV2(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusV2) { // not required
		return nil
	}

	if m.StatusV2 != nil {
		if err := m.StatusV2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_v2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_v2")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponent) validateTerraformWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformWorkspace) { // not required
		return nil
	}

	if m.TerraformWorkspace != nil {
		if err := m.TerraformWorkspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app install component based on the context it is used
func (m *AppInstallComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHelmChart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallDeploys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusV2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraformWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallComponent) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {

		if swag.IsZero(m.Component) { // not required
			return nil
		}

		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponent) contextValidateHelmChart(ctx context.Context, formats strfmt.Registry) error {

	if m.HelmChart != nil {

		if swag.IsZero(m.HelmChart) { // not required
			return nil
		}

		if err := m.HelmChart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm_chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm_chart")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponent) contextValidateInstallDeploys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstallDeploys); i++ {

		if m.InstallDeploys[i] != nil {

			if swag.IsZero(m.InstallDeploys[i]) { // not required
				return nil
			}

			if err := m.InstallDeploys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_deploys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_deploys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstallComponent) contextValidateStatusV2(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusV2 != nil {

		if swag.IsZero(m.StatusV2) { // not required
			return nil
		}

		if err := m.StatusV2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_v2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_v2")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponent) contextValidateTerraformWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformWorkspace != nil {

		if swag.IsZero(m.TerraformWorkspace) { // not required
			return nil
		}

		if err := m.TerraformWorkspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraform_workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraform_workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallComponent) UnmarshalBinary(b []byte) error {
	var res AppInstallComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
