package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func NftDeposit(ctx context.Context, cli client.Client, sender, receiver, token, tokenID, bridge string, isWrapped bool) string {
	depositResp, err := cli.TransactionSend(ctx, sender, token, []base.Action{
		action.NewNftTransferCall(action.NftTransferArgs{
			ReceiverId: bridge,
			TokenID:    tokenID,
			Msg: action.TransferArgs{
				Token:     token,
				Receiver:  receiver,
				Chain:     types.NetworkTestnet,
				IsWrapped: isWrapped,
			},
		}, GetGasPrice(ctx, cli), types.OneYocto),
		//action.NewNftTransferCall(action.NftTransferArgs{
		//	ReceiverId: cfg.BridgeAddress,
		//	TokenID:    "2",
		//	Msg: action.TransferArgs{
		//		Token:     cfg.NftAddressOriginal,
		//		Receiver:  cfg.AccountID,
		//		Chain:     types.NetworkTestnet,
		//		IsWrapped: false,
		//	},
		//}, types.YoctoToNEAR(gas.GasPrice), types.OneYocto),
	})
	if err != nil {
		panic(err)
	}
	return depositResp.String()
}
