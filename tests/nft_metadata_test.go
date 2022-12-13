package tests

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/pkg/client/block"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func TestNftMetadata(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	res := map[string]string{"token_id": cfg.TokenID}
	r, _ := json.Marshal(res)

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.NftAddressOriginal, action.NftGetMethod, base64.StdEncoding.EncodeToString(r), block.FinalityFinal())
	var z *types.NftView
	json.Unmarshal(x.Result, &z)
	fmt.Println(string(x.Result))
}
