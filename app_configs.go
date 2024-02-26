package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetAppConfigTemplate(ctx context.Context, appID string, typ models.ServiceAppTemplateType) (*models.ServiceAppTemplate, error) {
	resp, err := c.genClient.Operations.GetAppConfigTemplate(&operations.GetAppConfigTemplateParams{
		AppID:   appID,
		Type:    string(typ),
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateAppConfig(ctx context.Context, appID string, req *models.ServiceCreateAppConfigRequest) (*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.CreateAppConfig(&operations.CreateAppConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppLatestConfig(ctx context.Context, appID string) (*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.GetAppLatestConfig(&operations.GetAppLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppConfigs(ctx context.Context, appID string) ([]*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.GetAppConfigs(&operations.GetAppConfigsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
