//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/scripts"
)

var chainID = "Near"

func printExplorerURL(t *testing.T, msg string, hash string, receiptID *string) {
	t.Logf("%s: https://explorer.testnet.near.org/transactions/%s", msg, hash)
	if receiptID != nil {
		t.Logf("%s bridge receipt: https://explorer.testnet.near.org/transactions/%s#%s", msg, hash, *receiptID)
	}
}

func TestNft(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	depositHash, depositedEventID := scripts.NftDeposit(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddress,
		cfg.TokenID,
		cfg.BridgeAddress,
		false,
	)
	printExplorerURL(t, "Deposited NFT", depositHash, &depositedEventID)

	withdrawHash := scripts.NftWithdraw(
		cfg.Ctx,
		cfg.Client,
		depositHash,
		depositedEventID,
		cfg.AccountID,
		cfg.AccountID,
		chainID,
		chainID,
		cfg.NftAddressWrapped,
		cfg.TokenID,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		true,
	)
	printExplorerURL(t, "Withdraw wrapped NFT", withdrawHash, nil)

	depositBackwardHash, depositedBackwardEventID := scripts.NftDeposit(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressWrapped,
		cfg.TokenID,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn Deposited NFT", depositBackwardHash, &depositedBackwardEventID)

	withdrawBackwardHash := scripts.NftWithdraw(
		cfg.Ctx,
		cfg.Client,
		depositBackwardHash,
		depositedBackwardEventID,
		cfg.AccountID,
		cfg.AccountID,
		chainID,
		chainID,
		cfg.NftAddress,
		cfg.TokenID,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		false,
	)
	printExplorerURL(t, "Unlock original NFT", withdrawBackwardHash, nil)
}
