package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) RenderInstaller(ctx context.Context, id string) (*models.ServiceRenderedInstaller, error) {
	resp, err := c.genClient.Operations.RenderInstaller(&operations.RenderInstallerParams{
		Context:     ctx,
		InstallerID: id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallers(ctx context.Context) ([]*models.AppInstaller, error) {
	resp, err := c.genClient.Operations.GetInstallers(&operations.GetInstallersParams{
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstaller(ctx context.Context, appInstallerID string) (*models.AppInstaller, error) {
	resp, err := c.genClient.Operations.GetInstaller(&operations.GetInstallerParams{
		Context:     ctx,
		InstallerID: appInstallerID,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateInstaller(ctx context.Context, req *models.ServiceCreateInstallerRequest) (*models.AppInstaller, error) {
	resp, err := c.genClient.Operations.CreateInstaller(&operations.CreateInstallerParams{
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateInstaller(ctx context.Context, installerID string, req *models.ServiceUpdateInstallerRequest) (*models.AppInstaller, error) {
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
