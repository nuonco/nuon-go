package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) CreateAppCloudFormationStackConfig(ctx context.Context, appID string, req *models.ServiceCreateAppCloudFormationStackConfigRequest) (*models.AppAppCloudFormationStackConfig, error) {
	resp, err := c.genClient.Operations.CreateAppCloudFormationStackConfig(&operations.CreateAppCloudFormationStackConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppCloudFormationStackConfig(ctx context.Context, appID, configID string) (*models.AppAppCloudFormationStackConfig, error) {
	params := &operations.GetAppCloudFormationStackConfigParams{
		AppID:    appID,
		ConfigID: configID,
		Context:  ctx,
	}

	resp, err := c.genClient.Operations.GetAppCloudFormationStackConfig(params, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
