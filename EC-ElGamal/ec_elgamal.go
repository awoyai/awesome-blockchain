package EC_ElGamal

import (
	"crypto/rand"
	"errors"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
)

type ECElGamal struct {
	K     *big.Int
	G     mathT.Point
	KG    mathT.Point
	Curve curve.EllipticCurve
}

func (e *ECElGamal) Encrypt(m mathT.Point) (ans [2]mathT.Point, err error) {
	// 校验是否在曲线上
	if !e.Curve.OnCurve(m) {
		return ans, errors.New("the point doesn't on curve")
	}
	k, _ := rand.Int(rand.Reader, e.Curve.P)
	ans[0] = e.Curve.Mul(e.G, k)
	ans[1] = e.Curve.Add(m, e.Curve.Mul(e.KG, k))
	return
}

func (e *ECElGamal) Decrypt(cm [2]mathT.Point) (string, error) {
	if e.K == nil {
		return "", errors.New("have no private Key")
	}
	resPoint := e.Curve.Sub(cm[1], e.Curve.Mul(cm[0], e.K))
	return string(resPoint.X.Bytes()), nil
}

func NewECElGamalByPrivateKey(K *big.Int, G mathT.Point, curve curve.EllipticCurve) ECElGamal {
	return ECElGamal{K: K, G: G, KG: curve.Mul(G, K), Curve: curve}
}

func NewECElGamalByPublicKey(G, KG mathT.Point, curve curve.EllipticCurve) ECElGamal {
	return ECElGamal{G: G, KG: KG, Curve: curve}
}

func (e *ECElGamal) TransferMessage2Point(message string) mathT.Point {
	x := new(big.Int).SetBytes([]byte(message))
	y := mathT.Sqrt(e.Curve.GetY2ByX(x), e.Curve.P)
	return mathT.Point{
		X: x,
		Y: y,
	}
}
