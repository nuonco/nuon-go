package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetActionWorkflows(ctx context.Context, appID string) ([]*models.AppActionWorkflow, error) {
	resp, err := c.genClient.Operations.GetActionWorkflows(&operations.GetActionWorkflowsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetActionWorkflow(ctx context.Context, actionWorkflowID string) (*models.AppActionWorkflow, error) {
	resp, err := c.genClient.Operations.GetActionWorkflow(&operations.GetActionWorkflowParams{
		ActionWorkflowID: actionWorkflowID,
		Context:          ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateActionWorkflow(ctx context.Context, appID string, req *models.ServiceCreateAppActionWorkflowRequest) (*models.AppActionWorkflow, error) {
	resp, err := c.genClient.Operations.CreateAppActionWorkflow(&operations.CreateAppActionWorkflowParams{
		AppID:   appID,
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateActionWorkflow(ctx context.Context, actionWorkflowID string, req *models.ServiceUpdateActionWorkflowRequest) (*models.AppActionWorkflow, error) {
	resp, err := c.genClient.Operations.UpdateAppActionWorkflow(&operations.UpdateAppActionWorkflowParams{
		ActionWorkflowID: actionWorkflowID,
		Req:              req,
		Context:          ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteActionWorkflow(ctx context.Context, actionWorkflowID string) error {
	_, err := c.genClient.Operations.DeleteActionWorkflow(&operations.DeleteActionWorkflowParams{
		ActionWorkflowID: actionWorkflowID,
		Context:          ctx,
	}, c.getOrgIDAuthInfo())

	return err
}

func (c *client) GetActionWorkflowConfigs(ctx context.Context, actionWorkflowID string) ([]*models.AppActionWorkflowConfig, error) {
	resp, err := c.genClient.Operations.GetActionWorkflowConfigs(&operations.GetActionWorkflowConfigsParams{
		ActionWorkflowID: actionWorkflowID,
		Context:          ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetActionWorkflowConfig(ctx context.Context, actionWorkflowConfigID string) (*models.AppActionWorkflowConfig, error) {
	resp, err := c.genClient.Operations.GetActionWorkflowConfig(&operations.GetActionWorkflowConfigParams{
		ActionWorkflowConfigID: actionWorkflowConfigID,
		Context:                ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateActionWorkflowConfig(ctx context.Context, actionWorkflowID string, req *models.ServiceCreateActionWorkflowConfigRequest) (*models.AppActionWorkflowConfig, error) {
	resp, err := c.genClient.Operations.CreateActionWorkflowConfig(&operations.CreateActionWorkflowConfigParams{
		ActionWorkflowID: actionWorkflowID,
		Req:              req,
		Context:          ctx,
	}, c.getOrgIDAuthInfo())

	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
