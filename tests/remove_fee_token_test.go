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

func TestRemoveFeeToken(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())
	tokenAddr := "ft_test_fee.napalmpapalam.testnet"

	hash := scripts.ManageFeeToken(
		cfg.Ctx,
		action.FeeRemoveFeeToken,
		cfg.Client,
		cfg.AccountID,
		cfg.FeerAddress,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		action.FeeToken{
			TokenAddr: &tokenAddr,
			TokenType: action.TokenType_FT,
			Fee:       types.Balance(uint128.From64(10000)),
		},
	)
	printExplorerURL(t, "Removed fee token", hash, nil)
}
