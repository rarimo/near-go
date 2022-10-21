package action

import "gitlab.com/rarify-protocol/near-bridge-go/pkg/types"

type NativeDepositArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	Chain      string          `json:"chain,required"`
}

type NativeWithdrawArgs struct {
	WithdrawArgs
	Amount types.Balance `json:"amount,required"`
}
