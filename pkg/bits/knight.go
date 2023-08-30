package bits

func Knight(n int) uint64 {
	var pos uint64 = 1 << n
	return nGH&(pos<<6|pos>>10) |
		nH&(pos<<15|pos>>17) |
		nA&(pos<<17|pos>>15) |
		nAB&(pos<<10|pos>>6)
}
