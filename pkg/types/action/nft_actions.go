package action

import (
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types/action/base"
)

func NewNftTransferCall(params NftTransferArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(NftTransferMethod, mustMarshalArgs(params), gas, deposit)
}

func NewNftWithdrawCall(params NftWithdrawArgs, gas types.Gas, deposit types.Balance) base.Action {
	return base.NewFunctionCall(NftWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}
