package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func FtDeposit(ctx context.Context, cli client.Client, sender, receiver, token string, amount string, bridge string, isWrapped bool) (string, string) {
	amn := parseAmount(amount)

	depositResp, err := cli.TransactionSendAwait(ctx, sender, token, []base.Action{
		action.NewFtDepositCall(action.FtDepositArgs{
			ReceiverId: bridge,
			Amount:     amn,
			Msg:        action.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped),
		}, MaxGas),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, client.LogEventTypeFtDeposited, bridge, &token, nil, &amn)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}

	return depositResp.Transaction.Hash.String(), eventID.String()
}
