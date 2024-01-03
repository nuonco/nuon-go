package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// install components
func (c *client) GetInstallComponents(ctx context.Context, installID string) ([]*models.AppInstallComponent, error) {
	resp, err := c.genClient.Operations.GetInstallComponents(&operations.GetInstallComponentsParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallComponentDeploys(ctx context.Context, installID string, componentID string) ([]*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetInstallComponentDeploys(&operations.GetInstallComponentDeploysParams{
		ComponentID: componentID,
		InstallID:   installID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallComponentLatestDeploy(ctx context.Context, installID string, componentID string) (*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.GetInstallComponentLatestDeploy(&operations.GetInstallComponentLatestDeployParams{
		ComponentID: componentID,
		InstallID:   installID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallDeployPlan(ctx context.Context, installID string, deployID string) (*models.Planv1Plan, error) {
	resp, err := c.genClient.Operations.GetInstallDeployPlan(&operations.GetInstallDeployPlanParams{
		Context:   ctx,
		InstallID: installID,
		DeployID:  deployID,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) TeardownInstallComponent(ctx context.Context, installID, componentID string) (*models.AppInstallDeploy, error) {
	resp, err := c.genClient.Operations.TeardownInstallComponent(&operations.TeardownInstallComponentParams{
		InstallID:   installID,
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
