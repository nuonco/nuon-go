package nuon

import (
	"context"

	"github.com/nuonco/nuon-go/client/operations"
	"github.com/nuonco/nuon-go/models"
)

func (c *client) GetLatestRunnerHeartBeat(ctx context.Context, runnerID string) (models.ServiceLatestRunnerHeartBeats, error) {
	resp, err := c.genClient.Operations.
		GetLatestRunnerHeartBeat(&operations.GetLatestRunnerHeartBeatParams{
			RunnerID: runnerID,
			Context:  ctx,
		}, c.getOrgIDAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
