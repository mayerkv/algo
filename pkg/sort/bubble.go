package sort

import "sort"

func BubbleSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for j := n - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func BubbleSortT[T sort.Interface](arr T) {
	n := arr.Len()
	if n < 2 {
		return
	}
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr.Less(j+1, j) {
				arr.Swap(j+1, j)
			}
		}
	}
}
