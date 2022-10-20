package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

type FtDepositArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	Amount     types.Balance   `json:"amount,required"`
	Msg        string          `json:"msg,required"` // TransferArgs
}

type FtWithdrawArgs struct {
	WithdrawArgs
	Token     types.AccountID `json:"token,required"`
	Amount    types.Balance   `json:"amount,required"`
	IsWrapped bool            `json:"is_wrapped,required"`
}
