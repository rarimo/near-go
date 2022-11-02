package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

type NftDepositArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	TokenID    string          `json:"token_id,required"`
	Msg        string          `json:"msg,required"` // TransferArgs
}

type NftWithdrawArgs struct {
	WithdrawArgs
	Token         types.AccountID        `json:"token,required"`
	TokenID       string                 `json:"token_id,required"`
	TokenMetadata *types.NftMetadataView `json:"token_metadata,omitempty"`
	IsWrapped     bool                   `json:"is_wrapped,required"`
}
