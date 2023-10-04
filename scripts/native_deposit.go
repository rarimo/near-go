package scripts

import (
	"context"
	nearclient2 "github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/client/models"
	"github.com/rarimo/near-go/common"
)

func NativeDeposit(ctx context.Context, cli nearclient2.Client, sender, receiver, chainTo, amount, bridge string) (string, string) {
	amnt, err := common.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	depositResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []common.Action{
		common.NewNativeDepositCall(common.NativeDepositArgs{
			ReceiverId: receiver,
			Chain:      chainTo,
		}, MaxGas, amnt),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	eventID, err := GetDepositedReceiptID(depositResp, models.LogEventTypeNativeDeposited, bridge, nil, nil, &amnt)
	if err != nil {
		panic(err)
	}
	if eventID == nil {
		panic("eventID is nil")
	}
	return depositResp.Transaction.Hash.String(), eventID.String()
}
