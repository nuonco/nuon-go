// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppInstallerMetadata app installer metadata
//
// swagger:model app.InstallerMetadata
type AppInstallerMetadata struct {

	// community url
	CommunityURL string `json:"community_url,omitempty"`

	// copyright markdown
	CopyrightMarkdown string `json:"copyright_markdown,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// demo url
	DemoURL string `json:"demo_url,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// documentation url
	DocumentationURL string `json:"documentation_url,omitempty"`

	// favicon url
	FaviconURL string `json:"favicon_url,omitempty"`

	// footer markdown
	FooterMarkdown string `json:"footer_markdown,omitempty"`

	// formatted demo url
	FormattedDemoURL string `json:"formatted_demo_url,omitempty"`

	// github url
	GithubURL string `json:"github_url,omitempty"`

	// homepage url
	HomepageURL string `json:"homepage_url,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// installer id
	InstallerID string `json:"installer_id,omitempty"`

	// logo url
	LogoURL string `json:"logo_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// og image url
	OgImageURL string `json:"og_image_url,omitempty"`

	// post install markdown
	PostInstallMarkdown string `json:"post_install_markdown,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app installer metadata
func (m *AppInstallerMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app installer metadata based on context it is used
func (m *AppInstallerMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallerMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallerMetadata) UnmarshalBinary(b []byte) error {
	var res AppInstallerMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
