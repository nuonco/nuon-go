package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
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
		ComponentID: &componentID,
		AppID:       &appID,
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
