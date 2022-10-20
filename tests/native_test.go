package tests

//
//import (
//	"context"
//	"gitlab.com/distributed_lab/kit/kv"
//	"gitlab.com/rarify-protocol/near-bridge-go/scripts"
//	"testing"
//)
//
//func TestNative(t *testing.T) {
//	cfg := NewConfig(context.Background(), kv.MustFromEnv())
//
//	depositHash, depositedEventID := scripts.NativeDeposit(
//		cfg.Ctx,
//		cfg.Client,
//		cfg.AccountID,
//		cfg.AccountID,
//		cfg.Amount,
//		cfg.BridgeAddress,
//	)
//	printExplorerURL(t, "Deposited Native", depositHash, &depositedEventID)
//
//	withdrawHash := scripts.NativeWithdraw(
//		cfg.Ctx,
//		cfg.Client,
//		depositHash,
//		depositedEventID,
//		cfg.AccountID,
//		cfg.AccountID,
//		chainID,
//		chainID,
//		cfg.Amount,
//		cfg.BridgeAddress,
//		cfg.SignerPrivateKey,
//	)
//	printExplorerURL(t, "Withdraw wrapped Native", withdrawHash, nil)
//
//	depositBackwardHash, depositedBackwardEventID := scripts.NativeDeposit(
//		cfg.Ctx,
//		cfg.Client,
//		cfg.AccountID,
//		cfg.AccountID,
//		cfg.Amount,
//		cfg.BridgeAddress,
//	)
//	printExplorerURL(t, "Burn wrapped Native", depositBackwardHash, &depositedBackwardEventID)
//
//	withdrawBackwardHash := scripts.NativeWithdraw(
//		cfg.Ctx,
//		cfg.Client,
//		depositBackwardHash,
//		depositedBackwardEventID,
//		cfg.AccountID,
//		cfg.AccountID,
//		chainID,
//		chainID,
//		cfg.Amount,
//		cfg.BridgeAddress,
//		cfg.SignerPrivateKey,
//	)
//	printExplorerURL(t, "Unlock Native", withdrawBackwardHash, nil)
//}
