package scripts

import (
	"context"
	"crypto/elliptic"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation/bundle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation/data"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation/origin"
)

func NftWithdraw(ctx context.Context, cli client.Client, txHash, eventID, sender, receiver, chainFrom, chainTo, token, tokenID, bridge, privateKey string, isWrapped bool) string {
	//fmt.Println("Token: " + hexutil.Encode([]byte(token)))
	//fmt.Println("Token ID: " + hexutil.Encode([]byte(tokenID)))
	//fmt.Println("nftMediaHash: " + hexutil.Encode([]byte(nftMediaHash)))

	builder := data.NewTransferDataBuilder().
		SetAddress(hexutil.Encode([]byte(token))).
		SetId(hexutil.Encode([]byte(tokenID))).
		SetName(nftMetadata[isWrapped].Title).
		SetImageURI(nftMedia).
		SetImageHash(hexutil.Encode(mustDecodeBase64(nftMediaHash)))

	c := builder.Build().GetContent()

	targetContent := operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash(txHash).SetOpId(eventID).SetCurrentNetwork(chainFrom).Build().GetOrigin(),
		TargetNetwork:  chainTo,
		Receiver:       []byte(receiver),
		TargetContract: []byte(bridge),
		Data:           c,
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	mt := merkle.NewTree(crypto.Keccak256, targetContent, content1, content2)
	path, ok := mt.Path(targetContent)
	if !ok {
		panic("path not found")
	}

	prvKey, err := base58.Decode(privateKey)
	if err != nil {
		panic(err)
	}

	pk, err := crypto.ToECDSA(prvKey)
	if err != nil {
		panic(err)
	}

	puk := elliptic.Marshal(secp256k1.S256(), pk.X, pk.Y)
	fmt.Println("PUB KEY: " + base58.Encode(puk[1:]))

	signature, err := crypto.Sign(mt.Root(), pk)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature %s\n", base58.Encode(signature[:64]))
	fmt.Printf("Root msg %s\n", base58.Encode(mt.Root()))

	recoveredKey, err := secp256k1.RecoverPubkey(mt.Root(), signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey[1:]))

	act := action.NftWithdrawArgs{
		Token:         token,
		TokenID:       tokenID,
		ReceiverID:    receiver,
		Chain:         targetNetwork,
		IsWrapped:     isWrapped,
		Origin:        hexutil.Encode(targetContent.Origin[:]),
		Path:          make([][32]byte, len(path)),
		Signatures:    []string{hexutil.Encode(signature[:64])},
		RecoveryID:    signature[64],
		TokenMetadata: nftMetadata[isWrapped],
	}

	for i, hash := range path {
		copy(act.Path[i][:], hash)
	}

	fmt.Println("Content hash: " + base58.Encode(targetContent.CalculateHash()))

	deposit := types.OneYocto
	if isWrapped {
		deposit = types.NEARToYocto(1)
	}

	withdrawResp, err := cli.TransactionSendAwait(
		ctx,
		sender,
		bridge,
		[]base.Action{action.NewNftWithdrawCall(act, MaxGas, deposit)},
		client.WithLatestBlock(),
	)

	if err != nil {
		panic(err)
	}
	return withdrawResp.Transaction.Hash.String()
}
