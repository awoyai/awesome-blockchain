package math_tools

import "math/big"

func Mod(rat *big.Rat, p *big.Int) *big.Int {
	if rat.Denom().Int64() == 1 {
		return new(big.Int).Mod(rat.Num(), p)
	}
	res := big.NewInt(1)
	fastPow := FastPow(rat.Denom(), new(big.Int).Sub(p, big.NewInt(2)), p)
	res.Mod(rat.Num(), p).Mul(res, fastPow).Mod(res, p)
	return res
}
