package action

import (
	"encoding/json"
	"gitlab.com/rarify-protocol/near-bridge-go/pkg/types"
)

type TransferArgs struct {
	Token      types.AccountID `json:"token,required"`
	Sender     types.AccountID `json:"sender,required"`
	Receiver   types.AccountID `json:"receiver,required"`
	ChainTo    string          `json:"chain_to,required"`
	IsWrapped  bool            `json:"is_wrapped,required"`
	BundleData []byte          `json:"bundle_data,omitempty"`
	BundleSalt [32]byte        `json:"bundle_salt,omitempty"`
}

func NewTransferArgs(token string, sender, receiver types.AccountID, chainTo string, isWrapped bool, bundleData []byte, bundleSalt *[32]byte) string {
	args := TransferArgs{
		Token:     token,
		Sender:    sender,
		Receiver:  receiver,
		ChainTo:   chainTo,
		IsWrapped: isWrapped,
	}

	if bundleData != nil {
		args.BundleData = bundleData
	}
	if bundleSalt != nil {
		args.BundleSalt = *bundleSalt
	}

	result, err := json.Marshal(&args)
	if err != nil {
		panic(err)
	}

	return string(result)
}
