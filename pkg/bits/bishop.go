package bits

func Bishop(n int) uint64 {
	yM := n % 9
	//xM := n % 7

	//fmt.Println((a1h8diagonal >> yM))
	//fmt.Println((diagonalLeft >> xM))

	var a uint64
	if yM > 8-yM {
		a = a1h8diagonal >> (yM * 8)
	} else {
		a = a1h8diagonal << (yM * 8)
	}

	return a
}

/*
0000000100000000000000000000000000000000000000000000000000000000
0000001000000001000000000000000000000000000000000000000000000000
0000010000000010000000010000000000000000000000000000000000000000
0000100000000100000000100000000100000000000000000000000000000000
0001000000001000000001000000001000000001000000000000000000000000
0010000000010000000010000000010000000010000000010000000000000000
0100000000100000000100000000100000000100000000100000000100000000
1000000001000000001000000001000000001000000001000000001000000001
0000000010000000010000000010000000010000000010000000010000000010

0000000100000000100000000100000000100000000100000000100000000100 --
0000000000000000100000000100000000100000000100000000100000000100

0000000000000000000000001000000001000000001000000001000000001000
0000000000000000000000000000000010000000010000000010000000010000
0000000000000000000000000000000000000000100000000100000000100000
0000000000000000000000000000000000000000000000001000000001000000
0000000000000000000000000000000000000000000000000000000010000000
*/
