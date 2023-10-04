package models

import (
	"github.com/rarimo/near-go/common"
)

type ShardView struct {
	ShardID                  common.ShardID                `json:"shard_id"`
	Chunk                    *ShardChunkView               `json:"chunk"`
	ReceiptExecutionOutcomes []ExecutionOutcomeWithReceipt `json:"receipt_execution_outcomes"`
	StateChanges             []ShardStateChangeView        `json:"state_changes"`
}

type ShardChunkView struct {
	Author       common.AccountID            `json:"author"`
	Header       ChunkHeaderView             `json:"header"`
	Receipts     []ReceiptView               `json:"receipts"`
	Transactions []ShardChunkTransactionView `json:"transactions"`
}

type ShardChunkTransactionView struct {
	Outcome     ExecutionOutcomeWithReceipt `json:"outcome"`
	Transaction SignedTransactionView       `json:"transaction"`
}

type ShardStateChangeView struct {
	Type   string                     `json:"type"`
	Cause  ShardStateChangeCauseView  `json:"cause"`
	Change ShardStateChangeChangeView `json:"change"`
}

type ShardStateChangeCauseView struct {
	Type   string      `json:"type"`
	TxHash common.Hash `json:"tx_hash"`
}

type ShardStateChangeChangeView struct {
	AccountID     common.AccountID        `json:"account_id"`
	AccessKey     *AccessKeyView          `json:"access_key,omitempty"`
	PublicKey     *common.Base58PublicKey `json:"public_key,omitempty"`
	Amount        *common.Balance         `json:"amount,omitempty"`
	Locked        *common.Balance         `json:"locked,omitempty"`
	CodeHash      *common.Hash            `json:"code_hash,omitempty"`
	StorageUsage  *common.StorageUsage    `json:"storage_usage,omitempty"`
	StoragePaidAt *common.BlockHeight     `json:"storage_paid_at,omitempty"`
}
