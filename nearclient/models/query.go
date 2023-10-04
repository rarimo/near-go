package models

import (
	"github.com/rarimo/near-go/common"
)

type QueryResponse struct {
	BlockHeight common.BlockHeight `json:"block_height"`
	BlockHash   common.Hash        `json:"block_hash"`
	Error       *string            `json:"error"`
	Logs        []interface{}      `json:"logs"` // TODO: use correct type
}
