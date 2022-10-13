package action

import "gitlab.com/rarify-protocol/near-bridge-go/pkg/types"

type FtTransferArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	Amount     types.Balance   `json:"amount,required"`
	Msg        TransferArgs    `json:"msg,required"`
}

type FtWithdrawArgs struct {
	Token      types.AccountID `json:"token,required"`
	Amount     types.Balance   `json:"amount,required"`
	ReceiverID types.AccountID `json:"receiver_id,required"`
	IsWrapped  bool            `json:"is_wrapped,required"`
	Chain      string          `json:"chain,required"`
	Origin     string          `json:"origin,required"`
	Path       [][32]byte      `json:"path,required"`
	Signatures [][]byte        `json:"signatures,required"`
	RecoveryID byte            `json:"recovery_id,required"`
}
