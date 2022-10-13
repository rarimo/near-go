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
	xcrypto "gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
)

func NativeWithdraw(ctx context.Context, cli client.Client, txHash string, sender, receiver, amount, bridge, privateKey string) string {
	amnt, err := types.BalanceFromString(amount)
	if err != nil {
		panic(err)
	}

	targetContent := xcrypto.HashContent{
		// TODO: fix event id
		Origin:         origin.NewDefaultOrigin(txHash, types.NetworkTestnet, "eventID").GetOrigin(),
		Receiver:       []byte(receiver),
		TargetNetwork:  types.NetworkTestnet,
		TargetContract: []byte(bridge),
		Data: operation.NewTransferOperation(
			"",
			"",
			fmt.Sprint(amount), "").GetContent(),
	}

	mt := merkle.NewTree(crypto.Keccak256, content1, targetContent, content2)
	path, _ := mt.Path(targetContent)

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

	recoveredKey, err := secp256k1.RecoverPubkey(mt.Root(), signature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey[1:]))

	act := action.NativeWithdrawArgs{
		Amount:     amnt,
		ReceiverID: receiver,
		Origin:     hexutil.Encode(targetContent.Origin[:]),
		Path:       make([][32]byte, len(path)),
		Signatures: [][]byte{signature[:64]},
		RecoveryID: 1,
	}

	fmt.Println("Content hash: " + base58.Encode(targetContent.CalculateHash()))
	for i, hash := range path {
		copy(act.Path[i][:], hash)
	}

	withdrawResp, err := cli.TransactionSend(ctx, sender, bridge, []base.Action{
		action.NewNativeWithdrawCall(act, GetGasPrice(ctx, cli), amnt),
	})
	if err != nil {
		panic(err)
	}
	return withdrawResp.String()
}
