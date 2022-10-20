package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func NativeDeposit(ctx context.Context, cli client.Client, sender, receiver, amount, bridge string) (string, string) {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []base.Action{
		action.NewNativeDepositCall(action.NativeDepositArgs{
			ReceiverId: receiver,
			Amount:     amnt,
		}, MaxGas, amnt),
	})
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, client.LogEventTypeNativeDeposited, bridge, bridge, nil, &amnt)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}
	return depositResp.Transaction.Hash.String(), eventID.String()
}
