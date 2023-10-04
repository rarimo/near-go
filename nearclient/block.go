package nearclient

import (
	"context"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient/models"
)

// BlockDetails https://docs.near.org/docs/api/rpc#block-details
func (c *Client) BlockDetails(ctx context.Context, block BlockCharacteristic) (resp models.BlockView, err error) {
	_, err = c.doRPC(ctx, &resp, "block", block, map[string]interface{}{})

	return
}

// BlockChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#changes-in-block
func (c *Client) BlockChanges(ctx context.Context, block BlockCharacteristic) (res *common.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes_in_block", block, map[string]interface{}{})

	return
}
