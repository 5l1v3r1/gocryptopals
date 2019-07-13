package Set1

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/c-sto/gocryptopals/asset"

	"github.com/c-sto/gocryptopals/pkg/padding"
)

/*

Detect AES in ECB mode
In this file are a bunch of hex-encoded ciphertexts.

One of them has been encrypted with ECB.

Detect it.

Remember that the problem with ECB is that it is stateless and deterministic; the same 16 byte plaintext block will always produce the same 16 byte ciphertext.

*/

func Challenge8() {
	fmt.Println("Test 8 Begin")
	content, err := asset.Challenge("challenge8.txt")
	if err != nil {
		panic("file load error")
	}
	lines := strings.Split(string(content), "\n")
	found := false
	for i, x := range lines {
		s, _ := hex.DecodeString(x)
		if y, _ := padding.HasRepeatedBlocks(s, 16); y {
			fmt.Println("Identified repeated block on line:", i)
			found = true
		}
	}
	if !found {
		panic("No duplicate block found!?!?")
	}
	fmt.Println("Challenge 8 complete")
}
