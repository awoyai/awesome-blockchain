package curve

import (
	"fmt"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
	"testing"
)

func Test_CalculateThird(t *testing.T) {
	A := big.NewInt(0)
	B := big.NewInt(7)
	order, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	curve := NewEllipticCurve(A, B, order)

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

	x, b := new(big.Int).SetString("89565891926547004231252920425935692360644145829622209833684329913297188986597", 10)
	if !b {
		panic("cant reverse")
	}
	y, b := new(big.Int).SetString("83121579216557378445487899878180864668798711284981320763518679672151497189239", 10)
	if !b {
		panic("cant reverse")
	}
	G := mathT.Point{
		X: x,
		Y: y,
	}
	println(curve.OnCurve(G))
}



func Test_TestCurve(t *testing.T) {
	order := big.NewInt(43)
	c := NewEllipticCurve(big.NewInt(0), big.NewInt(1), order)
	G := mathT.Point{
		X: big.NewInt(2),
		Y: big.NewInt(3),
	}
	fmt.Println(c.OnCurve(G))
	z := big.NewInt(22332)
	z.Mod(z, order)
	da := big.NewInt(23232311)
	da.Mod(da, order)
	k := big.NewInt(3214141241)
	k.Mod(k, order)

	P := c.Mul(G, k)
	fmt.Println("P",P)
	r := new(big.Int).Mod(P.X, order)
	num := new(big.Int).Mul(r, da)
	num.Add(num, z)
	s := mathT.Mod(new(big.Rat).SetFrac(num, k), order)
	fmt.Println("R", r)
	fmt.Println("S", s)

	negS := mathT.Mod(new(big.Rat).SetFrac(big.NewInt(1), s), order)

	p1 := c.Mul(G, new(big.Int).Mul(negS, z))
	p2 := c.Mul(c.Mul(G, da), new(big.Int).Mul(negS, r))
	p3 := c.Add(p1, p2)
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
	fmt.Println("p3", p3)
}

func Test_TestCurve2(t *testing.T) {
	order := big.NewInt(23)
	c := NewEllipticCurve(big.NewInt(0), big.NewInt(1), order)
	G := mathT.Point{
		X: big.NewInt(2),
		Y: big.NewInt(3),
	}
	m := big.NewInt(24)
	p2 := c.Mul(G, m.Mod(m, order))
	fmt.Println(p2)
}