package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// components
func (c *client) GetAllComponents(ctx context.Context) ([]*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetOrgComponents(&operations.GetOrgComponentsParams{
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppComponents(ctx context.Context, appID string) ([]*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetAppComponents(&operations.GetAppComponentsParams{
		AppID:   appID,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppComponent(ctx context.Context, appID, nameOrID string) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetAppComponent(&operations.GetAppComponentParams{
		AppID:             appID,
		ComponentNameOrID: nameOrID,
		Context:           ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateComponent(ctx context.Context, appID string, req *models.ServiceCreateComponentRequest) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.CreateComponent(&operations.CreateComponentParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetComponent(ctx context.Context, componentID string) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.GetComponent(&operations.GetComponentParams{
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateComponent(ctx context.Context, componentID string, req *models.ServiceUpdateComponentRequest) (*models.AppComponent, error) {
	resp, err := c.genClient.Operations.UpdateComponent(&operations.UpdateComponentParams{
		Req:         req,
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteComponent(ctx context.Context, componentID string) (bool, error) {
	resp, err := c.genClient.Operations.DeleteComponent(&operations.DeleteComponentParams{
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return false, err
	}

	return resp.Payload, nil
}

// component configs
func (c *client) CreateTerraformModuleComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateTerraformModuleComponentConfigRequest) (*models.AppTerraformModuleComponentConfig, error) {
	resp, err := c.genClient.Operations.CreateTerraformModuleComponentConfig(&operations.CreateTerraformModuleComponentConfigParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateHelmComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateHelmComponentConfigRequest) (*models.AppHelmComponentConfig, error) {
	resp, err := c.genClient.Operations.CreateHelmComponentConfig(&operations.CreateHelmComponentConfigParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateDockerBuildComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateDockerBuildComponentConfigRequest) (*models.AppDockerBuildComponentConfig, error) {
	resp, err := c.genClient.Operations.CreateDockerBuildComponentConfig(&operations.CreateDockerBuildComponentConfigParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateExternalImageComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateExternalImageComponentConfigRequest) (*models.AppExternalImageComponentConfig, error) {
	resp, err := c.genClient.Operations.CreateExternalImageComponentConfig(&operations.CreateExternalImageComponentConfigParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to create external image component config: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) CreateJobComponentConfig(ctx context.Context, componentID string, req *models.ServiceCreateJobComponentConfigRequest) (*models.AppJobComponentConfig, error) {
	resp, err := c.genClient.Operations.CreateJobComponentConfig(&operations.CreateJobComponentConfigParams{
		ComponentID: componentID,
		Req:         req,
		Context:     ctx,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create job component config: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentConfigs(ctx context.Context, componentID string) ([]*models.AppComponentConfigConnection, error) {
	resp, err := c.genClient.Operations.GetComponentConfigs(&operations.GetComponentConfigsParams{
		ComponentID: componentID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, fmt.Errorf("unable to get component configs: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetComponentLatestConfig(ctx context.Context, componentID string) (*models.AppComponentConfigConnection, error) {
	resp, err := c.genClient.Operations.GetComponentLatestConfig(&operations.GetComponentLatestConfigParams{
		ComponentID: componentID,
		Context:     ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
