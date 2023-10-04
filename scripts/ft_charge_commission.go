package scripts

import (
	"context"
	"github.com/rarimo/near-go/common"
	nearclient2 "github.com/rarimo/near-go/nearclient"
)

func FtChargeCommission(
	ctx context.Context,
	cli nearclient2.Client,
	feeTokenAddr,
	tokenAddr,
	sender,
	receiver,
	amount,
	feer string,
) (string, string) {
	rawFeeLog := common.FeerDepositArgs{
		FeeTokenAddr: &feeTokenAddr,
		TokenAddr:    &tokenAddr,
		TokenType:    common.TokenType_FT,
		TransferType: common.FeerTransferType_Fee,
		Receiver:     receiver,
		ChainTo:      targetNetwork,
		IsWrapped:    false,
	}

	feeLog, err := rawFeeLog.String()
	if err != nil {
		panic(err)
	}

	depositLog, err := rawFeeLog.WithTransferType(common.FeerTransferType_Deposit).String()
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []common.Action{
		common.NewFtTransferCall(common.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount(amount),
			Msg:        depositLog,
		}, MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	feeResp, err := cli.TransactionSendAwait(ctx, sender, feeTokenAddr, []common.Action{
		common.NewFtTransferCall(common.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount("1000"),
			Msg:        feeLog,
		}, MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return feeResp.Transaction.Hash.String(), depositResp.Transaction.Hash.String()
}
