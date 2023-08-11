package lucky

func getNextArr(prevArr []int) []int {
	newLen := len(prevArr) + 9
	arr := make([]int, newLen)
	for i := 0; i < newLen; i++ {
		var q int
		for j := 0; j < 10; j++ {
			idx := i - j
			if idx >= 0 && idx < len(prevArr) {
				q += prevArr[idx]
			}
		}
		arr[i] = q
	}
	return arr
}

func LuckyTicketsCount(n int) int {
	var result int
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 0; i < n-1; i++ {
		arr = getNextArr(arr)
	}
	for _, i := range arr {
		result += i * i
	}
	return result
}
