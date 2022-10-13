package scripts

import (
	"context"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

func GetGasPrice(ctx context.Context, cli client.Client) (gasPrice uint64) {
	gas, err := cli.GasPriceView(ctx, nil)
	if err != nil {
		panic(err)
	}

	return types.YoctoToNEAR(gas.GasPrice)
}
