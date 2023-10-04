package models

import (
	"encoding/json"
	"github.com/rarimo/near-go/common"
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
	SignerID   common.AccountID       `json:"signer_id"`
	PublicKey  common.Base58PublicKey `json:"public_key"`
	Nonce      common.Nonce           `json:"nonce"`
	ReceiverID common.AccountID       `json:"receiver_id"`
	Actions    []common.Action        `json:"actions"`
	Signature  common.Base58Signature `json:"signature"`
	Hash       common.Hash            `json:"hash"`
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
	PredecessorID common.AccountID `json:"predecessor_id"`
	ReceiverID    common.AccountID `json:"receiver_id"`
	ReceiptID     common.Hash      `json:"receipt_id"`
	Receipt       json.RawMessage  `json:"receipt"` // TODO: needs a type!
}

type ExecutionOutcomeView struct {
	Logs        []string          `json:"logs"`
	ReceiptIDs  []common.Hash     `json:"receipt_ids"`
	GasBurnt    common.Gas        `json:"gas_burnt"`
	TokensBurnt common.Balance    `json:"tokens_burnt"`
	ExecutorID  common.AccountID  `json:"executor_id"`
	Status      TransactionStatus `json:"status"`
}

type LogEvent struct {
	Standard string                  `json:"standard,required"`
	Version  string                  `json:"version,required"`
	Event    LogEventType            `json:"event,required"`
	Data     []LogEventDepositedData `json:"data,required"`
}

type LogEventDepositedData struct {
	Token *common.AccountID `json:"token,required"`
	// Empty if fungible token
	TokenID *string `json:"token_id,omitempty"`
	// Empty if non fungible token
	Amount    *common.Balance `json:"amount,omitempty"`
	Chain     string          `json:"chain,required"`
	IsWrapped bool            `json:"is_wrapped,required"`
}

type MerklePathItem struct {
	Hash      common.Hash `json:"hash"`
	Direction string      `json:"direction"` // TODO: enum type, either 'Left' or 'Right'
}

type MerklePath = []MerklePathItem

type ExecutionOutcomeWithIdView struct {
	Proof     MerklePath           `json:"proof"`
	BlockHash common.Hash          `json:"block_hash"`
	ID        common.Hash          `json:"id"`
	Outcome   ExecutionOutcomeView `json:"outcome"`
}

type ExecutionOutcomeWithReceipt struct {
	ExecutionOutcome ExecutionOutcomeWithIdView `json:"execution_outcome"`
	Receipt          *ReceiptView               `json:"receipt"`
}
