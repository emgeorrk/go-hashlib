package photon

const (
	n       = 128 // output size in bits
	r       = 16  // bitrate in bits
	rPrime  = 16  // squeezing rate in bits
	t       = 144 // state size in bits
	d       = 6   // matrix dimension
	nRounds = 12
)

var (
	sbox4          = [16]byte{0xc, 0x5, 0x6, 0xb, 0x9, 0x0, 0xa, 0xd, 0x3, 0xe, 0xf, 0x8, 0x4, 0x7, 0x1, 0x2}
	roundConstants = [12]byte{1, 3, 7, 14, 13, 11, 6, 12, 9, 2, 5, 10}
	ICd            = [d]byte{0, 1, 3, 7, 6, 4}
	Z              = [d]byte{1, 2, 8, 5, 8, 2}
	gfPoly         = byte(0x13)
)
