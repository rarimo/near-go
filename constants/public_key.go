package constants

import "github.com/rarimo/near-go/common"

var KeyTypes = map[byte]common.PublicKeyType{
	common.RawKeyTypeED25519:   common.KeyTypeED25519,
	common.RawKeyTypeSECP256K1: common.KeyTypeSECP256K1,
}

var ReverseKeyTypeMapping = map[string]byte{
	string(common.KeyTypeED25519):   common.RawKeyTypeED25519,
	string(common.KeyTypeSECP256K1): common.RawKeyTypeSECP256K1,
}
