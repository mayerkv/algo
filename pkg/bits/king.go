package bits

func King(n int) uint64 {
	var pos uint64 = 1 << n
	leftPos := pos & leftMask
	rightPos := pos & rightMask
	return (leftPos << 7) | (pos << 8) | (rightPos << 9) |
		(leftPos >> 1) | (rightPos << 1) |
		(leftPos >> 9) | (pos >> 8) | (rightPos >> 7)
}
