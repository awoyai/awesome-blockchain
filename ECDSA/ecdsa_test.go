package ECDSA

import (
	"crypto/rand"
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
	"testing"
)

func TestECDSA_Sign(t *testing.T) {
	c := curve.NewSecp256k1Curve()
	G := curve.NewSecp256k1G()
	message := "hello world"
	da, _ := rand.Int(rand.Reader, c.P)
	fmt.Println(c.OnCurve(G))
	ecdsa := NewECDSAByPrivateKey(da, G, c)
	r, s := ecdsa.Sign(message)
	b := ecdsa.Valid(message, r, s)
	fmt.Println("R", r)
	fmt.Println("S", s)
	println(b)
}

func TestECDSA_SignX3(t *testing.T) {
	c := curve.NewEllipticCurveByStr("0", "9", "93", "97")
	G := mathT.Point{
		X: big.NewInt(3),
		Y: big.NewInt(6),
	}
	message := "hello world"
	da, _ := rand.Int(rand.Reader, c.P)
	da = big.NewInt(43)
	fmt.Println(c.OnCurve(G))
	ecdsa := NewECDSAByPrivateKey(da, G, c)
	r, s := ecdsa.Sign(message)
	b := ecdsa.Valid(message, r, s)
	fmt.Println("R", r)
	fmt.Println("S", s)
	println(b)
}
