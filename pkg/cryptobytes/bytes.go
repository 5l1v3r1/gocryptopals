package cryptobytes

import "github.com/c-sto/gocryptopals/pkg/lang"

//Chunker splits a byte array into chunks of given size, returns an array of arrays
func Chunker(b []byte, n int) (chunks [][]byte) {
	//this is better, thanks swarlz
	if n == 0 {
		return [][]byte{}
	}
	//while i is lower than len of b
	for i := 0; i < len(b); i += n {
		//find last boundary
		nn := i + n
		//chunk last section to end of array
		if nn > len(b) {
			nn = len(b)
		}
		//append chunks as it goes
		chunks = append(chunks, b[i:nn])
	}
	return chunks
}

func Transpose(ss [][]byte) [][]byte {
	// first block will have the total size we need
	ts := make([][]byte, len(ss[0]))
	//iterate over each block
	for _, block := range ss {
		// iterate over each byte in the block
		for i, b := range block {
			//place byte in appropriate transposed block segment
			ts[i] = append(ts[i], b)
		}
	}
	return ts
}

func NormalisedHamming(s []byte, l int) float32 {
	hamming_sum := float32(0)
	for i := 0; i < (len(s)/l - 1); i++ {
		hamming_sum += float32(lang.Hamming(string(s[i*l:(i+1)*l]), string(s[(i+1)*l:(i+2)*l])))
	}
	ham_avg := hamming_sum / float32(len(s)/l-1)
	norm_ham := ham_avg / float32(l)
	return norm_ham
}
