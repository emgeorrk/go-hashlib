package examples

import (
	"fmt"
	"testing"

	"github.com/emgeorrk/go-hashlib/hash/tjuilik"
)

func TestExampleTjuilik(t *testing.T) {
	h := tjuilik.New()

	str := "Hello World!"

	hashed := h.Hash([]byte(str))

	fmt.Printf("%s -> %v\n", str, hashed)
}
