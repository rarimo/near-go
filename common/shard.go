package common

type ShardView struct {
	ShardID                  ShardID                       `json:"shard_id"`
	Chunk                    *ShardChunkView               `json:"chunk"`
	ReceiptExecutionOutcomes []ExecutionOutcomeWithReceipt `json:"receipt_execution_outcomes"`
	StateChanges             []ShardStateChangeView        `json:"state_changes"`
}

type ShardChunkView struct {
	Author       AccountID                   `json:"author"`
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
	Type   string `json:"type"`
	TxHash Hash   `json:"tx_hash"`
}

type ShardStateChangeChangeView struct {
	AccountID     AccountID        `json:"account_id"`
	AccessKey     *AccessKeyView   `json:"access_key,omitempty"`
	PublicKey     *Base58PublicKey `json:"public_key,omitempty"`
	Amount        *Balance         `json:"amount,omitempty"`
	Locked        *Balance         `json:"locked,omitempty"`
	CodeHash      *Hash            `json:"code_hash,omitempty"`
	StorageUsage  *StorageUsage    `json:"storage_usage,omitempty"`
	StoragePaidAt *BlockHeight     `json:"storage_paid_at,omitempty"`
}
