package tjuilik

func permute(input []byte) []byte {
	hashState := bytesToState(input)
	const rounds = 32
	const domain = 2
	key := [4][4][4]byte{}

	rcs0, rcs1 := makeRCs(rounds, domain)

	for r := 0; r < rounds; r++ {
		applySBox(&hashState)
		applyLinearLayer(&hashState)
		applySBox(&hashState)

		if r%2 == 0 {
			_SRsheet(&hashState)
			applyLinearLayer(&hashState)
			invSRsheet(&hashState)
			addRoundConstants(&hashState, rcs0[r], rcs1[r])
			addKey(&hashState, (*state)(&key))
		} else {
			_SRslice(&hashState)
			applyLinearLayer(&hashState)
			invSRslice(&hashState)
			addRoundConstants(&hashState, rcs0[r], rcs1[r])
			addKeyRotated(&hashState, (*state)(&key))
		}
	}

	return stateToBytes(hashState)
}
