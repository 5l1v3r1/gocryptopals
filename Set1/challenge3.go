package Set1

import (
	"encoding/hex"
	"strings"
)

/*
Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

Achievement Unlocked
You now have our permission to make "ETAOIN SHRDLU" jokes on Twitter.
*/

//this file needs refactoring

func singleByteXorNTest(ciphertext string) (bestChar string, plaintext string, score float32) {
	hexString := ciphertext
	byteString, err := hex.DecodeString(hexString)
	if err != nil {
		return "", "", 99999
	}
	lowestVal := float32(999999999)
	lowestChar := ""
	plain := ""
	for i := 0; i < 255; i++ {
		xorCandidate := strings.Repeat(string(i), len(byteString))
		hexXorCandidate := hex.EncodeToString([]byte(xorCandidate))
		decodedXor, _ := hex.DecodeString(xorHexStrings(hexString, hexXorCandidate))
		if score := scorePlaintext(string(decodedXor)); score < lowestVal {
			lowestVal = score
			lowestChar = string(i)
			plain = string(decodedXor)
		}
	}
	return lowestChar, plain, lowestVal
}

func scorePlaintext(candidate string) float32 {
	unig := map[string]float32{
		"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 12.702,
		"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
		"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
		"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
		"u": 2.758, "v": 0.978, "w": 2.360, "x": 0.150, "y": 1.974,
		"z": 0.074, " ": 55.00, "\n": 0.01, ".": 6.53, ",": 6.16,
		";": 0.32, ":": 0.34, "!": 0.33, "?": 0.56, "'": 2.43,
		"-": 1.53,
	}
	score := float32(0)
	lowCand := strings.ToLower(candidate)
	//check unigrams
	for k, v := range unig {
		score += chiSquare(lowCand, k, v)
	}
	//bump score for non-printables and uncommon symbols
	for _, letter := range candidate {
		if letter == '\n' {
			continue
		}
		if _, ok := unig[string(letter)]; ok {
			continue
		}
		letterDec := int(letter)
		if letterDec < 32 || letterDec > 126 {
			score += chiSquare(lowCand, string(letter), 0.001)
		} else if letterDec < 65 || letterDec > 122 || (letterDec < 97 && letterDec > 90) {
			score += chiSquare(lowCand, string(letter), 0.01)
		}
	}
	return score
}

// like strings.count but counts overlapping
func countOverlapping(text string, substring string) int {
	c := 0
	for i := 0; i < (len(text) - len(substring)); i++ {
		window := text[i : i+len(substring)]
		if window == substring {
			c++
		}
	}
	return c
}

func chiSquare(text string, substring string, freq float32) float32 {
	count := float32(countOverlapping(text, substring))
	expected := freq * 0.01 * float32(len(text))
	return ((count - expected) * (count - expected)) / expected
}
