package tjuilik

func tag256(t0 []byte) []byte {
	t1 := saturninPermute(t0)
	t2 := saturninPermute(t1)
	return append(t1[:16], t2[:16]...)
}
