package Set1

import (
	"fmt"

	"github.com/c-sto/gocryptopals/pkg/xor"
)

/*
Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

Achievement Unlocked
You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
*/

//this file needs refactoring

func Challenge3() {
	fmt.Println("Test 3 Begin")

	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	lowestChar, plain, score := xor.SingleByteXorNTest(hexString)
	fmt.Printf("%v, %v, %v\n", lowestChar, plain, score)
	fmt.Println("Challenge 3 complete")
}
