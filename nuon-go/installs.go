package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// installs
func (c *client) CreateInstall(ctx context.Context, appID string, req *models.ServiceCreateInstallRequest) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.CreateInstall(&operations.CreateInstallParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppInstalls(ctx context.Context, appID string) ([]*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetAppInstalls(&operations.GetAppInstallsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAllInstalls(ctx context.Context) ([]*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetOrgInstalls(&operations.GetOrgInstallsParams{
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstall(ctx context.Context, installID string) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.GetInstall(&operations.GetInstallParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateInstall(ctx context.Context, installID string, req *models.ServiceUpdateInstallRequest) (*models.AppInstall, error) {
	resp, err := c.genClient.Operations.UpdateInstall(&operations.UpdateInstallParams{
		Req:       req,
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to update install: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) DeleteInstall(ctx context.Context, installID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteInstall(&operations.DeleteInstallParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}

func (c *client) ReprovisionInstall(ctx context.Context, installID string) error {
	resp, err := c.genClient.Operations.ReprovisionInstall(&operations.ReprovisionInstallParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return err
	}

	if resp.Payload != "ok" {
		return statusErr{resp.Payload}
	}

	return nil
}

func (c *client) DeprovisionInstall(ctx context.Context, installID string) error {
	resp, err := c.genClient.Operations.DeprovisionInstall(&operations.DeprovisionInstallParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return err
	}

	if resp.Payload != "ok" {
		return statusErr{resp.Payload}
	}

	return nil
}
