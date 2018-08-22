package cryptoauthlib_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	element "github.com/riddleandcode/go-cryptoauthlib"
)

func TestRandom(t *testing.T) {

	zeroBytes := make([]byte, 32)

	const numLoops = 10

	fmt.Println("hello")
	for i := 0; i < 10; i++ {
		randomBytes := element.Random()

		fmt.Printf("randomBytes:\n%s\n", hex.Dump(randomBytes))

		if len(randomBytes) != 32 {
			t.Errorf("The length is incorrect, got: %d, want: %d.", len(randomBytes), 32)
		}

		if bytes.Equal(zeroBytes, randomBytes) {
			t.Errorf("The random bytes are all zero. Not very random, is it?\n%s", hex.Dump(randomBytes))
		}
    }

	var res int
	var random_number = Random()
	fmt.Printf("Here's your random: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(random_number))


	var digest = Random()
	fmt.Printf("Here's your digest: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(digest))

	var public_key []byte
	res, public_key = GetPublicKey()
	fmt.Printf("Here's your public key: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(public_key))

	var signature []byte
	res, signature = SignDigest(digest)
	fmt.Printf("Here's your signature: %d\n", res)
	fmt.Printf("%s\n", hex.Dump(signature))

	var verified bool
	res, verified = VerifySignedDigest( digest, signature, public_key )
	fmt.Printf("Here's your verification: %d\n", res)
	fmt.Printf("%t \n",verified )
	if !verified {
		t.Errorf("The verification of the singing process failed?")		
	}
}
