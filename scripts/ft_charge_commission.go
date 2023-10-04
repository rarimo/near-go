package scripts

import (
	"context"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func FtChargeCommission(
	ctx context.Context,
	cli client.Client,
	feeTokenAddr,
	tokenAddr,
	sender,
	receiver,
	amount,
	feer string,
) (string, string) {
	rawFeeLog := action.FeerDepositArgs{
		FeeTokenAddr: &feeTokenAddr,
		TokenAddr:    &tokenAddr,
		TokenType:    action.TokenType_FT,
		TransferType: action.FeerTransferType_Fee,
		Receiver:     receiver,
		ChainTo:      targetNetwork,
		IsWrapped:    false,
	}

	feeLog, err := rawFeeLog.String()
	if err != nil {
		panic(err)
	}

	depositLog, err := rawFeeLog.WithTransferType(action.FeerTransferType_Deposit).String()
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []action.Action{
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount(amount),
			Msg:        depositLog,
		}, MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	feeResp, err := cli.TransactionSendAwait(ctx, sender, feeTokenAddr, []action.Action{
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount("1000"),
			Msg:        feeLog,
		}, MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return feeResp.Transaction.Hash.String(), depositResp.Transaction.Hash.String()
}
