package tjuilik

const (
	poly0 = 0x1002D
	poly1 = 0x10053
)

func makeRCs(rounds int, domain byte) ([]uint16, []uint16) {
	rc0 := uint32(0xFE00 | (rounds << 4) | int(domain))
	rc1 := rc0
	rcs0 := make([]uint16, rounds)
	rcs1 := make([]uint16, rounds)

	for i := 0; i < rounds; i++ {
		rcs0[i] = uint16(rc0 & 0xFFFF)
		rcs1[i] = uint16(rc1 & 0xFFFF)
		rc0 = clockLFSR(rc0, poly0)
		rc1 = clockLFSR(rc1, poly1)
	}
	return rcs0, rcs1
}

func clockLFSR(rc uint32, poly uint32) uint32 {
	if (rc & 0x8000) == 0 {
		return rc << 1
	}
	return (rc << 1) ^ poly
}
