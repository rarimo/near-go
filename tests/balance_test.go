package tests

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client/block"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action"
	"testing"
)

func TestFtBalance(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	res := map[string]string{"account_id": "napalmpapalam.testnet"}
	r, _ := json.Marshal(res)

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.FtAddressOriginal, action.FtBalanceOfMethod, base64.StdEncoding.EncodeToString(r), block.FinalityFinal())
	var result string

	json.Unmarshal(x.Result, &result)
	fmt.Println(result)
}

func TestNftBalance(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	res := map[string]string{"account_id": "napalmpapalam.testnet"}
	r, _ := json.Marshal(res)

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.NftAddressOriginal, action.NftTokensForOwnerMethod, base64.StdEncoding.EncodeToString(r), block.FinalityFinal())
	fmt.Println(x)
}
