package Set2

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/c-sto/cryptochallenges_golang/cryptolib"
)

/*
Byte-at-a-time ECB decryption (Harder)
Take your oracle function from #12. Now generate a random count of random bytes and prepend this string to every plaintext. You are now doing:

AES-128-ECB(random-prefix || attacker-controlled || target-bytes, random-key)
Same goal: decrypt the target-bytes.

Stop and think for a second.
What's harder than challenge #12 about doing this? How would you overcome that obstacle? The hint is: you're using all the tools you already have; no crazy math is required.

Think "STIMULUS" and "RESPONSE".
*/
var randomPrefix = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func Challenge14() {
	key = cryptolib.RandomKey()
	//Identify what a block of A's looks like (so we know where the injection is)
	ct := Challenge14Oracle([]byte(strings.Repeat("A", 48)))

	knownBlock := []byte{}
	if repeated, index := cryptolib.HasRepeatedBlocks(ct, 16); repeated {
		knownBlock = cryptolib.Chunker(ct, 16)[index]
	}
	fmt.Println(knownBlock)
	fmt.Println(cryptolib.Chunker(ct, 16))
}

//GetIndexOfBlock given a slice of byteslices, will find the index that the given block is at
func GetIndexOfBlock(block []byte, blocks [][]byte) int {
	return -1
}

func Challenge14Oracle(in []byte) []byte {
	secret := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	return cryptolib.AESECBEncrypt(append(randomPrefix, append(in, decodedSecret...)...), key)
}
