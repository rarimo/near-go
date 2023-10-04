package common

import (
	"encoding/json"
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/rarimo/near-go/constants"
	"github.com/rarimo/near-go/errors"
	"strings"
)

type Base58Signature struct {
	Type  SignatureType
	Value string

	//sig Signature
}

func NewBase58Signature(raw string) (pk Base58Signature, err error) {
	split := strings.SplitN(raw, ":", 2)
	if len(split) != 2 {
		return pk, common.ErrInvalidSignature
	}

	sigTypeRaw := split[0]
	encodedSig := split[1]

	sigType, ok := constants.ReverseSignatureMapping[sigTypeRaw]
	if !ok {
		return pk, common.ErrInvalidSignatureType
	}

	decoded, err := base58.Decode(encodedSig)
	if err != nil {
		return pk, fmt.Errorf("failed to decode signature: %w", err)
	}

	pk.Type = constants.SignatureTypes[sigType]
	pk.Value = encodedSig

	// TODO
	_ = decoded

	return
}

func (sig Base58Signature) String() string {
	return fmt.Sprintf("%s:%s", sig.Type, sig.Value)
}

func (sig Base58Signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(sig.String())
}

func (sig *Base58Signature) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	*sig, err = NewBase58Signature(s)
	return
}
