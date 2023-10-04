package scripts

import (
	"context"
	"encoding/json"
	"github.com/rarimo/near-go/common"
	nearclient2 "github.com/rarimo/near-go/nearclient"
)

func NftChargeCommission(ctx context.Context, cli nearclient2.Client, feeTokenAddr, tokenAddr, sender, receiver, tokenId string, feer string) (string, string) {
	rawLog := common.FeerDepositArgs{
		FeeTokenAddr: &feeTokenAddr,
		TokenAddr:    &tokenAddr,
		TokenType:    common.TokenType_NFT,
		Receiver:     receiver,
		ChainTo:      targetNetwork,
		IsWrapped:    false,
	}

	rawFeeLog := rawLog
	rawDepositLog := rawLog
	rawFeeLog.TransferType = common.FeerTransferType_Fee
	rawDepositLog.TransferType = common.FeerTransferType_Deposit

	feeLog, _ := json.Marshal(rawFeeLog)
	depositLog, _ := json.Marshal(rawDepositLog)

	depositResp, err := cli.TransactionSendAwait(ctx, sender, tokenAddr, []common.Action{
		common.NewNftTransferCall(common.NftTransferArgs{
			ReceiverId: feer,
			TokenID:    tokenId,
			Msg:        string(depositLog),
		}, MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	feeResp, err := cli.TransactionSendAwait(ctx, sender, feeTokenAddr, []common.Action{
		common.NewFtTransferCall(common.FtTransferArgs{
			ReceiverId: feer,
			Amount:     parseAmount("1000"),
			Msg:        string(feeLog),
		}, MaxGas),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return feeResp.Transaction.Hash.String(), depositResp.Transaction.Hash.String()
}
