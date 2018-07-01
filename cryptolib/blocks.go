package cryptolib

func PKCS7(inblocks []byte, length int) []byte {
	padsize := length - (len(inblocks) % length)
	out := inblocks
	for i := 0; i < padsize; i++ {
		out = append(out, byte(padsize))
	}
	return out
}

func HasRepeatedBlocks(ciphertext []byte, blocksize int) bool {
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
