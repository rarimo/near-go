package common

import "crypto/ed25519"

type SignatureType string

const (
	RawSignatureTypeED25519 byte = iota
	RawSignatureTypeSECP256K1
)

const (
	SignatureTypeED25519   = SignatureType("ed25519")
	SignatureTypeSECP256K1 = SignatureType("secp256k1")
)

// TODO: SECP256K1 support
type Signature [1 + ed25519.SignatureSize]byte

func NewSignatureED25519(data []byte) Signature {
	var buf Signature
	buf[0] = RawSignatureTypeED25519
	copy(buf[1:], data[0:ed25519.SignatureSize])
	return buf
}

func (s Signature) Type() byte {
	return s[0]
}

func (s Signature) Value() []byte {
	return s[1:]
}
