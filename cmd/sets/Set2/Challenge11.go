package Set2

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"strings"

	caes "github.com/c-sto/gocryptopals/pkg/aes"
	"github.com/c-sto/gocryptopals/pkg/padding"
)

/*
An ECB/CBC detection oracle
Now that you have ECB and CBC working:

Write a function to generate a random AES key; that's just 16 random bytes.

Write a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.

The function should look like:

encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]
Under the hood, have the function append 5-10 bytes (count chosen randomly) before the plaintext and 5-10 bytes after the plaintext.

Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC the other half (just use random IVs each time for CBC). Use rand(2) to decide which to use.

Detect the block cipher mode the function is using each time. You should end up with a piece of code that, pointed at a black box that might be encrypting ECB or CBC, tells you which one is happening.
*/

func Challenge11() {
	fmt.Println("Begin Test 11")
	//if it detects ECB, DoCBCorECB should return true
	wins := 0
	for x := 0; x < 1000; x++ {
		guess := false //default not ecb guess
		ct, confirm := DoCBCorECB([]byte(strings.Repeat("a", 100)))
		if x, _ := padding.HasRepeatedBlocks(ct, aes.BlockSize); x {
			guess = true
		}
		if guess == confirm {
			wins++
		}
	}
	if wins != 1000 {
		panic(fmt.Sprintf("Fail, detection sucks: %v", wins))
	}
	fmt.Println("ECB is dumb (hacked)")
}

func DoCBCorECB(plain []byte) ([]byte, bool) {
	isECB := true
	b := make([]byte, 1)
	rand.Read(b)
	if int(b[0])%2 == 0 {
		isECB = false
	}
	rand.Read(b)
	appendAmount := int(b[0])%5 + 5
	rand.Read(b)
	prependAmount := int(b[0])%5 + 5
	preBytes := make([]byte, prependAmount)
	postBytes := make([]byte, appendAmount)
	rand.Read(preBytes)
	rand.Read(postBytes)
	plain = append(preBytes, plain...)
	plain = append(plain, postBytes...)
	if isECB {
		b = caes.AESECBEncrypt(padding.PKCS7(plain, 16), caes.RandomKey())
	} else {
		b = caes.AESCBCEncrypt(padding.PKCS7(plain, 16), caes.RandomKey(), caes.RandomKey())
	}
	return b, isECB
}
