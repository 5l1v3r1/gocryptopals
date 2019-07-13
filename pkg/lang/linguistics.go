package lang

import (
	"fmt"
	"strings"
)

//expected occurences of each individual character in a typical english block of text
var unig = map[string]float32{
	"a": 8.167, "b": 1.492, "c": 2.782, "d": 4.253, "e": 12.702,
	"f": 2.228, "g": 2.015, "h": 6.094, "i": 6.966, "j": 0.153,
	"k": 0.772, "l": 4.025, "m": 2.406, "n": 6.749, "o": 7.507,
	"p": 1.929, "q": 0.095, "r": 5.987, "s": 6.327, "t": 9.056,
	"u": 2.758, "v": 0.978, "w": 2.360, "x": 0.150, "y": 1.974,
	"z": 0.074, " ": 55.00, "\n": 0.01, ".": 6.53, ",": 6.16,
	";": 0.32, ":": 0.34, "!": 0.33, "?": 0.56, "'": 2.43,
	"-": 1.53,
}

//hamming calculates the hamming distance between two given strings
func Hamming(s1 string, s2 string) int {
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

func ScorePlaintext(candidate string) float32 {
	checked := map[rune]bool{}
	score := float32(0)
	//make it all lowercase to avoid dealing with too much garbage I guess
	lowCand := strings.ToLower(candidate)
	//check each letter for expected occurence
	for k, v := range unig {
		score += chiSquare(lowCand, k, v)
	}

	//bump score for non-printables and uncommon symbols
	for _, letter := range candidate {
		if _, ok := checked[letter]; ok {
			continue //already checked it
		}
		if letter == '\n' { //don't score newlines
			continue
		}
		//check if the candidate is in our test block
		if _, ok := unig[string(letter)]; ok {
			continue
		}
		//get the decimal representation of the byte
		letterDec := int(letter)
		//if it's non-ascii, we don't expect it to be there at all (0.001% chance)
		if letterDec < 32 || letterDec > 126 {
			score += chiSquare(lowCand, string(letter), 0.001)
			//it's ascii, but not common
		} else if letterDec < 65 || //symbols, numbers etc I think?
			letterDec > 122 || //
			(letterDec < 97 && letterDec > 90) { //
			//expect to see it (0.01%) of the time
			score += chiSquare(lowCand, string(letter), 0.01)
		}
		checked[letter] = true
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
