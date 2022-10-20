package tests

import (
	"context"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/key"
)

type Config struct {
	Ctx                  context.Context `fig:"-"` // context with key pair
	Client               client.Client   `fig:"-"`
	RpcURL               string          `fig:"rpc_url,required"`
	PrivateKey           string          `fig:"private_key,required"`
	SignerPrivateKey     string          `fig:"signer_private_key,required"`
	TokenID              string          `fig:"token_id"`
	AccountID            types.AccountID `fig:"account_id,required"`
	BridgeAddress        types.AccountID `fig:"bridge_address,required"`
	FtAddressOriginal    types.AccountID `fig:"ft_address_original,required"`
	FtAddressWrapped     types.AccountID `fig:"ft_address_wrapped,required"`
	NftAddressOriginal   types.AccountID `fig:"nft_address_original,required"`
	NftAddressWrapped    types.AccountID `fig:"nft_address_wrapped,required"`
	NativeAddressWrapped types.AccountID `fig:"native_address_wrapped,required"`
	Amount               string          `fig:"amount"`
}

func NewConfig(ctx context.Context, getter kv.Getter) Config {
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

	cfg.Ctx = client.ContextWithKeyPair(ctx, keyPair)
	cfg.Client, err = client.NewClient(cfg.RpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create rpc client"))
	}
	return cfg
}
