package Set2

import (
	//"encoding/base64"
	//"encoding/hex"

	"fmt"

	//"io/ioutil"
	"reflect"
	"testing"

	"github.com/c-sto/gocryptopals/pkg/padding"
)

func Test9(t *testing.T) {
	fmt.Println("Begin Test 9")
	x := padding.PKCS7([]byte("YELLOW SUBMARINE"), 20)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")) {
		t.Error("Bad padding:", x)
	}
	x = padding.PKCS7([]byte("YELLOW SUBMARINE"), 19)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x03\x03\x03")) {
		t.Error("Bad padding:", x)
	}
	x = padding.PKCS7([]byte("YELLOW SUBMARINE"), 15)
	if !reflect.DeepEqual(x, []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e")) {
		t.Error("Bad padding:", x, []byte("YELLOW SUBMARINE\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e\x0e"))
	}

	Challenge9()

}

func Test10(t *testing.T) {
	Challenge10()
}

func Test11(t *testing.T) {
	Challenge11()
}

func Test12(t *testing.T) {
	Challenge12()
}

func Test13(t *testing.T) {
	Challenge13()
}

func Test14(t *testing.T) {
	Challenge14()
}

func Test15(t *testing.T) {
	Challenge15()
}

func Test16(t *testing.T) {
	Challenge16()
}
