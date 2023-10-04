package tests

import (
	"context"
	nearclient2 "github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/common"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Config struct {
	Ctx                  context.Context    `fig:"-"` // context with key pair
	Client               nearclient2.Client `fig:"-"`
	RpcURL               string             `fig:"rpc_url,required"`
	PrivateKey           string             `fig:"private_key,required"`
	SignerPrivateKey     string             `fig:"signer_private_key,required"`
	TokenID              string             `fig:"token_id"`
	AccountID            common.AccountID   `fig:"account_id,required"`
	BridgeAddress        common.AccountID   `fig:"bridge_address,required"`
	FtAddress            common.AccountID   `fig:"ft_address_original,required"`
	FtAddressWrapped     common.AccountID   `fig:"ft_address_wrapped,required"`
	NftAddress           common.AccountID   `fig:"nft_address_original,required"`
	NftAddressWrapped    common.AccountID   `fig:"nft_address_wrapped,required"`
	NativeAddressWrapped common.AccountID   `fig:"native_address_wrapped,required"`
	FeerAddress          common.AccountID   `fig:"feer_address,required"`
	FeeTokenAddress      common.AccountID   `fig:"fee_token_address,required"`
	Amount               string             `fig:"amount"`
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

	keyPair, err := common.NewBase58KeyPair(cfg.PrivateKey)
	if err != nil {
		panic(errors.Wrap(err, "failed to get key pair"))
	}

	cfg.Ctx = nearclient2.ContextWithKeyPair(ctx, keyPair)
	cfg.Client, err = nearclient2.NewClient(cfg.RpcURL)
	if err != nil {
		panic(errors.Wrap(err, "failed to create rpc client"))
	}
	return cfg
}
