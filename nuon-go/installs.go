package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// installs
func (c *client) CreateInstall(ctx context.Context, appID string, req *models.ServiceCreateInstallRequest) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.PostV1AppsAppIDInstalls(&operations.PostV1AppsAppIDInstallsParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInstalls(ctx context.Context, appID string) ([]*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDInstalls(&operations.GetV1AppsAppIDInstallsParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAllInstalls(ctx context.Context) ([]*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetV1Installs(&operations.GetV1InstallsParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstall(ctx context.Context, installID string) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallID(&operations.GetV1InstallsInstallIDParams{
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateInstall(ctx context.Context, installID string, req *models.ServiceUpdateInstallRequest) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.PatchV1InstallsInstallID(&operations.PatchV1InstallsInstallIDParams{
		Req:       req,
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to update install: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) DeleteInstall(ctx context.Context, installID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteV1InstallsInstallID(&operations.DeleteV1InstallsInstallIDParams{
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}
