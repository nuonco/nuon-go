package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetOrg(ctx context.Context) (*models.AppOrg, error) {
	resp, err := c.genClient.Operations.GetOrg(&operations.GetOrgParams{
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetOrgs(ctx context.Context) ([]*models.AppOrg, error) {
	resp, err := c.genClient.Operations.GetOrgs(&operations.GetOrgsParams{
		Context: ctx,
	}, c.getApiKeyAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) DeleteOrg(ctx context.Context) (bool, error) {
	resp, err := c.genClient.Operations.DeleteOrg(&operations.DeleteOrgParams{
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return false, err
	}

	return resp.Payload, nil
}

func (c *client) CreateOrg(ctx context.Context, req *models.ServiceCreateOrgRequest) (*models.AppOrg, error) {
	resp, err := c.genClient.Operations.CreateOrg(&operations.CreateOrgParams{
		Req:     req,
		Context: ctx,
	}, c.getApiKeyAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateOrg(ctx context.Context, req *models.ServiceUpdateOrgRequest) (*models.AppOrg, error) {
	resp, err := c.genClient.Operations.UpdateOrg(&operations.UpdateOrgParams{
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
