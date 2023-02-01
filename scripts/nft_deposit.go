package scripts

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action/base"
)

func NftDeposit(ctx context.Context, cli client.Client, sender, receiver, token, tokenID, bridge string, isWrapped bool) (string, string) {
	depositResp, err := cli.TransactionSendAwait(ctx, sender, token, []base.Action{
		action.NewNftDepositCall(action.NftDepositArgs{
			ReceiverId: bridge,
			TokenID:    tokenID,
			Msg:        action.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped),
		}, MaxGas/2),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, client.LogEventTypeNftDeposited, bridge, &token, &tokenID, nil)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}

	return depositResp.Transaction.Hash.String(), eventID.String()
}
