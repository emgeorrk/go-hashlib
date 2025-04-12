package examples

import (
	"fmt"
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestExampleSpongent(t *testing.T) {
	h := hashlib.NewSpongent()

	str := "Hello World!"

	hashed := h.Hash([]byte(str))

	fmt.Printf("%s -> %v\n", str, hashed)
}
