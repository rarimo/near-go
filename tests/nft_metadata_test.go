//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/rarimo/near-go/client"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/constants"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
)

func TestNftMetadata(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	res := map[string]string{"token_id": cfg.TokenID}
	r, _ := json.Marshal(res)

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.NftAddress, constants.ContractNftGet, base64.StdEncoding.EncodeToString(r), nearclient.FinalityFinal())
	var z *common.NftView
	json.Unmarshal(x.Result, &z)
	fmt.Println(string(x.Result))
}
