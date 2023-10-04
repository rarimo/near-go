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

func TestWithdrawFeeToken(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())
	tokenAddr := "ft_test_fee.napalmpapalam.testnet"

	hash := scripts.ManageFeeToken(
		cfg.Ctx,
		common.FeeWithdraw,
		cfg.Client,
		cfg.AccountID,
		cfg.FeerAddress,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		common.FeeToken{
			TokenAddr: &tokenAddr,
			TokenType: common.TokenType_FT,
			Fee:       common.Balance(uint128.From64(1000)),
		},
	)
	printExplorerURL(t, "Withdraw fee token", hash, nil)
}
