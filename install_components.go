package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

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
