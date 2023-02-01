package types

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"lukechampine.com/uint128"
)

var (
	tenPower24 = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
	ZeroNEAR   = Balance(uint128.From64(0))
	OneYocto   = Balance(uint128.From64(1))
)

// Balance holds amount of yoctoNEAR
type Balance uint128.Uint128

func (bal *Balance) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	val := big.Int{}
	if _, ok := val.SetString(s, 10); !ok {
		return fmt.Errorf("unable to parse '%s'", s)
	}

	*bal = Balance(uint128.FromBig(&val))

	return nil
}

func (bal Balance) MarshalJSON() ([]byte, error) {
	return json.Marshal(bal.String())
}

func (bal Balance) String() string {
	return uint128.Uint128(bal).String()
}

func (bal Balance) Uint64() uint64 {
	return uint128.Uint128(bal).Big().Uint64()
}

func (bal Balance) Div64(div uint64) Balance {
	return Balance(uint128.Uint128(bal).Div64(div))
}

func (bal Balance) Equals(other Balance) bool {
	return uint128.Uint128(bal).Equals(uint128.Uint128(other))
}

func (bal Balance) Empty() bool {
	return uint128.Uint128(bal).IsZero()
}

func (bal Balance) Big() *big.Int {
	return uint128.Uint128(bal).Big()
}

func NEARToYocto(near uint64) Balance {
	if near == 0 {
		return ZeroNEAR
	}

	return Balance(uint128.From64(near).Mul(tenPower24))
}

func YoctoToNEAR(yocto Balance) uint64 {
	div := uint128.Uint128(yocto).Div(tenPower24)
	if h := div.Hi; h != 0 {
		panic(fmt.Errorf("yocto div failed, remaining: %d", h))
	}

	return div.Lo
}

func scaleToYocto(f *big.Float) (r *big.Int) {
	// Convert reference 1 NEAR to big.Float
	base := new(big.Float).SetPrec(128).SetInt(uint128.Uint128(NEARToYocto(1)).Big())

	// Multiply base using the supplied float
	// XXX: small precision issues here will haunt me forever
	bigf2 := new(big.Float).SetPrec(128).SetMode(big.ToZero).Mul(base, f)

	// Convert it to big.Int
	r, _ = bigf2.Int(nil)
	return
}

func BalanceFromString(s string) (bal Balance, err error) {
	var bigf *big.Float
	bigf, _, err = big.ParseFloat(s, 10, 128, big.ToZero)
	if err != nil {
		return
	}

	bal = Balance(uint128.FromBig(scaleToYocto(bigf)))
	return
}
