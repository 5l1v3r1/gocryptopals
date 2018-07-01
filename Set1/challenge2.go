package Set1

import (
	"encoding/hex"
	"fmt"

	"github.com/c-sto/cryptochallenges_golang/cryptolib"
)

/*
!!!! see test file for confirmation that this works

Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179
*/

func Challenge2() {
	fmt.Println("Test 2 Begin")
	v := cryptolib.XorHexStrings("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if v == "746865206b696420646f6e277420706c6179" {
		fmt.Println("Challenge 2 complete")
		s, _ := hex.DecodeString(v)
		fmt.Println(string(s))
	} else {
		panic("String output does not match")
	}
}
