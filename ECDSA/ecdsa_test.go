package ECDSA

import (
	"crypto/rand"
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	"testing"
)

func TestECDSA_Sign(t *testing.T) {
	c := curve.NewSecp256k1Curve()
	G := curve.NewSecp256k1G()
	message := "hello world"
	da, _ := rand.Int(rand.Reader, c.Order)
	fmt.Println(c.OnCurve(G))
	ecdsa := NewECDSAByPrivateKey(da, G, c)
	r, s := ecdsa.Sign(message)
	b := ecdsa.Valid(message, r, s)
	fmt.Println("R", r)
	fmt.Println("S", s)
	println(b)
}
