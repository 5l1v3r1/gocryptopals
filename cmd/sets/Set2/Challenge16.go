package Set2

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/c-sto/gocryptopals/pkg/aes"
	"github.com/c-sto/gocryptopals/pkg/cryptobytes"
	"github.com/c-sto/gocryptopals/pkg/padding"
	"github.com/c-sto/gocryptopals/pkg/xor"
)

/*
CBC bitflipping attacks
Generate a random AES key.

Combine your padding code and CBC code to write two functions.

The first function should take an arbitrary input string, prepend the string:

"comment1=cooking%20MCs;userdata="
.. and append the string:

";comment2=%20like%20a%20pound%20of%20bacon"
The function should quote out the ";" and "=" characters.

The function should then pad out the input to the 16-byte AES block length and encrypt it under the random AES key.

The second function should decrypt the string and look for the characters ";admin=true;" (or, equivalently, decrypt, split the string on ";", convert each resulting string into 2-tuples, and look for the "admin" tuple).

Return true or false based on whether the string exists.

If you've written the first function properly, it should not be possible to provide user input to it that will generate the string the second function is looking for. We'll have to break the crypto to do that.

Instead, modify the ciphertext (without knowledge of the AES key) to accomplish this.

You're relying on the fact that in CBC mode, a 1-bit error in a ciphertext block:

Completely scrambles the block the error occurs in
Produces the identical 1-bit error(/edit) in the next ciphertext block.
Stop and think for a second.
Before you implement this attack, answer this question: why does CBC mode have this property?
*/

func Challenge16() {
	key = aes.RandomKey()
	//blocksize := aes.BlockSize
	x := Challenge16_Function1(strings.Repeat("A", 256))
	y := false
	var err error
	//get count of ciphertext
	chunks := cryptobytes.Chunker(x, aes.BlockSize)
	//blockCount := len(chunks)
	for i := 0; i < 256; i++ {
		//most of the encrypted string is known plaintext.
		//Taking a gamble, but we assume block 10 and 11 will be known plaintext
		//to decrypt block 11 into a known value, we decrypt the block, then xor it against the previous block (10)
		//this means that p = aesdec(11) ^ block[10]
		//which means that if we flip the bits in the last byte of block 10 and send block 11 as the last block
		//we should be able to control the ciphertext
		chunks[10] = xor.XorBytes(
			[]byte(strings.Repeat(string(byte(i)), 16)),
			chunks[10],
		)
		//11th chunk should be all 1's at this point
		//xor everything with 1's to make it 0's
		chunks[10] = xor.XorBytes(
			[]byte(strings.Repeat(string(byte(1)), 16)),
			chunks[10],
		)
		//we want to complete the block, so xor with the desired value to make the thing do the thing
		chunks[10] = xor.XorBytes(
			padding.PKCS7([]byte(";admin=true;"), 16),
			chunks[10],
		)

		y, err = Challenge16_Function2(bytes.Join(chunks[:12], nil))
		if err == nil {
			//we know i will cause plaintex to decrypt to \x01
			break
		}
	}
	if y {
		fmt.Println("Wow hacked!", err)
	} else {
		fmt.Println(y, err)
	}

}

func Challenge16_Function2(input []byte) (bool, error) {
	//decrypt
	plaintext := aes.AESCBCDecrypt(input[16:], key, input[:16])
	//unpad
	plaintext, err := padding.PKCS7Unpad(plaintext, 16)
	if err != nil {
		return false, err
	}
	//check for string
	return strings.Contains(string(plaintext), ";admin=true;"), nil

}

func Challenge16_Function1(input string) []byte {
	//get rid of naughty vals
	input = strings.Replace(input, "=", "", -1)
	input = strings.Replace(input, ";", "", -1)
	//concat
	input = "comment1=cooking%20MCs;userdata=" + input + ";comment2=%20like%20a%20pound%20of%20bacon"
	//pad
	padded := padding.PKCS7([]byte(input), 16)
	//fmt.Println(padded)
	//encrypt
	iv := aes.RandomKey()
	encrypted := aes.AESCBCEncrypt(padded, key, iv)
	encrypted = append(iv, encrypted...)
	return encrypted
}
