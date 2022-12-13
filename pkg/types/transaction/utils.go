package transaction

import "gitlab.com/rarimo/near-bridge-go/pkg/types/key"

func SignAndSerializeTransaction(keyPair key.KeyPair, txn Transaction) (blob string, err error) {
	var stxn SignedTransaction
	if stxn, err = NewSignedTransaction(keyPair, txn); err != nil {
		return
	}

	blob, err = stxn.Serialize()
	return
}
