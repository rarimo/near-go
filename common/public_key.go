package common

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/rarimo/near-go/errors"
)

type PublicKeyType string

const (
	KeyTypeED25519   PublicKeyType = "ed25519"
	KeyTypeSECP256K1 PublicKeyType = "secp256k1"
)

const (
	RawKeyTypeED25519 byte = iota
	RawKeyTypeSECP256K1
)

var KeyTypes = map[byte]PublicKeyType{
	RawKeyTypeED25519:   KeyTypeED25519,
	RawKeyTypeSECP256K1: KeyTypeSECP256K1,
}

var ReverseKeyTypeMapping = map[string]byte{
	string(KeyTypeED25519):   RawKeyTypeED25519,
	string(KeyTypeSECP256K1): RawKeyTypeSECP256K1,
}

// TODO: SECP256K1
type PublicKey [33]byte

func (p PublicKey) Hash() string {
	return hex.EncodeToString(p[1:])
}

func (p PublicKey) TypeByte() byte {
	return p[0]
}

func (p PublicKey) Value() []byte {
	return p[1:]
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}

func (p PublicKey) String() string {
	return fmt.Sprintf("%s:%s", KeyTypes[p.TypeByte()], base58.Encode(p.Value()))
}

func (p *PublicKey) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	dec, err := base58.Decode(s)
	if err != nil {
		return err
	}

	*p = PublicKey{}
	copy(p[:], dec)
	return nil
}

func (p *PublicKey) Verify(data []byte, signature Signature) (ok bool, err error) {
	keyType := p.TypeByte()
	if signature.Type() != keyType {
		return false, fmt.Errorf("cannot verify signature type %d with key type %d", signature.Type(), p.TypeByte())
	}

	switch keyType {
	case RawKeyTypeED25519:
		ok = ed25519.Verify(ed25519.PublicKey(p.Value()), data, signature.Value())
	case RawKeyTypeSECP256K1:
		// TODO!
		return false, fmt.Errorf("SECP256K1 is not supported yet")
	}

	return
}

func (p *PublicKey) ToBase58PublicKey() Base58PublicKey {
	return Base58PublicKey{
		Type:  KeyTypes[p[0]],
		Value: base58.Encode(p[1:]),
		Key:   *p,
	}
}

func PublicKeyFromBytes(b []byte) (pk PublicKey, err error) {
	f := b[0]
	l := len(b) - 1
	switch f {
	case RawKeyTypeED25519:
		if l != ed25519.PublicKeySize {
			return pk, common.ErrInvalidPublicKey
		}
		copy(pk[:], b)
		return
	case RawKeyTypeSECP256K1:
		// TODO!
		return pk, fmt.Errorf("SECP256K1 is not supported yet")
	}

	return pk, common.ErrInvalidKeyType
}

func WrapRawKey(keyType PublicKeyType, key []byte) (pk PublicKey, err error) {
	switch keyType {
	case KeyTypeED25519:
		if len(key) != ed25519.PublicKeySize {
			return pk, common.ErrInvalidPublicKey
		}

		pk[0] = RawKeyTypeED25519
		copy(pk[1:], key[0:ed25519.PublicKeySize])
		return
	case KeyTypeSECP256K1:
		// TODO!
		return pk, fmt.Errorf("SECP256K1 is not supported yet")
	}

	return pk, common.ErrInvalidKeyType
}

func WrapED25519(key ed25519.PublicKey) PublicKey {
	if pk, err := WrapRawKey(KeyTypeED25519, key); err != nil {
		panic(err)
	} else {
		return pk
	}
}
