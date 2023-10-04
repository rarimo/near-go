package common

import "errors"

var (
	ErrInvalidSignatureType = errors.New("invalid signature type")
	ErrInvalidSignature     = errors.New("invalid signature")
)
