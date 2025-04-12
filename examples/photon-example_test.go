package examples

import (
	"fmt"
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestExamplePhoton128(t *testing.T) {
	h := hashlib.NewPhoton128()

	str := "Hello World!"

	hashed := h.Hash([]byte(str))

	fmt.Printf("%s -> %v\n", str, hashed)
}
