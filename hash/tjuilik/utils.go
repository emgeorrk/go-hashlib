package tjuilik

func padTo128(input []byte) []byte {
	out := make([]byte, 16)
	n := copy(out, input)
	if n < 16 {
		out[n] = 0x80
	}
	return out
}

func padTo32Chunks(input []byte) [][]byte {
	var out [][]byte
	for len(input) > 0 {
		chunk := make([]byte, 4)
		n := copy(chunk, input)
		if n < 4 {
			chunk[n] = 0x80
		}
		out = append(out, chunk)
		if len(input) > 4 {
			input = input[4:]
		} else {
			break
		}
	}
	return out
}

func xorBytes(a, b []byte) []byte {
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return out
}
