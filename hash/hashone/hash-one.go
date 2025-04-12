package hashone

import (
	"math"
	"math/big"

	"github.com/emgeorrk/go-hashlib/internal/utils"
)

// HashOne implements the "Hash-One" algorithm
type HashOne struct {
	state [161]bool // 161-bit sponge state
}

// New creates a new HashOne instance with state initialized from Pi
func New() *HashOne {
	h := &HashOne{}
	h.initStateFromPi()
	return h
}

// Hash implements the Hash interface
func (h *HashOne) Hash(data []byte) []byte {
	// Create a copy of the hasher to preserve the original state
	hasher := *h

	// Convert bytes to bits
	bits := utils.BytesToBits(data)

	// Absorbing phase
	for l := 0; l < len(bits); l++ {
		if l == 0 || l == len(bits)-1 {
			// First and last bits: XOR and update 324 times
			hasher.state[160] = utils.BoolXor(hasher.state[160], bits[l])
			for i := 0; i < 324; i++ {
				hasher.updateState()
			}
		} else {
			// Intermediate bits: XOR and update 162 times
			hasher.state[160] = utils.BoolXor(hasher.state[160], bits[l])
			for i := 0; i < 162; i++ {
				hasher.updateState()
			}
		}
	}

	// Squeezing phase
	outputBits := make([]bool, 160)

	// Get the first bit
	outputBits[0] = hasher.state[160]

	// Get the remaining 159 bits
	for i := 1; i < 160; i++ {
		hasher.updateState()
		outputBits[i] = hasher.state[160]
	}

	// Convert bits to hex string
	return utils.BitsToBytes(outputBits)
}

// initStateFromPi initializes the sponge state with the first 161 bits of Pi
func (h *HashOne) initStateFromPi() {
	// Generate binary digits of Pi
	piBits := make([]bool, 161)

	// An approximation of Pi with enough precision for our needs
	pi := big.NewFloat(math.Pi)
	pi.SetPrec(512) // Set precision high enough to get 161 bits

	// Extract 161 bits from Pi
	for i := 0; i < 161; i++ {
		pi = new(big.Float).Mul(pi, big.NewFloat(2))
		intPart, _ := pi.Int(nil)
		bit := intPart.Int64() == 1
		piBits[i] = bit
		pi = new(big.Float).Sub(pi, new(big.Float).SetInt(intPart))
	}

	// Set the state
	copy(h.state[:], piBits)
}

// pf is the non-linear update function for register P
func pf(p0, p11, q23, p55 bool) bool {
	return utils.BoolXor(p0 && p11, p11 && q23, q23 && p55, p55 && p0)
}

// qf is the non-linear update function for register Q
func qf(q0, q25, q41, p48 bool) bool {
	return utils.BoolXor(q0 && q25, q25 && q41, q41 && p48, p48 && q0)
}

// lf is the linear function used in both register updates
func lf(p1, q1, p50 bool) bool {
	return utils.BoolXor(p1, q1, p50)
}

// updateState performs one round of the permutation function
func (h *HashOne) updateState() {
	// Extract registers P and Q from the sponge state
	var P [80]bool
	var Q [81]bool

	copy(P[:], h.state[:80])
	copy(Q[:], h.state[80:])

	// Calculate new rightmost bits
	t1 := utils.BoolXor(pf(P[0], P[11], Q[23], P[55]), lf(P[1], Q[1], P[50]))
	t2 := utils.BoolXor(qf(Q[0], Q[25], Q[41], P[48]), lf(P[1], Q[1], P[50]))

	// Shift registers and update with new bits
	for i := 0; i < 79; i++ {
		h.state[i] = P[i+1]
	}
	h.state[79] = t1

	for i := 0; i < 80; i++ {
		h.state[i+80] = Q[i+1]
	}
	h.state[160] = t2
}
