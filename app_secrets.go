package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetAppSecrets(ctx context.Context, appID string) ([]*models.AppAppSecret, error) {
	resp, err := c.genClient.Operations.GetAppSecrets(&operations.GetAppSecretsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateAppSecret(ctx context.Context, appID string, req *models.ServiceCreateAppSecretRequest) (*models.AppAppSecret, error) {
	resp, err := c.genClient.Operations.CreateAppSecret(&operations.CreateAppSecretParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteAppSecret(ctx context.Context, appID, secretID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteAppSecret(&operations.DeleteAppSecretParams{
		AppID:    appID,
		SecretID: secretID,
		Context:  ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}
