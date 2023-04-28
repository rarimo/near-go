package tests

import (
	"context"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/scripts"
	"testing"
)

func TestChargeNativeCommission(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	feeHash, depositHash := scripts.NativeChargeCommission(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		"2100",
		cfg.FeerAddress,
	)

	printExplorerURL(t, "Native Charged fee", feeHash, nil)
	printExplorerURL(t, "Deposited token", depositHash, nil)
}
