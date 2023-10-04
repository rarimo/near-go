package constants

import (
	"github.com/rarimo/near-go/common"
	"lukechampine.com/uint128"
	"math"
	"math/big"
)

var (
	TenPower24 = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
	ZeroNEAR   = common.Balance(uint128.From64(0))
	OneYocto   = common.Balance(uint128.From64(1))
	One        = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(24), nil)
)
