package schnorr

import (
	"crypto/rand"
	"crypto/sha1"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
)

type Schnorr struct {
	Curve curve.EllipticCurve
	Sk    *big.Int
	G     mathT.Point
	Pk    mathT.Point
}

func NewSchnorr(ellipticCurve curve.EllipticCurve, g mathT.Point, sk *big.Int) Schnorr {
	return Schnorr{
		Curve: ellipticCurve,
		Sk:    sk,
		G:     g,
		Pk:    ellipticCurve.Mul(g, sk),
	}
}

func (sc *Schnorr) Sign(message string) (c, s *big.Int) {
	k, _ := rand.Int(rand.Reader, sc.Curve.Order)
	r := sc.Curve.Mul(sc.G, k)
	c = sc.hash(sc.Pk, r, message)
	// S = K + C * Sk
	s = new(big.Int).Mul(c, sc.Sk)
	s.Add(k, s).Mod(s ,sc.Curve.N)
	return
}

func (sc *Schnorr) hash(p, r mathT.Point, message string) *big.Int {
	str := p.X.String() + p.Y.String() + message + r.X.String() + r.Y.String()
	h := sha1.New()
	h.Write([]byte(str))
	c := new(big.Int).SetBytes(h.Sum(nil))
	return c.Mod(c, sc.Curve.Order)
}

func (sc *Schnorr) Verify(c, s *big.Int, message string) bool {
	// R' = S * G - C * PK
	r := sc.Curve.Sub(sc.Curve.Mul(sc.G, s), sc.Curve.Mul(sc.Pk, c))
	// C' = hash(Pk, R', m)
	c_ := sc.hash(sc.Pk, r, message)
	return c_.Cmp(c) == 0
}
