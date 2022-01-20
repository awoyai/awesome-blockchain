package schnorr

import (
	"crypto/rand"
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"testing"
)

func TestSchnorr_Sign(t *testing.T) {
	curve2561k1 := curve.NewSecp256k1Curve()
	G2561k1 := curve.NewSecp256k1G()

	ts := []tests{
		{
			Curve:    curve2561k1,
			G:        G2561k1,
			testText: "hello world",
		},
		{
			Curve:    curve2561k1,
			G:        G2561k1,
			testText: "hello llj",
		},
	}
	for _, v := range ts {
		sk, _ := rand.Int(rand.Reader, v.Curve.P)
		schnorr := NewSchnorr(v.Curve, v.G, sk)
		c,s := schnorr.Sign(v.testText)
		b := schnorr.Verify(c, s, v.testText)
		if !b {
			t.Errorf("wrong")
		}
		fmt.Println("success")
	}
}

type tests struct {
	Curve    curve.EllipticCurve
	G        mathT.Point
	testText string
}

func TestSchnorr_Verify(t *testing.T) {
	curve2561k1 := curve.NewSecp256k1Curve()
	G2561k1 := curve.NewSecp256k1G()
	sk, _ := rand.Int(rand.Reader, curve2561k1.P)
	schnorr := NewSchnorr(curve2561k1, G2561k1, sk)


	testText := "hello world"
	c, r := schnorr.Sign(testText)
	b := schnorr.Verify(c, r, testText)
	fmt.Println(b)
}