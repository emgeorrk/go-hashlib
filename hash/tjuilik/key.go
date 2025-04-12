package tjuilik

func addRoundConstants(s *state, rc0, rc1 uint16) {
	for i := 0; i < 16; i++ {
		if (rc0>>i)&1 == 1 {
			s[(4*i/64)%4][(4*i/16)%4][(4*i)%4] ^= 1
		}
		if (rc1>>i)&1 == 1 {
			s[((4*i+2)/64)%4][((4*i+2)/16)%4][((4*i + 2) % 4)] ^= 1
		}
	}
}

func addKey(s *state, key *state) {
	for z := 0; z < 4; z++ {
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				s[z][y][x] ^= key[z][y][x]
			}
		}
	}
}

func addKeyRotated(s *state, key *state) {
	for i := 0; i < 64; i++ {
		from := (i + 20) % 64
		z1, y1, x1 := (i>>4)&3, (i>>2)&3, i&3
		z2, y2, x2 := (from>>4)&3, (from>>2)&3, from&3
		s[z1][y1][x1] ^= key[z2][y2][x2]
	}
}
