package main

import "encoding/hex"

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

func xorHexStrings(arg1 string, arg2 string) string {
	//turn args to bytes
	bytes1, _ := hex.DecodeString(arg1)
	bytes2, _ := hex.DecodeString(arg2)

	//use shortest string
	n := len(bytes1)
	if len(bytes1) > len(bytes2) {
		n = len(bytes2)
	}
	ret := make([]byte, n)
	//xor the things
	for i := 0; i < n; i++ {
		ret[i] = bytes1[i] ^ bytes2[i]
	}
	//turn return bytes back to hex
	hexret := hex.EncodeToString(ret)
	return string(hexret)
}
