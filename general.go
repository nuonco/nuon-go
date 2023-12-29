package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// general methods
func (c *client) GetCurrentUser(ctx context.Context) (*models.AppUserToken, error) {
	resp, err := c.genClient.Operations.GetCurrentUser(&operations.GetCurrentUserParams{
		Context: ctx,
	}, c.getApiKeyAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) PublishMetrics(ctx context.Context, req []*models.ServicePublishMetricInput) error {
	_, err := c.genClient.Operations.PublishMetrics(&operations.PublishMetricsParams{
		Req:     req,
		Context: ctx,
	}, c.getApiKeyAuthInfo())
	if err != nil {
		return err
	}

	return nil
}

func (c *client) GetCLIConfig(ctx context.Context) (*models.ServiceCLIConfig, error) {
	resp, err := c.genClient.Operations.GetCLIConfig(&operations.GetCLIConfigParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
