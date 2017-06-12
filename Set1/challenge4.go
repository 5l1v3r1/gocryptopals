package Set1

/*

One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)

*/

func multiSingleByteXorNTest(rows []string) (bestchar string, plaintext string, bestScore float32) {
	bestScore = float32(99999999)
	bestchar = "aa"
	plaintext = ""

	for _, line := range rows {
		char, plain, score := singleByteXorNTest(line)
		if score < bestScore {
			bestScore = score
			bestchar = char
			plaintext = plain
		}
	}
	return bestchar, plaintext, bestScore
}
