package Set2

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/c-sto/cryptochallenges_golang/cryptolib"
)

/*

Implement CBC mode
CBC mode is a block cipher mode that allows us to encrypt irregularly-sized messages, despite the fact that a block cipher natively only transforms individual blocks.

In CBC mode, each ciphertext block is added to the next plaintext block before the next call to the cipher core.

The first plaintext block, which has no associated previous ciphertext block, is added to a "fake 0th ciphertext block" called the initialization vector, or IV.

Implement CBC mode by hand by taking the ECB function you wrote earlier, making it encrypt instead of decrypt (verify this by decrypting whatever you encrypt to test), and using your XOR function from the previous exercise to combine them.

The file here is intelligible (somewhat) when CBC decrypted against "YELLOW SUBMARINE" with an IV of all ASCII 0 (\x00\x00\x00 &c)

Don't cheat.
Do not use OpenSSL's CBC code to do CBC mode, even to verify your results. What's the point of even doing this stuff if you aren't going to learn from it?

*/

func Challenge10() {
	fmt.Println("Begin Test 10")
	content, err := ioutil.ReadFile("./resources/challenge10.txt")
	if err != nil {
		panic("file load error")
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(string(content))
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	key := []byte("YELLOW SUBMARINE")
	r := cryptolib.AESCBCDecrypt(ciphertext, key, iv)
	fmt.Println(string(r))
	fmt.Println("If the above is garbage, it failed")

}
