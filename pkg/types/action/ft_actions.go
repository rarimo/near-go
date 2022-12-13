package action

import (
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
	"gitlab.com/rarimo/near-bridge-go/pkg/types/action/base"
)

func NewFtDepositCall(params FtDepositArgs, gas types.Gas) base.Action {
	return base.NewFunctionCall(FtTransferMethod, mustMarshalArgs(params), gas, types.OneYocto)
}

func NewFtWithdrawCall(params FtWithdrawArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(FtWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}
