package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	challenge1()
}

func Test2(t *testing.T) {

	if xorHexStrings("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965") == "746865206b696420646f6e277420706c6179" {
		fmt.Println("Challenge 2 complete")
	} else {
		t.Error("String output does not match")
	}

}

func Test3(t *testing.T) {
	// wow this is a gross test, PLS REFACTOR ME
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	byteString, _ := hex.DecodeString(hexString)
	lowestVal := float32(99999)
	lowestChar := "aa"
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
	fmt.Printf("%v, %v, %v\n", lowestVal, lowestChar, plain)

}
