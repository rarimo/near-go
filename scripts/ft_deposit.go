package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func FtDeposit(ctx context.Context, cli client.Client, sender, receiver, token, amount, bridge string, isWrapped bool) (string, string) {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, token, []base.Action{
		action.NewFtDepositCall(action.FtDepositArgs{
			ReceiverId: bridge,
			Amount:     amnt,
			Msg:        action.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped),
		}, MaxGas),
	})
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, client.LogEventTypeFtDeposited, bridge, token, nil, &amnt)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}

	return depositResp.Transaction.Hash.String(), eventID.String()
}
