package nearclient

import (
	"context"
	"github.com/rarimo/near-go/nearclient/models"
)

// https://docs.near.org/docs/api/rpc#gas-price
func (c *Client) GasPriceView(ctx context.Context, block BlockCharacteristic) (res models.GasPrice, err error) {
	_, err = c.doRPC(ctx, &res, "gas_price", nil, blockIDArrayParams(block))

	return
}
