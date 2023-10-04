//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"github.com/rarimo/near-go/scripts"
	"gitlab.com/distributed_lab/kit/kv"
	"testing"
)

func TestNftChargeCommission(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	feeHash, depositHash := scripts.NftChargeCommission(
		cfg.Ctx,
		cfg.Client,
		cfg.FeeTokenAddress,
		cfg.NftAddress,
		cfg.AccountID,
		cfg.AccountID,
		"6",
		cfg.FeerAddress,
	)

	printExplorerURL(t, "Charged fee", feeHash, nil)
	printExplorerURL(t, "Deposited token", depositHash, nil)
}
