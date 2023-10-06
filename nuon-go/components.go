package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// components
func (c *client) GetAllComponents(ctx context.Context) ([]*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetV1Components(&operations.GetV1ComponentsParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppComponents(ctx context.Context, appID string) ([]*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetV1AppsAppIDComponents(&operations.GetV1AppsAppIDComponentsParams{
		AppID:   appID,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateComponent(ctx context.Context, appID string, req *models.ServiceCreateComponentRequest) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.PostV1AppsAppIDComponents(&operations.PostV1AppsAppIDComponentsParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponent(ctx context.Context, componentID string) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentID(&operations.GetV1ComponentsComponentIDParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateComponent(ctx context.Context, componentID string, req *models.ServiceUpdateComponentRequest) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.PatchV1ComponentsComponentID(&operations.PatchV1ComponentsComponentIDParams{
		Req:         req,
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteComponent(ctx context.Context, componentID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteV1ComponentsComponentID(&operations.DeleteV1ComponentsComponentIDParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return false, err
	}

	return resp.Payload, nil
}

// component configs
func (c *client) CreateTerraformModuleComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateTerraformModuleComponentConfigRequest) (*models.AppTerraformModuleComponentConfig, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDConfigsTerraformModule(&operations.PostV1ComponentsComponentIDConfigsTerraformModuleParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateHelmComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateHelmComponentConfigRequest) (*models.AppHelmComponentConfig, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDConfigsHelm(&operations.PostV1ComponentsComponentIDConfigsHelmParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateDockerBuildComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateDockerBuildComponentConfigRequest) (*models.AppDockerBuildComponentConfig, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDConfigsDockerBuild(&operations.PostV1ComponentsComponentIDConfigsDockerBuildParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateExternalImageComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateExternalImageComponentConfigRequest) (*models.AppExternalImageComponentConfig, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDConfigsExternalImage(&operations.PostV1ComponentsComponentIDConfigsExternalImageParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create external image component config: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) CreateJobComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateJobComponentConfigRequest) (*models.AppJobComponentConfig, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDConfigsJob(&operations.PostV1ComponentsComponentIDConfigsJobParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create job component config: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentConfigs(ctx context.Context, componentID string) ([]*models.AppComponentConfigConnection, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDConfigs(&operations.GetV1ComponentsComponentIDConfigsParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get component configs: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentLatestConfig(ctx context.Context, componentID string) (*models.AppComponentConfigConnection, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDConfigsLatest(&operations.GetV1ComponentsComponentIDConfigsLatestParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// builds
func (c *client) CreateComponentBuild(ctx context.Context, componentID string, req *models.ServiceCreateComponentBuildRequest) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.PostV1ComponentsComponentIDBuilds(&operations.PostV1ComponentsComponentIDBuildsParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentBuilds(ctx context.Context, componentID string) ([]*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDBuilds(&operations.GetV1ComponentsComponentIDBuildsParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentLatestBuild(ctx context.Context, componentID string) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDBuildsLatest(&operations.GetV1ComponentsComponentIDBuildsLatestParams{
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponentBuild(ctx context.Context, componentID string, buildID string) (*models.AppComponentBuild, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDBuildsBuildID(&operations.GetV1ComponentsComponentIDBuildsBuildIDParams{
		BuildID:     buildID,
		ComponentID: componentID,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get component build: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentBuildLogs(ctx context.Context, componentID string, buildID string) ([]models.ServiceBuildLog, error) {
	resp, err := c.genClient.Operations.GetV1ComponentsComponentIDBuildsBuildIDLogs(&operations.GetV1ComponentsComponentIDBuildsBuildIDLogsParams{
		ComponentID: componentID,
		BuildID:     buildID,
		Context:     ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get build logs: %w", err)
	}

	return resp.Payload, nil
}
