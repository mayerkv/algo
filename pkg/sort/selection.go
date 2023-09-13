package sort

func SelectionSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := n - 1; i > 0; i-- {
		maxIdx := findMax(arr, i)
		arr[maxIdx], arr[i] = arr[i], arr[maxIdx]
	}
}

func findMax[T Ordered](arr []T, j int) int {
	max := 0
	for i := 1; i <= j; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}
