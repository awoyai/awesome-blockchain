package EC_ElGamal

import (
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	mathT "github.com/wuyedebianhua/awesome-blockchain/math-tools"
	"math/big"
)

type ECElGamal struct {
	K  *big.Int
	G  mathT.Point
	KG mathT.Point
	y  curve.EllipticCurve
}

func (e *ECElGamal) Encrypt(m mathT.Point) (ans [2]mathT.Point, err error) {
	// 校验是否在曲线上
	if !e.y.OnCurve(m) {
		return ans, errors.New("the point doesn't on curve")
	}
	k, _ := rand.Int(rand.Reader,e.y.Order)
	ans[0] = e.y.Mul(e.G, k)
	ans[1] = e.y.Add(m, e.y.Mul(e.KG, k))
	return
}

func (e *ECElGamal) Decrypt(cm [2]mathT.Point) (mathT.Point, error) {
	if e.K == nil {
		return mathT.Point{}, errors.New("have no private Key")
	}
	return e.y.Sub(cm[1],e.y.Mul(cm[0], e.K)), nil
}

func NewECElGamalByPrivateKey(K *big.Int, G mathT.Point, y curve.EllipticCurve) ECElGamal {
	return ECElGamal{K: K, G: G, KG: y.Mul(G, K), y: y}
}

func NewECElGamalByPublicKey(G, KG mathT.Point, y curve.EllipticCurve) ECElGamal {
	return ECElGamal{G: G, KG: KG, y: y}
}

func (e *ECElGamal) TransferMessage2Point(message string) mathT.Point {
	h := sha1.New()
	h.Write([]byte(message))
	xx := h.Sum(nil)
	return e.y.Mul(e.G, new(big.Int).SetBytes(xx))
}
