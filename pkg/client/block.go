package client

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/client/block"
	"gitlab.com/rarimo/near-bridge-go/pkg/jsonrpc"
)

// BlockDetails https://docs.near.org/docs/api/rpc#block-details
func (c *Client) BlockDetails(ctx context.Context, block block.BlockCharacteristic) (resp BlockView, err error) {
	_, err = c.doRPC(ctx, &resp, "block", block, map[string]interface{}{})

	return
}

// BlockChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#changes-in-block
func (c *Client) BlockChanges(ctx context.Context, block block.BlockCharacteristic) (res *jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes_in_block", block, map[string]interface{}{})

	return
}
