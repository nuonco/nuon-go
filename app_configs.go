package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetAppConfigTemplate(ctx context.Context, appID string, typ models.ServiceAppConfigTemplateType) (*models.ServiceAppConfigTemplate, error) {
	resp, err := c.genClient.Operations.GetAppConfigTemplate(&operations.GetAppConfigTemplateParams{
		AppID:   appID,
		Type:    string(typ),
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppConfig(ctx context.Context, appID, appConfigID string) (*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.GetAppConfig(&operations.GetAppConfigParams{
		AppID:       appID,
		AppConfigID: appConfigID,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateAppConfig(ctx context.Context, appID string, req *models.ServiceCreateAppConfigRequest) (*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.CreateAppConfig(&operations.CreateAppConfigParams{
		Req:     req,
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppLatestConfig(ctx context.Context, appID string) (*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.GetAppLatestConfig(&operations.GetAppLatestConfigParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetAppConfigs(ctx context.Context, appID string, query *models.GetAppConfigsQuery) ([]*models.AppAppConfig, bool, error) {
	params := &operations.GetAppConfigsParams{
		AppID:   appID,
		Context: ctx,
	}

	if query != nil {
		offset := int64(query.Offset)
		limit := int64(query.Limit)
		params.Offset = &offset
		params.Limit = &limit
		params.XNuonPaginationEnabled = &query.PaginationEnabled
	}

	resp, err := c.genClient.Operations.GetAppConfigs(&operations.GetAppConfigsParams{
		AppID:   appID,
		Context: ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, false, err
	}

	if query != nil {
		items, hasMore := handlePagination(resp.Payload, int64(query.Offset), int64(query.Limit))
		return items, hasMore, nil
	}

	return resp.Payload, false, nil
}

func (c *client) UpdateAppConfig(ctx context.Context, appID, appConfigID string, req *models.ServiceUpdateAppConfigRequest) (*models.AppAppConfig, error) {
	resp, err := c.genClient.Operations.UpdateAppConfig(&operations.UpdateAppConfigParams{
		AppID:       appID,
		AppConfigID: appConfigID,
		Req:         req,
		Context:     ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
