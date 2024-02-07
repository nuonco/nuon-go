package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) RenderAppInstaller(ctx context.Context, slug string) (*models.ServiceAppInstaller, error) {
	resp, err := c.genClient.Operations.RenderAppInstaller(&operations.RenderAppInstallerParams{
		Context:       ctx,
		InstallerSlug: slug,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallers(ctx context.Context) ([]*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.GetInstallers(&operations.GetInstallersParams{
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInstaller(ctx context.Context, appInstallerID string) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.GetAppInstaller(&operations.GetAppInstallerParams{
		Context:     ctx,
		InstallerID: appInstallerID,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateAppInstaller(ctx context.Context, req *models.ServiceCreateAppInstallerRequest) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.CreateAppInstaller(&operations.CreateAppInstallerParams{
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateAppInstaller(ctx context.Context, installerID string, req *models.ServiceUpdateAppInstallerRequest) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.UpdateAppInstaller(&operations.UpdateAppInstallerParams{
		InstallerID: installerID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteAppInstaller(ctx context.Context, installerID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteAppInstaller(&operations.DeleteAppInstallerParams{
		InstallerID: installerID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}

func (c *client) CreateInstallFromInstaller(ctx context.Context, req *models.ServiceCreateInstallRequest) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.InstallerCreateInstall(&operations.InstallerCreateInstallParams{
		Context: ctx,
		Req:     req,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
