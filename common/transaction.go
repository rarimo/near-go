package common

import (
	"encoding/json"
	"github.com/near/borsh-go"
)

type LogEventType = string

const (
	LogEventTypeNftDeposited    LogEventType = "nft_deposited"
	LogEventTypeFtDeposited     LogEventType = "ft_deposited"
	LogEventTypeNativeDeposited LogEventType = "native_deposited"
)

type TransactionStatus struct {
	SuccessValue     string          `json:"SuccessValue"`
	SuccessReceiptID string          `json:"SuccessReceiptId"`
	Failure          json.RawMessage `json:"Failure"` // TODO
}

type SignedTransactionView struct {
	SignerID   AccountID       `json:"signer_id"`
	PublicKey  Base58PublicKey `json:"public_key"`
	Nonce      Nonce           `json:"nonce"`
	ReceiverID AccountID       `json:"receiver_id"`
	Actions    []Action        `json:"actions"`
	Signature  Base58Signature `json:"signature"`
	Hash       Hash            `json:"hash"`
}

type FinalExecutionOutcomeView struct {
	Status             TransactionStatus            `json:"status"`
	Transaction        SignedTransactionView        `json:"transaction"`
	TransactionOutcome ExecutionOutcomeWithIdView   `json:"transaction_outcome"`
	ReceiptsOutcome    []ExecutionOutcomeWithIdView `json:"receipts_outcome"`
}

type FinalExecutionOutcomeWithReceiptView struct {
	FinalExecutionOutcomeView
	Receipts []ReceiptView `json:"receipts"`
}

type ReceiptView struct {
	PredecessorID AccountID       `json:"predecessor_id"`
	ReceiverID    AccountID       `json:"receiver_id"`
	ReceiptID     Hash            `json:"receipt_id"`
	Receipt       json.RawMessage `json:"receipt"` // TODO: needs a type!
}

type ExecutionOutcomeView struct {
	Logs        []string          `json:"logs"`
	ReceiptIDs  []Hash            `json:"receipt_ids"`
	GasBurnt    Gas               `json:"gas_burnt"`
	TokensBurnt Balance           `json:"tokens_burnt"`
	ExecutorID  AccountID         `json:"executor_id"`
	Status      TransactionStatus `json:"status"`
}

type LogEvent struct {
	Standard string                  `json:"standard,required"`
	Version  string                  `json:"version,required"`
	Event    LogEventType            `json:"event,required"`
	Data     []LogEventDepositedData `json:"data,required"`
}

type LogEventDepositedData struct {
	Token *AccountID `json:"token,required"`
	// Empty if fungible token
	TokenID *string `json:"token_id,omitempty"`
	// Empty if non fungible token
	Amount    *Balance `json:"amount,omitempty"`
	Chain     string   `json:"chain,required"`
	IsWrapped bool     `json:"is_wrapped,required"`
}

type MerklePathItem struct {
	Hash      Hash   `json:"hash"`
	Direction string `json:"direction"` // TODO: enum type, either 'Left' or 'Right'
}

type MerklePath = []MerklePathItem

type ExecutionOutcomeWithIdView struct {
	Proof     MerklePath           `json:"proof"`
	BlockHash Hash                 `json:"block_hash"`
	ID        Hash                 `json:"id"`
	Outcome   ExecutionOutcomeView `json:"outcome"`
}

type ExecutionOutcomeWithReceipt struct {
	ExecutionOutcome ExecutionOutcomeWithIdView `json:"execution_outcome"`
	Receipt          *ReceiptView               `json:"receipt"`
}

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
