package scripts

import (
	"context"
	nearclient2 "github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/common"
	"lukechampine.com/uint128"
)

func ManageFeeToken(ctx context.Context, operationType common.FeeManageOperationType, cli nearclient2.Client, sender, receiver, bridgeAddr, privateKey string, token common.FeeToken) string {
	feeAmount := token.Fee.Big().String()

	origin, signature, path, recoveryID := getFeeManageOperationSignArgs(
		operationType,
		token,
		feeAmount,
		privateKey,
		receiver,
		bridgeAddr,
	)

	op := common.FeeManageOperationArgs{
		Operation: common.FeeManageOperation{
			SignArgs: common.SignArgs{
				Origin:     origin,
				Signature:  signature,
				Path:       path,
				RecoveryID: recoveryID,
			},
			Token: token,
		},
	}

	actions := make([]common.Action, 0, 1)

	switch operationType {
	case actions.FeeAddFeeToken:
		actions = append(actions, actions.NewFeeTokenAddCall(op, MaxGas))
	case actions.FeeRemoveFeeToken:
		actions = append(actions, actions.NewFeeTokenRemoveCall(op, MaxGas))
	case actions.FeeUpdateFeeToken:
		actions = append(actions, actions.NewFeeTokenUpdateCall(op, MaxGas))
	case actions.FeeWithdraw:
		actions = append(actions, actions.NewFeeTokenWithdrawCall(op, sender, common.Balance(uint128.From64(1000)), MaxGas))
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, receiver, actions, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return withdrawResp.Transaction.Hash.String()
}
