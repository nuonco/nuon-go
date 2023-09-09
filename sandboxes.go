package nuon

import (
	"context"
	"fmt"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// sandbox methods
func (c *client) GetSandboxes(ctx context.Context) ([]*models.AppSandbox, error) {
	resp, err := c.genClient.Operations.GetV1Sandboxes(&operations.GetV1SandboxesParams{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get sandboxes: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetSandbox(ctx context.Context, sandboxID string) (*models.AppSandbox, error) {
	resp, err := c.genClient.Operations.GetV1SandboxesSandboxID(&operations.GetV1SandboxesSandboxIDParams{
		SandboxID: sandboxID,
		Context:   ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get sandbox: %w", err)
	}

	return resp.Payload, nil
}

func (c *client) GetSandboxReleases(ctx context.Context, sandboxID string) ([]*models.AppSandboxRelease, error) {
	resp, err := c.genClient.Operations.GetV1SandboxesSandboxIDReleases(&operations.GetV1SandboxesSandboxIDReleasesParams{
		SandboxID: sandboxID,
		Context:   ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to get sandbox releases: %w", err)
	}

	return resp.Payload, nil
}
