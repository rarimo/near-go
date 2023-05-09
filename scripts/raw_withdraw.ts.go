package scripts

import (
	"context"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func RawNFTWithdraw(ctx context.Context, cli *client.Client, sender, bridge, token, tokenID, receiver, chainTo, origin, signature string, path [][32]byte, recoveryID uint8, isWrapped bool) string {
	act := action.NftWithdrawArgs{
		Token:     token,
		TokenID:   tokenID,
		IsWrapped: isWrapped,
		WithdrawArgs: action.WithdrawArgs{
			ReceiverID: receiver,
			Chain:      chainTo,
			SignArgs: action.SignArgs{
				Origin:     origin,
				Path:       path,
				Signature:  signature,
				RecoveryID: recoveryID,
			},
		},
	}

	deposit := types.OneYocto
	if isWrapped {
		act.TokenMetadata = nftMetadata[isWrapped]
		deposit = types.NEARToYocto(0.2)
	}

	withdrawResp, err := cli.TransactionSendAwait(
		ctx,
		sender,
		bridge,
		[]action.Action{action.NewNftWithdrawCall(act, MaxGas, deposit)},
		client.WithLatestBlock(),
	)
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}

func RawFTWithdraw(ctx context.Context, cli *client.Client, sender, bridge, token, amount, receiver, chainTo, origin, signature string, path [][32]byte, recoveryID uint8, isWrapped bool) string {
	amnt := parseAmount(amount)

	act := action.FtWithdrawArgs{
		Token:     token,
		Amount:    amnt,
		IsWrapped: isWrapped,
		WithdrawArgs: action.WithdrawArgs{
			ReceiverID: receiver,
			Chain:      chainTo,
			SignArgs: action.SignArgs{
				Origin:     origin,
				Path:       path,
				Signature:  signature,
				RecoveryID: recoveryID,
			},
		},
	}

	deposit := types.OneYocto

	if isWrapped {
		deposit = types.ZeroNEAR
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []action.Action{
		action.NewFtWithdrawCall(act, MaxGas, deposit),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}

func RawNativeWithdraw(ctx context.Context, cli *client.Client, sender, bridge, amount, receiver, chainTo, origin, signature string, path [][32]byte, recoveryID uint8) string {
	amnt := parseAmount(amount)

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
