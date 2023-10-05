package nearprovider

import (
	"context"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var ErrUnknownTx = errors.New("unknown tx")

func (p *provider) GetTransaction(ctx context.Context, hash common.Hash, sender common.AccountID) (*common.FinalExecutionOutcomeWithReceiptView, error) {
	resp, err := p.getTx(ctx, hash, sender)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tx")
	}

	return resp, nil
}

func (p *provider) getTx(ctx context.Context, hash common.Hash, sender common.AccountID) (*common.FinalExecutionOutcomeWithReceiptView, error) {
	tx, err := p.tryToGetTxFromRPC(ctx, p.c, hash, sender)

	if err != nil {
		if errors.Cause(err) == ErrUnknownTx {
			return p.tryToGetTxFromRPC(ctx, p.hc, hash, sender)
		}

		return nil, errors.Wrap(err, "failed to get tx")
	}

	return tx, nil
}

func (p *provider) tryToGetTxFromRPC(ctx context.Context, cli *nearclient.Client, hash common.Hash, sender common.AccountID) (*common.FinalExecutionOutcomeWithReceiptView, error) {
	tx, err := cli.TransactionStatusWithReceipts(ctx, hash, sender)
	if err != nil {
		if isErrUnknownTx(err) {
			return nil, ErrUnknownTx
		}
		return nil, errors.Wrap(err, "failed to get tx from rpc")
	}

	return &tx, nil
}

func isErrUnknownTx(err error) bool {
	e, ok := err.(*common.JsonRpcError)
	return ok && e.Cause.Name == "UNKNOWN_TRANSACTION"
}
