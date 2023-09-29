package sort

import "sort"

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

func heapify[T Ordered](arr []T, root, size int) {
	x := root
	l := 2*x + 1
	r := 2*x + 2
	if l < size && arr[l] > arr[x] {
		x = l
	}
	if r < size && arr[r] > arr[x] {
		x = r
	}
	if x == root {
		return
	}
	swap(arr, root, x)
	heapify(arr, x, size)
}

func HeapSortT[T sort.Interface](arr T) {
	n := arr.Len()
	if n < 2 {
		return
	}
	for h := n/2 - 1; h >= 0; h-- {
		heapifyT(arr, h, n)
	}
	for j := n - 1; j > 0; j-- {
		arr.Swap(0, j)
		heapifyT(arr, 0, j)
	}
}

func heapifyT[T sort.Interface](arr T, root, n int) {
	x := root
	l := 2*x + 1
	r := 2*x + 2
	if l < n && arr.Less(x, l) {
		x = l
	}
	if r < n && arr.Less(x, r) {
		x = r
	}
	if x == root {
		return
	}
	arr.Swap(root, x)
	heapifyT(arr, x, n)
}
