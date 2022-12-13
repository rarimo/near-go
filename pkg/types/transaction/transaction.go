package transaction

import (
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action/base"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/hash"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/key"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/signature"
)

type Transaction struct {
	SignerID   types.AccountID
	PublicKey  key.PublicKey
	Nonce      types.Nonce
	ReceiverID types.AccountID
	BlockHash  hash.CryptoHash
	Actions    []base.Action
}

func (t Transaction) Hash() (txnHash hash.CryptoHash, serialized []byte, err error) {
	// Serialize into Borsh
	serialized, err = borsh.Serialize(t)
	if err != nil {
		return
	}
	txnHash = hash.NewCryptoHash(serialized)
	return
}

func (t Transaction) HashAndSign(keyPair key.KeyPair) (txnHash hash.CryptoHash, serialized []byte, sig signature.Signature, err error) {
	txnHash, serialized, err = t.Hash()
	if err != nil {
		return
	}

	sig = keyPair.Sign(txnHash[:])
	return
}
