package action

import "gitlab.com/rarify-protocol/near-bridge-go/pkg/types"

type NativeDepositArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	Amount     types.Balance   `json:"amount,required"`
}

type NativeWithdrawArgs struct {
	Amount     types.Balance   `json:"amount,required"`
	ReceiverID types.AccountID `json:"receiver_id,required"`
	Origin     string          `json:"origin,required"`
	Chain      string          `json:"chain,required"`
	Path       [][32]byte      `json:"path,required"`
	Signatures [][]byte        `json:"signatures,required"`
	RecoveryID byte            `json:"recovery_id,required"`
}
