package sort

import "math/rand"

const seed = 12345

func Random(n int) []int {
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rnd.Intn(n)
	}
	return arr
}

func Sorted(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func Reversed(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - 1 - i
	}
	return arr
}

func Digits(n int) []int {
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rnd.Intn(10)
	}
	return arr
}

func binarySearch[T Ordered](arr []T, key T, low, high int) int {
	if high <= low {
		if key >= arr[low] {
			return low + 1
		}
		return low
	}
	mid := (low + high) / 2
	if key >= arr[mid] {
		return binarySearch(arr, key, mid+1, high)
	}
	return binarySearch(arr, key, low, mid-1)
}

func swap[T Ordered](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func more[T Ordered](arr []T, i, j int) bool {
	return arr[i] > arr[j]
}

// L=2X+1
// R=2X+2
// P=(x-1)/2
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
