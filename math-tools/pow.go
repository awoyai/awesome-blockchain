package math_tools

import "math/big"

func Pow(a *big.Int, p int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(0); i < p; i++ {
		res.Mul(res, a)
	}
	return res
}

func FastPow(a1, b1, p *big.Int) *big.Int {
	var res = big.NewInt(1)
	a := new(big.Int).Set(a1)
	b := new(big.Int).Set(b1)
	a.Mod(a, p)
	for b.Int64() != 0 {
		if b.Int64()&1 == 1 {
			res.Mul(res, a).Mod(res, p)
		}
		b = b.Rsh(b, 1)
		a.Mul(a, a).Mod(a, p)
	}
	return res
}