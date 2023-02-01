package action

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func mustMarshalArgs(params interface{}) []byte {
	args, err := json.Marshal(params)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal args"))
	}
	return args
}
