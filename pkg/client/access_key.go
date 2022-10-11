package client

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client/block"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/jsonrpc"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/key"
)

// AccessKeyView https://docs.near.org/docs/api/rpc#view-access-key
func (c *Client) AccessKeyView(ctx context.Context, accountID types.AccountID, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (resp AccessKeyView, err error) {
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
func (c *Client) AccessKeyViewList(ctx context.Context, accountID types.AccountID, block block.BlockCharacteristic) (resp AccessKeyList, err error) {
	_, err = c.doRPC(ctx, &resp, "query", block, map[string]interface{}{
		"request_type": "view_access_key_list",
		"account_id":   accountID,
	})

	return
}

// AccessKeyViewChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#view-access-key-changes-single
func (c *Client) AccessKeyViewChanges(ctx context.Context, accountID types.AccountID, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (res *jsonrpc.Response, err error) {
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
func (c *Client) AccessKeyViewChangesAll(ctx context.Context, accountIDs []types.AccountID, block block.BlockCharacteristic) (res *jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids":  accountIDs,
	})

	return
}
