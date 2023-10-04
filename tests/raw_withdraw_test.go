package tests

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarimo/near-bridge-go/pkg/client"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/key"
	"gitlab.com/rarimo/near-bridge-go/scripts"
	"testing"
)

func printExplorerURL(t *testing.T, msg string, hash string, receiptID *string) {
	t.Logf("%s: https://explorer.testnet.near.org/transactions/%s", msg, hash)
	if receiptID != nil {
		t.Logf("%s bridge receipt: https://explorer.testnet.near.org/transactions/%s#%s", msg, hash, *receiptID)
	}
}

func getClient() (context.Context, *client.Client) {
	keyPair, err := key.NewBase58KeyPair("ed25519:privatekey")
	if err != nil {
		panic(errors.Wrap(err, "failed to get key pair"))
	}

	ctx := client.ContextWithKeyPair(context.Background(), keyPair)

	cli, err := client.NewClient("https://rpc.testnet.near.org")
	if err != nil {
		panic(errors.Wrap(err, "failed to create rpc client"))
	}

	return ctx, &cli
}

func TestRawNftWithdraw(t *testing.T) {
	ctx, cli := getClient()

	hash := scripts.RawNFTWithdraw(ctx,
		cli,
		"napalmpapalam.testnet",
		"bridge.rarimo.testnet",
		"nft.rarimo.testnet",
		"1",
		"napalmpapalam.testnet",
		"origin",
		"0x",
		[][32]byte{},
		0,
		false,
	)

	printExplorerURL(t, "Withdraw NFT", hash, nil)
}

func TestRawFtWithdraw(t *testing.T) {
	ctx, cli := getClient()

	hash := scripts.RawFTWithdraw(ctx,
		cli,
		"napalmpapalam.testnet",
		"bridge.rarimo.testnet",
		"ft.rarimo.testnet",
		"10000",
		"napalmpapalam.testnet",
		"ethereum",
		"0x",
		[][32]byte{},
		0,
		false,
	)

	printExplorerURL(t, "Withdraw FT", hash, nil)
}

func TestRawNativeWithdraw(t *testing.T) {
	ctx, cli := getClient()

	hash := scripts.RawNativeWithdraw(ctx,
		cli,
		"napalmpapalam.testnet",
		"bridge.rarimo.testnet",
		"10000",
		"napalmpapalam.testnet",
		"ethereum",
		"0x",
		[][32]byte{},
		0,
	)

	printExplorerURL(t, "Withdraw Native", hash, nil)
}
