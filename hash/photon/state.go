package photon

func permutation(state []byte) []byte {
	S := toMatrix(state)
	for round := 0; round < nRounds; round++ {
		addConstants(S, round)
		subCells(S)
		shiftRows(S)
		mixColumnsSerial(S)
	}
	return fromMatrix(S)
}

func toMatrix(state []byte) [][]byte {
	S := make([][]byte, d)
	for i := range S {
		S[i] = make([]byte, d)
		for j := 0; j < d; j++ {
			byteIndex := (i*d + j) / 2
			if (i*d+j)%2 == 0 {
				S[i][j] = state[byteIndex] >> 4
			} else {
				S[i][j] = state[byteIndex] & 0x0F
			}
		}
	}
	return S
}
