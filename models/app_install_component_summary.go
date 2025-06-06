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

// AppInstallComponentSummary app install component summary
//
// swagger:model app.InstallComponentSummary
type AppInstallComponentSummary struct {

	// build status
	BuildStatus AppComponentBuildStatus `json:"build_status,omitempty"`

	// build status description
	BuildStatusDescription string `json:"build_status_description,omitempty"`

	// component config
	ComponentConfig *AppComponentConfigConnection `json:"component_config,omitempty"`

	// component id
	ComponentID string `json:"component_id,omitempty"`

	// component name
	ComponentName string `json:"component_name,omitempty"`

	// dependencies
	Dependencies []*AppComponent `json:"dependencies"`

	// deploy status
	DeployStatus AppInstallDeployStatus `json:"deploy_status,omitempty"`

	// deploy status description
	DeployStatusDescription string `json:"deploy_status_description,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this app install component summary
func (m *AppInstallComponentSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponentConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallComponentSummary) validateBuildStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildStatus) { // not required
		return nil
	}

	if err := m.BuildStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("build_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("build_status")
		}
		return err
	}

	return nil
}

func (m *AppInstallComponentSummary) validateComponentConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentConfig) { // not required
		return nil
	}

	if m.ComponentConfig != nil {
		if err := m.ComponentConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponentSummary) validateDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.Dependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstallComponentSummary) validateDeployStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.DeployStatus) { // not required
		return nil
	}

	if err := m.DeployStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploy_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploy_status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app install component summary based on the context it is used
func (m *AppInstallComponentSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponentConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeployStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallComponentSummary) contextValidateBuildStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.BuildStatus) { // not required
		return nil
	}

	if err := m.BuildStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("build_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("build_status")
		}
		return err
	}

	return nil
}

func (m *AppInstallComponentSummary) contextValidateComponentConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ComponentConfig != nil {

		if swag.IsZero(m.ComponentConfig) { // not required
			return nil
		}

		if err := m.ComponentConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstallComponentSummary) contextValidateDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dependencies); i++ {

		if m.Dependencies[i] != nil {

			if swag.IsZero(m.Dependencies[i]) { // not required
				return nil
			}

			if err := m.Dependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstallComponentSummary) contextValidateDeployStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DeployStatus) { // not required
		return nil
	}

	if err := m.DeployStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploy_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deploy_status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallComponentSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallComponentSummary) UnmarshalBinary(b []byte) error {
	var res AppInstallComponentSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
