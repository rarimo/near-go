package models

import (
	"encoding/json"
	"github.com/rarimo/near-go/common"
)

type ChunkView struct {
	Author       common.AccountID        `json:"author"`
	Header       ChunkHeaderView         `json:"header"`
	Transactions []SignedTransactionView `json:"transactions"`
	Receipts     []ReceiptView           `json:"receipts"`
}

type ChunkHeaderView struct {
	ChunkHash            common.Hash            `json:"chunk_hash"`
	PrevBlockHash        common.Hash            `json:"prev_block_hash"`
	OutcomeRoot          common.Hash            `json:"outcome_root"`
	PrevStateRoot        json.RawMessage        `json:"prev_state_root"` // TODO: needs a type!
	EncodedMerkleRoot    common.Hash            `json:"encoded_merkle_root"`
	EncodedLength        uint64                 `json:"encoded_length"`
	HeightCreated        common.BlockHeight     `json:"height_created"`
	HeightIncluded       common.BlockHeight     `json:"height_included"`
	ShardID              common.ShardID         `json:"shard_id"`
	GasUsed              common.Gas             `json:"gas_used"`
	GasLimit             common.Gas             `json:"gas_limit"`
	RentPaid             common.Balance         `json:"rent_paid"`        // TODO: deprecated
	ValidatorReward      common.Balance         `json:"validator_reward"` // TODO: deprecated
	BalanceBurnt         common.Balance         `json:"balance_burnt"`
	OutgoingReceiptsRoot common.Hash            `json:"outgoing_receipts_root"`
	TxRoot               common.Hash            `json:"tx_root"`
	ValidatorProposals   []ValidatorStakeView   `json:"validator_proposals"`
	Signature            common.Base58Signature `json:"signature"`
}
