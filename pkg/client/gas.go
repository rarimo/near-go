package client

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/client/block"
)

// https://docs.near.org/docs/api/rpc#gas-price
func (c *Client) GasPriceView(ctx context.Context, block block.BlockCharacteristic) (res GasPrice, err error) {
	_, err = c.doRPC(ctx, &res, "gas_price", nil, blockIDArrayParams(block))

	return
}
