package Set1

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/c-sto/cryptochallenges_golang/cryptolib"
)

/*

One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)

*/

func Challenge4() {
	fmt.Println("Test 4 Begin")
	//load lines
	content, err := ioutil.ReadFile("./resources/challenge4.txt")
	if err != nil {
		panic("file load error")
	}
	lines := strings.Split(string(content), "\n")
	lowestChar, plain, score := cryptolib.MultiSingleByteXorNTest(lines)
	if plain != "Now that the party is jumping\n" {
		panic(fmt.Sprintf("Incorrect output: %v %v", plain, score))
	}
	fmt.Println("Lowest char, score")
	fmt.Printf("%v, %v \n", lowestChar, score)
	fmt.Println("Plaintext: ", plain)
	fmt.Println("Challenge 4 complete")
}
