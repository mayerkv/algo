package sort

func MergeSort[T Ordered](arr []T) {
	n := len(arr)
	if n < 2 {
		return
	}
	msort(arr, 0, n-1)
}

func msort[T Ordered](arr []T, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	msort(arr, l, m)
	msort(arr, m+1, r)
	merge(arr, l, m, r)
}

func merge[T Ordered](arr []T, l int, m int, r int) {
	t := make([]T, r-l+1)
	a := l
	b := m + 1
	idx := 0
	for a <= m && b <= r {
		if more(arr, a, b) {
			t[idx] = arr[b]
			b++
		} else {
			t[idx] = arr[a]
			a++
		}
		idx++
	}
	for a <= m {
		t[idx] = arr[a]
		a++
		idx++
	}
	for b <= r {
		t[idx] = arr[b]
		b++
		idx++
	}
	for i := l; i <= r; i++ {
		arr[i] = t[i-l]
	}
}
