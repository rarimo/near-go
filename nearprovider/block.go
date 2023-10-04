package nearprovider

import (
	"context"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *provider) GetLastKnownBlockHeight(ctx context.Context) (*common.BlockHeight, error) {
	resp, err := p.c.BlockDetails(ctx, nearclient.FinalityFinal())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get last known block height")
	}

	return &resp.Header.Height, nil
}
