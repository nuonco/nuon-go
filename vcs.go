package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// vcs connections
func (c *client) CreateVCSConnection(ctx context.Context, req *models.ServiceCreateConnectionRequest) (*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.CreateVCSConnection(&operations.CreateVCSConnectionParams{
		Req:     req,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateVCSConnectionCallback(ctx context.Context, req *models.ServiceCreateConnectionCallbackRequest) (*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.CreateVCSConnectionCallback(&operations.CreateVCSConnectionCallbackParams{
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetVCSConnections(ctx context.Context) ([]*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.GetOrgVCSConnections(&operations.GetOrgVCSConnectionsParams{
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetVCSConnection(ctx context.Context, connID string) (*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.GetVCSConnection(&operations.GetVCSConnectionParams{
		ConnectionID: connID,
		Context:      ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAllVCSConnectedRepos(ctx context.Context) ([]*models.ServiceRepository, error) {
	resp, err := c.genClient.Operations.GetAllVCSConnectedRepos(&operations.GetAllVCSConnectedReposParams{
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
