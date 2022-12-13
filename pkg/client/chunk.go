package client

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/types/hash"
)

// ChunkDetails https://docs.near.org/docs/api/rpc#chunk-details
func (c *Client) ChunkDetails(ctx context.Context, chunkHash hash.CryptoHash) (res ChunkView, err error) {
	_, err = c.doRPC(ctx, &res, "chunk", nil, []string{chunkHash.String()})

	return
}
