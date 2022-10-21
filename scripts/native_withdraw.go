package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation/data"
)

func NativeWithdraw(ctx context.Context, cli client.Client, txHash, eventID string, sender, receiver, chainFrom, chainTo, amount, bridge, privateKey string) string {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	content := data.NewTransferDataBuilder().SetAmount(amnt.String()).Build().GetContent()

	origin, signature, path, recoveryID := getContent(
		privateKey,
		txHash,
		eventID,
		receiver,
		bridge,
		chainFrom,
		chainTo,
		content,
	)

	act := action.NativeWithdrawArgs{
		Amount: amnt,
		WithdrawArgs: action.WithdrawArgs{
			Chain:      chainTo,
			ReceiverID: receiver,
			Origin:     origin,
			Path:       path,
			Signatures: []string{signature},
			RecoveryID: recoveryID,
		},
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []base.Action{
		action.NewNativeWithdrawCall(act, MaxGas, types.OneYocto),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}
