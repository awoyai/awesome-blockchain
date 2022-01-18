package EC_ElGamal

import (
	"fmt"
	"github.com/wuyedebianhua/awesome-blockchain/curve"
	"math/big"
	"testing"
)

func TestECElGamal_TransferMessage2Point(t *testing.T) {
	c := curve.NewSecp256k1Curve()
	G := curve.NewSecp256k1G()
	eceg := NewECElGamalByPrivateKey(big.NewInt(114514), G, c)
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
}
