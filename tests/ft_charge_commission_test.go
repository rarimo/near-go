package tests

import (
	"context"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/scripts"
	"testing"
)

func TestFtChargeCommission(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	feeHash, depositHash := scripts.FtChargeCommission(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		"2000",
		cfg.FeerAddress,
	)

	printExplorerURL(t, "Charged fee", feeHash, nil)
	printExplorerURL(t, "Deposited token", depositHash, nil)
}
