package asset

import "github.com/gobuffalo/packr"

func Challenge(s string) ([]byte, error) {
	box := packr.NewBox("./")
	val, err := box.Find(s)
	return val, err
}
