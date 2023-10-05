package scripts

import (
	"context"
	"github.com/rarimo/near-go/common"
	nearclient2 "github.com/rarimo/near-go/nearclient"
)

func NftDeposit(ctx context.Context, cli nearclient2.Client, sender, receiver, token, tokenID, bridge string, isWrapped bool) (string, string) {
	transferArgs := common.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped)
	msg, err := transferArgs.String()
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, token, []common.Action{
		common.NewNftTransferCall(common.NftTransferArgs{
			ReceiverId: bridge,
			TokenID:    tokenID,
			Msg:        msg,
		}, MaxGas/2),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, common.LogEventTypeNftDeposited, bridge, &token, &tokenID, nil)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}

	return depositResp.Transaction.Hash.String(), eventID.String()
}
