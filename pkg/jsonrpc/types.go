package jsonrpc

import "encoding/json"

type EndpointSetup struct {
	JsonRpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
}

type Request struct {
	EndpointSetup
	Params interface{} `json:"params,omitempty"`
}

type Response struct {
	EndpointSetup
	Error  *Error          `json:"error"`
	Result json.RawMessage `json:"result"`
}
