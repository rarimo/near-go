package client

import (
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/hash"
)

type AccountView struct {
	Amount        types.Balance      `json:"amount"`
	Locked        types.Balance      `json:"locked"`
	CodeHash      hash.CryptoHash    `json:"code_hash"`
	StorageUsage  types.StorageUsage `json:"storage_usage"`
	StoragePaidAt types.BlockHeight  `json:"storage_paid_at"`

	QueryResponse
}
