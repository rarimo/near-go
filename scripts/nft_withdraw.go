package scripts

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
)

func NftWithdraw(ctx context.Context, cli client.Client, txHash, eventID, sender, receiver, chainFrom, chainTo, token, tokenID, bridge, privateKey string, isWrapped bool) string {
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

	act := action.NftWithdrawArgs{
		Token:     token,
		TokenID:   tokenID,
		IsWrapped: isWrapped,
		WithdrawArgs: action.WithdrawArgs{
			ReceiverID: receiver,
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
		deposit = types.NEARToYocto(1)
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

func to32Bytes(arr []byte) []byte {
	if len(arr) == 32 || len(arr) == 0 {
		return arr
	}

	result := make([]byte, 32-len(arr))
	return append(result, arr...)

}
