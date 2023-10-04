package common

import (
	"encoding/base64"
	"github.com/near/borsh-go"
)

type SignedTransaction struct {
	Transaction Transaction
	Signature   Signature

	SerializedTransaction []byte `borsh_skip:"true"`
	hash                  Hash   `borsh_skip:"true"`
	size                  int    `borsh_skip:"true"`
}

func NewSignedTransaction(keyPair KeyPair, transaction Transaction) (stxn SignedTransaction, err error) {
	stxn.Transaction = transaction
	stxn.hash, stxn.SerializedTransaction, stxn.Signature, err = transaction.HashAndSign(keyPair)
	if err != nil {
		return
	}

	stxn.size = len(stxn.SerializedTransaction)
	return
}

func (st *SignedTransaction) Hash() Hash {
	return st.hash
}

func (st *SignedTransaction) Size() int {
	return st.size
}

func (st *SignedTransaction) Verify(pubKey PublicKey) (ok bool, err error) {
	var txnHash Hash
	txnHash, _, err = st.Transaction.Hash()
	if err != nil {
		return
	}

	return pubKey.Verify(txnHash[:], st.Signature)
}

func (st SignedTransaction) Serialize() (serialized string, err error) {
	var blob []byte

	blob, err = borsh.Serialize(st)
	if err != nil {
		return
	}

	serialized = base64.StdEncoding.EncodeToString(blob)

	return
}
