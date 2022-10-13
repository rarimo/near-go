package tests

import (
	"context"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarify-protocol/near-bridge-go/scripts"
	"testing"
)

func printExplorerURL(t *testing.T, hash string, msg string) {
	t.Logf("%s: https://explorer.testnet.near.org/transactions/%s", msg, hash)
}

func TestAll(t *testing.T) {
	cfg, ctx, client := NewConfig(context.Background(), kv.MustFromEnv())

	// NFT
	nftDepositHash := scripts.NftDeposit(
		ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressOriginal,
		cfg.TokenID,
		cfg.BridgeAddress,
		false,
	)
	printExplorerURL(t, "Deposited NFT", nftDepositHash)

	nftWithdrawHash := scripts.NftWithdraw(
		ctx,
		client,
		nftDepositHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressWrapped,
		cfg.TokenID,
		cfg.BridgeAddress,
		cfg.PrivateKey, // ????
		true,
	)
	printExplorerURL(t, "Withdraw wrapped NFT", nftWithdrawHash)

	nftDepositBackwardHash := scripts.NftDeposit(ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressWrapped,
		cfg.TokenID,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn wrapped NFT", nftDepositBackwardHash)

	nftWithdrawBackwardHash := scripts.NftWithdraw(
		ctx,
		client,
		nftDepositBackwardHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressOriginal,
		cfg.TokenID,
		cfg.BridgeAddress,
		cfg.PrivateKey,
		false,
	)
	printExplorerURL(t, "Unlock original NFT", nftWithdrawBackwardHash)

	// FT
	ftDepositHash := scripts.FtDeposit(ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.FtAddressOriginal,
		cfg.Amount,
		cfg.BridgeAddress,
		false,
	)
	printExplorerURL(t, "Deposited FT", ftDepositHash)

	ftWithdrawHash := scripts.FtWithdraw(
		ctx,
		client,
		nftDepositHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.PrivateKey,
		true,
	)
	printExplorerURL(t, "Withdraw wrapped FT", ftWithdrawHash)

	ftDepositBackwardHash := scripts.FtDeposit(ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn wrapped NFT", ftDepositBackwardHash)

	ftWithdrawBackwardHash := scripts.FtWithdraw(
		ctx,
		client,
		nftDepositBackwardHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressOriginal,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.PrivateKey,
		false,
	)
	printExplorerURL(t, "Unlock original FT", ftWithdrawBackwardHash)

}
