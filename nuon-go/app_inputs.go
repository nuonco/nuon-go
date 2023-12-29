package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) CreateAppInputConfig(ctx context.Context, appID string, req *models.ServiceCreateAppInputConfigRequest) (*models.AppAppInputConfig, error) {
	resp, err := c.genClient.Operations.CreateAppInputConfig(&operations.CreateAppInputConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInputLatestConfig(ctx context.Context, appID string) (*models.AppAppInputConfig, error) {
	resp, err := c.genClient.Operations.GetAppInputLatestConfig(&operations.GetAppInputLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInputConfigs(ctx context.Context, appID string) ([]*models.AppAppInputConfig, error) {
	resp, err := c.genClient.Operations.GetAppInputConfigs(&operations.GetAppInputConfigsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
