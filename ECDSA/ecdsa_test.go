package ECDSA

import (
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	"math/big"
	"testing"
)

func TestECDSA_Sign(t *testing.T) {
	c := curve.NewSecp256k1Curve()
	G := curve.NewSecp256k1G()
	message := "hello world"
	// todo random da
	//k, _ := rand.Int(rand.Reader, c.Order)
	da := big.NewInt(123)
	fmt.Println(c.OnCurve(G))
	ecdsa := NewECDSAByPrivateKey(da, G, c)
	r, s := ecdsa.Sign(message)
	b := ecdsa.Valid(message, r, s)
	fmt.Println("R", r)
	fmt.Println("S", s)
	println(b)
}
