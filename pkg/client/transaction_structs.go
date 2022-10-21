package client

import (
	"encoding/json"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"

	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/hash"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/key"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/signature"
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
	SignerID   types.AccountID           `json:"signer_id"`
	PublicKey  key.Base58PublicKey       `json:"public_key"`
	Nonce      types.Nonce               `json:"nonce"`
	ReceiverID types.AccountID           `json:"receiver_id"`
	Actions    []base.Action             `json:"actions"`
	Signature  signature.Base58Signature `json:"signature"`
	Hash       hash.CryptoHash           `json:"hash"`
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
	PredecessorID types.AccountID `json:"predecessor_id"`
	ReceiverID    types.AccountID `json:"receiver_id"`
	ReceiptID     hash.CryptoHash `json:"receipt_id"`
	Receipt       json.RawMessage `json:"receipt"` // TODO: needs a type!
}

type ExecutionOutcomeView struct {
	Logs        []string          `json:"logs"`
	ReceiptIDs  []hash.CryptoHash `json:"receipt_ids"`
	GasBurnt    types.Gas         `json:"gas_burnt"`
	TokensBurnt types.Balance     `json:"tokens_burnt"`
	ExecutorID  types.AccountID   `json:"executor_id"`
	Status      TransactionStatus `json:"status"`
}

type LogEvent struct {
	Standard string                  `json:"standard,required"`
	Version  string                  `json:"version,required"`
	Event    LogEventType            `json:"event,required"`
	Data     []LogEventDepositedData `json:"data,required"`
}

type LogEventDepositedData struct {
	Token *types.AccountID `json:"token,required"`
	// Empty if fungible token
	TokenID *string `json:"token_id,omitempty"`
	// Empty if non fungible token
	Amount    *types.Balance `json:"amount,omitempty"`
	Chain     string         `json:"chain,required"`
	IsWrapped bool           `json:"is_wrapped,required"`
}

type MerklePathItem struct {
	Hash      hash.CryptoHash `json:"hash"`
	Direction string          `json:"direction"` // TODO: enum type, either 'Left' or 'Right'
}

type MerklePath = []MerklePathItem

type ExecutionOutcomeWithIdView struct {
	Proof     MerklePath           `json:"proof"`
	BlockHash hash.CryptoHash      `json:"block_hash"`
	ID        hash.CryptoHash      `json:"id"`
	Outcome   ExecutionOutcomeView `json:"outcome"`
}
