package scripts

import (
	"context"
	nearclient2 "github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/constants"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
)

func NativeWithdraw(ctx context.Context, cli nearclient2.Client, txHash, eventID string, sender, receiver, chainFrom, chainTo, amount, bridge, privateKey string) string {
	amnt, err := common.BalanceFromString(amount)
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

	act := common.NativeWithdrawArgs{
		Amount: amnt,
		WithdrawArgs: common.WithdrawArgs{
			ReceiverID: receiver,
			SignArgs: common.SignArgs{
				Origin:     origin,
				Path:       path,
				Signature:  signature,
				RecoveryID: recoveryID,
			},
		},
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []common.Action{
		common.NewNativeWithdrawCall(act, MaxGas, constants.OneYocto),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}
