package scripts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action/base"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
	"lukechampine.com/uint128"
)

func FtWithdraw(ctx context.Context, cli client.Client, txHash, eventID, sender, receiver, chainFrom, chainTo, token, amount, bridge, privateKey string, isWrapped bool) string {
	amnt := parseAmount(amount)

	content := data.NewTransferDataBuilder().
		SetAddress(hexutil.Encode([]byte(token))).
		SetAmount(amnt.String()).Build().
		GetContent()

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

	deposit := types.OneYocto

	if isWrapped {
		deposit = types.ZeroNEAR
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []base.Action{
		action.NewFtWithdrawCall(act, MaxGas, deposit),
	}, client.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}

func parseAmount(amount string) types.Balance {
	bigAmount, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		panic("invalid amount")
	}

	av := uint128.FromBig(bigAmount)
	return types.Balance(av)
}
