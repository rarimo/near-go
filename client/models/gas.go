package models

import (
	"github.com/rarimo/near-go/common"
)

type GasPrice struct {
	GasPrice common.Balance `json:"gas_price"`
}
