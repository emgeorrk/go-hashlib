package tests

import (
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestPhoton128(t *testing.T) {
	h := hashlib.NewPhoton128()

	HashConsistency(t, h)

	HashCollisions(t, h)
}
