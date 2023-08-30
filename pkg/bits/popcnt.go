package bits

func PopCnt1(m uint64) int {
	cnt := 0
	for m > 0 {
		cnt += int(m & 1)
		m >>= 1
	}
	return cnt
}

func PopCnt2(m uint64) int {
	cnt := 0
	for m > 0 {
		cnt++
		m &= m - 1
	}
	return cnt
}

func PopCnt3(m uint64) int {
	cnt := 0
	for m > 0 {
		cnt += bits[m&255]
		m >>= 8
	}
	return cnt
}
