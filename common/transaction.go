package common

import (
	"github.com/near/borsh-go"
)

type Transaction struct {
	SignerID   AccountID
	PublicKey  PublicKey
	Nonce      Nonce
	ReceiverID AccountID
	BlockHash  Hash
	Actions    []Action
}

func (t Transaction) Hash() (txnHash Hash, serialized []byte, err error) {
	// Serialize into Borsh
	serialized, err = borsh.Serialize(t)
	if err != nil {
		return
	}
	txnHash = NewCryptoHash(serialized)
	return
}

func (t Transaction) HashAndSign(keyPair KeyPair) (txnHash Hash, serialized []byte, sig Signature, err error) {
	txnHash, serialized, err = t.Hash()
	if err != nil {
		return
	}

	sig = keyPair.Sign(txnHash[:])
	return
}

func SignAndSerializeTransaction(keyPair KeyPair, txn Transaction) (blob string, err error) {
	var stxn SignedTransaction
	if stxn, err = NewSignedTransaction(keyPair, txn); err != nil {
		return
	}

	blob, err = stxn.Serialize()
	return
}
