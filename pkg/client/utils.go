package client

import "gitlab.com/rarimo/near-bridge-go/pkg/client/block"

// HACK
func blockIDArrayParams(block block.BlockCharacteristic) []interface{} {
	params := []interface{}{nil}

	if block == nil {
		return params
	}
	p := map[string]interface{}{}

	block(p)
	if v, ok := p["block_id"]; ok {
		params[0] = v
	}

	return params
}
