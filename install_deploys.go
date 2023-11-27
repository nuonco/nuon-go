package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

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
