package scripts

import (
	"context"
	"encoding/json"
	nearclient2 "github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/common"
)

func NativeChargeCommission(ctx context.Context, cli nearclient2.Client, tokenAddr, sender, receiver, amount string, feer string) (string, string) {
	feeLog := common.FeerDepositArgs{
		TokenAddr:    &tokenAddr,
		TokenType:    common.TokenType_Native,
		TransferType: common.FeerTransferType_Fee,
		Receiver:     receiver,
		ChainTo:      targetNetwork,
		IsWrapped:    false,
	}

	rawDepositLog := feeLog.WithTransferType(common.FeerTransferType_Deposit)
	depositLog, _ := json.Marshal(rawDepositLog)

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []common.Action{
		common.NewFtTransferCall(common.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount(amount),
			Msg:        string(depositLog),
		}, MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	feeResp, err := cli.TransactionSendAwait(ctx, sender, feer, []common.Action{
		common.NewFeeChargeNativeCall(feeLog, parseAmount("1000"), MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return feeResp.Transaction.Hash.String(), depositResp.Transaction.Hash.String()
}
