package nearclient

import (
	"context"
	"github.com/rarimo/near-go/common"
)

// ContractViewState https://docs.near.org/docs/api/rpc#view-contract-state
func (c *Client) ContractViewState(ctx context.Context, accountID common.AccountID, prefixBase64 string, block BlockCharacteristic) (res common.ViewStateResult, err error) {
	_, err = c.doRPC(ctx, &res, "query", block, map[string]interface{}{
		"request_type":  "view_state",
		"account_id":    accountID,
		"prefix_base64": prefixBase64,
	})

	return
}

// ContractViewStateChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#view-contract-state-changes
func (c *Client) ContractViewStateChanges(ctx context.Context, accountIDs []common.AccountID, keyPrefixBase64 string, block BlockCharacteristic) (res *common.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type":      "data_changes",
		"account_ids":       accountIDs,
		"key_prefix_base64": keyPrefixBase64,
	})

	return
}

// ContractViewCodeChanges TODO: decode response
// https://docs.near.org/docs/api/rpc#view-contract-code-changes
func (c *Client) ContractViewCodeChanges(ctx context.Context, accountIDs []common.AccountID, block BlockCharacteristic) (res *common.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "contract_code_changes",
		"account_ids":  accountIDs,
	})

	return
}

// ContractViewCallFunction https://docs.near.org/docs/api/rpc#call-a-contract-function
func (c *Client) ContractViewCallFunction(ctx context.Context, accountID, methodName, argsBase64 string, block BlockCharacteristic) (res common.CallResult, err error) {
	_, err = c.doRPC(ctx, &res, "query", block, map[string]interface{}{
		"request_type": "call_function",
		"account_id":   accountID,
		"method_name":  methodName,
		"args_base64":  argsBase64,
	})

	return
}
