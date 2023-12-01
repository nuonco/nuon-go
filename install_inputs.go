package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetInstallCurrentInputs(ctx context.Context, installID string) (*models.AppInstallInputs, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDInputsCurrent(&operations.GetV1InstallsInstallIDInputsCurrentParams{
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallInputs(ctx context.Context, appID string) ([]*models.AppInstallInputs, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDInputs(&operations.GetV1InstallsInstallIDInputsParams{
		InstallID: appID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateInstallInputs(ctx context.Context, installID string, req *models.ServiceCreateInstallInputsRequest) (*models.AppInstallInputs, error) {
	resp, err := c.genClient.Operations.PostV1InstallsInstallIDInputs(&operations.PostV1InstallsInstallIDInputsParams{
		InstallID: installID,
		Req:       req,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
