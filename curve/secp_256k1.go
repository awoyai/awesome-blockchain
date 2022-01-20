package curve

import (
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
)

func NewSecp256k1Curve() EllipticCurve {
	A := big.NewInt(0)
	B := big.NewInt(7)
	P, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	N, _ := new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	return NewEllipticCurve(A, B, N, P)
}

func NewSecp256k1G() mathT.Point {
	x, b := new(big.Int).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	if !b {
		panic("cant reverse")
	}
	y, b := new(big.Int).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	if !b {
		panic("cant reverse")
	}
	return mathT.Point{
		X: x,
		Y: y,
	}
}