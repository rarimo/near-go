package scripts

import (
	"context"
	"github.com/rarimo/near-go/common"
	nearclient2 "github.com/rarimo/near-go/nearclient"
)

func RawNFTWithdraw(ctx context.Context, cli *nearclient2.Client, sender, bridge, token, tokenID, receiver, origin, signature string, path [][32]byte, recoveryID uint8, isWrapped bool) string {
	act := common.NftWithdrawArgs{
		Token:     token,
		TokenID:   tokenID,
		IsWrapped: isWrapped,
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

	deposit := common.OneYocto
	if isWrapped {
		act.TokenMetadata = nftMetadata[isWrapped]
		deposit = parseAmount("200000000000000000000000")
	}

	withdrawResp, err := cli.TransactionSendAwait(
		ctx,
		sender,
		bridge,
		[]common.Action{common.NewNftWithdrawCall(act, MaxGas, deposit)},
		nearclient2.WithLatestBlock(),
	)
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}

func RawFTWithdraw(ctx context.Context, cli *nearclient2.Client, sender, bridge, token, amount, receiver, origin, signature string, path [][32]byte, recoveryID uint8, isWrapped bool) string {
	amnt := parseAmount(amount)

	act := common.FtWithdrawArgs{
		Token:     token,
		Amount:    amnt,
		IsWrapped: isWrapped,
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

	deposit := common.OneYocto

	if isWrapped {
		deposit = parseAmount("1250000000000000000000")
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []common.Action{
		common.NewFtWithdrawCall(act, MaxGas, deposit),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}

func RawNativeWithdraw(ctx context.Context, cli *nearclient2.Client, sender, bridge, amount, receiver, origin, signature string, path [][32]byte, recoveryID uint8) string {
	amnt := parseAmount(amount)

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
		common.NewNativeWithdrawCall(act, MaxGas, common.OneYocto),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}

	return withdrawResp.Transaction.Hash.String()
}
