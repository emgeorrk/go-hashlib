package examples

import (
	"fmt"
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestExampleHashOne(t *testing.T) {
	h := hashlib.NewHashOne()

	str := "Hello World!"

	hashed := h.Hash([]byte(str))

	fmt.Printf("%s -> %v\n", str, hashed)
}
