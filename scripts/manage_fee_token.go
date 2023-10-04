package scripts

import (
	"context"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"lukechampine.com/uint128"
)

func ManageFeeToken(ctx context.Context, operationType action.FeeManageOperationType, cli client.Client, sender, receiver, bridgeAddr, privateKey string, token action.FeeToken) string {
	feeAmount := token.Fee.Big().String()

	origin, signature, path, recoveryID := getFeeManageOperationSignArgs(
		operationType,
		token,
		feeAmount,
		privateKey,
		receiver,
		bridgeAddr,
	)

	op := action.FeeManageOperationArgs{
		Operation: action.FeeManageOperation{
			SignArgs: action.SignArgs{
				Origin:     origin,
				Signature:  signature,
				Path:       path,
				RecoveryID: recoveryID,
			},
			Token: token,
		},
	}

	actions := make([]action.Action, 0, 1)

	switch operationType {
	case action.FeeAddFeeToken:
		actions = append(actions, action.NewFeeTokenAddCall(op, MaxGas))
	case action.FeeRemoveFeeToken:
		actions = append(actions, action.NewFeeTokenRemoveCall(op, MaxGas))
	case action.FeeUpdateFeeToken:
		actions = append(actions, action.NewFeeTokenUpdateCall(op, MaxGas))
	case action.FeeWithdraw:
		actions = append(actions, action.NewFeeTokenWithdrawCall(op, sender, types.Balance(uint128.From64(1000)), MaxGas))
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, receiver, actions, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return withdrawResp.Transaction.Hash.String()
}
