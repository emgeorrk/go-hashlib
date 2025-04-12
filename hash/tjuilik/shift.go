package tjuilik

func _SRsheet(s *state) {
	for z := 0; z < 4; z++ {
		for y := 0; y < 4; y++ {
			row := [4]byte{}
			for x := 0; x < 4; x++ {
				row[x] = s[z][y][x]
			}
			row = rotLeft(row, y)
			for x := 0; x < 4; x++ {
				s[z][y][x] = row[x]
			}
		}
	}
}

func invSRsheet(s *state) {
	for z := 0; z < 4; z++ {
		for y := 0; y < 4; y++ {
			row := [4]byte{}
			for x := 0; x < 4; x++ {
				row[x] = s[z][y][x]
			}
			row = rotRight(row, y)
			for x := 0; x < 4; x++ {
				s[z][y][x] = row[x]
			}
		}
	}
}

func _SRslice(s *state) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			col := [4]byte{}
			for z := 0; z < 4; z++ {
				col[z] = s[z][y][x]
			}
			col = rotLeft(col, x)
			for z := 0; z < 4; z++ {
				s[z][y][x] = col[z]
			}
		}
	}
}

func invSRslice(s *state) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			col := [4]byte{}
			for z := 0; z < 4; z++ {
				col[z] = s[z][y][x]
			}
			col = rotRight(col, x)
			for z := 0; z < 4; z++ {
				s[z][y][x] = col[z]
			}
		}
	}
}

func rotLeft(a [4]byte, n int) [4]byte {
	var out [4]byte
	for i := 0; i < 4; i++ {
		out[i] = a[(i+n)%4]
	}
	return out
}

func rotRight(a [4]byte, n int) [4]byte {
	var out [4]byte
	for i := 0; i < 4; i++ {
		out[i] = a[(i+4-n)%4]
	}
	return out
}
