package tests

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"

	"github.com/mr-tron/base58"
	"gitlab.com/rarimo/near-bridge-go/scripts"
)

func TestGenKey(t *testing.T) {
	pbk, prk := scripts.GenKey()
	t.Logf("Public key base58: %s", base58.Encode(pbk))
	t.Logf("Private key base58: %s", base58.Encode(prk))
	t.Logf("Public key hex: %s", hexutil.Encode(pbk))
	t.Logf("Private key hex: %s", hexutil.Encode(prk))
}
