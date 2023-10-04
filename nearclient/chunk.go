package nearclient

import (
	"context"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient/models"
)

// ChunkDetails https://docs.near.org/docs/api/rpc#chunk-details
func (c *Client) ChunkDetails(ctx context.Context, chunkHash common.Hash) (res models.ChunkView, err error) {
	_, err = c.doRPC(ctx, &res, "chunk", nil, []string{chunkHash.String()})

	return
}
