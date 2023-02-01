package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/rarimo/near-bridge-go/pkg/client/block"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action"
)

func TestNftContractMetadata(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.NftAddressOriginal, action.NftContractMetadataMethod, "", block.FinalityFinal())
	var z types.NftContractMetadataView
	json.Unmarshal(x.Result, &z)
	fmt.Println(string(x.Result))
}
