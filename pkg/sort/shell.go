package sort

func Shell[T Ordered](arr []T) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for j := gap; j < n; j++ {
			for i := j; i >= gap && arr[i-gap] > arr[i]; i -= gap {
				arr[i-gap], arr[i] = arr[i], arr[i-gap]
			}
		}
	}
}
