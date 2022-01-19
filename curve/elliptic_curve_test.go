package curve

import (
	"fmt"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
	"testing"
)

func Test_CalculateThird(t *testing.T) {
	curve := NewSecp256k1Curve()
	G := NewSecp256k1G()
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
	curve := NewSecp256k1Curve()
	G := NewSecp256k1G()
	println(curve.OnCurve(G))
}