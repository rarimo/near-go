package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

type NftTransferArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	TokenID    string          `json:"token_id,required"`
	Msg        TransferArgs    `json:"msg,required"`
}

type NftWithdrawArgs struct {
	Token         types.AccountID `json:"token,required"`
	TokenID       string          `json:"token_id,required"`
	ReceiverID    types.AccountID `json:"receiver_id,required"`
	IsWrapped     bool            `json:"is_wrapped,required"`
	Chain         string          `json:"chain,required"`
	Origin        string          `json:"origin,required"`
	Path          [][32]byte      `json:"path,required"`
	Signatures    [][]byte        `json:"signatures,required"`
	RecoveryID    byte            `json:"recovery_id,required"`
	TokenMetadata *NftMetadata    `json:"token_metadata,omitempty"`
}
