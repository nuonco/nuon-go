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

// ConfigAppPolicyType config app policy type
//
// swagger:model config.AppPolicyType
type ConfigAppPolicyType string

func NewConfigAppPolicyType(value ConfigAppPolicyType) *ConfigAppPolicyType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigAppPolicyType.
func (m ConfigAppPolicyType) Pointer() *ConfigAppPolicyType {
	return &m
}

const (

	// ConfigAppPolicyTypeKubernetesCluster captures enum value "kubernetes_cluster"
	ConfigAppPolicyTypeKubernetesCluster ConfigAppPolicyType = "kubernetes_cluster"

	// ConfigAppPolicyTypeRunnerJobTerraformDeploy captures enum value "runner_job_terraform_deploy"
	ConfigAppPolicyTypeRunnerJobTerraformDeploy ConfigAppPolicyType = "runner_job_terraform_deploy"

	// ConfigAppPolicyTypeRunnerJobHelmDeploy captures enum value "runner_job_helm_deploy"
	ConfigAppPolicyTypeRunnerJobHelmDeploy ConfigAppPolicyType = "runner_job_helm_deploy"

	// ConfigAppPolicyTypeRunnerJobActionWorkflow captures enum value "runner_job_action_workflow"
	ConfigAppPolicyTypeRunnerJobActionWorkflow ConfigAppPolicyType = "runner_job_action_workflow"
)

// for schema
var configAppPolicyTypeEnum []interface{}

func init() {
	var res []ConfigAppPolicyType
	if err := json.Unmarshal([]byte(`["kubernetes_cluster","runner_job_terraform_deploy","runner_job_helm_deploy","runner_job_action_workflow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configAppPolicyTypeEnum = append(configAppPolicyTypeEnum, v)
	}
}

func (m ConfigAppPolicyType) validateConfigAppPolicyTypeEnum(path, location string, value ConfigAppPolicyType) error {
	if err := validate.EnumCase(path, location, value, configAppPolicyTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config app policy type
func (m ConfigAppPolicyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigAppPolicyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config app policy type based on context it is used
func (m ConfigAppPolicyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
