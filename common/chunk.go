package common

import (
	"encoding/json"
)

type ChunkView struct {
	Author       AccountID               `json:"author"`
	Header       ChunkHeaderView         `json:"header"`
	Transactions []SignedTransactionView `json:"transactions"`
	Receipts     []ReceiptView           `json:"receipts"`
}

type ChunkHeaderView struct {
	ChunkHash            Hash                 `json:"chunk_hash"`
	PrevBlockHash        Hash                 `json:"prev_block_hash"`
	OutcomeRoot          Hash                 `json:"outcome_root"`
	PrevStateRoot        json.RawMessage      `json:"prev_state_root"` // TODO: needs a type!
	EncodedMerkleRoot    Hash                 `json:"encoded_merkle_root"`
	EncodedLength        uint64               `json:"encoded_length"`
	HeightCreated        BlockHeight          `json:"height_created"`
	HeightIncluded       BlockHeight          `json:"height_included"`
	ShardID              ShardID              `json:"shard_id"`
	GasUsed              Gas                  `json:"gas_used"`
	GasLimit             Gas                  `json:"gas_limit"`
	RentPaid             Balance              `json:"rent_paid"`        // TODO: deprecated
	ValidatorReward      Balance              `json:"validator_reward"` // TODO: deprecated
	BalanceBurnt         Balance              `json:"balance_burnt"`
	OutgoingReceiptsRoot Hash                 `json:"outgoing_receipts_root"`
	TxRoot               Hash                 `json:"tx_root"`
	ValidatorProposals   []ValidatorStakeView `json:"validator_proposals"`
	Signature            Base58Signature      `json:"signature"`
}
