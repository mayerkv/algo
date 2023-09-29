package sort

func RadixSort[T Int](arr []T) {
	n := len(arr)

	var max T
	for i := 0; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	for exp := 1; max/T(exp) > 0; exp *= 10 {
		countSortExp(arr, exp)
	}
}

func countSortExp[T Int](arr []T, exp int) {
	n := len(arr)
	output := make([]T, n)
	counts := make([]int, 10)

	for _, v := range arr {
		p := v / T(exp) % 10
		counts[p]++
	}

	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		p := arr[i] / T(exp) % 10
		output[counts[p]-1] = arr[i]
		counts[p]--
	}

	copy(arr, output)
}
