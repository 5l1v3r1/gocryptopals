package Set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/c-sto/cryptochallenges_golang/cryptolib"
)

/*
The Base64-encoded content in this file has been encrypted via AES-128 in ECB mode under the key

"YELLOW SUBMARINE".
(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.

Do this with code.
You can obviously decrypt this using the OpenSSL command-line tool, but we're having you get ECB working in code for a reason. You'll need it a lot later on, and not just for attacking ECB.

*/

func Challenge7() {

	fmt.Println("Test 7 Begin")
	content, err := ioutil.ReadFile("./resources/challenge7.txt")
	if err != nil {
		panic("file load error")
	}
	key := []byte("YELLOW SUBMARINE")

	ciphertext, _ := base64.StdEncoding.DecodeString(string(content))
	plain := cryptolib.AESECBDecrypt(ciphertext, key)

	if plain[0] == 0 || plain[40] == 0 {
		panic("Bad decrypt: " + string(plain))
	}

	fmt.Println(string(plain))
	fmt.Println("Challenge 7 complete")
}
