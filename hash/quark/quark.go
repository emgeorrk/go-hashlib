package quark

type DQuark struct{}

func NewDQuark() *DQuark {
	return &DQuark{}
}

func (DQuark) Hash(data []byte) []byte {
	var out [digest]byte
	var state hashState
	initState(&state)
	update(&state, data)
	finalize(&state, out[:])
	return out[:]
}
