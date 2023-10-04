package scripts

import (
	"context"
	"encoding/json"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func NftChargeCommission(ctx context.Context, cli client.Client, feeTokenAddr, tokenAddr, sender, receiver, tokenId string, feer string) (string, string) {
	rawLog := action.FeerDepositArgs{
		FeeTokenAddr: &feeTokenAddr,
		TokenAddr:    &tokenAddr,
		TokenType:    action.TokenType_NFT,
		Receiver:     receiver,
		ChainTo:      targetNetwork,
		IsWrapped:    false,
	}

	rawFeeLog := rawLog
	rawDepositLog := rawLog
	rawFeeLog.TransferType = action.FeerTransferType_Fee
	rawDepositLog.TransferType = action.FeerTransferType_Deposit

	feeLog, _ := json.Marshal(rawFeeLog)
	depositLog, _ := json.Marshal(rawDepositLog)

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []action.Action{
		action.NewNftTransferCall(action.NftTransferArgs{
			ReceiverId: feer,
			TokenID:    tokenId,
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
