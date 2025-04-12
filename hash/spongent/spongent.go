package spongent

type Spongent struct{}

func New() *Spongent {
	return &Spongent{}
}

func (Spongent) Hash(data []byte) []byte {
	var out [hashSize / 8]byte
	spongentHash(data, len(data)*8, &out)
	return out[:]
}

func spongentHash(data []byte, databitlen int, hashval *[hashSize / 8]byte) int {
	var state hashState
	_init(&state, hashval[:])
	for databitlen >= rate {
		copy(state.messageblock[:], data[:rSizeInBytes])
		absorb(&state)
		databitlen -= rate
		data = data[rSizeInBytes:]
	}
	if databitlen > 0 {
		copy(state.messageblock[:], data[:(databitlen+7)/8])
		state.remainingbitlen = databitlen
	} else {
		for i := range state.messageblock {
			state.messageblock[i] = 0
		}
		state.remainingbitlen = 0
	}
	pad(&state)
	absorb(&state)
	state.hashbitlen += rate
	for state.hashbitlen < hashSize {
		squeeze(&state)
		copy(hashval[state.hashbitlen/8:], state.messageblock[:])
		state.hashbitlen += rate
	}
	copy(hashval[hashSize/8-len(state.value):], state.value[:])
	return success
}
