package nearclient

import (
	"context"
	"github.com/rarimo/near-go/common"
)

type rpcContext int

const (
	clientCtx = rpcContext(iota)
	keyPairCtx
)

func ContextWithKeyPair(ctx context.Context, keyPair common.KeyPair) context.Context {
	kp := keyPair
	return context.WithValue(ctx, keyPairCtx, &kp)
}

func getKeyPair(ctx context.Context) *common.KeyPair {
	v, ok := ctx.Value(keyPairCtx).(*common.KeyPair)
	if ok {
		return v
	}

	return nil
}
