package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetInstallSandboxRuns(ctx context.Context, installID string) ([]*models.AppInstallSandboxRun, error) {
	resp, err := c.genClient.Operations.GetInstallSandboxRuns(&operations.GetInstallSandboxRunsParams{
		InstallID: installID,
		Context:   ctx,
	}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
