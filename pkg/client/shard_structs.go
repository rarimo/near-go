package client

import (
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/hash"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/key"
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
	AccountID     types.AccountID      `json:"account_id"`
	AccessKey     *AccessKeyView       `json:"access_key,omitempty"`
	PublicKey     *key.Base58PublicKey `json:"public_key,omitempty"`
	Amount        *types.Balance       `json:"amount,omitempty"`
	Locked        *types.Balance       `json:"locked,omitempty"`
	CodeHash      *hash.CryptoHash     `json:"code_hash,omitempty"`
	StorageUsage  *types.StorageUsage  `json:"storage_usage,omitempty"`
	StoragePaidAt *types.BlockHeight   `json:"storage_paid_at,omitempty"`
}
