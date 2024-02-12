package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
	"github.com/powertoolsdev/mono/pkg/generics"
)

// builds
func (c *client) CreateComponentBuild(ctx context.Context, componentID string, req *models.ServiceCreateComponentBuildRequest) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.CreateComponentBuild(&operations.CreateComponentBuildParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentBuilds(ctx context.Context, componentID, appID string, limit *int64) ([]*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetComponentBuilds(&operations.GetComponentBuildsParams{
		ComponentID: generics.ToPtr(componentID),
		AppID:       generics.ToPtr(appID),
		Limit:       limit,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentLatestBuild(ctx context.Context, componentID string) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetComponentLatestBuild(&operations.GetComponentLatestBuildParams{
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentBuild(ctx context.Context, componentID string, buildID string) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetComponentBuild(&operations.GetComponentBuildParams{
		BuildID:     buildID,
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to get component build: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetBuild(ctx context.Context, buildID string) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetBuild(&operations.GetBuildParams{
		BuildID: buildID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to get build: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentBuildLogs(ctx context.Context, componentID string, buildID string) ([]models.ServiceBuildLog, error) {
	resp, err := c.genClient.Operations.GetComponentBuildLogs(&operations.GetComponentBuildLogsParams{
		ComponentID: componentID,
		BuildID:     buildID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to get build logs: %w", err)
	}

	// need to assert the type of each slice item individually
	response := make([]models.ServiceBuildLog, len(resp.Payload))
	for idx, item := range resp.Payload {
		response[idx] = item.(models.ServiceBuildLog)
	}

	return response, nil
}

func (c *client) GetComponentBuildPlan(ctx context.Context, componentID, buildID string) (*models.Planv1Plan, error) {
	resp, err := c.genClient.Operations.GetComponentBuildPlan(&operations.GetComponentBuildPlanParams{
		ComponentID: componentID,
		BuildID:     buildID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to get build plan: %w", err)
	}

	return resp.Payload, nil
}
