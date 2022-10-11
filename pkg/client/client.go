package client

import (
	"context"
	"encoding/json"

	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client/block"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/jsonrpc"
)

type Client struct {
	RPCClient *jsonrpc.Client
}

func NewClient(rpcUrl string) (client Client, err error) {
	client.RPCClient, err = jsonrpc.NewClient(rpcUrl)
	if err != nil {
		return
	}

	return
}

func (c *Client) NetworkAddr() string {
	return c.RPCClient.URL
}

func (c *Client) doRPC(ctx context.Context, result interface{}, method string, block block.BlockCharacteristic, params interface{}) (res *jsonrpc.Response, err error) {
	if block != nil {
		if mapv, ok := params.(map[string]interface{}); ok {
			block(mapv)
		}
	}

	res, err = c.RPCClient.CallRPC(ctx, method, params)
	if err != nil {
		return
	}

	// If JSON-RPC error happens, conveniently set it as err to avoid duplicating code
	// XXX: using plain assignment makes `err != nil` true for some reason
	if err := res.Error; err != nil {
		return res, err
	}

	if result != nil {
		if err = json.Unmarshal(res.Result, result); err != nil {
			return
		}
	}

	return
}
