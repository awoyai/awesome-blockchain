package EC_ElGamal

import (
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
	"testing"
)

func TestECElGamal_TransferMessage2Point(t *testing.T) {
	A := big.NewInt(0)
	B := big.NewInt(7)
	order, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	c := curve.NewEllipticCurve(A, B, order)

	x, b := new(big.Int).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	if !b {
		panic("cant reverse")
	}
	y, b := new(big.Int).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	if !b {
		panic("cant reverse")
	}
	G := mathT.Point{
		X: x,
		Y: y,
	}
	eceg := NewECElGamalByPrivateKey(big.NewInt(10), G, c)
	mPoint := eceg.TransferMessage2Point("hello")
	fmt.Printf("%s,%s", mPoint.X, mPoint.Y)
	println(c.OnCurve(mPoint))
}
