package Set2

import (
	//"encoding/base64"
	//"encoding/hex"
	"fmt"
	//"io/ioutil"
	"reflect"
	"testing"
)

func Test9(t *testing.T) {
	fmt.Println("Begin Test 9")
	x := PKCS7([]byte("YELLOW SUBMARINE"), 20)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		t.Error("Bad padding:", x)
	}
	x = PKCS7([]byte("YELLOW SUBMARINE"), 19)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x03\x03\x03")) {
		t.Error("Bad padding:", x)
	}
	x = PKCS7([]byte("YELLOW SUBMARINE"), 15)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")) {
		t.Error("Bad padding:", x, []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e"))
	}

}

func test10(t *testing.T) {
	content, err := ioutil.ReadFile("../resources/challenge10.txt")
	if err != nil {
		t.Error("file load error")
	}
	contentBytes, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		t.Error("b64 decode error")
	}

	key := []byte("YELLOW SUBMARINE")
}
