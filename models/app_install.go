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

// AppInstall app install
//
// swagger:model app.Install
type AppInstall struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// app runner config
	AppRunnerConfig *AppAppRunnerConfig `json:"app_runner_config,omitempty"`

	// app sandbox config
	AppSandboxConfig *AppAppSandboxConfig `json:"app_sandbox_config,omitempty"`

	// aws account
	AwsAccount *AppAWSAccount `json:"aws_account,omitempty"`

	// azure account
	AzureAccount *AppAzureAccount `json:"azure_account,omitempty"`

	// component statuses
	ComponentStatuses map[string]string `json:"component_statuses,omitempty"`

	// composite component status
	CompositeComponentStatus string `json:"composite_component_status,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install components
	InstallComponents []*AppInstallComponent `json:"install_components"`

	// install events
	InstallEvents []*AppInstallEvent `json:"install_events"`

	// install inputs
	InstallInputs []*AppInstallInputs `json:"install_inputs"`

	// install number
	InstallNumber int64 `json:"install_number,omitempty"`

	// install sandbox runs
	InstallSandboxRuns []*AppInstallSandboxRun `json:"install_sandbox_runs"`

	// name
	Name string `json:"name,omitempty"`

	// runner status
	RunnerStatus string `json:"runner_status,omitempty"`

	// sandbox status
	SandboxStatus string `json:"sandbox_status,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install
func (m *AppInstall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppRunnerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppSandboxConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallSandboxRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstall) validateAppRunnerConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AppRunnerConfig) { // not required
		return nil
	}

	if m.AppRunnerConfig != nil {
		if err := m.AppRunnerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_runner_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app_runner_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) validateAppSandboxConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AppSandboxConfig) { // not required
		return nil
	}

	if m.AppSandboxConfig != nil {
		if err := m.AppSandboxConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_sandbox_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app_sandbox_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) validateAwsAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsAccount) { // not required
		return nil
	}

	if m.AwsAccount != nil {
		if err := m.AwsAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_account")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) validateAzureAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureAccount) { // not required
		return nil
	}

	if m.AzureAccount != nil {
		if err := m.AzureAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_account")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) validateInstallComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallComponents) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallComponents); i++ {
		if swag.IsZero(m.InstallComponents[i]) { // not required
			continue
		}

		if m.InstallComponents[i] != nil {
			if err := m.InstallComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstall) validateInstallEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallEvents); i++ {
		if swag.IsZero(m.InstallEvents[i]) { // not required
			continue
		}

		if m.InstallEvents[i] != nil {
			if err := m.InstallEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstall) validateInstallInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallInputs) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallInputs); i++ {
		if swag.IsZero(m.InstallInputs[i]) { // not required
			continue
		}

		if m.InstallInputs[i] != nil {
			if err := m.InstallInputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstall) validateInstallSandboxRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallSandboxRuns) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallSandboxRuns); i++ {
		if swag.IsZero(m.InstallSandboxRuns[i]) { // not required
			continue
		}

		if m.InstallSandboxRuns[i] != nil {
			if err := m.InstallSandboxRuns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_sandbox_runs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_sandbox_runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app install based on the context it is used
func (m *AppInstall) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppRunnerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppSandboxConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallSandboxRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstall) contextValidateAppRunnerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AppRunnerConfig != nil {

		if swag.IsZero(m.AppRunnerConfig) { // not required
			return nil
		}

		if err := m.AppRunnerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_runner_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app_runner_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) contextValidateAppSandboxConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AppSandboxConfig != nil {

		if swag.IsZero(m.AppSandboxConfig) { // not required
			return nil
		}

		if err := m.AppSandboxConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_sandbox_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("app_sandbox_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) contextValidateAwsAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsAccount != nil {

		if swag.IsZero(m.AwsAccount) { // not required
			return nil
		}

		if err := m.AwsAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_account")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) contextValidateAzureAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureAccount != nil {

		if swag.IsZero(m.AzureAccount) { // not required
			return nil
		}

		if err := m.AzureAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure_account")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstall) contextValidateInstallComponents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstallComponents); i++ {

		if m.InstallComponents[i] != nil {

			if swag.IsZero(m.InstallComponents[i]) { // not required
				return nil
			}

			if err := m.InstallComponents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_components" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstall) contextValidateInstallEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstallEvents); i++ {

		if m.InstallEvents[i] != nil {

			if swag.IsZero(m.InstallEvents[i]) { // not required
				return nil
			}

			if err := m.InstallEvents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstall) contextValidateInstallInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstallInputs); i++ {

		if m.InstallInputs[i] != nil {

			if swag.IsZero(m.InstallInputs[i]) { // not required
				return nil
			}

			if err := m.InstallInputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppInstall) contextValidateInstallSandboxRuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstallSandboxRuns); i++ {

		if m.InstallSandboxRuns[i] != nil {

			if swag.IsZero(m.InstallSandboxRuns[i]) { // not required
				return nil
			}

			if err := m.InstallSandboxRuns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_sandbox_runs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_sandbox_runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstall) UnmarshalBinary(b []byte) error {
	var res AppInstall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
