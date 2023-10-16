// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest github com powertoolsdev mono services ctl api internal app installs service create install request
//
// swagger:model github_com_powertoolsdev_mono_services_ctl-api_internal_app_installs_service.CreateInstallRequest
type GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest struct {

	// aws account
	// Required: true
	AwsAccount *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount `json:"aws_account"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this github com powertoolsdev mono services ctl api internal app installs service create install request
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) validateAwsAccount(formats strfmt.Registry) error {

	if err := validate.Required("aws_account", "body", m.AwsAccount); err != nil {
		return err
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

func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this github com powertoolsdev mono services ctl api internal app installs service create install request based on the context it is used
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) contextValidateAwsAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsAccount != nil {

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

// MarshalBinary interface implementation
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest) UnmarshalBinary(b []byte) error {
	var res GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount github com powertoolsdev mono services ctl API internal app installs service create install request aws account
//
// swagger:model GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount
type GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount struct {

	// iam role arn
	// Required: true
	IamRoleArn *string `json:"iam_role_arn"`

	// region
	// Enum: [us-east-1 us-east-2 us-west-2]
	Region string `json:"region,omitempty"`
}

// Validate validates this github com powertoolsdev mono services ctl API internal app installs service create install request aws account
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIamRoleArn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) validateIamRoleArn(formats strfmt.Registry) error {

	if err := validate.Required("aws_account"+"."+"iam_role_arn", "body", m.IamRoleArn); err != nil {
		return err
	}

	return nil
}

var githubComPowertoolsdevMonoServicesCtlApiInternalAppInstallsServiceCreateInstallRequestAwsAccountTypeRegionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["us-east-1","us-east-2","us-west-2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComPowertoolsdevMonoServicesCtlApiInternalAppInstallsServiceCreateInstallRequestAwsAccountTypeRegionPropEnum = append(githubComPowertoolsdevMonoServicesCtlApiInternalAppInstallsServiceCreateInstallRequestAwsAccountTypeRegionPropEnum, v)
	}
}

const (

	// GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccountRegionUsDashEastDash1 captures enum value "us-east-1"
	GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccountRegionUsDashEastDash1 string = "us-east-1"

	// GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccountRegionUsDashEastDash2 captures enum value "us-east-2"
	GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccountRegionUsDashEastDash2 string = "us-east-2"

	// GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccountRegionUsDashWestDash2 captures enum value "us-west-2"
	GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccountRegionUsDashWestDash2 string = "us-west-2"
)

// prop value enum
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) validateRegionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, githubComPowertoolsdevMonoServicesCtlApiInternalAppInstallsServiceCreateInstallRequestAwsAccountTypeRegionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	// value enum
	if err := m.validateRegionEnum("aws_account"+"."+"region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this github com powertoolsdev mono services ctl API internal app installs service create install request aws account based on context it is used
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount) UnmarshalBinary(b []byte) error {
	var res GithubComPowertoolsdevMonoServicesCtlAPIInternalAppInstallsServiceCreateInstallRequestAwsAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
