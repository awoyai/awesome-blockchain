package ECDSA

import (
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
	//k, _ := rand.Int(rand.Reader, e.Curve.Order)
	k := big.NewInt(4)
	r := new(big.Int).Mod(e.Curve.Mul(e.G, k).X, e.Curve.Order)
	num := new(big.Int)
	num = num.Add(new(big.Int).Mul(e.DA, r), z).Mod(num, e.Curve.Order)
	s := mathT.Mod(new(big.Rat).SetFrac(num, k), e.Curve.Order)
	return r, s
}

func (e *ECDSA) Valid(message string, r, s *big.Int) bool {
	negS := mathT.Mod(new(big.Rat).SetFrac(big.NewInt(1), s), e.Curve.Order)
	fmt.Println("negS", negS)
	z := e.sha1Message(message)
	point1 := new(big.Int).Mod(new(big.Int).Mul(z, negS), e.Curve.Order)
	point2 := new(big.Int).Mod(new(big.Int).Mul(r, negS), e.Curve.Order)
	P := e.Curve.Add(e.Curve.Mul(e.G, point1), e.Curve.Mul(e.QA, point2))
	fmt.Println("P", P)
	return P.X.Cmp(r) == 0
}

func (e *ECDSA) sha1Message(message string) *big.Int {
	h := sha1.New()
	h.Write([]byte(message))
	z := new(big.Int).SetBytes(h.Sum(nil))
	return z.Mod(z, e.Curve.Order)
}
