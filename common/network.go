package common

import (
	"encoding/json"
	"time"
)

// NetworkInfo holds network information
type NetworkInfo struct {
	ActivePeers         []FullPeerInfo  `json:"active_peers"`
	NumActivePeers      uint            `json:"num_active_peers"`
	PeerMaxCount        uint32          `json:"peer_max_count"`
	HighestHeightPeers  []FullPeerInfo  `json:"highest_height_peers"`
	SentBytesPerSec     uint64          `json:"sent_bytes_per_sec"`
	ReceivedBytesPerSec uint64          `json:"received_bytes_per_sec"`
	KnownProducers      []KnownProducer `json:"known_producers"`
	MetricRecorder      MetricRecorder  `json:"metric_recorder"`
	PeerCounter         uint            `json:"peer_counter"`
}

type FullPeerInfo struct {
	PeerInfo  PeerInfo      `json:"peer_info"`
	ChainInfo PeerChainInfo `json:"chain_info"`
	EdgeInfo  EdgeInfo      `json:"edge_info"`
}

// PeerInfo holds peer information
type PeerInfo struct {
	ID        PeerID     `json:"id"`
	Addr      *string    `json:"addr"`
	AccountID *AccountID `json:"account_id"`
}

// PeerChainInfo contains peer chain information. This is derived from PeerCHainInfoV2 in nearcore
type PeerChainInfo struct {
	// ChainTo Id and hash of genesis block.
	GenesisID GenesisID `json:"genesis_id"`
	// Last known chain height of the peer.
	Height BlockHeight `json:"height"`
	// Shards that the peer is tracking.
	TrackedShards []ShardID `json:"tracked_shards"`
	// Denote if a node is running in archival mode or not.
	Archival bool `json:"archival"`
}

// EdgeInfo contains information that will be ultimately used to create a new edge. It contains nonce proposed for the edge with signature from peer.
type EdgeInfo struct {
	Nonce     Nonce     `json:"nonce"`
	Signature Signature `json:"signature"`
}

// KnownProducer is basically PeerInfo, but AccountID is known
type KnownProducer struct {
	AccountID AccountID `json:"account_id"`
	Addr      *string   `json:"addr"`
	PeerID    PeerID    `json:"peer_id"`
}

// TODO: chain/network/src/recorder.rs
type MetricRecorder = json.RawMessage

type GenesisID struct {
	// ChainTo Id
	ChainID string `json:"chain_id"`
	// Hash of genesis block
	Hash Hash `json:"hash"`
}

type StatusResponse struct {
	// Binary version
	Version NodeVersion `json:"version"`
	// Unique chain id.
	ChainID string `json:"chain_id"`
	// Currently active protocol version.
	ProtocolVersion uint32 `json:"protocol_version"`
	// Latest protocol version that this client supports.
	LatestProtocolVersion uint32 `json:"latest_protocol_version"`
	// Address for RPC server.
	RPCAddr string `json:"rpc_addr"`
	// Current epoch validators.
	Validators []ValidatorInfo `json:"validators"`
	// Sync status of the node.
	SyncInfo StatusSyncInfo `json:"sync_info"`
	// Validator id of the node
	ValidatorAccountID *AccountID `json:"validator_account_id"`
}

type NodeVersion struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

type ValidatorInfo struct {
	AccountID AccountID `json:"account_id"`
	Slashed   bool      `json:"is_slashed"`
}

type StatusSyncInfo struct {
	LatestBlockHash   Hash        `json:"latest_block_hash"`
	LatestBlockHeight BlockHeight `json:"latest_block_height"`
	LatestBlockTime   time.Time   `json:"latest_block_time"`
	Syncing           bool        `json:"syncing"`
}

type ValidatorsResponse struct {
	CurrentValidators []CurrentEpochValidatorInfo `json:"current_validator"`
}

type CurrentEpochValidatorInfo struct {
	ValidatorInfo
	PublicKey         Base58PublicKey `json:"public_key"`
	Stake             Balance         `json:"stake"`
	Shards            []ShardID       `json:"shards"`
	NumProducedBlocks NumBlocks       `json:"num_produced_blocks"`
	NumExpectedBlocks NumBlocks       `json:"num_expected_blocks"`
}
