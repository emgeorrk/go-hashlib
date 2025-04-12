package photon

func pad(data []byte) []byte {
	padLen := r/8 - (len(data)+1)%(r/8)
	padded := append(data, 0x80)
	padded = append(padded, make([]byte, padLen)...)
	return padded
}

func fromMatrix(S [][]byte) []byte {
	state := make([]byte, t/8)
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			byteIndex := (i*d + j) / 2
			if (i*d+j)%2 == 0 {
				state[byteIndex] &= 0x0F
				state[byteIndex] |= S[i][j] << 4
			} else {
				state[byteIndex] &= 0xF0
				state[byteIndex] |= S[i][j]
			}
		}
	}
	return state
}

func addConstants(S [][]byte, round int) {
	for i := 0; i < d; i++ {
		S[i][0] ^= roundConstants[round] ^ ICd[i]
	}
}

func subCells(S [][]byte) {
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			S[i][j] = sbox4[S[i][j]]
		}
	}
}

func shiftRows(S [][]byte) {
	for i := 0; i < d; i++ {
		row := make([]byte, d)
		for j := 0; j < d; j++ {
			row[j] = S[i][(j+i)%d]
		}
		copy(S[i], row)
	}
}

func mixColumnsSerial(S [][]byte) {
	for j := 0; j < d; j++ {
		for k := 0; k < d; k++ {
			val := byte(0)
			for l := 0; l < d; l++ {
				val ^= gfMult(Z[(k-l+d)%d], S[l][j])
			}
			S[k][j] = val
		}
	}
}

func gfMult(a, b byte) byte {
	res := byte(0)
	for i := 0; i < 4; i++ {
		if (b & (1 << i)) != 0 {
			res ^= a << i
		}
	}
	for i := 7; i >= 4; i-- {
		if (res & (1 << i)) != 0 {
			res ^= gfPoly << (i - 4)
		}
	}
	return res & 0x0F
}
