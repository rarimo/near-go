package client

import (
	"context"

	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client/block"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/jsonrpc"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

// AccountView https://docs.near.org/docs/api/rpc#view-account
func (c *Client) AccountView(ctx context.Context, accountID types.AccountID, block block.BlockCharacteristic) (res AccountView, err error) {
	_, err = c.doRPC(ctx, &res, "query", block, map[string]interface{}{
		"request_type": "view_account",
		"account_id":   accountID,
	})

	return
}

// AccountViewChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#view-account-changes
func (c *Client) AccountViewChanges(ctx context.Context, accountIDs []types.AccountID, block block.BlockCharacteristic) (res *jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "account_changes",
		"account_ids":  accountIDs,
	})

	return
}
