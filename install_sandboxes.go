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
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallSandboxRunLogs(ctx context.Context, installID, runID string) ([]interface{}, error) {
	resp, err := c.genClient.Operations.GetInstallSandboxRunLogs(&operations.GetInstallSandboxRunLogsParams{
		InstallID: installID,
		RunID:     runID,
		Context:   ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
