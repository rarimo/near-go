package common

type QueryResponse struct {
	BlockHeight BlockHeight   `json:"block_height"`
	BlockHash   Hash          `json:"block_hash"`
	Error       *string       `json:"error"`
	Logs        []interface{} `json:"logs"` // TODO: use correct type
}
