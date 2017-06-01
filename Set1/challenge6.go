package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

/*

Break repeating-key XOR
It is officially on, now.
This challenge isn't conceptually hard, but it involves actual error-prone coding. The other challenges in this set are there to bring you up to speed. This one is there to qualify you. If you can do this one, you're probably just fine up to Set 6.

There's a file here. It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

Here's how:

Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the number of differing bits. The distance between:
this is a test
and
wokka wokka!!!
is 37. Make sure your code agrees before you proceed.
For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4 KEYSIZE blocks instead of 2 and average the distances.
Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
Solve each block as if it was single-character XOR. You already have code to do this.
For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.
This code is going to turn out to be surprisingly useful later on. Breaking repeating-key XOR ("Vigenere") statistically is obviously an academic exercise, a "Crypto 101" thing. But more people "know how" to break it than can actually break it, and a similar technique breaks something much more important.

No, that's not a mistake.
We get more tech support questions for this challenge than any of the other ones. We promise, there aren't any blatant errors in this text. In particular: the "wokka wokka!!!" edit distance really is 37.

*/

func hamming(s1 string, s2 string) int {
	// This function assumes string inputs of same length
	total := 0
	//for each byte
	for i, x := range s1 {
		//xor byte with same position in s2
		xord := byte(x) ^ byte(s2[i])
		//convert result to binary
		bin := fmt.Sprintf("%b", xord)
		total += strings.Count(bin, "1")
	}
	return total

}

func breakRepeatingKeyXor(s1 string) string {

	//get keysize
	ks := getKeysize(s1)
	// break ciphertext into blocks	of keysize length
	chunks := chunker(s1, ks)
	// transpose the blocks: one block of all first bytes, one of all second bytes, etc
	transposed := transpose(chunks)
	// send each block to the single byte xor solver. Get the output from that to have the key
	//extractedKey := make([]string, ks)
	for _, transposedBlock := range transposed {
		fmt.Println(hex.EncodeToString([]byte(transposedBlock)))
		// singleByteXorNTest(transposedBlock)
		x, _, _ := singleByteXorNTest(hex.EncodeToString([]byte(transposedBlock)))
		fmt.Println(x)
	}
	return ""
}

func transpose(ss []string) []string {
	// first block will have the total size we need
	ts := make([]string, len(ss[0]))
	//iterate over each block
	for _, block := range ss {
		// iterate over each byte in the block
		for i, b := range block {
			//place byte in appropriate transposed block segment
			ts[i] += string(b)
		}
	}
	return ts
}

func getKeysize(s string) int {
	//key range
	var fromKey = 4
	var toKey = 400
	min := float32(9999999)
	size := -1
	//for each keysize
	for i := fromKey; i <= toKey; i++ {
		//take the first keysize of bytes
		k1 := s[0:i]
		//and the second keysize of bytes
		k2 := s[i : i*2]
		//get hamming distance
		dist := float32(hamming(string(k1), string(k2)))
		//divide this by keysize (to normalize)
		dist = float32(dist) / float32(i)
		//remember the lowest hamming distance (it's probably the key)
		if dist < min {
			min = dist
			size = i
		}
	}
	return size

}

func chunker(s string, chunksize int) []string {
	r := make([]string, (len(s)/chunksize)+1)
	j := 0
	for i := 0; i < len(s)-1; i += chunksize {
		r[j] = s[i : i+chunksize]
		j++
	}
	r[j] = s[(len(s)/chunksize)*chunksize:]
	return r
}
