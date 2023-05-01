package scripts

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func FtDeposit(ctx context.Context, cli client.Client, sender, receiver, token string, amount string, bridge string, isWrapped bool) (string, string) {
	amn := parseAmount(amount)

	transferArgs := action.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped)
	msg, err := transferArgs.String()
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, token, []action.Action{
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: bridge,
			Amount:     amn,
			Msg:        msg,
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
