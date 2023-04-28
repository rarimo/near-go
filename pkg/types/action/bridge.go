package action

import (
	"encoding/json"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
)

type SignArgs struct {
	Origin     string     `json:"origin,required"`
	Path       [][32]byte `json:"path,required"`
	Signature  string     `json:"signature,required"`
	RecoveryID byte       `json:"recovery_id,required"`
}

type WithdrawArgs struct {
	SignArgs
	ReceiverID types.AccountID `json:"receiver_id,required"`
	Chain      string          `json:"chain,required"`
}

type TransferArgs struct {
	Token      types.AccountID `json:"token,required"`
	Sender     types.AccountID `json:"sender,required"`
	Receiver   types.AccountID `json:"receiver,required"`
	ChainTo    string          `json:"chain_to,required"`
	IsWrapped  bool            `json:"is_wrapped,required"`
	BundleData string          `json:"bundle_data,omitempty"`
	BundleSalt string          `json:"bundle_salt,omitempty"`
}

func NewTransferArgs(token string, sender, receiver types.AccountID, chainTo string, isWrapped bool) string {
	args := TransferArgs{
		Token:     token,
		Sender:    sender,
		Receiver:  receiver,
		ChainTo:   chainTo,
		IsWrapped: isWrapped,
	}

	result, err := json.Marshal(&args)
	if err != nil {
		panic(err)
	}

	return string(result)
}

type FtTransferArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	Amount     types.Balance   `json:"amount,required"`
	Msg        string          `json:"msg,required"` // TransferArgs | FeeDepositArgs
}

func NewFtTransferCall(params FtTransferArgs, gas types.Gas) Action {
	return NewFunctionCall(FtTransferMethod, mustMarshalArgs(params), gas, types.OneYocto)
}

type FtWithdrawArgs struct {
	WithdrawArgs
	Token     types.AccountID `json:"token,required"`
	Amount    types.Balance   `json:"amount,required"`
	IsWrapped bool            `json:"is_wrapped,required"`
}

func NewFtWithdrawCall(params FtWithdrawArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(FtWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}

type NftTransferArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	TokenID    string          `json:"token_id,required"`
	Msg        string          `json:"msg,required"` // // TransferArgs | FeeDepositArgs
}

func NewNftTransferCall(params NftTransferArgs, gas types.Gas) Action {
	return NewFunctionCall(NftTransferMethod, mustMarshalArgs(params), gas, types.OneYocto)
}

type NftWithdrawArgs struct {
	WithdrawArgs
	Token         types.AccountID        `json:"token,required"`
	TokenID       string                 `json:"token_id,required"`
	TokenMetadata *types.NftMetadataView `json:"token_metadata,omitempty"`
	IsWrapped     bool                   `json:"is_wrapped,required"`
}

func NewNftWithdrawCall(params NftWithdrawArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(NftWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}

type NativeDepositArgs struct {
	ReceiverId types.AccountID `json:"receiver_id,required"`
	Chain      string          `json:"chain,required"`
}

func NewNativeDepositCall(params NativeDepositArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(BridgeNativeDepositMethod, mustMarshalArgs(params), gas, deposit)
}

type NativeWithdrawArgs struct {
	WithdrawArgs
	Amount types.Balance `json:"amount,required"`
}

func NewNativeWithdrawCall(params NativeWithdrawArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(BridgeNativeWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}
