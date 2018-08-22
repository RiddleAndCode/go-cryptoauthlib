/*
Generate 32 bytes of true, hardware-generated, as-good-as-it-gets,
finest-entropy randomness.
*/
package cryptoauthlib

//#cgo CFLAGS: -I./inc
//#cgo LDFLAGS: -L. -lcryptoauth
//#include "./wrapper.h"
import "C"
import (
	"unsafe"
)

/*
Random generates 32 bytes of true, hardware-generated, as-good-as-it-gets,
finest-entropy randomness.
*/
func Random() []byte {
	randomBytes := make([]byte, C.RANDOM_NUM_SIZE)

	C.getRandomNumber((*C.uint8_t)(unsafe.Pointer(&randomBytes[0])))
	
	return randomBytes
}

/*
Returns the public key belonging to the RIDDLE&CODE Secure Element.
*/
func GetPublicKey() (int, []byte) {
	var result C.int
	pub_key := make([]byte, C.PUBLIC_KEY_SIZE)

	result = C.getPublicKey((*C.uint8_t)(unsafe.Pointer(&pub_key[0])))

    return int(result), pub_key
}

/*
Signs the given digest with the private key of the RIDDLE&CODE Secure Element.
*/
func SignDigest( digest []byte ) (int, []byte) {
	var result C.int
	signature :=make([]byte, C.SIGNATURE_SIZE)

	result = C.sign( (*C.uint8_t)(unsafe.Pointer(&digest[0])), 
					(*C.uint8_t)(unsafe.Pointer(&signature[0])))

	return int(result), signature
}

/*
Verifies if the given signature has been created by the combintion of the given
digest and the given identity associated with the public_key.
*/
func VerifySignedDigest( digest []byte, 
				signature []byte, public_key []byte ) (int, bool){
	var result C.int
	var verified C.bool

	result = C.verify_extern( (*C.uint8_t)(unsafe.Pointer(&digest[0])), 
				(*C.uint8_t)(unsafe.Pointer(&signature[0])),
				(*C.uint8_t)(unsafe.Pointer(&public_key[0])), &verified )

	return int(result), bool(verified)
}

