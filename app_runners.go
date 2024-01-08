package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) CreateAppRunnerConfig(ctx context.Context, appID string, req *models.ServiceCreateAppRunnerConfigRequest) (*models.AppAppRunnerConfig, error) {
	resp, err := c.genClient.Operations.CreateAppRunnerConfig(&operations.CreateAppRunnerConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppRunnerLatestConfig(ctx context.Context, appID string) (*models.AppAppRunnerConfig, error) {
	resp, err := c.genClient.Operations.GetAppRunnerLatestConfig(&operations.GetAppRunnerLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppRunnerConfigs(ctx context.Context, appID string) ([]*models.AppAppRunnerConfig, error) {
	resp, err := c.genClient.Operations.GetAppRunnerConfigs(&operations.GetAppRunnerConfigsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
