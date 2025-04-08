// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppAppSecretConfig app app secret config
//
// swagger:model app.AppSecretConfig
type AppAppSecretConfig struct {

	// app config id
	AppConfigID string `json:"app_config_id,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// app secrets config id
	AppSecretsConfigID string `json:"app_secrets_config_id,omitempty"`

	// cloudformation stack name
	CloudformationStackName string `json:"cloudformation_stack_name,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// kubernetes secret name
	KubernetesSecretName string `json:"kubernetes_secret_name,omitempty"`

	// for syncing into kubernetes
	KubernetesSecretNamespace string `json:"kubernetes_secret_namespace,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app app secret config
func (m *AppAppSecretConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app app secret config based on context it is used
func (m *AppAppSecretConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppAppSecretConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppSecretConfig) UnmarshalBinary(b []byte) error {
	var res AppAppSecretConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
