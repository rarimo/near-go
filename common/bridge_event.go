package common

type BridgeEventType string

const (
	EventTypeFTDeposited     BridgeEventType = "ft_deposited"
	EventTypeFTWithdrawn     BridgeEventType = "ft_withdrawn"
	EventTypeNFTDeposited    BridgeEventType = "nft_deposited"
	EventTypeNFTWithdrawn    BridgeEventType = "nft_withdrawn"
	EventTypeNativeDeposited BridgeEventType = "native_deposited"
	EventTypeNativeWithdrawn BridgeEventType = "native_withdrawn"
)

type BridgeBaseEventData struct {
	Sender    AccountID  `json:"sender"`
	Receiver  string     `json:"receiver"`
	Token     *AccountID `json:"token,omitempty"`
	TokenID   *string    `json:"token_id,omitempty"`
	Amount    *string    `json:"amount,omitempty"`
	IsWrapped *bool      `json:"is_wrapped,omitempty"`
}

type BridgeDepositedEventData struct {
	BridgeBaseEventData
	ChainTo    string  `json:"chain_to"`
	BundleData *string `json:"bundle_data,omitempty"`
	BundleSalt *string `json:"bundle_salt,omitempty"`
}

type BridgeWithdrawnEventData struct {
	BridgeBaseEventData
	Origin     string   `json:"origin"`
	Signature  string   `json:"signature"`
	Path       []string `json:"path"`
	RecoveryID uint8    `json:"recovery_id"`
}

type BridgeEvent struct {
	Standard string `json:"standard"`
	Version  string `json:"version"`
	Event    string `json:"event"`
}

type BridgeDepositedEvent struct {
	BridgeEvent
	Data []BridgeDepositedEventData `json:"data"`
}

type BridgeWithdrawnEvent struct {
	BridgeEvent
	Data []BridgeWithdrawnEventData `json:"data"`
}
