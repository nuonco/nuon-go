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
	}, c.getOrgIDAuthInfo())
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

func (c *client) GetVCSConnections(ctx context.Context, query *models.GetVCSConnectionsQuery) ([]*models.AppVCSConnection, bool, error) {
	params := &operations.GetOrgVCSConnectionsParams{
		Context: ctx,
	}

	if query != nil {
		offset := int64(query.Offset)
		limit := int64(query.Limit)
		params.Offset = &offset
		params.Limit = &limit
		params.XNuonPaginationEnabled = &query.PaginationEnabled
	}

	resp, err := c.genClient.Operations.GetOrgVCSConnections(params, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, false, err
	}

	if query != nil {
		items, hasMore := handlePagination(resp.Payload, int64(query.Offset), int64(query.Limit))
		return items, hasMore, nil
	}

	return resp.Payload, false, nil
}

func (c *client) GetVCSConnection(ctx context.Context, connID string) (*models.AppVCSConnection, error) {
	resp, err := c.genClient.Operations.GetVCSConnection(&operations.GetVCSConnectionParams{
		ConnectionID: connID,
		Context:      ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAllVCSConnectedRepos(ctx context.Context, query *models.GetAllVCSConnectedReposQuery) ([]*models.ServiceRepository, bool, error) {
	params := &operations.GetAllVCSConnectedReposParams{
		Context: ctx,
	}

	if query != nil {
		offset := int64(query.Offset)
		limit := int64(query.Limit)
		params.Offset = &offset
		params.Limit = &limit
		params.XNuonPaginationEnabled = &query.PaginationEnabled
	}

	resp, err := c.genClient.Operations.GetAllVCSConnectedRepos(params, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, false, err
	}

	if query != nil {
		items, hasMore := handlePagination(resp.Payload, int64(query.Offset), int64(query.Limit))
		return items, hasMore, nil
	}

	return resp.Payload, false, nil
}
