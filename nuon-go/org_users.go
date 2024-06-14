package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetOrgInvites(ctx context.Context, limit *int64) ([]*models.AppOrgInvite, error) {
	resp, err := c.genClient.Operations.GetOrgInvites(&operations.GetOrgInvitesParams{
		Limit:   limit,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateOrgInvite(ctx context.Context, req *models.ServiceCreateOrgInviteRequest) (*models.AppOrgInvite, error) {
	resp, err := c.genClient.Operations.CreateOrgInvite(&operations.CreateOrgInviteParams{
		Req:     req,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
