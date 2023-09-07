package bits

const (
	leftMask       uint64 = 0xfefefefefefefefe
	rightMask             = 0x7f7f7f7f7f7f7f7f
	horizontalMask        = 0xff
	verticalMask          = 0x101010101010101
	nA                    = leftMask
	nAB                   = 0xfcfcfcfcfcfcfcfc
	nH                    = rightMask
	nGH                   = 0x3f3f3f3f3f3f3f3f
	diagonalLeft          = 0x102040810204080
	a1h8diagonal          = 0x8040201008040201
)

var bits [256]int

func init() {
	for i := 0; i < 256; i++ {
		bits[i] = PopCnt2(uint64(i))
	}
}
