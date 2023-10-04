package nearclient

import (
	"context"
	"github.com/rarimo/near-go/nearclient/models"
)

// NetworkInfo https://docs.near.org/docs/api/rpc#network-info
func (c *Client) NetworkInfo(ctx context.Context) (res models.NetworkInfo, err error) {
	_, err = c.doRPC(ctx, &res, "network_info", nil, []string{})

	return
}

// NetworkStatusValidators https://docs.near.org/docs/api/rpc#general-validator-status
func (c *Client) NetworkStatusValidators(ctx context.Context) (res models.StatusResponse, err error) {
	_, err = c.doRPC(ctx, &res, "status", nil, []string{})

	return
}

// NetworkStatusValidatorsDetailed https://docs.near.org/docs/api/rpc#detailed-validator-status
func (c *Client) NetworkStatusValidatorsDetailed(ctx context.Context, block BlockCharacteristic) (res models.ValidatorsResponse, err error) {
	_, err = c.doRPC(ctx, nil, "validators", nil, blockIDArrayParams(block))

	return
}
