package client

import "gitlab.com/rarify-protocol/near-bridge-go/pkg/types"

type ShardView struct {
	ShardID                  types.ShardID                 `json:"shard_id"`
	Chunk                    *ChunkView                    `json:"chunk"`
	ReceiptExecutionOutcomes []ExecutionOutcomeWithReceipt `json:"receipt_execution_outcomes"`
}
