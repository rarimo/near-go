package client

import "gitlab.com/rarimo/near-bridge-go/pkg/types"

type GasPrice struct {
	GasPrice types.Balance `json:"gas_price"`
}
