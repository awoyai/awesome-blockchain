package EC_ElGamal

import (
	"crypto/rand"
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	"testing"
)

func TestECElGamal_TransferMessage2Point(t *testing.T) {
	c := curve.NewSecp256k1Curve()
	G := curve.NewSecp256k1G()
	sk, _ := rand.Int(rand.Reader, c.Order)
	eceg := NewECElGamalByPrivateKey(sk, G, c)
	mPoint := eceg.TransferMessage2Point("hello")

	res, err := eceg.Encrypt(mPoint)
	if err != nil {
		panic(err)
	}

	resPoint, err := eceg.Decrypt(res)
	if err != nil {
		panic(err)
	}
	fmt.Println("M", mPoint)
	fmt.Println("K", res[0])
	fmt.Println("C", res[1])
	fmt.Println("RES", resPoint)
	if mPoint.X.Cmp(resPoint.X) != 0 && mPoint.Y.Cmp(resPoint.Y) != 0 {
		t.Errorf("wrong")
	}
}
