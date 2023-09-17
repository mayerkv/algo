package sort

func QuickSort[T Ordered](arr []T) {
	n := len(arr)
	if n < 2 {
		return
	}
	qsort(arr, 0, n-1)
}

func qsort[T Ordered](arr []T, l, r int) {
	if l >= r {
		return
	}
	m := split(arr, l, r)
	qsort(arr, l, m-1)
	qsort(arr, m+1, r)
}

func split[T Ordered](arr []T, l, r int) int {
	p := arr[r]
	m := l - 1
	for j := l; j <= r; j++ {
		if arr[j] <= p {
			m += 1
			swap(arr, m, j)
		}
	}
	return m
}
