package tests

import (
	"context"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/key"
)

type Config struct {
	RpcURL             string `fig:"rpc_url,required"`
	PrivateKey         string `fig:"private_key,required"`
	SignerPrivateKey   string `fig:"signer_private_key,required"`
	AccountID          string `fig:"account_id,required"`
	BridgeAddress      string `fig:"bridge_address,required"`
	FtAddressOriginal  string `fig:"ft_address_original,required"`
	FtAddressWrapped   string `fig:"ft_address_wrapped,required"`
	NftAddressOriginal string `fig:"nft_address_original,required"`
	NftAddressWrapped  string `fig:"nft_address_wrapped,required"`
	TokenID            string `fig:"token_id"`
	Amount             string `fig:"amount"`
}

func NewConfig(ctx context.Context, getter kv.Getter) (Config, context.Context, client.Client) {
	var cfg Config

	err := figure.
		Out(&cfg).
		From(kv.MustGetStringMap(getter, "tests")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure config"))
	}

	keyPair, err := key.NewBase58KeyPair(cfg.PrivateKey)
	if err != nil {
		panic(errors.Wrap(err, "failed to get key pair"))
	}

	newCtx := client.ContextWithKeyPair(ctx, keyPair)
	rpc, err := client.NewClient(cfg.RpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create rpc client"))
	}
	return cfg, newCtx, rpc
}
