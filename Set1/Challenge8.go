package main

/*

Detect AES in ECB mode
In this file are a bunch of hex-encoded ciphertexts.

One of them has been encrypted with ECB.

Detect it.

Remember that the problem with ECB is that it is stateless and deterministic; the same 16 byte plaintext block will always produce the same 16 byte ciphertext.

*/

func testRepeatedBlocks(ciphertext []byte, blocksize int) bool {
	mapset := make(map[string]bool)
	//iterate over each block
	for i := 0; i < len(ciphertext)-blocksize; i += blocksize {
		//check for membership in the set-map thing
		blk := string(ciphertext[i : i+blocksize])
		if _, ok := mapset[blk]; ok {
			return true
		} else {
			mapset[blk] = true
		}
	}
	return false
}
