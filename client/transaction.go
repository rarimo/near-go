package nearclient

import (
	"context"
	"github.com/rarimo/near-go/client/models"
	"github.com/rarimo/near-go/common"
)

// RPCTransactionSend https://docs.near.org/docs/api/rpc#send-transaction-async
func (c *Client) RPCTransactionSend(ctx context.Context, signedTxnBase64 string) (resp common.Hash, err error) {
	_, err = c.doRPC(ctx, &resp, "broadcast_tx_async", nil, []string{signedTxnBase64})

	return
}

// RPCTransactionSendAwait https://docs.near.org/docs/api/rpc#send-transaction-await
func (c *Client) RPCTransactionSendAwait(ctx context.Context, signedTxnBase64 string) (resp models.FinalExecutionOutcomeView, err error) {
	_, err = c.doRPC(ctx, &resp, "broadcast_tx_commit", nil, []string{signedTxnBase64})

	return
}

// TransactionStatus https://docs.near.org/docs/api/rpc#transaction-status
func (c *Client) TransactionStatus(ctx context.Context, tx common.Hash, sender common.AccountID) (resp models.FinalExecutionOutcomeView, err error) {
	_, err = c.doRPC(ctx, &resp, "tx", nil, []string{
		tx.String(), sender,
	})

	return
}

// TransactionStatusWithReceipts https://docs.near.org/docs/api/rpc#transaction-status-with-receipts
func (c *Client) TransactionStatusWithReceipts(ctx context.Context, tx common.Hash, sender common.AccountID) (resp models.FinalExecutionOutcomeWithReceiptView, err error) {
	_, err = c.doRPC(ctx, &resp, "EXPERIMENTAL_tx_status", nil, []string{
		tx.String(), sender,
	})

	return
}
