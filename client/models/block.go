package models

import (
	"github.com/rarimo/near-go/common"
)

type ChallengesResult = []SlashedValidator

type SlashedValidator struct {
	AccountID    common.AccountID `json:"account_id"`
	IsDoubleSign bool             `json:"is_double_sign"`
}

// ValidatorStakeView is based on ValidatorStakeV1 struct in nearcore
type ValidatorStakeView struct {
	AccountID common.AccountID       `json:"account_id"`
	PublicKey common.Base58PublicKey `json:"public_key"`
	Stake     common.Balance         `json:"stake"`
}

type BlockView struct {
	Author common.AccountID  `json:"author"`
	Header BlockHeaderView   `json:"header"`
	Chunks []ChunkHeaderView `json:"chunks"`
}

type BlockHeaderView struct {
	Height                common.BlockHeight        `json:"height"`
	EpochID               common.Hash               `json:"epoch_id"`
	NextEpochID           common.Hash               `json:"next_epoch_id"`
	Hash                  common.Hash               `json:"hash"`
	PrevHash              common.Hash               `json:"prev_hash"`
	PrevStateRoot         common.Hash               `json:"prev_state_root"`
	ChunkReceiptsRoot     common.Hash               `json:"chunk_receipts_root"`
	ChunkHeadersRoot      common.Hash               `json:"chunk_headers_root"`
	ChunkTxRoot           common.Hash               `json:"chunk_tx_root"`
	OutcomeRoot           common.Hash               `json:"outcome_root"`
	ChunksIncluded        uint64                    `json:"chunks_included"`
	ChallengesRoot        common.Hash               `json:"challenges_root"`
	Timestamp             uint64                    `json:"timestamp"`         // milliseconds
	TimestampNanosec      common.TimeNanos          `json:"timestamp_nanosec"` // nanoseconds, uint128
	RandomValue           common.Hash               `json:"random_value"`
	ValidatorProposals    []ValidatorStakeView      `json:"validator_proposals"`
	ChunkMask             []bool                    `json:"chunk_mask"`
	GasPrice              common.Balance            `json:"gas_price"`
	RentPaid              common.Balance            `json:"rent_paid"`        // NOTE: deprecated - 2021-05-14
	ValidatorReward       common.Balance            `json:"validator_reward"` // NOTE: deprecated - 2021-05-14
	TotalSupply           common.Balance            `json:"total_supply"`
	ChallengesResult      ChallengesResult          `json:"challenges_result"`
	LastFinalBlock        common.Hash               `json:"last_final_block"`
	LastDSFinalBlock      common.Hash               `json:"last_ds_final_block"`
	NextBPHash            common.Hash               `json:"next_bp_hash"`
	BlockMerkleRoot       common.Hash               `json:"block_merkle_root"`
	Approvals             []*common.Base58Signature `json:"approvals"`
	Signature             common.Base58Signature    `json:"signature"`
	LatestProtocolVersion uint64                    `json:"latest_protocol_version"`
}
