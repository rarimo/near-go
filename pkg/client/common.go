package client

import (
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/hash"
)

type QueryResponse struct {
	BlockHeight types.BlockHeight `json:"block_height"`
	BlockHash   hash.CryptoHash   `json:"block_hash"`
	Error       *string           `json:"error"`
	Logs        []interface{}     `json:"logs"` // TODO: use correct type
}
