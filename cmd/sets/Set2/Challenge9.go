package Set2

import (
	"fmt"
	"reflect"

	"github.com/c-sto/gocryptopals/pkg/padding"
)

/*

Implement PKCS#7 padding
A block cipher transforms a fixed-sized block (usually 8 or 16 bytes) of plaintext into ciphertext. But we almost never want to transform a single block; we encrypt irregularly-sized messages.

One way we account for irregularly-sized messages is by padding, creating a plaintext that is an even multiple of the blocksize. The most popular padding scheme is called PKCS#7.

So: pad any block to a specific block length, by appending the number of bytes of padding to the end of the block. For instance,

"YELLOW SUBMARINE"
... padded to 20 bytes would be:

"YELLOW SUBMARINE\x04\x04\x04\x04"

*/

func Challenge9() {
	fmt.Println("Begin Test 9")
	x := padding.PKCS7([]byte("YELLOW SUBMARINE"), 20)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		panic(fmt.Sprintf("Bad padding: %v", x))
	}
	x = padding.PKCS7([]byte("YELLOW SUBMARINE"), 19)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x03\x03\x03")) {
		panic(fmt.Sprintf("Bad padding: %v", x))
	}
	x = padding.PKCS7([]byte("YELLOW SUBMARINE"), 15)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")) {
		panic(fmt.Sprintf("Bad padding: %v %v", x, []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")))
	}
	fmt.Println("Test9 Complete!")
}
