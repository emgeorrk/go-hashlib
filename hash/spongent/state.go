package spongent

type hashState struct {
	value           [nSBox]byte
	messageblock    [rSizeInBytes]byte
	hashbitlen      int
	remainingbitlen int
}

func _init(state *hashState, hashval []byte) int {
	for i := range state.value {
		state.value[i] = 0
	}
	state.hashbitlen = 0
	state.remainingbitlen = 0
	for i := range hashval {
		hashval[i] = 0
	}
	return success
}

func absorb(state *hashState) int {
	for i := 0; i < rSizeInBytes; i++ {
		state.value[i] ^= state.messageblock[i]
	}
	permute(state)
	return success
}

func pad(state *hashState) int {
	byteind := state.remainingbitlen / 8
	bitpos := state.remainingbitlen % 8
	if bitpos > 0 {
		state.messageblock[byteind] &= 0xFF << (8 - bitpos)
		state.messageblock[byteind] |= 0x80 >> bitpos
	} else {
		state.messageblock[byteind] = 0x80
	}
	for i := byteind + 1; i < rSizeInBytes; i++ {
		state.messageblock[i] = 0
	}
	return success
}

func squeeze(state *hashState) int {
	copy(state.messageblock[:], state.value[:rSizeInBytes])
	permute(state)
	return success
}

func permute(state *hashState) {
	IV := byte(0x05)
	for i := 0; i < nRounds; i++ {
		state.value[0] ^= IV
		INV_IV := reverseByte(IV)
		state.value[nSBox-1] ^= INV_IV
		IV = lCounter(IV)
		for j := range state.value {
			state.value[j] = sBoxLayer[state.value[j]]
		}
		pLayer(state)
	}
}
