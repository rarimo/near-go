package tests

import (
	"context"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/scripts"
	"testing"
)

func TestNftChargeCommission(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	feeHash, depositHash := scripts.NftChargeCommission(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		"6",
		cfg.FeerAddress,
	)

	printExplorerURL(t, "Charged fee", feeHash, nil)
	printExplorerURL(t, "Deposited token", depositHash, nil)
}
