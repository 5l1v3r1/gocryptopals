package cryptolib

import (
	"bytes"
	"errors"
)

func PKCS7(inblocks []byte, length int) []byte {
	padsize := length - (len(inblocks) % length)
	out := inblocks
	for i := 0; i < padsize; i++ {
		out = append(out, byte(padsize))
	}
	return out
}

func PKCS7Unpad(inval []byte, blocksize int) ([]byte, error) {
	x := int(inval[len(inval)-1])
	if x == 0 || blocksize < x {
		return nil, errors.New("Padding error")
	}
	count := 0
	for i := len(inval) - 1; i >= 0 && count < x; i-- {
		if int(inval[i]) == x {
			count++
			continue
		} else {
			return nil, errors.New("Padding error")
		}
	}
	//return unpadded valz
	r := inval[:len(inval)-x]
	return r, nil
}

//CompareBlocks will compare blocks indexed at x and y, and return true if they are identical
func CompareBlocks(ciphertext []byte, blocksize, x, y int) bool {
	chunks := Chunker(ciphertext, blocksize)
	if bytes.Compare(chunks[x], chunks[y]) == 0 {
		return true
	}
	return false
}

func HasRepeatedBlocks(ciphertext []byte, blocksize int) (bool, int) {
	chunks := Chunker(ciphertext, blocksize)
	prev := []byte{}
	for i, x := range chunks {
		if bytes.Compare(prev, x) == 0 {
			return true, i
		}
		prev = x
	}
	return false, -1
}

//PopFromBlock removes one byte from the beginning of the first array while maintaining block segments
func PopFromBlock(blocks [][]byte) (r [][]byte) {
	bb := bytes.Join(blocks, nil)
	bb = bb[1:len(bb)]
	r = Chunker(bb, len(blocks[0]))
	return
}
