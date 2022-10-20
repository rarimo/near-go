package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation/data"
)

func FtWithdraw(ctx context.Context, cli client.Client, txHash, eventID, sender, receiver, chainFrom, chainTo, token, amount, bridge, privateKey string, isWrapped bool) string {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	content := data.NewTransferDataBuilder().SetAddress(token).SetAmount(amnt.String()).Build().GetContent()

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

	act := action.FtWithdrawArgs{
		Token:     token,
		Amount:    amnt,
		IsWrapped: isWrapped,
		WithdrawArgs: action.WithdrawArgs{
			ReceiverID: receiver,
			Chain:      chainTo,
			Origin:     origin,
			Path:       path,
			Signatures: []string{signature},
			RecoveryID: recoveryID,
		},
	}

	withdrawResp, err := cli.TransactionSend(ctx, sender, bridge, []base.Action{
		action.NewFtWithdrawCall(act, MaxGas),
	})
	if err != nil {
		panic(err)
	}
	return withdrawResp.String()
}
