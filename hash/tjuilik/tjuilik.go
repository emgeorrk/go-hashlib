package tjuilik

type Tjuilik struct{}

func New() *Tjuilik {
	return &Tjuilik{}
}

func (t *Tjuilik) Hash(data []byte) []byte {
	if len(data) == 0 {
		iv := make([]byte, 64)
		iv[0] ^= 1
		return tag256(iv)
	}
	if len(data) <= 16 {
		iv := make([]byte, 64)
		copy(iv, padTo128(data))
		if len(data) < 16 {
			iv[16] ^= 1
		} else {
			iv[16] ^= 2
		}
		return tag256(iv)
	}

	m1 := data[:16]
	rest := padTo32Chunks(data[16:])
	iv := make([]byte, 64)
	copy(iv[:16], m1)

	for _, block := range rest {
		yz := saturninPermute(iv)
		y := xorBytes(yz[:4], block)
		iv = append(y, yz[4:]...)
	}
	iv[16] ^= 2
	return tag256(iv)
}
