package quark

type hashState struct {
	pos int
	x   [width * 8]uint32
}

func initState(state *hashState) {
	for i := 0; i < 8*width; i++ {
		state.x[i] = uint32((iv[i/8] >> (7 - (i % 8))) & 1)
	}
	state.pos = 0
}

func update(state *hashState, data []byte) {
	for _, b := range data {
		for i := 0; i < 8; i++ {
			state.x[8*(width-rate)+state.pos*8+i] ^= uint32((b >> i) & 1)
		}
		state.pos++
		if state.pos == rate {
			permuteD(&state.x)
			state.pos = 0
		}
	}
}

func finalize(state *hashState, out []byte) {
	state.x[8*(width-rate)+state.pos*8] ^= 1
	permuteD(&state.x)

	for i := range out {
		out[i] = 0
	}

	outBytes := 0
	for outBytes < digest {
		for i := 0; i < 8; i++ {
			bit := state.x[8*(width-rate)+i+8*(outBytes%rate)] & 1
			out[outBytes] ^= byte(bit << (7 - i))
		}
		outBytes++
		if outBytes == digest {
			break
		}
		if outBytes%rate == 0 {
			permuteD(&state.x)
		}
	}
}
