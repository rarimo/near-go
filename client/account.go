package nearclient

import (
	"context"
	"github.com/rarimo/near-go/common"
)

// AccountView https://docs.near.org/docs/api/rpc#view-account
func (c *Client) AccountView(ctx context.Context, accountID common.AccountID, block BlockCharacteristic) (res common.AccountView, err error) {
	_, err = c.doRPC(ctx, &res, "query", block, map[string]interface{}{
		"request_type": "view_account",
		"account_id":   accountID,
	})

	return
}

// AccountViewChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#view-account-changes
func (c *Client) AccountViewChanges(ctx context.Context, accountIDs []common.AccountID, block BlockCharacteristic) (res *common.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "account_changes",
		"account_ids":  accountIDs,
	})

	return
}
