//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"testing"

	"github.com/rarimo/near-go/scripts"
	"gitlab.com/distributed_lab/kit/kv"
)

func TestFt(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	depositHash, depositedEventID := scripts.FtDeposit(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.FtAddress,
		cfg.Amount,
		cfg.BridgeAddress,
		false,
	)
	printExplorerURL(t, "Deposited FT", depositHash, &depositedEventID)

	withdrawHash := scripts.FtWithdraw(
		cfg.Ctx,
		cfg.Client,
		depositHash,
		depositedEventID,
		cfg.AccountID,
		cfg.AccountID,
		chainID,
		chainID,
		cfg.FtAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		true,
	)
	printExplorerURL(t, "Withdraw wrapped FT", withdrawHash, nil)

	depositBackwardHash, depositedBackwardEventID := scripts.FtDeposit(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.FtAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn wrapped FT", depositBackwardHash, &depositedBackwardEventID)

	ftWithdrawBackwardHash := scripts.FtWithdraw(
		cfg.Ctx,
		cfg.Client,
		depositBackwardHash,
		depositedBackwardEventID,
		cfg.AccountID,
		cfg.AccountID,
		chainID,
		chainID,
		cfg.FtAddress,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		false,
	)

	printExplorerURL(t, "Unlock original FT", ftWithdrawBackwardHash, nil)
}
