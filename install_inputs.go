package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetInstallCurrentInputs(ctx context.Context, installID string) (*models.AppInstallInputs, error) {
	resp, err := c.genClient.Operations.GetCurrentInstallInputs(&operations.GetCurrentInstallInputsParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallInputs(ctx context.Context, appID string) ([]*models.AppInstallInputs, error) {
	resp, err := c.genClient.Operations.GetInstallInputs(&operations.GetInstallInputsParams{
		InstallID: appID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateInstallInputs(ctx context.Context, installID string, req *models.ServiceCreateInstallInputsRequest) (*models.AppInstallInputs, error) {
	resp, err := c.genClient.Operations.CreateInstallInputs(&operations.CreateInstallInputsParams{
		InstallID: installID,
		Req:       req,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
