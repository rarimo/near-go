package scripts

import (
	"crypto/elliptic"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/mr-tron/base58"
	merkle "github.com/rarimo/go-merkle"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/bundle"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
)

func getContent(
	privateKey,
	txHash,
	eventID,
	receiver,
	targetContract,
	chainFrom,
	chainTo string,
	data operation.ContentData,
) (
	originHash string,
	signature string,
	resultPath [][32]byte,
	recoveryID byte,
) {
	targetContent := operation.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash(txHash).SetOpId(eventID).SetCurrentNetwork(chainFrom).Build().GetOrigin(),
		TargetNetwork:  chainTo,
		Receiver:       []byte(receiver),
		TargetContract: []byte(targetContract),
		Data:           data,
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle("").SetSalt("").Build().GetBundle(),
	}

	originHash = hexutil.Encode(targetContent.Origin[:])

	mt := merkle.NewTree(crypto.Keccak256, targetContent)
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

	sign, err := crypto.Sign(mt.Root(), pk)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature %s\n", base58.Encode(sign[:64]))

	recoveredKey, err := secp256k1.RecoverPubkey(mt.Root(), sign)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recovered pub key " + base58.Encode(recoveredKey[1:]))

	signature = hexutil.Encode(sign[:64])
	recoveryID = sign[64]
	resultPath = make([][32]byte, len(path))
	for i, hash := range path {
		copy(resultPath[i][:], hash)
	}

	fmt.Println("Content hash: " + base58.Encode(targetContent.CalculateHash()))
	return
}
