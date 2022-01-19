package ECDSA

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
)

type ECDSA struct {
	DA    *big.Int
	QA    mathT.Point
	G     mathT.Point
	Curve curve.EllipticCurve
}

func NewECDSAByPrivateKey(da *big.Int, g mathT.Point, curve curve.EllipticCurve) ECDSA {
	return ECDSA{
		DA:    da,
		QA:    curve.Mul(g, da),
		G:     g,
		Curve: curve,
	}
}

func (e *ECDSA) Sign(message string) (*big.Int, *big.Int) {
	z := e.sha1Message(message)
	k, _ := rand.Int(rand.Reader, e.Curve.N)
	k = big.NewInt(2)
	negK := new(big.Int).Exp(k, big.NewInt(-1), e.Curve.N)
	r := e.Curve.Mul(e.G, k).X
	s := new(big.Int).Add(new(big.Int).Mul(e.DA, r), z)
	s.Mul(s, negK).Mod(s, e.Curve.N)
	return r, s
}

func (e *ECDSA) Valid(message string, r, s *big.Int) bool {
	negS := new(big.Int).Exp(s, big.NewInt(-1), e.Curve.N)
	z := e.sha1Message(message)
	// P = KG  KG = S^-1 (Z*G + R*QA)
	P := e.Curve.Mul(e.Curve.Add(e.Curve.Mul(e.G, z), e.Curve.Mul(e.QA, r)), negS)
	fmt.Println("P.x", P.X)
	return P.X.Cmp(r) == 0
}

func (e *ECDSA) sha1Message(message string) *big.Int {
	h := sha1.New()
	h.Write([]byte(message))
	z := new(big.Int).SetBytes(h.Sum(nil))
	return z.Mod(z, e.Curve.Order)
}
