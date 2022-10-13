package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func NewNativeDepositCall(params NativeDepositArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(NativeDepositMethod, mustMarshalArgs(params), gas, deposit)
}

func NewNativeWithdrawCall(params NativeWithdrawArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(NativeWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}
