package main

import (
	"context"
	"fmt"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client/block"
)

func main() {
	rpc, err := client.NewClient("https://rpc.mainnet.near.org")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	rpc.TransactionSendAwait()

	res, err := rpc.BlockDetails(ctx, block.FinalityFinal())
	if err != nil {
		panic(err)
	}

	fmt.Println("latest block: ", res.Header.Hash)
}
