package math_tools

import "math/big"

// Sqrt y2 = y^2 p为模 res为y
// p must module 4 = 3
func Sqrt(y2, p *big.Int) *big.Int {
	n := new(big.Int).Add(p, big.NewInt(1))
	n.Div(n, big.NewInt(4))
	return FastPow(y2, n, p)
}