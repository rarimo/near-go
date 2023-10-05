package common

import (
	"encoding/json"
	"fmt"
	"lukechampine.com/uint128"
	"math"
	"math/big"
)

var (
	TenPower24 = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
	ZeroNEAR   = Balance(uint128.From64(0))
	OneYocto   = Balance(uint128.From64(1))
	One        = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(24), nil)
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

	return Balance(uint128.From64(near).Mul(TenPower24))
}

func YoctoToNEAR(yocto Balance) uint64 {
	div := uint128.Uint128(yocto).Div(TenPower24)
	if h := div.Hi; h != 0 {
		panic(fmt.Errorf("yocto div failed, remaining: %d", h))
	}

	return div.Lo
}

func BalanceFromString(s string) (bal Balance, err error) {
	var f, o, r big.Rat

	_, ok := f.SetString(s)
	if !ok {
		return bal, fmt.Errorf("cannot parse amount: %s", s)
	}

	o.SetInt(One)
	r.Mul(&f, &o)

	is := r.FloatString(0)
	amount := big.NewInt(0)
	amount.SetString(is, 10)

	bal = Balance(uint128.FromBig(amount))
	return
}

func MustBalanceFromString(s string) Balance {
	bal, err := BalanceFromString(s)
	if err != nil {
		panic(err)
	}

	return bal
}
