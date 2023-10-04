package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JsonRpcError struct {
	Name  string            `json:"name"`
	Cause JsonRpcErrorCause `json:"cause"`

	// Legacy - do not rely on them
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (err *JsonRpcError) Error() string {
	return fmt.Sprintf("RPC error %s (%s)", err.Name, err.Cause.String())
}

type JsonRpcErrorCause struct {
	Name string          `json:"name"`
	Info json.RawMessage `json:"info"`

	message *JsonRpcErrorCauseMessage
}

type JsonRpcErrorCauseMessage struct {
	ErrorMessage string `json:"error_message"`
}

func (cause *JsonRpcErrorCause) UnmarshalJSON(b []byte) (err error) {
	var data struct {
		Name string          `json:"name"`
		Info json.RawMessage `json:"info"`
	}

	if err = json.Unmarshal(b, &data); err != nil {
		err = fmt.Errorf(string(b))
		return
	}

	var info map[string]interface{}
	if err = json.Unmarshal(data.Info, &info); err != nil {
		err = fmt.Errorf(string(b))
		return
	}

	var message *JsonRpcErrorCauseMessage
	if v, ok := info["error_message"]; ok {
		message = &JsonRpcErrorCauseMessage{
			ErrorMessage: v.(string),
		}
	}

	*cause = JsonRpcErrorCause{
		Name:    data.Name,
		Info:    data.Info,
		message: message,
	}

	return
}

func (cause JsonRpcErrorCause) String() string {
	if cause.message != nil {
		return fmt.Sprintf("name=%s, message=%s", cause.Name, cause.message.ErrorMessage)
	}
	return fmt.Sprintf("name=%s, info=%s", cause.Name, strconv.Quote(string(cause.Info)))
}
