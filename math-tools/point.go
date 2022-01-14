package math_tools

import "math/big"


type Point struct {
	X, Y *big.Int
}

func (p *Point) Equal(point Point) bool {
	return point.X.Cmp(p.X) == 0 && point.Y.Cmp(p.Y) == 0
}
