package photon

type Photon128 struct{}

func New() *Photon128 {
	return &Photon128{}
}

func (Photon128) Hash(data []byte) []byte {
	msg := pad(data)
	state := make([]byte, t/8)
	copy(state[:t/8-3], make([]byte, t/8-3))
	state[t/8-3] = byte(n / 4)
	state[t/8-2] = byte(r)
	state[t/8-1] = byte(rPrime)

	for i := 0; i < len(msg); i += r / 8 {
		for j := 0; j < r/8; j++ {
			state[j] ^= msg[i+j]
		}
		state = permutation(state)
	}

	hash := make([]byte, 0, n/8)
	for len(hash) < n/8 {
		hash = append(hash, state[:rPrime/8]...)
		state = permutation(state)
	}
	return hash[:n/8]
}
