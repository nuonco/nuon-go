package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// internal methods
func (c *client) GetApp(ctx context.Context, appID string) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppID(&operations.GetV1AppsAppIDParams{
		Context: ctx,
		AppID:   appID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetApps(ctx context.Context) ([]*models.AppApp, error) {
	resp, err := c.genClient.Operations.GetV1Apps(&operations.GetV1AppsParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateApp(ctx context.Context, req *models.ServiceCreateAppRequest) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.PostV1Apps(&operations.PostV1AppsParams{
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateApp(ctx context.Context, appID string, req *models.ServiceUpdateAppRequest) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.PatchV1AppsAppID(&operations.PatchV1AppsAppIDParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteApp(ctx context.Context, appID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteV1AppsAppID(&operations.DeleteV1AppsAppIDParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}

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
