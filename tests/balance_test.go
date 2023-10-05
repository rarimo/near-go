//go:build manual_test
// +build manual_test

package tests

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
)

func TestFtBalance(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	res := map[string]string{"account_id": cfg.AccountID}
	r, _ := json.Marshal(res)

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.FtAddress, common.ContractFtBalanceOf, base64.StdEncoding.EncodeToString(r), nearclient.FinalityFinal())
	var result string

	json.Unmarshal(x.Result, &result)
	fmt.Println(result)
}

func TestNftBalance(t *testing.T) {
	cfg := NewConfig(context.Background(), kv.MustFromEnv())

	res := map[string]string{"account_id": cfg.AccountID}
	r, _ := json.Marshal(res)

	x, _ := cfg.Client.ContractViewCallFunction(context.Background(), cfg.NftAddress, common.ContractNftTokensForOwner, base64.StdEncoding.EncodeToString(r), nearclient.FinalityFinal())
	var z []common.NftView
	json.Unmarshal(x.Result, &z)
	fmt.Println(z)
}
