package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) CreateAppInputConfig(ctx context.Context, appID string, req *models.ServiceCreateAppInputConfigRequest) (*models.AppAppInputConfig, error) {
	resp, err := c.genClient.Operations.PostV1AppsAppIDInputConfig(&operations.PostV1AppsAppIDInputConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInputLatestConfig(ctx context.Context, appID string) (*models.AppAppInputConfig, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDInputLatestConfig(&operations.GetV1AppsAppIDInputLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInputConfigs(ctx context.Context, appID string) ([]*models.AppAppInputConfig, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDInputConfigs(&operations.GetV1AppsAppIDInputConfigsParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
