//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/scripts"
	"gitlab.com/distributed_lab/kit/kv"
	"lukechampine.com/uint128"
	"testing"
)

func TestRemoveFeeToken(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	hash := scripts.ManageFeeToken(
		cfg.Ctx,
		common.FeeRemoveFeeToken,
		cfg.Client,
		cfg.AccountID,
		cfg.FeerAddress,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		common.FeeToken{
			TokenAddr: &cfg.FeeTokenAddress,
			TokenType: common.TokenType_FT,
			Fee:       common.Balance(uint128.From64(10000)),
		},
	)
	printExplorerURL(t, "Removed fee token", hash, nil)
}
