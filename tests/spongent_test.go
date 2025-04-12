package tests

import (
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestSpongent(t *testing.T) {
	h := hashlib.NewSpongent()

	HashConsistency(t, h)

	HashCollisions(t, h)
}
