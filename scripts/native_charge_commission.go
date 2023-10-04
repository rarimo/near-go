package scripts

import (
	"context"
	"encoding/json"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func NativeChargeCommission(ctx context.Context, cli client.Client, tokenAddr, sender, receiver, amount string, feer string) (string, string) {
	feeLog := action.FeerDepositArgs{
		TokenAddr:    &tokenAddr,
		TokenType:    action.TokenType_Native,
		TransferType: action.FeerTransferType_Fee,
		Receiver:     receiver,
		ChainTo:      targetNetwork,
		IsWrapped:    false,
	}

	rawDepositLog := feeLog.WithTransferType(action.FeerTransferType_Deposit)
	depositLog, _ := json.Marshal(rawDepositLog)

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []action.Action{
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount(amount),
			Msg:        string(depositLog),
		}, MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	feeResp, err := cli.TransactionSendAwait(ctx, sender, feer, []action.Action{
		action.NewFeeChargeNativeCall(feeLog, parseAmount("1000"), MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return feeResp.Transaction.Hash.String(), depositResp.Transaction.Hash.String()
}
