package scripts

import (
	"context"
	"encoding/json"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func FtChargeCommission(ctx context.Context, cli client.Client, sender, receiver, amount string, feer string) (string, string) {
	amn := parseAmount(amount)

	feeTokenAddr := "ft_test_fee.napalmpapalam.testnet"
	tokenAddr := "ft_test.napalmpapalam.testnet"

	rawLog := action.FeerDepositArgs{
		FeeTokenAddr: &feeTokenAddr,
		TokenAddr:    &tokenAddr,
		TokenType:    action.TokenType_FT,
		Receiver:     receiver,
		ChainTo:      "Near",
		IsWrapped:    false,
	}

	rawFeeLog := rawLog
	rawDepositLog := rawLog
	rawFeeLog.TransferType = action.FeerTransferType_Fee
	rawDepositLog.TransferType = action.FeerTransferType_Deposit

	feeLog, _ := json.Marshal(rawFeeLog)
	depositLog, _ := json.Marshal(rawDepositLog)

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []action.Action{
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: feer,
			Amount:     amn,
			Msg:        string(depositLog),
		}, MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	feeResp, err := cli.TransactionSendAwait(ctx, sender, feeTokenAddr, []action.Action{
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount("1000"),
			Msg:        string(feeLog),
		}, MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return feeResp.Transaction.Hash.String(), depositResp.Transaction.Hash.String()
}
