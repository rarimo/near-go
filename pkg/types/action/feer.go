package action

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
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
	TokenAddr *types.AccountID `json:"token_addr"`
	TokenType TokenType        `json:"token_type"`
	Fee       types.Balance    `json:"fee"`
}

type FeeManageOperation struct {
	SignArgs
	Token FeeToken `json:"token,required"`
}

type FeeManageOperationArgs struct {
	Operation FeeManageOperation `json:"op,required"`
}

type FeeTokenWithBorsh struct {
	TokenAddr *types.AccountID `json:"token_addr"`
	TokenType string           `json:"token_type"`
	Fee       types.Balance    `json:"fee"`
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
	FeeTokenAddr *types.AccountID `json:"fee_token_addr"`
	TokenAddr    *types.AccountID `json:"token_addr,omitempty"`
	TokenType    TokenType        `json:"token_type,omitempty"`
	TransferType FeerTransferType `json:"transfer_type"`
	Receiver     string           `json:"receiver"`
	ChainTo      string           `json:"chain_to"`
	IsWrapped    bool             `json:"is_wrapped"`
	BundleData   *string          `json:"bundle_data,omitempty"`
	BundleSalt   *string          `json:"bundle_salt,omitempty"`
}

func (f FeerDepositArgs) String() (string, error) {
	bytes, err := json.Marshal(f)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal feer deposit args")
	}

	return string(bytes), nil
}

func (f FeerDepositArgs) WithTransferType(transferType FeerTransferType) FeerDepositArgs {
	return FeerDepositArgs{
		FeeTokenAddr: f.FeeTokenAddr,
		TokenAddr:    f.TokenAddr,
		TokenType:    f.TokenType,
		TransferType: transferType,
		Receiver:     f.Receiver,
		ChainTo:      f.ChainTo,
		IsWrapped:    f.IsWrapped,
		BundleData:   f.BundleData,
		BundleSalt:   f.BundleSalt,
	}
}

func NewFeeChargeNativeCall(params FeerDepositArgs, amount types.Balance, gas types.Gas) Action {
	return NewFunctionCall(FeerChargeNative, mustMarshalArgs(map[string]FeerDepositArgs{
		"deposit": params,
	}), gas, amount)
}
