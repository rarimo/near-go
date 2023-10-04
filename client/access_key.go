package nearclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/rarimo/near-go/client/models"
	"github.com/rarimo/near-go/common"
)

// AccessKeyView https://docs.near.org/docs/api/rpc#view-access-key
func (c *Client) AccessKeyView(ctx context.Context, accountID common.AccountID, publicKey common.Base58PublicKey, block BlockCharacteristic) (resp models.AccessKeyView, err error) {
	_, err = c.doRPC(ctx, &resp, "query", block, map[string]interface{}{
		"request_type": "view_access_key",
		"account_id":   accountID,
		"public_key":   publicKey,
	})

	if resp.Error != nil {
		err = fmt.Errorf("RPC returned an error: %w", errors.New(*resp.Error))
	}

	return
}

// AccessKeyViewList https://docs.near.org/docs/api/rpc#view-access-key-list
func (c *Client) AccessKeyViewList(ctx context.Context, accountID common.AccountID, block BlockCharacteristic) (resp models.AccessKeyList, err error) {
	_, err = c.doRPC(ctx, &resp, "query", block, map[string]interface{}{
		"request_type": "view_access_key_list",
		"account_id":   accountID,
	})

	return
}

// AccessKeyViewChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#view-access-key-changes-single
func (c *Client) AccessKeyViewChanges(ctx context.Context, accountID common.AccountID, publicKey common.Base58PublicKey, block BlockCharacteristic) (res *common.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "single_access_key_changes",
		"keys": map[string]interface{}{
			"account_id": accountID,
			"public_key": publicKey,
		},
	})

	return
}

// AccessKeyViewChangesAll TODO: decode response
// https://docs.near.org/docs/api/rpc#view-access-key-changes-all
func (c *Client) AccessKeyViewChangesAll(ctx context.Context, accountIDs []common.AccountID, block BlockCharacteristic) (res *common.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids":  accountIDs,
	})

	return
}
