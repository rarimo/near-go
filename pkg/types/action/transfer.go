package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

type TransferArgs struct {
	Token     types.AccountID `json:"token"`
	Receiver  types.AccountID `json:"receiver"`
	Chain     string          `json:"chain"`
	IsWrapped bool            `json:"is_wrapped"`
}
