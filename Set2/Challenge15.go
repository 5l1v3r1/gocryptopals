package Set2

import (
	"fmt"

	"github.com/c-sto/cryptochallenges_golang/cryptolib"
)

/*
PKCS#7 padding validation
Write a function that takes a plaintext, determines if it has valid PKCS#7 padding, and strips the padding off.

The string:

"ICE ICE BABY\x04\x04\x04\x04"
... has valid padding, and produces the result "ICE ICE BABY".

The string:

"ICE ICE BABY\x05\x05\x05\x05"
... does not have valid padding, nor does:

"ICE ICE BABY\x01\x02\x03\x04"
If you are writing in a language with exceptions, like Python or Ruby, make your function throw an exception on bad padding.

Crypto nerds know where we're going with this. Bear with us.
*/

func Challenge15() {
	s1 := []byte("ICE ICE BABY\x04\x04\x04\x04")
	s2 := []byte("ICE ICE BABY\x05\x05\x05\x05")
	s3 := []byte("ICE ICE BABY\x01\x02\x03\x04")

	_, err := cryptolib.PKCS7Unpad(s1, 16)
	if err != nil {
		panic("fail s1")
	}

	_, err = cryptolib.PKCS7Unpad(s2, 16)
	if err == nil {
		panic("fail s2")
	}

	_, err = cryptolib.PKCS7Unpad(s3, 16)
	if err == nil {
		panic("fail s3")
	}
	fmt.Println("passed tests")
}
