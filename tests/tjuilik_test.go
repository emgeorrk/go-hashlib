package tests

import (
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestTjuilik(t *testing.T) {
	h := hashlib.NewTjuilik()

	HashConsistency(t, h)

	HashCollisions(t, h)
}
