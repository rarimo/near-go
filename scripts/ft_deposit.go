package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func FtDeposit(ctx context.Context, cli client.Client, sender, receiver, token, amount, bridge string, isWrapped bool) string {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}
	depositResp, err := cli.TransactionSend(ctx, sender, token, []base.Action{
		action.NewFtDepositCall(action.FtDepositArgs{
			ReceiverId: bridge,
			Amount:     amnt,
			Msg:        action.NewTransferArgs(token, sender, receiver, targetNetwork, isWrapped, nil, nil),
		}, MaxGas),
	})
	if err != nil {
		panic(err)
	}
	return depositResp.String()
}
