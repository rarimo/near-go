package client

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/hash"
)

type ShardView struct {
	ShardID                  types.ShardID                 `json:"shard_id"`
	Chunk                    *ShardChunkView               `json:"chunk"`
	ReceiptExecutionOutcomes []ExecutionOutcomeWithReceipt `json:"receipt_execution_outcomes"`
	StateChanges             []ShardStateChangeView        `json:"state_changes"`
}

type ShardChunkView struct {
	Author       types.AccountID             `json:"author"`
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
	Type   string          `json:"type"`
	TxHash hash.CryptoHash `json:"tx_hash"`
}

type ShardStateChangeChangeView struct {
	AccountID     types.AccountID    `json:"account_id"`
	Amount        types.Balance      `json:"amount"`
	Locked        types.Balance      `json:"locked"`
	CodeHash      hash.CryptoHash    `json:"code_hash"`
	StorageUsage  types.StorageUsage `json:"storage_usage"`
	StoragePaidAt types.BlockHeight  `json:"storage_paid_at"`
}
