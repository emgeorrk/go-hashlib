package tjuilik

type state [4][4][4]byte // [z][y][x]

func bytesToState(data []byte) state {
	var s state
	for i := 0; i < 64; i++ {
		z := (i >> 4) & 0x3
		y := (i >> 2) & 0x3
		x := i & 0x3
		s[z][y][x] = data[i]
	}
	return s
}

func stateToBytes(s state) []byte {
	out := make([]byte, 64)
	for i := 0; i < 64; i++ {
		z := (i >> 4) & 0x3
		y := (i >> 2) & 0x3
		x := i & 0x3
		out[i] = s[z][y][x]
	}
	return out
}
