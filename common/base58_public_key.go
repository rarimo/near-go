package common

import (
	"encoding/json"
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/rarimo/near-go/errors"
	"strings"
)

type Base58PublicKey struct {
	Type  PublicKeyType
	Value string

	Key PublicKey
}

func NewBase58PublicKey(raw string) (pk Base58PublicKey, err error) {
	split := strings.SplitN(raw, ":", 2)
	if len(split) != 2 {
		return pk, common.ErrInvalidPublicKey
	}

	keyTypeRaw := split[0]
	encodedKey := split[1]

	keyType, ok := ReverseKeyTypeMapping[keyTypeRaw]
	if !ok {
		return pk, common.ErrInvalidKeyType
	}

	decoded, err := base58.Decode(encodedKey)
	if err != nil {
		return pk, fmt.Errorf("failed to decode public key: %w", err)
	}

	pk.Type = KeyTypes[keyType]
	pk.Value = encodedKey

	pk.Key, err = WrapRawKey(pk.Type, decoded)

	return
}

func (pk Base58PublicKey) String() string {
	return fmt.Sprintf("%s:%s", pk.Type, pk.Value)
}

func (pk Base58PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(pk.String())
}

func (pk *Base58PublicKey) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	*pk, err = NewBase58PublicKey(s)
	return
}

// Copies Base58PublicKey to PublicKey
func (pk *Base58PublicKey) ToPublicKey() PublicKey {
	var buf PublicKey
	copy(buf[:], pk.Key[:])
	return buf
}
