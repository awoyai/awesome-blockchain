package EC_ElGamal

import (
	"fmt"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"go.dedis.ch/kyber/v3/util/random"
	"os"
)

func ElEncrypt(group kyber.Group, pubkey kyber.Point, message []byte) (
	K, C kyber.Point, remainder []byte) {

	// Embed the message (or as much of it as will fit) into a curve point.
	M := group.Point().Embed(message, random.New())

	fmt.Printf("Message point:\t%s\n" , M.String())

	max := group.Point().EmbedLen()

	if max > len(message) {
		max = len(message)
	}
	remainder = message[max:]

	// ElGamal-encrypt the point to produce ciphertext (K,C).

	k := group.Scalar().Pick(random.New()) // ephemeral private key
	K = group.Point().Mul(k, nil)          // ephemeral DH public key
	S := group.Point().Mul(k, pubkey)      // ephemeral DH shared secret
	C = S.Add(S, M)                        // message blinded with secret
	return
}

func ElDencrypt(group kyber.Group, prikey kyber.Scalar, K, C kyber.Point) (
	message []byte, err error) {

	// ElGamal-decrypt the ciphertext (K,C) to reproduce the message.
	S := group.Point().Mul(prikey, K) // regenerate shared secret
	M := group.Point().Sub(C, S)      // use to un-blind the message
	message, err = M.Data()           // extract the embedded data
	return
}


func main() {

	M:="Testing"
	argCount := len(os.Args[1:])

	if (argCount>0) {M= string(os.Args[1])}

	suite := edwards25519.NewBlakeSHA256Ed25519()


	// Alice's key pair (a,A)
	a := suite.Scalar().Pick(suite.RandomStream())
	A := suite.Point().Mul(a, nil)


	fmt.Printf("Private key (Alice):\t%s\n" ,a.String())
	fmt.Printf("Public key (Alice):\t%s\n" , A.String())


	fmt.Printf("\n\n--- Now Bob will cipher the message for Alice\n")
	fmt.Printf("Bob's message:\t%s\n",M)

	m := []byte(M)
	K, C, _ := ElEncrypt(suite, A, m)

	fmt.Printf("\nBob cipher (K):\t%s\n" , K.String())
	fmt.Printf("Bob cipher (C):\t%s\n" , C.String())

	fmt.Printf("\n\n--- Now Alice will decipher the ciphertext (K,C) with her private key (a)\n")


	M_decrypted, err := ElDencrypt(suite, a, K, C)


	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	fmt.Printf("\nOutput:\t%s",string(M_decrypted))

}