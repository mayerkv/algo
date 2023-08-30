package bits

func Rook(n int) uint64 {
	yM := n % 8
	xM := n - yM
	return (horizontalMask << xM) ^ (verticalMask << yM)
}
