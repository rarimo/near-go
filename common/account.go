package common

import (
	"github.com/rarimo/near-go/nearclient/models"
)

type AccountView struct {
	Amount        Balance      `json:"amount"`
	Locked        Balance      `json:"locked"`
	CodeHash      Hash         `json:"code_hash"`
	StorageUsage  StorageUsage `json:"storage_usage"`
	StoragePaidAt BlockHeight  `json:"storage_paid_at"`

	models.QueryResponse
}
