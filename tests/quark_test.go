package tests

import (
	"testing"

	"github.com/emgeorrk/go-hashlib"
)

func TestQuark(t *testing.T) {
	h := hashlib.NewDQuark()

	HashConsistency(t, h)

	HashCollisions(t, h)
}
