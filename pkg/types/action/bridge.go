package action

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarimo/near-bridge-go/pkg/types"
)

type SignArgs struct {
	Origin     string     `json:"origin"`
	Path       [][32]byte `json:"path"`
	Signature  string     `json:"signature"`
	RecoveryID byte       `json:"recovery_id"`
}

type WithdrawArgs struct {
	SignArgs
	ReceiverID types.AccountID `json:"receiver_id"`
	Chain      string          `json:"chain"`
}

type TransferArgs struct {
	Token      types.AccountID `json:"token"`
	Sender     types.AccountID `json:"sender"`
	Receiver   types.AccountID `json:"receiver"`
	ChainTo    string          `json:"chain_to"`
	IsWrapped  bool            `json:"is_wrapped"`
	BundleData string          `json:"bundle_data,omitempty"`
	BundleSalt string          `json:"bundle_salt,omitempty"`
}

func (t *TransferArgs) String() (string, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal transfer args")
	}

	return string(bytes), nil
}

func NewTransferArgs(token string, sender, receiver types.AccountID, chainTo string, isWrapped bool) *TransferArgs {
	return &TransferArgs{
		Token:     token,
		Sender:    sender,
		Receiver:  receiver,
		ChainTo:   chainTo,
		IsWrapped: isWrapped,
	}
}

type FtTransferArgs struct {
	ReceiverId types.AccountID `json:"receiver_id"`
	Amount     types.Balance   `json:"amount"`
	Msg        string          `json:"msg"` // TransferArgs | FeeDepositArgs
}

func NewFtTransferCall(params FtTransferArgs, gas types.Gas) Action {
	return NewFunctionCall(FtTransferMethod, mustMarshalArgs(params), gas, types.OneYocto)
}

type FtWithdrawArgs struct {
	WithdrawArgs
	Token     types.AccountID `json:"token"`
	Amount    types.Balance   `json:"amount"`
	IsWrapped bool            `json:"is_wrapped"`
}

func NewFtWithdrawCall(params FtWithdrawArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(FtWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}

type NftTransferArgs struct {
	ReceiverId types.AccountID `json:"receiver_id"`
	TokenID    string          `json:"token_id"`
	Msg        string          `json:"msg"` // // TransferArgs | FeeDepositArgs
}

func NewNftTransferCall(params NftTransferArgs, gas types.Gas) Action {
	return NewFunctionCall(NftTransferMethod, mustMarshalArgs(params), gas, types.OneYocto)
}

type NftWithdrawArgs struct {
	WithdrawArgs
	Token         types.AccountID        `json:"token"`
	TokenID       string                 `json:"token_id"`
	TokenMetadata *types.NftMetadataView `json:"token_metadata,omitempty"`
	IsWrapped     bool                   `json:"is_wrapped"`
}

func NewNftWithdrawCall(params NftWithdrawArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(NftWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}

type NativeDepositArgs struct {
	ReceiverId types.AccountID `json:"receiver_id"`
	Chain      string          `json:"chain"`
}

func NewNativeDepositCall(params NativeDepositArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(BridgeNativeDepositMethod, mustMarshalArgs(params), gas, deposit)
}

type NativeWithdrawArgs struct {
	WithdrawArgs
	Amount types.Balance `json:"amount"`
}

func NewNativeWithdrawCall(params NativeWithdrawArgs, gas types.Gas, deposit types.Balance) Action {
	return NewFunctionCall(BridgeNativeWithdrawMethod, mustMarshalArgs(params), gas, deposit)
}
