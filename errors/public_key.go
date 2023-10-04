package common

import "errors"

var (
	ErrInvalidPublicKey  = errors.New("invalid public key")
	ErrInvalidPrivateKey = errors.New("invalid private key")
	ErrInvalidKeyType    = errors.New("invalid key type")
)
