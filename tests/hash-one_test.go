package tests

import (
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestHashOne(t *testing.T) {
	h := hashlib.NewHashOne()

	HashConsistency(t, h)

	HashCollisions(t, h)
}
