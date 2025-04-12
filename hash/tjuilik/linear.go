package tjuilik

func applyLinearLayer(s *state) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			var col [4]byte
			for z := 0; z < 4; z++ {
				col[z] = s[z][y][x]
			}
			col = mixColumn(col)
			for z := 0; z < 4; z++ {
				s[z][y][x] = col[z]
			}
		}
	}
}

func alpha(x byte) byte {
	x &= 0x0F
	x0 := (x >> 0) & 1
	x1 := (x >> 1) & 1
	x2 := (x >> 2) & 1
	x3 := (x >> 3) & 1

	y0 := x1
	y1 := x2
	y2 := x3
	y3 := x0 ^ x1

	return (y3 << 3) | (y2 << 2) | (y1 << 1) | y0
}

func mixColumn(col [4]byte) [4]byte {
	a, b, c, d := col[0], col[1], col[2], col[3]

	return [4]byte{
		alpha(alpha(a)) ^ alpha(alpha(b)) ^ alpha(c) ^ d,
		a ^ alpha(b) ^ b ^ alpha(alpha(c)) ^ c ^ alpha(alpha(d)) ^ d,
		a ^ b ^ alpha(alpha(c)) ^ alpha(alpha(d)) ^ alpha(d),
		alpha(alpha(a)) ^ a ^ alpha(alpha(b)) ^ alpha(b) ^ b ^ c ^ alpha(d) ^ d,
	}
}
