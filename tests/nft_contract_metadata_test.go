//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
)

func TestNftContractMetadata(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.NftAddress, common.ContractNftMetadata, "", nearclient.FinalityFinal())
	var z common.NftContractMetadataView
	json.Unmarshal(x.Result, &z)
	fmt.Println(string(x.Result))
}
