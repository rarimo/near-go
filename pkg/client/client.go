package client

import (
	"context"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/client/block"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/jsonrpc"
	"time"
)

type Client struct {
	RPCClient *jsonrpc.Client
	Log       *logan.Entry
}

func NewClient(rpcUrl string) (client Client, err error) {
	client.RPCClient, err = jsonrpc.NewClient(rpcUrl)
	if err != nil {
		return
	}

	client.Log = logan.New().WithField("entry", "near-rpc-client")

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

	timeoutContext, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	res, err = c.callRpcWithRetry(timeoutContext, method, params)
	if err != nil || res == nil {
		return
	}

	if err = json.Unmarshal(res.Result, result); err != nil {
		return
	}

	return
}

func (c *Client) callRpcWithRetry(ctx context.Context, method string, params interface{}) (res *jsonrpc.Response, err error) {
	try := 1

	for {
		log := c.Log.WithField("try", try)

		res, err = c.RPCClient.CallRPC(ctx, method, params)
		// If JSON-RPC error happens, conveniently set it as err to avoid duplicating code
		// XXX: using plain assignment makes `err != nil` true for some reason
		if res.Error != nil {
			err = res.Error
		}
		if err == nil {
			log.Info("successfully got result from rpc")
			return res, nil
		}

		select {
		case <-ctx.Done():
			return nil, errors.Wrap(err, "context timeout exceeded")
		case <-time.After(time.Second * 5):
			log.WithError(err).WithField("try", try).Info("failed to call rpc storage")
			log.Info("retrying...")
			try++
		}
	}
}
