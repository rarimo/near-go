//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
	"gitlab.com/rarimo/near-bridge-go/scripts"
	"lukechampine.com/uint128"
	"testing"
)

func TestUpdateFeeToken(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	hash := scripts.ManageFeeToken(
		cfg.Ctx,
		action.FeeUpdateFeeToken,
		cfg.Client,
		cfg.AccountID,
		cfg.FeerAddress,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		action.FeeToken{
			TokenAddr: &cfg.FeeTokenAddress,
			TokenType: action.TokenType_FT,
			Fee:       types.Balance(uint128.From64(10000)),
		},
	)
	printExplorerURL(t, "Updated fee token", hash, nil)
}
