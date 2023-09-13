package sort

func HeapSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for h := n/2 - 1; h >= 0; h-- {
		heapify(arr, h, n)
	}
	for j := n - 1; j > 0; j-- {
		swap(arr, 0, j)
		heapify(arr, 0, j)
	}
}
