package nearclient

import (
	"context"
	"errors"
	"github.com/rarimo/near-go/common"
)

type transactionCtx struct {
	txn         common.Transaction
	keyPair     *common.KeyPair
	keyNonceSet bool
}

type TransactionOpt func(context.Context, *transactionCtx) error

func (c *Client) prepareTransaction(ctx context.Context, from, to common.AccountID, actions []common.Action, txnOpts ...TransactionOpt) (ctx2 context.Context, blob string, err error) {
	ctx2 = context.WithValue(ctx, clientCtx, c)
	txn := common.Transaction{
		SignerID:   from,
		ReceiverID: to,
		Actions:    actions,
	}
	txnCtx := transactionCtx{
		txn:         txn,
		keyPair:     getKeyPair(ctx2),
		keyNonceSet: false,
	}

	for _, opt := range txnOpts {
		if err = opt(ctx2, &txnCtx); err != nil {
			return
		}
	}

	if txnCtx.keyPair == nil {
		err = errors.New("no keypair specified")
		return
	}

	txnCtx.txn.PublicKey = txnCtx.keyPair.PublicKey.ToPublicKey()

	// Query the access key nonce, if not specified
	if !txnCtx.keyNonceSet {
		var accessKey common.AccessKeyView
		accessKey, err = c.AccessKeyView(ctx2, txnCtx.txn.SignerID, txnCtx.keyPair.PublicKey, FinalityFinal())
		if err != nil {
			return
		}

		nonce := accessKey.Nonce

		// Increment nonce by 1
		txnCtx.txn.Nonce = nonce + 1
		txnCtx.keyNonceSet = true
	}

	blob, err = common.SignAndSerializeTransaction(*txnCtx.keyPair, txnCtx.txn)
	return
}

// TransactionSend https://docs.near.org/docs/api/rpc#send-transaction-async
func (c *Client) TransactionSend(ctx context.Context, from, to common.AccountID, actions []common.Action, txnOpts ...TransactionOpt) (res common.Hash, err error) {
	ctx2, blob, err := c.prepareTransaction(ctx, from, to, actions, txnOpts...)
	if err != nil {
		return
	}
	return c.RPCTransactionSend(ctx2, blob)
}

// TransactionSendAwait https://docs.near.org/docs/api/rpc#send-transaction-await
func (c *Client) TransactionSendAwait(ctx context.Context, from, to common.AccountID, actions []common.Action, txnOpts ...TransactionOpt) (res common.FinalExecutionOutcomeView, err error) {
	ctx2, blob, err := c.prepareTransaction(ctx, from, to, actions, txnOpts...)
	if err != nil {
		return
	}
	return c.RPCTransactionSendAwait(ctx2, blob)
}

func WithBlockCharacteristic(block BlockCharacteristic) TransactionOpt {
	return func(ctx context.Context, txnCtx *transactionCtx) (err error) {
		client := ctx.Value(clientCtx).(*Client)

		var res common.BlockView
		if res, err = client.BlockDetails(ctx, block); err != nil {
			return
		}

		txnCtx.txn.BlockHash = res.Header.Hash
		return
	}

}

// WithBlockHash sets block hash to attach this transaction to
func WithBlockHash(hash common.Hash) TransactionOpt {
	return func(_ context.Context, txnCtx *transactionCtx) (err error) {
		txnCtx.txn.BlockHash = hash
		return
	}
}

// WithLatestBlock is alias to `WithBlockCharacteristic(block.FinalityFinal())`
func WithLatestBlock() TransactionOpt {
	return WithBlockCharacteristic(FinalityFinal())
}

// WithKeyPair sets key pair to use sign this transaction with
func WithKeyPair(keyPair common.KeyPair) TransactionOpt {
	return func(_ context.Context, txnCtx *transactionCtx) (err error) {
		kp := keyPair
		txnCtx.keyPair = &kp
		return
	}
}

// WithKeyNonce sets key nonce to use with this transaction. If not set via this function, a RPC query will be done to query current nonce and
// (nonce+1) will be used
func WithKeyNonce(nonce common.Nonce) TransactionOpt {
	return func(_ context.Context, txnCtx *transactionCtx) (err error) {
		txnCtx.txn.Nonce = nonce
		txnCtx.keyNonceSet = true
		return
	}
}
