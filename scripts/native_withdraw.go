package scripts

import (
	"context"

	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
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
			SignArgs: action.SignArgs{
				Origin:     origin,
				Path:       path,
				Signature:  signature,
				RecoveryID: recoveryID,
			},
		},
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []action.Action{
		action.NewNativeWithdrawCall(act, MaxGas, types.OneYocto),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}
