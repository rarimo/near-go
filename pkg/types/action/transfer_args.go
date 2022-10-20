package action

import "gitlab.com/rarify-protocol/near-bridge-go/pkg/types"

type WithdrawArgs struct {
	ReceiverID types.AccountID `json:"receiver_id,required"`
	Chain      string          `json:"chain,required"`
	Origin     string          `json:"origin,required"`
	Path       [][32]byte      `json:"path,required"`
	Signatures []string        `json:"signatures,required"`
	RecoveryID byte            `json:"recovery_id,required"`
}
