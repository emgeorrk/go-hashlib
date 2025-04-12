package spongent

func reverseByte(b byte) byte {
	b = (b&0xF0)>>4 | (b&0x0F)<<4
	b = (b&0xCC)>>2 | (b&0x33)<<2
	b = (b&0xAA)>>1 | (b&0x55)<<1
	return b
}

func lCounter(b byte) byte {
	return ((b << 1) | (((b >> 5) ^ (b >> 4)) & 1)) & 0x3f
}

func pi(i int) int {
	if i != nBits-1 {
		return (i * nBits / 4) % (nBits - 1)
	}
	return nBits - 1
}

func pLayer(state *hashState) {
	var tmp [nSBox]byte
	for i := 0; i < nSBox*8; i++ {
		x := (state.value[i/8] >> (i % 8)) & 1
		j := pi(i)
		tmp[j/8] ^= x << (j % 8)
	}
	copy(state.value[:], tmp[:])
}
