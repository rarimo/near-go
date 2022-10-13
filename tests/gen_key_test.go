package tests

import (
	"gitlab.com/rarify-protocol/near-bridge-go/scripts"
	"testing"
)

func TestGenKey(t *testing.T) {
	pbk, prk := scripts.GenKey()
	t.Logf("Public key: %x", pbk)
	t.Logf("Private key: %x", prk)
}
