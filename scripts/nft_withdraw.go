package scripts

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/near-go/common"
	nearclient2 "github.com/rarimo/near-go/nearclient"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
)

func NftWithdraw(ctx context.Context, cli nearclient2.Client, txHash, eventID, sender, receiver, chainFrom, chainTo, token, tokenID, bridge, privateKey string, isWrapped bool) string {
	content := data.NewTransferDataBuilder().
		SetAddress(hexutil.Encode([]byte(token))).
		SetId(hexutil.Encode(to32Bytes([]byte(tokenID)))).
		SetName(nftMetadata[isWrapped].Title).
		SetImageURI(nftMedia).
		SetImageHash(hexutil.Encode(mustDecodeBase64(nftMediaHash))).
		Build().
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
		deposit = common.NEARToYocto(1)
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

func to32Bytes(arr []byte) []byte {
	if len(arr) == 32 || len(arr) == 0 {
		return arr
	}

	result := make([]byte, 32-len(arr))
	return append(result, arr...)

}
