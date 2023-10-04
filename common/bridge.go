package common

import (
	"encoding/json"
	"github.com/rarimo/near-go/constants"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type SignArgs struct {
	Origin     string     `json:"origin"`
	Path       [][32]byte `json:"path"`
	Signature  string     `json:"signature"`
	RecoveryID byte       `json:"recovery_id"`
}

type WithdrawArgs struct {
	SignArgs
	ReceiverID AccountID `json:"receiver_id"`
}

type TransferArgs struct {
	Token      AccountID `json:"token"`
	Sender     AccountID `json:"sender"`
	Receiver   AccountID `json:"receiver"`
	ChainTo    string    `json:"chain_to"`
	IsWrapped  bool      `json:"is_wrapped"`
	BundleData string    `json:"bundle_data,omitempty"`
	BundleSalt string    `json:"bundle_salt,omitempty"`
}

func NewTransferArgs(token string, sender, receiver AccountID, chainTo string, isWrapped bool) *TransferArgs {
	return &TransferArgs{
		Token:     token,
		Sender:    sender,
		Receiver:  receiver,
		ChainTo:   chainTo,
		IsWrapped: isWrapped,
	}
}

func (t *TransferArgs) String() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal transfer args")
	}

	return string(bytes), nil
}

type FtTransferArgs struct {
	ReceiverId AccountID `json:"receiver_id"`
	Amount     Balance   `json:"amount"`
	Msg        string    `json:"msg"` // TransferArgs | FeeDepositArgs
}

func NewFtTransferCall(params FtTransferArgs, gas Gas) Action {
	return NewFunctionCall(constants.ContractFtTransfer, MustMarshalJson(params), gas, constants.OneYocto)
}

type FtWithdrawArgs struct {
	WithdrawArgs
	Token     AccountID `json:"token"`
	Amount    Balance   `json:"amount"`
	IsWrapped bool      `json:"is_wrapped"`
}

func NewFtWithdrawCall(params FtWithdrawArgs, gas Gas, deposit Balance) Action {
	return NewFunctionCall(constants.ContractFtWithdraw, MustMarshalJson(params), gas, deposit)
}

type NftTransferArgs struct {
	ReceiverId AccountID `json:"receiver_id"`
	TokenID    string    `json:"token_id"`
	Msg        string    `json:"msg"` // // TransferArgs | FeeDepositArgs
}

func NewNftTransferCall(params NftTransferArgs, gas Gas) Action {
	return NewFunctionCall(constants.ContractNftTransfer, MustMarshalJson(params), gas, constants.OneYocto)
}

type NftWithdrawArgs struct {
	WithdrawArgs
	Token         AccountID        `json:"token"`
	TokenID       string           `json:"token_id"`
	TokenMetadata *NftMetadataView `json:"token_metadata,omitempty"`
	IsWrapped     bool             `json:"is_wrapped"`
}

func NewNftWithdrawCall(params NftWithdrawArgs, gas Gas, deposit Balance) Action {
	return NewFunctionCall(constants.ContractNftWithdraw, MustMarshalJson(params), gas, deposit)
}

type NativeDepositArgs struct {
	ReceiverId AccountID `json:"receiver_id"`
	Chain      string    `json:"chain"`
}

func NewNativeDepositCall(params NativeDepositArgs, gas Gas, deposit Balance) Action {
	return NewFunctionCall(constants.ContractBridgeNativeDeposit, MustMarshalJson(params), gas, deposit)
}

type NativeWithdrawArgs struct {
	WithdrawArgs
	Amount Balance `json:"amount"`
}

func NewNativeWithdrawCall(params NativeWithdrawArgs, gas Gas, deposit Balance) Action {
	return NewFunctionCall(constants.ContractBridgeNativeWithdraw, MustMarshalJson(params), gas, deposit)
}
