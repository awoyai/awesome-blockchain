package curve

import (
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
)

var Point0 = mathT.Point{X: big.NewInt(0), Y: big.NewInt(0)}

type EllipticCurve struct {
	A, B, N, P *big.Int
}

func NewEllipticCurve(A, B, N, P *big.Int) EllipticCurve {
	return EllipticCurve{
		A: A,
		B: B,
		N: N,
		P: P,
	}
}

func NewEllipticCurveByStr(a, b, n, p string) EllipticCurve {
	A, _ := new(big.Int).SetString(a, 10)
	B, _ := new(big.Int).SetString(b, 10)
	N, _ := new(big.Int).SetString(n, 10)
	P, _ := new(big.Int).SetString(p, 10)
	return EllipticCurve{
		A: A,
		B: B,
		N: N,
		P: P,
	}
}

func (e *EllipticCurve) Add(a, b mathT.Point) mathT.Point {
	switch {
	case e.Is0(a):
		return b
	case e.Is0(b):
		return a
	case e.MirrorByX(a, b):
		return Point0
	}

	x := big.NewInt(1)
	y := big.NewInt(1)
	k := e.CalculateK(a, b)

	x = x.Mul(x, mathT.Pow(k, 2)).Sub(x, a.X).Sub(x, b.X).Mod(x, e.P)
	y = y.Mul(new(big.Int).Sub(a.X, x), k).Sub(y, a.Y).Mod(y, e.P)
	return mathT.Point{X: x, Y: y}
}

func (e *EllipticCurve) Sub(a, b mathT.Point) mathT.Point {
	b.Y.Mul(b.Y, big.NewInt(-1)).Mod(b.Y, e.P)
	sub := e.Add(a, b)
	return sub
}

func (e *EllipticCurve) Mul(a mathT.Point, k *big.Int) mathT.Point {
	res := Point0
	{
		k := new(big.Int).Set(k)
		a := a
		for k.Int64() != 0 {
			if k.Int64()&1 == 1 {
				res = e.Add(res, a)
			}
			k.Rsh(k, 1)
			a = e.Add(a, a)
		}
	}
	return res
}

func (e *EllipticCurve) Is0(p mathT.Point) bool {
	return p.X.Cmp(Point0.X) == 0 && p.Y.Cmp(Point0.Y) == 0
}

func (e *EllipticCurve) MirrorByX(a, b mathT.Point) bool {
	mirrorY := new(big.Int).Mod(new(big.Int).Mul(b.Y, big.NewInt(-1)), e.P)
	return (a.X.Cmp(b.X) == 0) && (a.Y.Cmp(mirrorY)) == 0
}

func (e *EllipticCurve) CalculateK(a, b mathT.Point) *big.Int {
	k := new(big.Rat)
	if a.Equal(b) {
		num := big.NewInt(3)
		num.Mul(num, mathT.Pow(a.X, 2)).Add(num, e.A)
		denom := new(big.Int).Mul(big.NewInt(2), a.Y)
		k.SetFrac(num, denom)
	} else {
		k.SetFrac(new(big.Int).Sub(b.Y, a.Y), new(big.Int).Sub(b.X, a.X))
	}
	return mathT.Mod(k, e.P)
}

func (e *EllipticCurve) OnCurve(p mathT.Point) bool {
	py2 := new(big.Int).Mod(mathT.Pow(p.Y, 2), e.P)
	return e.GetY2ByX(p.X).Cmp(py2) == 0
}

func (e *EllipticCurve) GetY2ByX(x *big.Int) *big.Int {
	// x^3 + ax + b
	x3 := mathT.Pow(x, 3)
	x3AndB := new(big.Int).Add(x3, e.B)
	y := new(big.Int).Add(x3AndB, new(big.Int).Mul(e.A, x))
	return new(big.Int).Mod(y, e.P)
}
