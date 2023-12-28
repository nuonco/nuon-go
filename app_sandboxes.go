package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) CreateAppSandboxConfig(ctx context.Context, appID string, req *models.ServiceCreateAppSandboxConfigRequest) (*models.AppAppSandboxConfig, error) {
	resp, err := c.genClient.Operations.CreateAppSandboxConfig(&operations.CreateAppSandboxConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppSandboxLatestConfig(ctx context.Context, appID string) (*models.AppAppSandboxConfig, error) {
	resp, err := c.genClient.Operations.GetAppSandboxLatestConfig(&operations.GetAppSandboxLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppSandboxConfigs(ctx context.Context, appID string) ([]*models.AppAppSandboxConfig, error) {
	resp, err := c.genClient.Operations.GetAppSandboxConfigs(&operations.GetAppSandboxConfigsParams{
		AppID:   appID,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
