//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"github.com/rarimo/near-go/scripts"
	"gitlab.com/distributed_lab/kit/kv"
	"testing"
)

func TestFtChargeCommission(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	feeHash, depositHash := scripts.FtChargeCommission(
		cfg.Ctx,
		cfg.Client,
		cfg.FeeTokenAddress,
		cfg.FtAddress,
		cfg.AccountID,
		cfg.AccountID,
		"2000",
		cfg.FeerAddress,
	)

	printExplorerURL(t, "Charged fee", feeHash, nil)
	printExplorerURL(t, "Deposited token", depositHash, nil)
}
