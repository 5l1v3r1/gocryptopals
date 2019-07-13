package Set1

import (
	"fmt"
	"strings"

	"github.com/c-sto/gocryptopals/asset"
	"github.com/c-sto/gocryptopals/pkg/xor"
)

/*

One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)

*/

func Challenge4() {

	fmt.Println("Test 4 Begin")
	//load lines
	content, err := asset.Challenge("challenge4.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\x0d\n")
	lowestChar, plain, score := xor.MultiSingleByteXorNTest(lines)
	if plain != "Now that the party is jumping\n" {
		panic(plain)
	}
	fmt.Printf("%v %v %v \n", lowestChar, plain, score)
	fmt.Println("Challenge 4 complete")
}
