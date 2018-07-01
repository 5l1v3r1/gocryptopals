package cryptolib

import (
	"encoding/hex"
	"strings"
	"sync"
)

func XorBytes(b1, b2 []byte) []byte {
	//use shortest
	n := len(b1)
	if len(b1) > len(b2) {
		n = len(b2)
	}
	ret := make([]byte, n)
	//xor the things
	for i := 0; i < n; i++ {
		ret[i] = b1[i] ^ b2[i]
	}
	return ret
}

func XorHexStrings(arg1 string, arg2 string) string {
	//turn args to bytes
	bytes1, err := hex.DecodeString(arg1)
	if err != nil {
		panic(err)
	}
	bytes2, err := hex.DecodeString(arg2)
	if err != nil {
		panic(err)
	}

	ret := XorBytes(bytes1, bytes2)
	//turn return bytes back to hex
	hexret := hex.EncodeToString(ret)
	return hexret
}

//SingleByteXorNTest Checks a ciphertext for the most english-like plaintext given a single byte xor test
func SingleByteXorNTest(ciphertext string) (bestChar string, plaintext string, score float32) {
	hexString := ciphertext
	byteString, err := hex.DecodeString(hexString)
	if err != nil {
		return "", "", 99999
	}
	lowestVal := float32(999999999)
	lowestChar := ""
	plain := ""
	//concurrentize it imo; minimal performance gain so no (but see below!)
	for i := 0; i < 255; i++ {
		xorCandidate := strings.Repeat(string(i), len(byteString))
		hexXorCandidate := hex.EncodeToString([]byte(xorCandidate))
		decodedXor, _ := hex.DecodeString(XorHexStrings(hexString, hexXorCandidate))
		if score := ScorePlaintext(string(decodedXor)); score < lowestVal {
			lowestVal = score
			lowestChar = string(i)
			plain = string(decodedXor)
		}
	}
	return lowestChar, plain, lowestVal
}

func RepeatingKeyXOR(arg1 string, arg2 string) string {
	//turn args to bytes
	bytes1, _ := hex.DecodeString(arg1)
	bytes2, _ := hex.DecodeString(arg2)
	//set lengths
	shortbytes := bytes1
	longbytes := bytes2
	if len(bytes1) > len(bytes2) {
		shortbytes = bytes2
		longbytes = bytes1
	}
	ret := make([]byte, len(longbytes))
	//xor the things
	for i := 0; i < len(longbytes); i++ {
		ret[i] = longbytes[i] ^ shortbytes[i%len(shortbytes)]
	}

	//turn return bytes back to hex
	hexret := hex.EncodeToString(ret)
	return string(hexret)
}

type testResult struct {
	Char  string
	Plain string
	Score float32
}

//MultiSingleByteXorNTest performs a check against several ciphertexts to check if any give good english looking results
func MultiSingleByteXorNTest(rows []string) (bestchar string, plaintext string, bestScore float32) {
	bestScore = float32(99999999)
	bestchar = "aa"
	plaintext = ""
	resultChan := make(chan testResult, 10)
	count := 0
	wg := &sync.WaitGroup{}
	// Asyn for the lolz, it's way faster but also ugly
	go func() {
		defer func() {
			wg.Wait()
			close(resultChan)
		}()
		for _, line := range rows {
			count++
			wg.Add(1)
			go func(l string) {
				char, plain, score := SingleByteXorNTest(l)
				resultChan <- testResult{
					Char:  char,
					Plain: plain,
					Score: score,
				}
				wg.Done()
			}(line)
		}

	}()
	for x := range resultChan {
		if x.Score < bestScore {
			bestScore = x.Score
			bestchar = x.Char
			plaintext = x.Plain
		}
	}
	return bestchar, plaintext, bestScore
}

func BreakRepeatingKeyXor(s1 []byte) (string, []byte) {
	// s1 is raw ciphertext bytes
	hexS1 := hex.EncodeToString(s1)
	//get keysize
	ks := getKeysize(s1)
	// break ciphertext into blocks	of keysize length
	chunks := chunker(s1, ks)
	// transpose the blocks: one block of all first bytes, one of all second bytes, etc
	transposed := transpose(chunks)
	// send each block to the single byte xor solver. Get the output from that to have the key
	//extractedKey := make([]string, ks)
	k := []byte("")
	for _, transposedBlock := range transposed {
		// singleByteXorNTest(transposedBlock)
		x, _, _ := SingleByteXorNTest(hex.EncodeToString(transposedBlock))
		k = append(k, byte(x[0]))
	}
	hexK := hex.EncodeToString(k)
	bkn := RepeatingKeyXOR(hexK, hexS1) //outputs a hex string
	rawbkn, _ := hex.DecodeString(bkn)
	return string(rawbkn), k
}

//getkeysize tries to dtermine the most likely length for a repeating xor key
func getKeysize(s []byte) int {
	//key range
	var fromKey = 1
	var toKey = 50
	min := float32(9999999)
	size := -1
	//for each keysize
	for i := fromKey; i <= toKey; i++ {
		dist := normalisedHamming(s, i)
		//remember the lowest hamming distance (it's probably the key)
		if dist < min {
			min = dist
			size = i
		}
	}
	return size

}
