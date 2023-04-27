package scripts

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func NativeDeposit(ctx context.Context, cli client.Client, sender, receiver, chainTo, amount, bridge string) (string, string) {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []action.Action{
		action.NewNativeDepositCall(action.NativeDepositArgs{
			ReceiverId: receiver,
			Chain:      chainTo,
		}, MaxGas, amnt),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, client.LogEventTypeNativeDeposited, bridge, nil, nil, &amnt)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}
	return depositResp.Transaction.Hash.String(), eventID.String()
}
