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
		action.NewFtTransferCall(action.FtTransferArgs{
			ReceiverId: bridge,
			Amount:     amnt,
			Msg: action.TransferArgs{
				Token:     token,
				Receiver:  receiver,
				Chain:     types.NetworkTestnet,
				IsWrapped: isWrapped,
			},
		}, GetGasPrice(ctx, cli), types.OneYocto),
	})
	if err != nil {
		panic(err)
	}
	return depositResp.String()
}
