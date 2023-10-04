package scripts

import (
	"context"
	nearclient2 "github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
	"lukechampine.com/uint128"
)

func FtWithdraw(ctx context.Context, cli nearclient2.Client, txHash, eventID, sender, receiver, chainFrom, chainTo, token, amount, bridge, privateKey string, isWrapped bool) string {
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

	deposit := constants.OneYocto

	if isWrapped {
		deposit = constants.ZeroNEAR
	}

	withdrawResp, err := cli.TransactionSendAwait(ctx, sender, bridge, []common.Action{
		common.NewFtWithdrawCall(act, MaxGas, deposit),
	}, nearclient2.WithLatestBlock())
	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}

func parseAmount(amount string) common.Balance {
	bigAmount, ok := big.NewInt(0).SetString(amount, 10)
	if !ok {
		panic("invalid amount")
	}

	av := uint128.FromBig(bigAmount)
	return common.Balance(av)
}
