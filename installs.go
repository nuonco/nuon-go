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

// install deploys
func (c *client) GetInstallDeploys(ctx context.Context, installID string) ([]*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDDeploys(&operations.GetV1InstallsInstallIDDeploysParams{
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateInstallDeploy(ctx context.Context, installID string, req *models.ServiceCreateInstallDeployRequest) (*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.PostV1InstallsInstallIDDeploys(&operations.PostV1InstallsInstallIDDeploysParams{
		InstallID: installID,
		Req:       req,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallLatestDeploy(ctx context.Context, installID string) (*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDDeploysLatest(&operations.GetV1InstallsInstallIDDeploysLatestParams{
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallDeploy(ctx context.Context, installID string, deployID string) (*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDDeploysDeployID(&operations.GetV1InstallsInstallIDDeploysDeployIDParams{
		InstallID: installID,
		DeployID:  deployID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallDeployLogs(ctx context.Context, installID string, deployID string) ([]models.ServiceDeployLog, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDDeploysDeployIDLogs(&operations.GetV1InstallsInstallIDDeploysDeployIDLogsParams{
		InstallID: installID,
		DeployID:  deployID,
		Context:   ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get install deploy logs: %w", err)
	}

	// need to assert the type of each slice item individually
	response := make([]models.ServiceDeployLog, len(resp.Payload))
	for idx, item := range resp.Payload {
		response[idx] = item.(models.ServiceDeployLog)
	}

	return response, nil
}

// install components
func (c *client) GetInstallComponents(ctx context.Context, installID string) ([]*models.AppInstallComponent, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDComponents(&operations.GetV1InstallsInstallIDComponentsParams{
		InstallID: installID,
		Context:   ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallComponentDeploys(ctx context.Context, installID string, componentID string) ([]*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDComponentsComponentIDDeploys(&operations.GetV1InstallsInstallIDComponentsComponentIDDeploysParams{
		ComponentID: componentID,
		InstallID:   installID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallComponentLatestDeploy(ctx context.Context, installID string, componentID string) (*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDComponentsComponentIDDeploysLatest(&operations.GetV1InstallsInstallIDComponentsComponentIDDeploysLatestParams{
		ComponentID: componentID,
		InstallID:   installID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallDeployPlan(ctx context.Context, installID string, deployID string) (*models.Planv1Plan, error) {
	resp, err := c.genClient.Operations.GetV1InstallsInstallIDDeploysDeployIDPlan(&operations.GetV1InstallsInstallIDDeploysDeployIDPlanParams{
		Context:   ctx,
		InstallID: installID,
		DeployID:  deployID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
