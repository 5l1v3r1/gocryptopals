package Set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/c-sto/gocryptopals/asset"

	"github.com/c-sto/gocryptopals/pkg/padding"

	"github.com/c-sto/gocryptopals/pkg/aes"
	"github.com/c-sto/gocryptopals/pkg/cryptobytes"
	"github.com/c-sto/gocryptopals/pkg/lang"
	"github.com/c-sto/gocryptopals/pkg/xor"
)

func Test1(t *testing.T) {
	fmt.Println("Test 1 Begin")
	Challenge1()
}

func Test2(t *testing.T) {
	fmt.Println("Test 2 Begin")
	if xor.XorHexStrings("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965") == "746865206b696420646f6e277420706c6179" {
		fmt.Println("Challenge 2 complete")
	} else {
		t.Error("String output does not match")
	}
	Challenge2()
}

func Test3(t *testing.T) {
	fmt.Println("Test 3 Begin")
	//test our english thingers
	/*
		eng := "The english bit of text that is a bit longer than just a few words.\nThis should probably appear as enlgish in tests."
		notEng := "aasldfkjoivjaodvij f aldskjfqew;klsnc dfwarfe}|}d 349r-0429fds.,aa sdpoifjaefp dfj;ds;sc a;saldkf esaorkap sa;lfdkafp"

		if scorePlaintext(eng) > scorePlaintext(notEng) {
			t.Error("English test failed.")
		}
	*/
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	lowestChar, plain, score := xor.SingleByteXorNTest(hexString)
	fmt.Printf("%v, %v, %v\n", lowestChar, plain, score)
	fmt.Println("Challenge 3 complete")
	Challenge3()
}

func Test4(t *testing.T) {
	fmt.Println("Test 4 Begin")
	//load lines
	content, err := asset.Challenge("challenge4.txt")
	if err != nil {
		t.Error("file load error")
	}
	lines := strings.Split(string(content), "\x0d\n")
	lowestChar, plain, score := xor.MultiSingleByteXorNTest(lines)
	if plain != "Now that the party is jumping\n" {
		t.Error("Incorrect output")
	}
	fmt.Printf("%v %v %v \n", lowestChar, plain, score)
	fmt.Println("Challenge 4 complete")
	Challenge4()
}

func Test5(t *testing.T) {
	fmt.Println("Test 5 Begin")
	text := hex.EncodeToString([]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"))
	key := hex.EncodeToString([]byte("ICE"))
	ciphertext := xor.RepeatingKeyXOR(text, key)
	check := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if ciphertext != check {
		fmt.Printf("%v\n%v\n", ciphertext, check)
		t.Error("Ciphertext mismatch!")
	}
	fmt.Println("Challenge 5 complete")
	Challenge5()
}

func Test6(t *testing.T) {

	fmt.Println("Test 6 Begin")

	content, err := asset.Challenge("challenge6.txt")
	if err != nil {
		t.Error("file load error")
	}
	contentBytes, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		t.Error("b64 decode error")
	}
	if lang.Hamming("this is a test", "wokka wokka!!!") != 37 {
		t.Error("Hamming function incorrect")
	}

	if cryptobytes.NormalisedHamming([]byte("this is a testwokka wokka!!!"), 14) != 37.0/14 {
		t.Error("Norm Hamming incorrect", 37.0/14, cryptobytes.NormalisedHamming([]byte("this is a testwokkawokka!!!"), 14))
	}
	c := cryptobytes.Chunker([]byte("abacada"), 2)
	if len(c) == 4 {
		if !reflect.DeepEqual(c[0], []byte("ab")) {
			t.Error("Chunker fail")
		}

		if !reflect.DeepEqual(c[3], []byte("a")) {
			t.Error("Chunker fail")
		}
	} else {
		t.Error("Chunker fail")
	}

	lol := cryptobytes.Transpose(c)

	if len(lol) == 2 {
		if !reflect.DeepEqual(lol[0], []byte("aaaa")) {
			t.Error("transpose fail2")
		}
		if !reflect.DeepEqual(lol[1], []byte("bcd")) {
			t.Error("Transpose fail3")
		}
	} else {
		t.Error("Transpose fail1")
	}

	plaintext, key := xor.BreakRepeatingKeyXor(contentBytes)

	if len(key) != 29 {
		t.Error("Key length incorrect: ", key)
	}

	if plaintext == "" {
		t.Error("S blank!")
	}
	fmt.Println("Key:\n", string(key))
	fmt.Println("Plaintext:\n", plaintext)
	fmt.Println("Challenge 6 complete")

	Challenge6()
}

func Test7(t *testing.T) {
	fmt.Println("Test 7 Begin")
	content, err := asset.Challenge("challenge7.txt")
	if err != nil {
		t.Error("file load error")
	}
	key := []byte("YELLOW SUBMARINE")

	ciphertext, _ := base64.StdEncoding.DecodeString(string(content))
	plain := aes.AESECBDecrypt(ciphertext, key)

	if plain[0] == 0 || plain[40] == 0 {
		t.Error("Bad decrypt:", string(plain))
	}

	fmt.Println(string(plain))
	fmt.Println("Challenge 7 complete")

	Challenge7()
}

func Test8(t *testing.T) {
	fmt.Println("Test 8 Begin")
	content, err := asset.Challenge("challenge8.txt")
	if err != nil {
		t.Error("file load error")
	}
	lines := strings.Split(string(content), "\x0d\n")
	found := false
	fmt.Println(lines)
	for i, x := range lines {
		s, _ := hex.DecodeString(x)
		if ok, _ := padding.HasRepeatedBlocks(s, 16); ok {
			fmt.Println("Identified repeated block on line:", i)
			found = true
		}
	}
	if !found {
		t.Error("No duplicate block found!?!?")
	}
	fmt.Println("Challenge 8 complete")

	Challenge8()
}
