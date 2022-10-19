package tests

import (
	"context"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarify-protocol/near-bridge-go/scripts"
	"testing"
)

func printExplorerURL(t *testing.T, msg string, hash string, receiptID *string) {
	t.Logf("%s: https://explorer.testnet.near.org/transactions/%s", msg, hash)
	if receiptID != nil {
		t.Logf("%s bridge receipt: https://explorer.testnet.near.org/transactions/%s#%s", msg, hash, *receiptID)
	}
}

var chainID = "Near"

func TestAll(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	// NFT
	nftDepositHash, depositedEventID := scripts.NftDeposit(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressOriginal,
		cfg.TokenID,
		cfg.BridgeAddress,
		false,
	)
	printExplorerURL(t, "Deposited NFT", nftDepositHash, &depositedEventID)

	nftWithdrawHash := scripts.NftWithdraw(
		cfg.Ctx,
		cfg.Client,
		nftDepositHash,
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
	printExplorerURL(t, "Withdraw wrapped NFT", nftWithdrawHash, nil)

	nftDepositBackwardHash, nftDepositBackwardEventID := scripts.NftDeposit(
		cfg.Ctx,
		cfg.Client,
		cfg.AccountID,
		cfg.AccountID,
		cfg.NftAddressWrapped,
		cfg.TokenID,
		cfg.BridgeAddress,
		true,
	)
	printExplorerURL(t, "Burn Deposited NFT", nftDepositBackwardHash, &nftDepositBackwardEventID)

	nftWithdrawBackwardHash := scripts.NftWithdraw(
		cfg.Ctx,
		cfg.Client,
		nftDepositBackwardHash,
		nftDepositBackwardEventID,
		cfg.AccountID,
		cfg.AccountID,
		chainID,
		chainID,
		cfg.NftAddressOriginal,
		cfg.TokenID,
		cfg.BridgeAddress,
		cfg.SignerPrivateKey,
		false,
	)
	printExplorerURL(t, "Unlock original NFT", nftWithdrawBackwardHash, nil)

	//// FT
	//ftDepositHash := scripts.FtDeposit(
	//	cfg.Ctx,
	//	cfg.Client,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	cfg.FtAddressOriginal,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	false,
	//)
	//printExplorerURL(t, "Deposited FT", ftDepositHash)
	//
	//ftWithdrawHash := scripts.FtWithdraw(
	//	cfg.Ctx,
	//	cfg.Client,
	//	ftDepositHash,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	chainID,
	//	cfg.FtAddressWrapped,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	cfg.SignerPrivateKey,
	//	true,
	//	false,
	//)
	//printExplorerURL(t, "Withdraw wrapped FT", ftWithdrawHash)
	//
	//ftDepositBackwardHash := scripts.FtDeposit(
	//	cfg.Ctx,
	//	cfg.Client,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	cfg.FtAddressWrapped,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	true,
	//)
	//printExplorerURL(t, "Burn wrapped FT", ftDepositBackwardHash)
	//
	//ftWithdrawBackwardHash := scripts.FtWithdraw(
	//	cfg.Ctx,
	//	cfg.Client,
	//	ftDepositBackwardHash,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	chainID,
	//	cfg.FtAddressOriginal,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	cfg.SignerPrivateKey,
	//	false,
	//	false,
	//)
	//printExplorerURL(t, "Unlock original FT", ftWithdrawBackwardHash)
	//
	//// Native
	//nativeDepositHash := scripts.NativeDeposit(
	//	cfg.Ctx,
	//	cfg.Client,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//)
	//printExplorerURL(t, "Deposited Native", nativeDepositHash)
	//
	//nativeWithdrawHash := scripts.NativeWithdraw(
	//	cfg.Ctx,
	//	cfg.Client,
	//	nativeDepositHash,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	chainID,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	cfg.SignerPrivateKey,
	//)
	//printExplorerURL(t, "Withdraw wrapped Native", nativeWithdrawHash)
	//
	//nativeDepositBackwardHash := scripts.FtDeposit(
	//	cfg.Ctx,
	//	cfg.Client,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	cfg.NativeAddressWrapped,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	true,
	//)
	//printExplorerURL(t, "Burn wrapped Native", ftDepositBackwardHash)
	//
	//nativeWithdrawBackwardHash := scripts.NativeWithdraw(
	//	cfg.Ctx,
	//	cfg.Client,
	//	nativeDepositBackwardHash,
	//	cfg.AccountID,
	//	cfg.AccountID,
	//	chainID,
	//	cfg.Amount,
	//	cfg.BridgeAddress,
	//	cfg.SignerPrivateKey,
	//)
	//printExplorerURL(t, "Unlock Native", nativeWithdrawBackwardHash)
}
