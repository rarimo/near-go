package action

import (
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
)

type FeeManageOperationType int

const (
	FeeAddFeeToken FeeManageOperationType = iota + 1
	FeeRemoveFeeToken
	FeeUpdateFeeToken
	FeeWithdraw
)

type TokenType = string

const (
	TokenType_Native TokenType = "Native"
	TokenType_NFT              = "NFT"
	TokenType_FT               = "FT"
)

type FeerTransferType string

const (
	FeerTransferType_Fee     FeerTransferType = "Fee"
	FeerTransferType_Deposit FeerTransferType = "Deposit"
)

type FeeToken struct {
	TokenAddr *types.AccountID `json:"token_addr,required"`
	TokenType TokenType        `json:"token_type,required"`
	Fee       types.Balance    `json:"fee,required"`
}

type FeeManageOperation struct {
	SignArgs
	Token FeeToken `json:"token,required"`
}

type FeeManageOperationArgs struct {
	Operation FeeManageOperation `json:"op,required"`
}

type FeeTokenWithBorsh struct {
	TokenAddr *types.AccountID `json:"token_addr,required"`
	TokenType string           `json:"token_type,required"`
	Fee       types.Balance    `json:"fee,required"`
}

type FeeManageOperationWithBorsh struct {
	SignArgs
	Token FeeTokenWithBorsh `json:"token,required"`
}

type FeeManageOperationArgsWithBorsh struct {
	Operation FeeManageOperationWithBorsh `json:"op,required"`
}

func NewFeeTokenAddCall(params FeeManageOperationArgs, gas types.Gas) Action {
	return NewFunctionCall(FeerAddFeeToken, mustMarshalArgs(params), gas, types.ZeroNEAR)
}

func NewFeeTokenUpdateCall(params FeeManageOperationArgs, gas types.Gas) Action {
	return NewFunctionCall(FeerUpdateFeeToken, mustMarshalArgs(params), gas, types.ZeroNEAR)
}

func NewFeeTokenRemoveCall(params FeeManageOperationArgs, gas types.Gas) Action {
	return NewFunctionCall(FeerRemoveFeeToken, mustMarshalArgs(params), gas, types.ZeroNEAR)
}

func NewFeeTokenWithdrawCall(params FeeManageOperationArgs, receiver types.AccountID, amount types.Balance, gas types.Gas) Action {
	opts := map[string]interface{}{
		"op":       params.Operation,
		"receiver": receiver,
		"amount":   amount,
	}
	return NewFunctionCall(FeerWithdraw, mustMarshalArgs(opts), gas, types.ZeroNEAR)
}

type FeerDepositArgs struct {
	FeeTokenAddr *types.AccountID `json:"fee_token_addr,required"`
	TokenAddr    *types.AccountID `json:"token_addr,omitempty"`
	TokenType    TokenType        `json:"token_type,omitempty"`
	TransferType FeerTransferType `json:"transfer_type,required"`
	Receiver     string           `json:"receiver,required"`
	ChainTo      string           `json:"chain_to,required"`
	IsWrapped    bool             `json:"is_wrapped,required"`
	BundleData   *string          `json:"bundle_data,omitempty"`
	BundleSalt   *string          `json:"bundle_salt,omitempty"`
}

func NewFeeChargeNativeCall(params FeerDepositArgs, amount types.Balance, gas types.Gas) Action {
	return NewFunctionCall(FeerChargeNative, mustMarshalArgs(map[string]FeerDepositArgs{
		"deposit": params,
	}), gas, amount)
}
