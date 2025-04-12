package hashlib

import (
	"github.com/emgeorrk/go-hashlib/hash/hashone"
	"github.com/emgeorrk/go-hashlib/hash/photon"
	"github.com/emgeorrk/go-hashlib/hash/quark"
	"github.com/emgeorrk/go-hashlib/hash/spongent"
	"github.com/emgeorrk/go-hashlib/hash/tjuilik"
)

func NewHashOne() *hashone.HashOne {
	return hashone.New()
}

func NewPhoton128() *photon.Photon128 {
	return photon.New()
}

func NewDQuark() *quark.DQuark {
	return quark.NewDQuark()
}

func NewSpongent() *spongent.Spongent {
	return spongent.New()
}

func NewTjuilik() *tjuilik.Tjuilik {
	return tjuilik.New()
}
