package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) RenderInstaller(ctx context.Context, slug string) (*models.ServiceRenderedInstaller, error) {
	resp, err := c.genClient.Operations.RenderInstaller(&operations.RenderInstallerParams{
		Context:       ctx,
		InstallerSlug: slug,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) RenderInstallerInstall(ctx context.Context, slug, installID string) (*models.ServiceRenderedInstall, error) {
	resp, err := c.genClient.Operations.RenderInstallerInstall(&operations.RenderInstallerInstallParams{
		Context:       ctx,
		InstallerSlug: slug,
		InstallID:     installID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallerInstall(ctx context.Context, slug, installID string) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetInstallerInstall(&operations.GetInstallerInstallParams{
		Context:       ctx,
		InstallerSlug: slug,
		InstallID:     installID,
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

func (c *client) GetInstaller(ctx context.Context, appInstallerID string) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.GetInstaller(&operations.GetInstallerParams{
		Context:     ctx,
		InstallerID: appInstallerID,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateInstaller(ctx context.Context, req *models.ServiceCreateAppInstallerRequest) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.CreateInstaller(&operations.CreateInstallerParams{
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateInstaller(ctx context.Context, installerID string, req *models.ServiceUpdateInstallerRequest) (*models.AppAppInstaller, error) {
	resp, err := c.genClient.Operations.UpdateInstaller(&operations.UpdateInstallerParams{
		InstallerID: installerID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteInstaller(ctx context.Context, installerID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteInstaller(&operations.DeleteInstallerParams{
		InstallerID: installerID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}

func (c *client) CreateInstallerInstall(ctx context.Context, req *models.ServiceCreateInstallRequest) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.InstallerCreateInstall(&operations.InstallerCreateInstallParams{
		Context: ctx,
		Req:     req,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
