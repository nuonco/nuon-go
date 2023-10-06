package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// vcs connections
func (c *client) CreateVCSConnection(ctx context.Context, req *models.ServiceCreateConnectionRequest) (*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.PostV1VcsConnections(&operations.PostV1VcsConnectionsParams{
		Req:     req,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetVCSConnections(ctx context.Context) ([]*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.GetV1VcsConnections(&operations.GetV1VcsConnectionsParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetVCSConnection(ctx context.Context, connID string) (*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.GetV1VcsConnectionsConnectionID(&operations.GetV1VcsConnectionsConnectionIDParams{
		ConnectionID: connID,
		Context:      ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAllVCSConnectedRepos(ctx context.Context) ([]*models.ServiceRepository, error) {
	resp, err := c.genClient.Operations.GetV1VcsConnectedRepos(&operations.GetV1VcsConnectedReposParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
