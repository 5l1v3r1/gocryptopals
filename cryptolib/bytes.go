package cryptolib

//Chunker splits a byte array into chunks of given size, returns an array of arrays
func chunker(s []byte, chunksize int) [][]byte {
	r := make([][]byte, (len(s)/chunksize)+1)
	j := 0

	for i := 0; i < len(s)-1; i += chunksize {
		r[j] = s[i : i+chunksize]
		j++
	}
	if j < len(r) {
		r[j] = s[(len(s)/chunksize)*chunksize:]
	}

	//clean blank slices idk why (fix this better)
	re := [][]byte{}
	for _, x := range r {
		if len(x) > 0 {
			re = append(re, x)
		}
	}
	return re
}

func transpose(ss [][]byte) [][]byte {
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

func normalisedHamming(s []byte, l int) float32 {
	hamming_sum := float32(0)
	for i := 0; i < (len(s)/l - 1); i++ {
		hamming_sum += float32(Hamming(string(s[i*l:(i+1)*l]), string(s[(i+1)*l:(i+2)*l])))
	}
	ham_avg := hamming_sum / float32(len(s)/l-1)
	norm_ham := ham_avg / float32(l)
	return norm_ham
}
