package common

type BridgeEventType string

const (
	FTEventType     BridgeEventType = "ft_deposited"
	NFTEventType    BridgeEventType = "nft_deposited"
	NativeEventType BridgeEventType = "native_deposited"
)

type BridgeEventData struct {
	Sender     AccountID  `json:"sender"`
	Receiver   string     `json:"receiver"`
	ChainTo    string     `json:"chain_to"`
	Token      *AccountID `json:"token,omitempty"`
	TokenID    *string    `json:"token_id,omitempty"`
	Amount     *string    `json:"amount,omitempty"`
	IsWrapped  *bool      `json:"is_wrapped,omitempty"`
	BundleData *string    `json:"bundle_data,omitempty"`
	BundleSalt *string    `json:"bundle_salt,omitempty"`
}

type BridgeEvent struct {
	Standard string            `json:"standard"`
	Version  string            `json:"version"`
	Event    string            `json:"event"`
	Data     []BridgeEventData `json:"data"`
}
