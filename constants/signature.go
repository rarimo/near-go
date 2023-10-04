package constants

import "github.com/rarimo/near-go/common"

var SignatureTypes = map[byte]common.SignatureType{
	common.RawSignatureTypeED25519:   common.SignatureTypeED25519,
	common.RawSignatureTypeSECP256K1: common.SignatureTypeSECP256K1,
}

var ReverseSignatureMapping = map[string]byte{
	string(common.SignatureTypeED25519):   common.RawSignatureTypeED25519,
	string(common.SignatureTypeSECP256K1): common.RawSignatureTypeSECP256K1,
}
