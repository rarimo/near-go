package common

type ChallengesResult = []SlashedValidator

type SlashedValidator struct {
	AccountID    AccountID `json:"account_id"`
	IsDoubleSign bool      `json:"is_double_sign"`
}

// ValidatorStakeView is based on ValidatorStakeV1 struct in nearcore
type ValidatorStakeView struct {
	AccountID AccountID       `json:"account_id"`
	PublicKey Base58PublicKey `json:"public_key"`
	Stake     Balance         `json:"stake"`
}

type BlockView struct {
	Author AccountID         `json:"author"`
	Header BlockHeaderView   `json:"header"`
	Chunks []ChunkHeaderView `json:"chunks"`
}

type BlockHeaderView struct {
	Height                BlockHeight          `json:"height"`
	EpochID               Hash                 `json:"epoch_id"`
	NextEpochID           Hash                 `json:"next_epoch_id"`
	Hash                  Hash                 `json:"hash"`
	PrevHash              Hash                 `json:"prev_hash"`
	PrevStateRoot         Hash                 `json:"prev_state_root"`
	ChunkReceiptsRoot     Hash                 `json:"chunk_receipts_root"`
	ChunkHeadersRoot      Hash                 `json:"chunk_headers_root"`
	ChunkTxRoot           Hash                 `json:"chunk_tx_root"`
	OutcomeRoot           Hash                 `json:"outcome_root"`
	ChunksIncluded        uint64               `json:"chunks_included"`
	ChallengesRoot        Hash                 `json:"challenges_root"`
	Timestamp             uint64               `json:"timestamp"`         // milliseconds
	TimestampNanosec      TimeNanos            `json:"timestamp_nanosec"` // nanoseconds, uint128
	RandomValue           Hash                 `json:"random_value"`
	ValidatorProposals    []ValidatorStakeView `json:"validator_proposals"`
	ChunkMask             []bool               `json:"chunk_mask"`
	GasPrice              Balance              `json:"gas_price"`
	RentPaid              Balance              `json:"rent_paid"`        // NOTE: deprecated - 2021-05-14
	ValidatorReward       Balance              `json:"validator_reward"` // NOTE: deprecated - 2021-05-14
	TotalSupply           Balance              `json:"total_supply"`
	ChallengesResult      ChallengesResult     `json:"challenges_result"`
	LastFinalBlock        Hash                 `json:"last_final_block"`
	LastDSFinalBlock      Hash                 `json:"last_ds_final_block"`
	NextBPHash            Hash                 `json:"next_bp_hash"`
	BlockMerkleRoot       Hash                 `json:"block_merkle_root"`
	Approvals             []*Base58Signature   `json:"approvals"`
	Signature             Base58Signature      `json:"signature"`
	LatestProtocolVersion uint64               `json:"latest_protocol_version"`
}
