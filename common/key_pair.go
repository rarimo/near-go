package common

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"github.com/mr-tron/base58"
	"github.com/rarimo/near-go/errors"
	"io"
	"strings"
)

type KeyPair struct {
	Type PublicKeyType

	PublicKey  Base58PublicKey
	PrivateKey ed25519.PrivateKey
}

func NewKeyPair(keyType PublicKeyType, pub Base58PublicKey, priv ed25519.PrivateKey) KeyPair {
	return KeyPair{
		Type:       keyType,
		PublicKey:  pub,
		PrivateKey: priv,
	}
}

func (kp *KeyPair) Sign(data []byte) (sig Signature) {
	sigType := ReverseKeyTypeMapping[string(kp.Type)]

	switch sigType {
	//case RawPublicKeyTypeSECP256K1:
	case RawKeyTypeED25519:
		sig = NewSignatureED25519(ed25519.Sign(kp.PrivateKey, data))
	}
	return
}

func (kp *KeyPair) PrivateEncoded() string {
	return fmt.Sprintf("%s:%s", kp.Type, base58.Encode(kp.PrivateKey))
}

func (kp *KeyPair) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	*kp, err = NewBase58KeyPair(s)
	return
}

func GenerateKeyPair(keyType PublicKeyType, rand io.Reader) (kp KeyPair, err error) {
	if _, ok := ReverseKeyTypeMapping[string(keyType)]; !ok {
		return kp, common.ErrInvalidKeyType
	}

	var rawPub PublicKey

	switch keyType {
	case KeyTypeED25519:
		var pub ed25519.PublicKey
		var priv ed25519.PrivateKey

		pub, priv, err = ed25519.GenerateKey(rand)
		if err != nil {
			return
		}

		rawPub, err = WrapRawKey(keyType, pub)
		if err != nil {
			return
		}

		kp = NewKeyPair(keyType, rawPub.ToBase58PublicKey(), priv)
	case KeyTypeSECP256K1:
		// TODO
		return kp, fmt.Errorf("SECP256K1 is not supported yet")
	}

	return
}

func NewBase58KeyPair(raw string) (kp KeyPair, err error) {
	split := strings.SplitN(raw, ":", 2)
	if len(split) != 2 {
		return kp, common.ErrInvalidPrivateKey
	}

	keyTypeRaw := split[0]
	encodedKey := split[1]

	keyType, ok := ReverseKeyTypeMapping[keyTypeRaw]
	if !ok {
		return kp, common.ErrInvalidKeyType
	}

	// TODO
	if keyType == RawKeyTypeSECP256K1 {
		return kp, fmt.Errorf("SECP256K1 is not supported yet")
	}

	decoded, err := base58.Decode(encodedKey)
	if err != nil {
		return kp, fmt.Errorf("failed to decode private key: %w", err)
	}

	if len(decoded) != ed25519.PrivateKeySize {
		return kp, common.ErrInvalidPrivateKey
	}

	var pubKey PublicKey

	theKeyType := KeyTypes[keyType]
	privKey := ed25519.PrivateKey(decoded)
	pubKey, err = WrapRawKey(theKeyType, privKey[32:]) // See ed25519.Public()
	if err != nil {
		println("wraprawkey failed")
		return
	}

	kp = NewKeyPair(theKeyType, pubKey.ToBase58PublicKey(), privKey)

	return
}
