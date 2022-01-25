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
	sk, _ := rand.Int(rand.Reader, c.P)

	eceg := NewECElGamalByPrivateKey(sk, G, c)
	pubEceg := NewECElGamalByPublicKey(G, eceg.KG, c)

	message := "hello"
	mPoint := eceg.TransferMessage2Point(message)

	res, err := pubEceg.Encrypt(mPoint)
	if err != nil {
		panic(err)
	}


	resStr, err := eceg.Decrypt(res)
	if err != nil {
		panic(err)
	}
	fmt.Println("M", mPoint)
	fmt.Println("K", res[0])
	fmt.Println("C", res[1])
	fmt.Println("RES", resStr)
	if resStr != message {
		t.Errorf("wrong")
	}
}
