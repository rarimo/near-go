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
		cfg.SignerPrivateKey,
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
		cfg.SignerPrivateKey,
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
		ftDepositHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.FtAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		true,
		false,
	)
	printExplorerURL(t, "Withdraw wrapped FT", ftWithdrawHash)

	ftDepositBackwardHash := scripts.FtDeposit(ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.FtAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn wrapped FT", ftDepositBackwardHash)

	ftWithdrawBackwardHash := scripts.FtWithdraw(
		ctx,
		client,
		ftDepositBackwardHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.FtAddressOriginal,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		false,
		false,
	)
	printExplorerURL(t, "Unlock original FT", ftWithdrawBackwardHash)

	// Native
	nativeDepositHash := scripts.NativeDeposit(
		ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.Amount,
		cfg.BridgeAddress,
	)
	printExplorerURL(t, "Deposited Native", nativeDepositHash)

	nativeWithdrawHash := scripts.NativeWithdraw(
		ctx,
		client,
		nativeDepositHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
	)
	printExplorerURL(t, "Withdraw wrapped Native", nativeWithdrawHash)

	nativeDepositBackwardHash := scripts.FtDeposit(ctx,
		client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NativeAddressWrapped,
		cfg.Amount,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn wrapped Native", ftDepositBackwardHash)

	nativeWithdrawBackwardHash := scripts.NativeWithdraw(
		ctx,
		client,
		nativeDepositBackwardHash,
		cfg.AccountID,
		cfg.AccountID,
		cfg.Amount,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
	)
	printExplorerURL(t, "Unlock Native", nativeWithdrawBackwardHash)
}
