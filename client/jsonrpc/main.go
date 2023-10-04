package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/rarimo/near-go/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"net/http"
	"net/url"
	"sync/atomic"
)

const RpcVersion = "2.0"

type JsonRpcClient struct {
	URL       string
	client    *http.Client
	nextReqId uint64
}

var RequestTimeoutError = errors.New("request timeout")

func New(rpcUrl string) (*JsonRpcClient, error) {
	_, err := url.Parse(rpcUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse rpc url")
	}

	client := JsonRpcClient{
		URL:    rpcUrl,
		client: new(http.Client),
	}

	atomic.StoreUint64(&client.nextReqId, 0)
	return &client, nil
}

func (c *JsonRpcClient) nextId() uint64 {
	return atomic.AddUint64(&c.nextReqId, 1)
}

func (c *JsonRpcClient) CallRPC(ctx context.Context, method string, params interface{}) (*common.Response, error) {
	reqId := fmt.Sprintf("%d", c.nextId())
	body, err := json.Marshal(common.Request{
		common.EndpointSetup{RpcVersion, reqId, method},
		params,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if response.StatusCode == http.StatusRequestTimeout {
		return nil, RequestTimeoutError
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	return parseRPCBody(response)
}

func (c *JsonRpcClient) CallRPCRaw(ctx context.Context, method string, params interface{}) ([]byte, error) {
	reqId := fmt.Sprintf("%d", c.nextId())
	body, err := json.Marshal(common.Request{
		common.EndpointSetup{RpcVersion, reqId, method},
		params,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if response.StatusCode == http.StatusRequestTimeout {
		return nil, RequestTimeoutError
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return content, nil
}

func parseRPCBody(r *http.Response) (*common.Response, error) {
	body := r.Body
	if body == nil {
		return nil, errors.New("nil body")
	}
	defer body.Close()

	if r.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid content type")
	}

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	var res common.Response
	if err := decoder.Decode(&res); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return &res, nil
}
