package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func NewFtTransferCall(params FtTransferArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(FtTransferMethod, mustMarshalArgs(params), gas, deposit)
}

func NewFtWithdrawCall(params FtWithdrawArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(FtWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}
