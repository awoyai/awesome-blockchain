package curve


import (
	"fmt"
	"math/big"
	mathT "awesome-blockchain/math-tools"
	"testing"
)

func Test_CalculateThird(t *testing.T) {
	A := big.NewInt(0)
	B := big.NewInt(7)
	order, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	curve := NewEllipticCurve(A, B, order)

	x,b := new(big.Int).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	if !b {
		panic("cant reverse")
	}
	y,b := new(big.Int).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	if !b {
		panic("cant reverse")
	}
	G := mathT.Point{
		X: x,
		Y: y,
	}
	println(curve.OnCurve(G))
	G2 := curve.Add(G, G)
	Gsub := curve.Sub(G2, G)
	G10 := curve.Mul(G, big.NewInt(10))
	println(curve.OnCurve(G2))
	println(curve.OnCurve(G10))
	println(curve.OnCurve(Gsub))
	fmt.Printf("%s,%s\n", G.X.String(), G.Y.String())
	fmt.Printf("%s,%s\n", G2.X.String(), G2.Y.String())
	fmt.Printf("%s,%s\n", Gsub.X.String(), Gsub.Y.String())
	fmt.Printf("%s,%s\n", G10.X.String(), G10.Y.String())
}

func Test_ModRat(t *testing.T) {
	rat := new(big.Rat).SetFrac(big.NewInt(29), big.NewInt(12))
	k := mathT.Mod(rat, big.NewInt(97))
	println(k.Int64())
}

func Test_CheckOnCurve(t *testing.T) {
	A := big.NewInt(0)
	B := big.NewInt(7)
	order, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	curve := NewEllipticCurve(A, B, order)

	x,b := new(big.Int).SetString("89565891926547004231252920425935692360644145829622209833684329913297188986597", 10)
	if !b {
		panic("cant reverse")
	}
	y,b := new(big.Int).SetString("83121579216557378445487899878180864668798711284981320763518679672151497189239", 10)
	if !b {
		panic("cant reverse")
	}
	G := mathT.Point{
		X: x,
		Y: y,
	}
	println(curve.OnCurve(G))
}