package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// releases methods
func (c *client) GetRelease(ctx context.Context, releaseID string) (*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.GetV1ReleasesReleaseID(&operations.GetV1ReleasesReleaseIDParams{
		ReleaseID: releaseID,
		Context:   ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get release: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetReleaseSteps(ctx context.Context, releaseID string) ([]*models.AppComponentReleaseStep, error) {
	resp, err := c.genClient.Operations.GetV1ReleasesReleaseIDSteps(&operations.GetV1ReleasesReleaseIDStepsParams{
		ReleaseID: releaseID,
		Context:   ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get release steps: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetAppReleases(ctx context.Context, appID string) ([]*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDReleases(&operations.GetV1AppsAppIDReleasesParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get app releases: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentReleases(ctx context.Context, componentID string) ([]*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDReleases(&operations.GetV1ComponentsComponentIDReleasesParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get component releases: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) CreateComponentRelease(ctx context.Context, componentID string, req *models.ServiceCreateComponentReleaseRequest) (*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDReleases(&operations.PostV1ComponentsComponentIDReleasesParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create component release: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) CreateRelease(ctx context.Context, req *models.ServiceCreateComponentReleaseRequest) (*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.PostV1Releases(&operations.PostV1ReleasesParams{
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create release: %w", err)
	}

	return resp.Payload, nil
}
