package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("Test 1 Begin")
	challenge1()
}

func Test2(t *testing.T) {
	fmt.Println("Test 2 Begin")
	if xorHexStrings("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965") == "746865206b696420646f6e277420706c6179" {
		fmt.Println("Challenge 2 complete")
	} else {
		t.Error("String output does not match")
	}

}

func Test3(t *testing.T) {
	fmt.Println("Test 3 Begin")
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	lowestChar, plain, score := singleByteXorNTest(hexString)
	fmt.Printf("%v, %v, %v\n", lowestChar, plain, score)
	fmt.Println("Challenge 3 complete")
}

func Test4(t *testing.T) {
	fmt.Println("Test 4 Begin")
	//load lines
	content, err := ioutil.ReadFile("../resources/challenge4.txt")
	if err != nil {
		t.Error("file load error")
	}
	lines := strings.Split(string(content), "\x0d\n")
	lowestChar, plain, score := multiSingleByteXorNTest(lines)
	fmt.Printf("%v %v %v \n", lowestChar, plain, score)
	fmt.Println("Challenge 4 complete")

}
