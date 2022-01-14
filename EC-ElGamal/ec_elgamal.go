package EC_ElGamal

import (
	"crypto/sha1"
	"errors"
	"math/big"
	mathT "awesome-blockchain/math-tools"
	"awesome-blockchain/curve"
)

type ECElGamal struct {
	k  *big.Int
	g  mathT.Point
	kg mathT.Point
	y  curve.EllipticCurve
}

func (e *ECElGamal) Encrypt(m mathT.Point) (ans [2]mathT.Point, err error) {
	// 校验是否在曲线上
	if !e.y.OnCurve(m) {
		return ans, errors.New("the point doesn't on curve")
	}
	// TODO:// 替换成随机数
	r := big.NewInt(2)
	ans[0] = e.y.Mul(e.g, r)
	ans[1] = e.y.Add(m, e.kg)
	return
}

func (e *ECElGamal) Decrypt(cm [2]mathT.Point) (mathT.Point, error) {
	if e.k == nil {
		return mathT.Point{}, errors.New("have no private key")
	}
	return e.y.Sub(e.y.Mul(cm[1], e.k), cm[0]), nil
}

func NewECElGamalByPrivateKey(k *big.Int, g mathT.Point, y curve.EllipticCurve) ECElGamal {
	return ECElGamal{k: k, g: g, kg: y.Mul(g, k), y: y}
}

func NewECElGamalByPublicKey(g, kg mathT.Point, y curve.EllipticCurve) ECElGamal {
	return ECElGamal{g: g, kg: kg, y: y}
}

func (e *ECElGamal) TransferMessage2Point(message string) mathT.Point {
	h := sha1.New()
	h.Write([]byte(message))
	xx := h.Sum(nil)
	return mathT.Point{X: new(big.Int).SetBytes(xx)}
}
