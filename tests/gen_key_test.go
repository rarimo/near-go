package tests

import (
	"github.com/mr-tron/base58"
	"gitlab.com/rarify-protocol/near-bridge-go/scripts"
	"testing"
)

func TestGenKey(t *testing.T) {
	pbk, prk := scripts.GenKey()
	t.Logf("Public key: %s", base58.Encode(pbk))
	t.Logf("Private key: %s", base58.Encode(prk))
}
