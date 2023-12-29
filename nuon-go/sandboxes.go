package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

// sandbox methods
func (c *client) GetSandboxes(ctx context.Context) ([]*models.AppSandbox, error) {
	resp, err := c.genClient.Operations.GetSandboxes(&operations.GetSandboxesParams{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetSandbox(ctx context.Context, sandboxID string) (*models.AppSandbox, error) {
	resp, err := c.genClient.Operations.GetSandbox(&operations.GetSandboxParams{
		SandboxID: sandboxID,
		Context:   ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetSandboxReleases(ctx context.Context, sandboxID string) ([]*models.AppSandboxRelease, error) {
	resp, err := c.genClient.Operations.GetSandboxReleases(&operations.GetSandboxReleasesParams{
		SandboxID: sandboxID,
		Context:   ctx,
	}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
