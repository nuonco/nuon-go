package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) CreateAppSandboxConfig(ctx context.Context, appID string, req *models.ServiceCreateAppSandboxConfigRequest) (*models.AppAppSandboxConfig, error) {
	resp, err := c.genClient.Operations.PostV1AppsAppIDSandboxConfig(&operations.PostV1AppsAppIDSandboxConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppSandboxLatestConfig(ctx context.Context, appID string) (*models.AppAppSandboxConfig, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDSandboxLatestConfig(&operations.GetV1AppsAppIDSandboxLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppSandboxConfigs(ctx context.Context, appID string) ([]*models.AppAppSandboxConfig, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDSandboxConfigs(&operations.GetV1AppsAppIDSandboxConfigsParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
