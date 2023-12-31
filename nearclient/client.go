package nearclient

import (
	"context"
	"encoding/json"
	"github.com/rarimo/near-go/common"
	"github.com/rarimo/near-go/nearclient/jsonrpc"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Client struct {
	RPCClient *jsonrpc.JsonRpcClient
	Log       *logan.Entry
}

func New(rpcUrl string) (*Client, error) {
	cli, err := jsonrpc.New(rpcUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		RPCClient: cli,
		Log:       logan.New().WithField("entry", "near-rpc-client"),
	}, nil
}

func (c *Client) NetworkAddr() string {
	return c.RPCClient.URL
}

func (c *Client) doRPC(ctx context.Context, result interface{}, method string, block BlockCharacteristic, params interface{}) (res *common.Response, err error) {
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

func (c *Client) callRpcWithRetry(ctx context.Context, method string, params interface{}) (res *common.Response, err error) {
	try := 1

	for {
		log := c.Log.WithField("try", try)

		res, err = c.RPCClient.CallRPC(ctx, method, params)
		// If JSON-RPC error happens, conveniently set it as err to avoid duplicating code
		// XXX: using plain assignment makes `err != nil` true for some reason
		if res != nil && res.Error != nil {
			err = res.Error
		}
		if errors.Cause(err) != jsonrpc.RequestTimeoutError {
			return res, err
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
