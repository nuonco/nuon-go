package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) RenderAppInstaller(ctx context.Context, slug string) (*models.ServiceAppInstaller, error) {
	resp, err := c.genClient.Operations.GetV1InstallerInstallerSlugRender(&operations.GetV1InstallerInstallerSlugRenderParams{
		Context:       ctx,
		InstallerSlug: slug,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInstaller(ctx context.Context, appInstallerID string) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.GetV1InstallersInstallerID(&operations.GetV1InstallersInstallerIDParams{
		Context:     ctx,
		InstallerID: appInstallerID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateAppInstaller(ctx context.Context, appID string, req *models.ServiceCreateAppInstallerRequest) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.PostV1AppsAppIDInstaller(&operations.PostV1AppsAppIDInstallerParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateAppInstaller(ctx context.Context, installerID string, req *models.ServiceUpdateAppInstallerRequest) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.PatchV1InstallersInstallerID(&operations.PatchV1InstallersInstallerIDParams{
		InstallerID: installerID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteAppInstaller(ctx context.Context, installerID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteV1InstallersInstallerID(&operations.DeleteV1InstallersInstallerIDParams{
		InstallerID: installerID,
		Context:     ctx,
	})
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}
