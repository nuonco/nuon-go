package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// internal methods
func (c *client) GetApp(ctx context.Context, appID string) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppID(&operations.GetV1AppsAppIDParams{
		Context: ctx,
		AppID:   appID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetApps(ctx context.Context) ([]*models.AppApp, error) {
	resp, err := c.genClient.Operations.GetV1Apps(&operations.GetV1AppsParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateApp(ctx context.Context, req *models.ServiceCreateAppRequest) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.PostV1Apps(&operations.PostV1AppsParams{
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateApp(ctx context.Context, appID string, req *models.ServiceUpdateAppRequest) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.PatchV1AppsAppID(&operations.PatchV1AppsAppIDParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteApp(ctx context.Context, appID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteV1AppsAppID(&operations.DeleteV1AppsAppIDParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return false, err
	}

	return resp.IsSuccess(), nil
}

func (c *client) UpdateAppSandbox(ctx context.Context, appID string, req *models.ServiceUpdateAppSandboxRequest) (*models.AppApp, error) {
	resp, err := c.genClient.Operations.PutV1AppsAppIDSandbox(&operations.PutV1AppsAppIDSandboxParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
