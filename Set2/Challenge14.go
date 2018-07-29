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
var randomPrefix = []byte{1, 2}

func Challenge14() {
	key = cryptolib.RandomKey()
	//randomPrefix = cryptolib.RandomBytes()
	blockSize := 16
	padCount := 48
	//Identify what a block of A's looks like (so we know where the injection is)
	ct := Challenge14Oracle([]byte(strings.Repeat("A", padCount)))

	_, knownIndex := cryptolib.HasRepeatedBlocks(ct, 16)
	//reduce the repeat count until the duplicate disappears
	for { //go and it's stupid scopes
		padCount--
		ct = Challenge14Oracle([]byte(strings.Repeat("A", padCount)))

		y, _ := cryptolib.HasRepeatedBlocks(ct, 16)
		if !y || padCount < 1 {
			break
		}
	}
	//we know that two full blocks occur at current padcount+1, so the 'standing' prefix len will be
	//(padcount+1)-(blocksize*2)
	padCount++
	standingPrefix := []byte(strings.Repeat("A", (padCount)-(16*2)))

	unknown := len(cryptolib.Chunker(Challenge14Oracle([]byte(strings.Repeat("A", padCount))), 16)) - (knownIndex + 1)

	//BEGIN COPYPASTAD CODE FROM 12
	known := []byte{}
	for blockIndex := knownIndex; blockIndex < unknown+1; blockIndex++ {
		for unknownByteIndex := 0; unknownByteIndex < blockSize; unknownByteIndex++ {

			//generate our checker block
			checkerBlock := standingPrefix
			//if we don't have enough known bytes, the checker will be padded with known bytes (A's)
			lastIndex := len(known) - blockSize + 1
			if len(known) < blockSize {
				checkerBlock = append(checkerBlock, []byte(strings.Repeat("A", blockSize-len(known)-1))...)
				lastIndex = 0
			}
			if len(known) > 0 {
				substringKnown := known[lastIndex:]
				checkerBlock = append(checkerBlock, substringKnown...)
			}

			for candidate := 0; candidate < 256; candidate++ {
				//append to candidate
				checkerBlock = append(checkerBlock, byte(candidate))

				//if we need to add the padding bytes because we don't know enough yet, make it so
				checkerBlock = append(checkerBlock, []byte(strings.Repeat("A", blockSize-unknownByteIndex-1))...)

				ct := Challenge14Oracle(checkerBlock)
				//fmt.Println(checkerBlock, "\n", cryptolib.Chunker(ct, blockSize)[1], cryptolib.Chunker(ct, blockSize)[blockIndex])
				//check if our block[0] matches the target
				if cryptolib.CompareBlocks(ct, blockSize, 1, blockIndex) {
					fmt.Print(string(byte(candidate)))
					known = append(known, byte(candidate))
					break
				}

				//reset checkerboi
				checkerBlock = checkerBlock[:(blockSize+len(standingPrefix))-1]
			}

		}
	}

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
	v := append(randomPrefix, append(in, decodedSecret...)...)
	return cryptolib.AESECBEncrypt(v, key)
}
