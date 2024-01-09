package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// releases methods
func (c *client) GetRelease(ctx context.Context, releaseID string) (*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.GetRelease(&operations.GetReleaseParams{
		ReleaseID: releaseID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetReleaseSteps(ctx context.Context, releaseID string) ([]*models.AppComponentReleaseStep, error) {
	resp, err := c.genClient.Operations.GetReleaseSteps(&operations.GetReleaseStepsParams{
		ReleaseID: releaseID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppReleases(ctx context.Context, appID string) ([]*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.GetAppReleases(&operations.GetAppReleasesParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentReleases(ctx context.Context, componentID string) ([]*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.GetComponentReleases(&operations.GetComponentReleasesParams{
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateComponentRelease(ctx context.Context, componentID string, req *models.ServiceCreateComponentReleaseRequest) (*models.AppComponentRelease, error) {
	resp, err := c.genClient.Operations.CreateComponentRelease(&operations.CreateComponentReleaseParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
