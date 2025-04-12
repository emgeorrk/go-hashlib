package examples

import (
	"fmt"
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestExampleQuark(t *testing.T) {
	h := hashlib.NewDQuark()

	str := "Hello World!"

	hashed := h.Hash([]byte(str))

	fmt.Printf("%s -> %v\n", str, hashed)
}
