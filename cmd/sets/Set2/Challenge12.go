package Set2

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/c-sto/gocryptopals/pkg/aes"
	"github.com/c-sto/gocryptopals/pkg/cryptobytes"
	"github.com/c-sto/gocryptopals/pkg/padding"
)

/*Byte-at-a-time ECB decryption (Simple)
Copy your oracle function to a new function that encrypts buffers under ECB mode using a consistent but unknown key (for instance, assign a single random key, once, to a global variable).

Now take that same function and have it append to the plaintext, BEFORE ENCRYPTING, the following string:

Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK
Spoiler alert.
Do not decode this string now. Don't do it.

Base64 decode the string before appending it. Do not base64 decode the string by hand; make your code do it. The point is that you don't know its contents.

What you have now is a function that produces:

AES-128-ECB(your-string || unknown-string, random-key)
It turns out: you can decrypt "unknown-string" with repeated calls to the oracle function!

Here's roughly how:

Feed identical bytes of your-string to the function 1 at a time --- start with 1 byte ("A"), then "AA", then "AAA" and so on. Discover the block size of the cipher. You know it, but do this step anyway.
Detect that the function is using ECB. You already know, but do this step anyways.
Knowing the block size, craft an input block that is exactly 1 byte short (for instance, if the block size is 8 bytes, make "AAAAAAA"). Think about what the oracle function is going to put in that last byte position.
Make a dictionary of every possible last byte by feeding different strings to the oracle; for instance, "AAAAAAAA", "AAAAAAAB", "AAAAAAAC", remembering the first block of each invocation.
Match the output of the one-byte-short input to one of the entries in your dictionary. You've now discovered the first byte of unknown-string.
Repeat for the next byte.
Congratulations.
This is the first challenge we've given you whose solution will break real crypto. Lots of people know that when you encrypt something in ECB mode, you can see penguins through it. Not so many of them can decrypt the contents of those ciphertexts, and now you can. If our experience is any guideline, this attack will get you code execution in security tests about once a year.
*/

func Challenge12() {
	key := aes.RandomKey()
	//discover blocksize
	bigblock := Challenge12Oracle([]byte(strings.Repeat("A", 100)), key)
	blockSize := -1
	for i := 3; i < 99; i++ {
		if x, _ := padding.HasRepeatedBlocks(bigblock, i); x {
			blockSize = i
			break
		}

	}
	unknown := len(cryptobytes.Chunker(Challenge12Oracle([]byte{}, key), blockSize))
	fmt.Println("Detected ECB blocksize: ", blockSize)
	fmt.Println("Estimated sekret blocks:", unknown)

	//previous solution put A's in place of ciphertext. This solution will only use at most 2 blocks of A's
	//pad := []byte(strings.Repeat("A", blockSize))
	known := []byte{}

	for blockIndex := 0; blockIndex < unknown; blockIndex++ {
		for unknownByteIndex := 0; unknownByteIndex < blockSize; unknownByteIndex++ {

			//generate our checker block
			checkerBlock := []byte{}
			//if we don't have enough known bytes, the checker will be padded with known bytes (A's)
			lastIndex := len(known) - blockSize + 1
			if len(known) < blockSize {
				checkerBlock = []byte(strings.Repeat("A", blockSize-len(known)-1))
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

				ct := Challenge12Oracle(checkerBlock, key)

				//check if our block[0] matches the target
				if padding.CompareBlocks(ct, blockSize, 0, blockIndex+1) {
					//fmt.Println(string(checkerBlock))
					fmt.Print(string(byte(candidate)))
					known = append(known, byte(candidate))
					//fmt.Println("Hax", blockIndex, unknownByteIndex, byte(candidate))
					break
				}

				//reset checkerboi
				checkerBlock = checkerBlock[:blockSize-1]
			}

		}
	}
}

func Challenge12Oracle(in, key []byte) []byte {
	secret := `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	return aes.AESECBEncrypt(append(in, decodedSecret...), key)
}
